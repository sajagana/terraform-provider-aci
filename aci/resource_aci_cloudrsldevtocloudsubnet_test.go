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

func TestAccAciRelationfromCloudLDevtoCloudSubnet_Basic(t *testing.T) {
	var relationfrom_cloud_ldevto_cloud_subnet models.RelationfromCloudLDevtoCloudSubnet
	fv_tenant_name := acctest.RandString(5)
	cloud_ldev_name := acctest.RandString(5)
	cloud_rs_ldev_to_cloud_subnet_name := acctest.RandString(5)
	description := "relationfrom_cloud_ldevto_cloud_subnet created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciRelationfromCloudLDevtoCloudSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciRelationfromCloudLDevtoCloudSubnetConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetExists("aci_relationfrom_cloud_ldevto_cloud_subnet.foo_relationfrom_cloud_ldevto_cloud_subnet", &relationfrom_cloud_ldevto_cloud_subnet),
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetAttributes(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name, description, &relationfrom_cloud_ldevto_cloud_subnet),
				),
			},
		},
	})
}

func TestAccAciRelationfromCloudLDevtoCloudSubnet_Update(t *testing.T) {
	var relationfrom_cloud_ldevto_cloud_subnet models.RelationfromCloudLDevtoCloudSubnet
	fv_tenant_name := acctest.RandString(5)
	cloud_ldev_name := acctest.RandString(5)
	cloud_rs_ldev_to_cloud_subnet_name := acctest.RandString(5)
	description := "relationfrom_cloud_ldevto_cloud_subnet created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciRelationfromCloudLDevtoCloudSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciRelationfromCloudLDevtoCloudSubnetConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetExists("aci_relationfrom_cloud_ldevto_cloud_subnet.foo_relationfrom_cloud_ldevto_cloud_subnet", &relationfrom_cloud_ldevto_cloud_subnet),
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetAttributes(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name, description, &relationfrom_cloud_ldevto_cloud_subnet),
				),
			},
			{
				Config: testAccCheckAciRelationfromCloudLDevtoCloudSubnetConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetExists("aci_relationfrom_cloud_ldevto_cloud_subnet.foo_relationfrom_cloud_ldevto_cloud_subnet", &relationfrom_cloud_ldevto_cloud_subnet),
					testAccCheckAciRelationfromCloudLDevtoCloudSubnetAttributes(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name, description, &relationfrom_cloud_ldevto_cloud_subnet),
				),
			},
		},
	})
}

func testAccCheckAciRelationfromCloudLDevtoCloudSubnetConfig_basic(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "foo_tenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_l4-l7device" "foo_cloud_l4-l7device" {
		name 		= "%s"
		description = "cloud_l4-l7device created while acceptance testing"
		tenant_dn = aci_tenant.foo_tenant.id
	}

	resource "aci_relationfrom_cloud_ldevto_cloud_subnet" "foo_relationfrom_cloud_ldevto_cloud_subnet" {
		name 		= "%s"
		description = "relationfrom_cloud_ldevto_cloud_subnet created while acceptance testing"
		cloud_l4-l7device_dn = aci_cloud_l4-l7device.foo_cloud_l4-l7device.id
	}

	`, fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name)
}

func testAccCheckAciRelationfromCloudLDevtoCloudSubnetExists(name string, relationfrom_cloud_ldevto_cloud_subnet *models.RelationfromCloudLDevtoCloudSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Relation from Cloud LDev to Cloud Subnet %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Relation from Cloud LDev to Cloud Subnet dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		relationfrom_cloud_ldevto_cloud_subnetFound := models.RelationfromCloudLDevtoCloudSubnetFromContainer(cont)
		if relationfrom_cloud_ldevto_cloud_subnetFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Relation from Cloud LDev to Cloud Subnet %s not found", rs.Primary.ID)
		}
		*relationfrom_cloud_ldevto_cloud_subnet = *relationfrom_cloud_ldevto_cloud_subnetFound
		return nil
	}
}

func testAccCheckAciRelationfromCloudLDevtoCloudSubnetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aci_relationfrom_cloud_ldevto_cloud_subnet" {
			cont, err := client.Get(rs.Primary.ID)
			relationfrom_cloud_ldevto_cloud_subnet := models.RelationfromCloudLDevtoCloudSubnetFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Relation from Cloud LDev to Cloud Subnet %s Still exists", relationfrom_cloud_ldevto_cloud_subnet.DistinguishedName)
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckAciRelationfromCloudLDevtoCloudSubnetAttributes(fv_tenant_name, cloud_ldev_name, cloud_rs_ldev_to_cloud_subnet_name, description string, relationfrom_cloud_ldevto_cloud_subnet *models.RelationfromCloudLDevtoCloudSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if cloud_rs_ldev_to_cloud_subnet_name != GetMOName(relationfrom_cloud_ldevto_cloud_subnet.DistinguishedName) {
			return fmt.Errorf("Bad cloudrs_ldev_to_cloud_subnet %s", GetMOName(relationfrom_cloud_ldevto_cloud_subnet.DistinguishedName))
		}

		if cloud_ldev_name != GetMOName(GetParentDn(relationfrom_cloud_ldevto_cloud_subnet.DistinguishedName, relationfrom_cloud_ldevto_cloud_subnet.Rn)) {
			return fmt.Errorf(" Bad cloudldev %s", GetMOName(GetParentDn(relationfrom_cloud_ldevto_cloud_subnet.DistinguishedName, relationfrom_cloud_ldevto_cloud_subnet.Rn)))
		}
		if description != relationfrom_cloud_ldevto_cloud_subnet.Description {
			return fmt.Errorf("Bad relationfrom_cloud_ldevto_cloud_subnet Description %s", relationfrom_cloud_ldevto_cloud_subnet.Description)
		}
		return nil
	}
}
