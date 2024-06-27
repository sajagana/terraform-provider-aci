// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvIpAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvIpAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "no"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvIpAttrAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "yes"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvIpAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvIpAttrResetDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "no"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_epg_useg_ip_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "no"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvIpAttrChildrenDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "no"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_epg_useg_ip_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "use_fv_subnet", "no"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvIpAttrChildrenRemoveFromConfigDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvIpAttrChildrenRemoveOneDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvIpAttrChildrenRemoveAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_ip_attribute.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvIpAttrMinDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  ip = "131.107.1.200"
  name = "131"
}
`

const testConfigFvIpAttrAllDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "131"
  annotation = "annotation"
  description = "description"
  ip = "0.0.0.0"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
  use_fv_subnet = "yes"
}
`

const testConfigFvIpAttrResetDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "131"
  annotation = "orchestrator:terraform"
  description = ""
  ip = "131.107.1.200"
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  use_fv_subnet = "no"
}
`
const testConfigFvIpAttrChildrenDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  ip = "131.107.1.200"
  name = "131"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigFvIpAttrChildrenRemoveFromConfigDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  ip = "131.107.1.200"
  name = "131"
}
`

const testConfigFvIpAttrChildrenRemoveOneDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  ip = "131.107.1.200"
  name = "131"
  annotations = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigFvIpAttrChildrenRemoveAllDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  ip = "131.107.1.200"
  name = "131"
  annotations = []
  tags = []
}
`
