// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvIdGroupAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvIdGroupAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvIdGroupAttrAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", "owner_tag"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvIdGroupAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvIdGroupAttrResetDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_epg_useg_criterion_identity_group_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigFvIdGroupAttrChildrenDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_epg_useg_criterion_identity_group_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "selector", "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvIdGroupAttrChildrenRemoveFromConfigDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvIdGroupAttrChildrenRemoveOneDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvIdGroupAttrChildrenRemoveAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_criterion_identity_group_attribute.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvIdGroupAttrMinDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
}
`

const testConfigFvIdGroupAttrAllDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
}
`

const testConfigFvIdGroupAttrResetDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigFvIdGroupAttrChildrenDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
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

const testConfigFvIdGroupAttrChildrenRemoveFromConfigDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
}
`

const testConfigFvIdGroupAttrChildrenRemoveOneDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
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

const testConfigFvIdGroupAttrChildrenRemoveAllDependencyWithFvCrtrn = testConfigFvCrtrnMin + `
resource "aci_epg_useg_criterion_identity_group_attribute" "test" {
  parent_dn = aci_epg_useg_criterion.test.id
  selector = "adepg/authsvr-common-sg1-ISE_1/grpcont/dom-cisco.com/grp-Eng"
  annotations = []
  tags = []
}
`
