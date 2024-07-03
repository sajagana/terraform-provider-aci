// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvSCrtrnWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvSCrtrnDataSourceDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "name", "sub_criterion"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "match", "any"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigFvSCrtrnNotExistingFvCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_sub_block_statement data source"),
			},
		},
	})
}
func TestAccDataSourceFvSCrtrnWithFvSCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvSCrtrnDataSourceDependencyWithFvSCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "name", "sub_criterion"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "match", "any"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_sub_block_statement.test1", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigFvSCrtrnNotExistingFvSCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_sub_block_statement data source"),
			},
		},
	})
}

const testConfigFvSCrtrnDataSourceDependencyWithFvCrtrn = testConfigFvSCrtrnMinDependencyWithFvCrtrn + `
data "aci_epg_useg_sub_block_statement" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "sub_criterion"
  depends_on = [aci_epg_useg_sub_block_statement.test]
}
`

const testConfigFvSCrtrnNotExistingFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
data "aci_epg_useg_sub_block_statement" "test_non_existing" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "sub_criterion_non_existing"
}
`
const testConfigFvSCrtrnDataSourceDependencyWithFvSCrtrn = testConfigFvSCrtrnMinDependencyWithFvSCrtrn + `
data "aci_epg_useg_sub_block_statement" "test1" {
  parent_dn = aci_epg_useg_sub_block_statement.test.id
  name = "sub_criterion"
  depends_on = [aci_epg_useg_sub_block_statement.test1]
}
`

const testConfigFvSCrtrnNotExistingFvSCrtrn = testConfigFvSCrtrnMinDependencyWithFvCrtrn + `
data "aci_epg_useg_sub_block_statement" "test_non_existing" {
  parent_dn = aci_epg_useg_sub_block_statement.test.id
  name = "sub_criterion_non_existing"
}
`