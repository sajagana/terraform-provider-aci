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

func TestAccAciAnnotationtocapturerbacinfo_Basic(t *testing.T) {
	var annotationtocapturerbacinfo models.Annotationtocapturerbacinfo
	monitor_telemetry_pol_name := acctest.RandString(5)
	telemetry_interface_flow_filter_name := acctest.RandString(5)
	monitor_telemetry_pol_name := acctest.RandString(5)
	hostprot_pol_name := acctest.RandString(5)
	coop_pol_name := acctest.RandString(5)
	aaa_rbac_annotation_name := acctest.RandString(5)
	description := "annotationtocapturerbacinfo created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAnnotationtocapturerbacinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAnnotationtocapturerbacinfoConfig_basic(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAnnotationtocapturerbacinfoExists("aci_annotationtocapturerbacinfo.foo_annotationtocapturerbacinfo", &annotationtocapturerbacinfo),
					testAccCheckAciAnnotationtocapturerbacinfoAttributes(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name, description, &annotationtocapturerbacinfo),
				),
			},
		},
	})
}

func TestAccAciAnnotationtocapturerbacinfo_Update(t *testing.T) {
	var annotationtocapturerbacinfo models.Annotationtocapturerbacinfo
	monitor_telemetry_pol_name := acctest.RandString(5)
	telemetry_interface_flow_filter_name := acctest.RandString(5)
	monitor_telemetry_pol_name := acctest.RandString(5)
	hostprot_pol_name := acctest.RandString(5)
	coop_pol_name := acctest.RandString(5)
	aaa_rbac_annotation_name := acctest.RandString(5)
	description := "annotationtocapturerbacinfo created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAnnotationtocapturerbacinfoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAnnotationtocapturerbacinfoConfig_basic(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAnnotationtocapturerbacinfoExists("aci_annotationtocapturerbacinfo.foo_annotationtocapturerbacinfo", &annotationtocapturerbacinfo),
					testAccCheckAciAnnotationtocapturerbacinfoAttributes(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name, description, &annotationtocapturerbacinfo),
				),
			},
			{
				Config: testAccCheckAciAnnotationtocapturerbacinfoConfig_basic(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAnnotationtocapturerbacinfoExists("aci_annotationtocapturerbacinfo.foo_annotationtocapturerbacinfo", &annotationtocapturerbacinfo),
					testAccCheckAciAnnotationtocapturerbacinfoAttributes(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name, description, &annotationtocapturerbacinfo),
				),
			},
		},
	})
}

func testAccCheckAciAnnotationtocapturerbacinfoConfig_basic(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name string) string {
	return fmt.Sprintf(`

	resource "aci_access_node_configuration" "foo_access_node_configuration" {
		name 		= "%s"
		description = "access_node_configuration created while acceptance testing"

	}

	resource "aci_logical_modelconfigholderforperinterface_flow_telemetry_filter" "foo_logical_modelconfigholderforperinterface_flow_telemetry_filter" {
		name 		= "%s"
		description = "logical_modelconfigholderforperinterface_flow_telemetry_filter created while acceptance testing"
		access_node_configuration_dn = aci_access_node_configuration.foo_access_node_configuration.id
	}

	resource "aci_access_node_configuration" "foo_access_node_configuration" {
		name 		= "%s"
		description = "access_node_configuration created while acceptance testing"
		logical_modelconfigholderforperinterface_flow_telemetry_filter_dn = aci_logical_modelconfigholderforperinterface_flow_telemetry_filter.foo_logical_modelconfigholderforperinterface_flow_telemetry_filter.id
	}

	resource "aci_host_protection_domain_policy" "foo_host_protection_domain_policy" {
		name 		= "%s"
		description = "host_protection_domain_policy created while acceptance testing"
		access_node_configuration_dn = aci_access_node_configuration.foo_access_node_configuration.id
	}

	resource "aci_coopgroup_policy" "foo_coopgroup_policy" {
		name 		= "%s"
		description = "coopgroup_policy created while acceptance testing"
		host_protection_domain_policy_dn = aci_host_protection_domain_policy.foo_host_protection_domain_policy.id
	}

	resource "aci_annotationtocapturerbacinfo" "foo_annotationtocapturerbacinfo" {
		name 		= "%s"
		description = "annotationtocapturerbacinfo created while acceptance testing"
		coopgroup_policy_dn = aci_coopgroup_policy.foo_coopgroup_policy.id
	}

	`, monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name)
}

func testAccCheckAciAnnotationtocapturerbacinfoExists(name string, annotationtocapturerbacinfo *models.Annotationtocapturerbacinfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Annotation to capture rbac info %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Annotation to capture rbac info dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		annotationtocapturerbacinfoFound := models.AnnotationtocapturerbacinfoFromContainer(cont)
		if annotationtocapturerbacinfoFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Annotation to capture rbac info %s not found", rs.Primary.ID)
		}
		*annotationtocapturerbacinfo = *annotationtocapturerbacinfoFound
		return nil
	}
}

func testAccCheckAciAnnotationtocapturerbacinfoDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aci_annotationtocapturerbacinfo" {
			cont, err := client.Get(rs.Primary.ID)
			annotationtocapturerbacinfo := models.AnnotationtocapturerbacinfoFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Annotation to capture rbac info %s Still exists", annotationtocapturerbacinfo.DistinguishedName)
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckAciAnnotationtocapturerbacinfoAttributes(monitor_telemetry_pol_name, telemetry_interface_flow_filter_name, monitor_telemetry_pol_name, hostprot_pol_name, coop_pol_name, aaa_rbac_annotation_name, description string, annotationtocapturerbacinfo *models.Annotationtocapturerbacinfo) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if aaa_rbac_annotation_name != GetMOName(annotationtocapturerbacinfo.DistinguishedName) {
			return fmt.Errorf("Bad aaarbac_annotation %s", GetMOName(annotationtocapturerbacinfo.DistinguishedName))
		}

		if coop_pol_name != GetMOName(GetParentDn(annotationtocapturerbacinfo.DistinguishedName, annotationtocapturerbacinfo.Rn)) {
			return fmt.Errorf(" Bad cooppol %s", GetMOName(GetParentDn(annotationtocapturerbacinfo.DistinguishedName, annotationtocapturerbacinfo.Rn)))
		}
		if description != annotationtocapturerbacinfo.Description {
			return fmt.Errorf("Bad annotationtocapturerbacinfo Description %s", annotationtocapturerbacinfo.Description)
		}
		return nil
	}
}
