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

func TestAccAciCloudL4L7LogicalInterface_Basic(t *testing.T) {
	var cloud_l4_l7_logical_interface models.CloudL4L7LogicalInterface
	fv_tenant_name := acctest.RandString(5)
	cloud_ldev_name := acctest.RandString(5)
	cloud_lif_name := acctest.RandString(5)
	description := "cloud_l4_l7_logical_interface created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudL4L7LogicalInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudL4L7LogicalInterfaceConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_lif_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7LogicalInterfaceExists("aci_cloud_l4_l7_logical_interface.foo_cloud_l4_l7_logical_interface", &cloud_l4_l7_logical_interface),
					testAccCheckAciCloudL4L7LogicalInterfaceAttributes(fv_tenant_name, cloud_ldev_name, cloud_lif_name, description, &cloud_l4_l7_logical_interface),
				),
			},
		},
	})
}

func TestAccAciCloudL4L7LogicalInterface_Update(t *testing.T) {
	var cloud_l4_l7_logical_interface models.CloudL4L7LogicalInterface
	fv_tenant_name := acctest.RandString(5)
	cloud_ldev_name := acctest.RandString(5)
	cloud_lif_name := acctest.RandString(5)
	description := "cloud_l4_l7_logical_interface created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudL4L7LogicalInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudL4L7LogicalInterfaceConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_lif_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7LogicalInterfaceExists("aci_cloud_l4_l7_logical_interface.foo_cloud_l4_l7_logical_interface", &cloud_l4_l7_logical_interface),
					testAccCheckAciCloudL4L7LogicalInterfaceAttributes(fv_tenant_name, cloud_ldev_name, cloud_lif_name, description, &cloud_l4_l7_logical_interface),
				),
			},
			{
				Config: testAccCheckAciCloudL4L7LogicalInterfaceConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_lif_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudL4L7LogicalInterfaceExists("aci_cloud_l4_l7_logical_interface.foo_cloud_l4_l7_logical_interface", &cloud_l4_l7_logical_interface),
					testAccCheckAciCloudL4L7LogicalInterfaceAttributes(fv_tenant_name, cloud_ldev_name, cloud_lif_name, description, &cloud_l4_l7_logical_interface),
				),
			},
		},
	})
}

func testAccCheckAciCloudL4L7LogicalInterfaceConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_lif_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "foo_tenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_l4_l7_device" "foo_cloud_l4_l7_device" {
		name 		= "%s"
		description = "cloud_l4_l7_device created while acceptance testing"
		tenant_dn = aci_tenant.foo_tenant.id
	}

	resource "aci_cloud_l4_l7_logical_interface" "foo_cloud_l4_l7_logical_interface" {
		name 		= "%s"
		description = "cloud_l4_l7_logical_interface created while acceptance testing"
		cloud_l4_l7_device_dn = aci_cloud_l4_l7_device.foo_cloud_l4_l7_device.id
	}

	`, fv_tenant_name, cloud_ldev_name, cloud_lif_name)
}

func testAccCheckAciCloudL4L7LogicalInterfaceExists(name string, cloud_l4_l7_logical_interface *models.CloudL4L7LogicalInterface) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud L4-L7 Logical Interface %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud L4-L7 Logical Interface dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_l4_l7_logical_interfaceFound := models.CloudL4L7LogicalInterfaceFromContainer(cont)
		if cloud_l4_l7_logical_interfaceFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud L4-L7 Logical Interface %s not found", rs.Primary.ID)
		}
		*cloud_l4_l7_logical_interface = *cloud_l4_l7_logical_interfaceFound
		return nil
	}
}

func testAccCheckAciCloudL4L7LogicalInterfaceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aci_cloud_l4_l7_logical_interface" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_l4_l7_logical_interface := models.CloudL4L7LogicalInterfaceFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud L4-L7 Logical Interface %s Still exists", cloud_l4_l7_logical_interface.DistinguishedName)
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckAciCloudL4L7LogicalInterfaceAttributes(fv_tenant_name, cloud_ldev_name, cloud_lif_name, description string, cloud_l4_l7_logical_interface *models.CloudL4L7LogicalInterface) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if cloud_lif_name != GetMOName(cloud_l4_l7_logical_interface.DistinguishedName) {
			return fmt.Errorf("Bad cloudlif %s", GetMOName(cloud_l4_l7_logical_interface.DistinguishedName))
		}

		if cloud_ldev_name != GetMOName(GetParentDn(cloud_l4_l7_logical_interface.DistinguishedName, cloud_l4_l7_logical_interface.Rn)) {
			return fmt.Errorf(" Bad cloudldev %s", GetMOName(GetParentDn(cloud_l4_l7_logical_interface.DistinguishedName, cloud_l4_l7_logical_interface.Rn)))
		}
		if description != cloud_l4_l7_logical_interface.Description {
			return fmt.Errorf("Bad cloud_l4_l7_logical_interface Description %s", cloud_l4_l7_logical_interface.Description)
		}
		return nil
	}
}
