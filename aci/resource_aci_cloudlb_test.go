package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAciCloudL4L7Devices_Basic(t *testing.T) {
	var cloud_l4l7load_balancer models.CloudL4L7Devices
	fv_tenant_name := acctest.RandString(5)
	cloud_lb_name := acctest.RandString(5)
	description := "cloud_l4l7load_balancer created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudL4L7DevicesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudL4L7DevicesConfig_basic(fv_tenant_name, cloud_lb_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7DevicesExists("aci_cloud_l4l7load_balancer.foo_cloud_l4l7load_balancer", &cloud_l4l7load_balancer),
					testAccCheckAciCloudL4L7DevicesAttributes(fv_tenant_name, cloud_lb_name, description, &cloud_l4l7load_balancer),
				),
			},
		},
	})
}

func TestAccAciCloudL4L7Devices_Update(t *testing.T) {
	var cloud_l4l7load_balancer models.CloudL4L7Devices
	fv_tenant_name := acctest.RandString(5)
	cloud_lb_name := acctest.RandString(5)
	description := "cloud_l4l7load_balancer created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudL4L7DevicesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudL4L7DevicesConfig_basic(fv_tenant_name, cloud_lb_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7DevicesExists("aci_cloud_l4l7load_balancer.foo_cloud_l4l7load_balancer", &cloud_l4l7load_balancer),
					testAccCheckAciCloudL4L7DevicesAttributes(fv_tenant_name, cloud_lb_name, description, &cloud_l4l7load_balancer),
				),
			},
			{
				Config: testAccCheckAciCloudL4L7DevicesConfig_basic(fv_tenant_name, cloud_lb_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7DevicesExists("aci_cloud_l4l7load_balancer.foo_cloud_l4l7load_balancer", &cloud_l4l7load_balancer),
					testAccCheckAciCloudL4L7DevicesAttributes(fv_tenant_name, cloud_lb_name, description, &cloud_l4l7load_balancer),
				),
			},
		},
	})
}

func testAccCheckAciCloudL4L7DevicesConfig_basic(fv_tenant_name, cloud_lb_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "foo_tenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_l4l7load_balancer" "foo_cloud_l4l7load_balancer" {
		name 		= "%s"
		description = "cloud_l4l7load_balancer created while acceptance testing"
		tenant_dn = aci_tenant.foo_tenant.id
	}

	`, fv_tenant_name, cloud_lb_name)
}

func testAccCheckAciCloudL4L7DevicesExists(name string, cloud_l4l7load_balancer *models.CloudL4L7Devices) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud L4L7 Load Balancer %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud L4L7 Load Balancer dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_l4l7load_balancerFound := models.CloudL4L7DevicesFromContainer(cont)
		if cloud_l4l7load_balancerFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud L4L7 Load Balancer %s not found", rs.Primary.ID)
		}
		*cloud_l4l7load_balancer = *cloud_l4l7load_balancerFound
		return nil
	}
}

func testAccCheckAciCloudL4L7DevicesDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aci_cloud_l4l7load_balancer" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_l4l7load_balancer := models.CloudL4L7DevicesFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud L4L7 Load Balancer %s Still exists", cloud_l4l7load_balancer.DistinguishedName)
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckAciCloudL4L7DevicesAttributes(fv_tenant_name, cloud_lb_name, description string, cloud_l4l7load_balancer *models.CloudL4L7Devices) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if cloud_lb_name != GetMOName(cloud_l4l7load_balancer.DistinguishedName) {
			return fmt.Errorf("Bad cloudlb %s", GetMOName(cloud_l4l7load_balancer.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloud_l4l7load_balancer.DistinguishedName, cloud_l4l7load_balancer.Rn)) {
			return fmt.Errorf(" Bad fvtenant %s", GetMOName(GetParentDn(cloud_l4l7load_balancer.DistinguishedName, cloud_l4l7load_balancer.Rn)))
		}
		if description != cloud_l4l7load_balancer.Description {
			return fmt.Errorf("Bad cloud_l4l7load_balancer Description %s", cloud_l4l7load_balancer.Description)
		}
		return nil
	}
}
