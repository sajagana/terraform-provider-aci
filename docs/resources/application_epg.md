---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_application_epg"
sidebar_current: "docs-aci-resource-aci_application_epg"
description: |-
  Manages ACI Application EPG
---

# aci_application_epg #

Manages ACI Application EPG



## API Information ##

* Class: [fvAEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAEPg/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-{name}/ap-{name}/epg-{name}`

## GUI Information ##

* Location: `Tenants -> Application Profiles -> Application EPGs`

## Example Usage ##

The configuration snippet below creates a Application EPG with only required attributes.

```hcl

resource "aci_application_epg" "example_application_profile" {
  parent_dn = aci_application_profile.example.id
  name      = "test_name"
}

```
The configuration snippet below shows all possible attributes of the Application EPG.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_application_epg" "full_example_application_profile" {
  parent_dn              = aci_application_profile.example.id
  annotation             = "annotation"
  description            = "description_1"
  contract_exception_tag = "contract_exception_tag_1"
  flood_in_encapsulation = "disabled"
  forwarding_control     = "proxy-arp"
  has_multicast_source   = "no"
  useg_epg               = "no"
  match_criteria         = "All"
  name                   = "test_name"
  name_alias             = "name_alias_1"
  intra_epg_isolation    = "enforced"
  preferred_group_member = "exclude"
  priority               = "level1"
  admin_state            = "no"
  epg_useg_block_statement = [
    {
      annotation  = "annotation_1"
      description = "description_1"
      match       = "all"
      name        = "criterion"
      name_alias  = "name_alias_1"
      owner_key   = "owner_key_1"
      owner_tag   = "owner_tag_1"
      precedence  = "1"
      scope       = "scope-bd"
    }
  ]
  relation_to_application_epg_monitoring_policy = [
    {
      annotation             = "annotation_1"
      monitoring_policy_name = aci_monitoring_policy.example.name
    }
  ]
  relation_to_bridge_domain = [
    {
      annotation         = "annotation_1"
      bridge_domain_name = aci_bridge_domain.example.name
    }
  ]
  relation_to_consumed_contracts = [
    {
      annotation    = "annotation_1"
      priority      = "level1"
      contract_name = aci_contract.example.name
    }
  ]
  relation_to_imported_contracts = [
    {
      annotation             = "annotation_1"
      priority               = "level1"
      imported_contract_name = aci_imported_contract.example.name
    }
  ]
  relation_to_custom_qos_policy = [
    {
      annotation             = "annotation_1"
      custom_qos_policy_name = aci_custom_qos_policy.example.name
    }
  ]
  relation_to_domains = [
    {
      annotation                    = "annotation_1"
      binding_type                  = "dynamicBinding"
      class_preference              = "encap"
      custom_epg_name               = "custom_epg_name_1"
      delimiter                     = "@"
      encapsulation                 = "vlan-100"
      encapsulation_mode            = "auto"
      epg_cos                       = "Cos0"
      epg_cos_pref                  = "disabled"
      deployment_immediacy          = "immediate"
      ipam_dhcp_override            = "10.0.0.2"
      ipam_enabled                  = "yes"
      ipam_gateway                  = "10.0.0.1"
      lag_policy_name               = "lag_policy_name_1"
      netflow_direction             = "both"
      enable_netflow                = "disabled"
      number_of_ports               = "1"
      port_allocation               = "elastic"
      primary_encapsulation         = "vlan-200"
      primary_encapsulation_inner   = "vlan-300"
      resolution_immediacy          = "immediate"
      secondary_encapsulation_inner = "vlan-400"
      switching_mode                = "AVE"
      target_dn                     = "uni/vmmp-VMware/dom-domain_1"
      untagged                      = "no"
      vnet_only                     = "no"
    }
  ]
  relation_to_data_plane_policing_policy = [
    {
      annotation                      = "annotation_1"
      data_plane_policing_policy_name = aci_data_plane_policing_policy.example.name
    }
  ]
  relation_to_fibre_channel_paths = [
    {
      annotation  = "annotation_1"
      description = "description_1"
      target_dn   = "topology/pod-1/paths-101/pathep-[eth1/1]"
      vsan        = "vsan-10"
      vsan_mode   = "native"
    }
  ]
  relation_to_intra_epg_contracts = [
    {
      annotation    = "annotation_1"
      contract_name = aci_contract.example.name
    }
  ]
  relation_to_static_leafs = [
    {
      annotation           = "annotation_1"
      description          = "description_1"
      encapsulation        = "vlan-100"
      deployment_immediacy = "immediate"
      mode                 = "native"
      target_dn            = "topology/pod-1/node-101"
    }
  ]
  relation_to_static_paths = [
    {
      annotation            = "annotation_1"
      description           = "description_1"
      encapsulation         = "vlan-202"
      deployment_immediacy  = "immediate"
      mode                  = "native"
      primary_encapsulation = "vlan-203"
      target_dn             = "topology/pod-1/paths-101/pathep-[eth1/1]"
    }
  ]
  relation_to_taboo_contracts = [
    {
      annotation          = "annotation_1"
      taboo_contract_name = aci_taboo_contract.example.name
    }
  ]
  relation_to_provided_contracts = [
    {
      annotation     = "annotation_1"
      match_criteria = "All"
      priority       = "level1"
      contract_name  = aci_contract.example.name
    }
  ]
  relation_to_contract_masters = [
    {
      annotation = "annotation_1"
      target_dn  = aci_application_epg.test_application_epg_0.id
    }
  ]
  relation_to_trust_control_policy = [
    {
      annotation                = "annotation_1"
      trust_control_policy_name = aci_trust_control_policy.example.name
    }
  ]
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}

```

All examples for the Application EPG resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_application_epg) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_application_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_profile) ([fvAp](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAp/overview))
* `name` (name) - (string) The name of the Application EPG object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Application EPG object.

### Optional ###
  
* `annotation` (annotation) - (string) The annotation of the Application EPG object.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the Application EPG object.
* `contract_exception_tag` (exceptionTag) - (string) The contract exception tag of the Application EPG object.
* `flood_in_encapsulation` (floodOnEncap) - (string) Flood L2 Multicast/Broadcast and Link Local Layer based on encapsulation.
  - Default: `disabled`
  - Valid Values: `disabled`, `enabled`.
* `forwarding_control` (fwdCtrl) - (string) The forwarding control of the Application EPG object.
  - Default: `none`
  - Valid Values: `none`, `proxy-arp`.
* `has_multicast_source` (hasMcastSource) - (string) The Application EPG object has a multicast source.
  - Default: `no`
  - Valid Values: `no`, `yes`.
* `useg_epg` (isAttrBasedEPg) - (string) The Application EPG object is microsegmented (uSeg).
  - Default: `no`
  - Valid Values: `no`, `yes`.
* `match_criteria` (matchT) - (string) The provider label match criteria.
  - Default: `AtleastOne`
  - Valid Values: `All`, `AtleastOne`, `AtmostOne`, `None`.
* `name_alias` (nameAlias) - (string) The name alias of the Application EPG object.
* `intra_epg_isolation` (pcEnfPref) - (string) Parameter used to determine whether communication between endpoints within the EPG is blocked.
  - Default: `unenforced`
  - Valid Values: `enforced`, `unenforced`.
* `pc_tag` (pcTag) - (string) The classification tag used for policy enforcement and zoning.
* `preferred_group_member` (prefGrMemb) - (string) Parameter used to determine whether the EPG is part of the preferred group. Members of this group are allowed to communicate without contracts.
  - Default: `exclude`
  - Valid Values: `exclude`, `include`.
* `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
  - Default: `unspecified`
  - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.
* `admin_state` (shutdown) - (string) Withdraw AEPg Configuration from all Nodes in the Fabric.
  - Default: `no`
  - Valid Values: `no`, `yes`.

* `epg_useg_block_statement` - (list) A list of EPG uSeg Block Statement (ACI object [fvCrtrn](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCrtrn/overview)). EPG uSeg Block Statement can also be configured using a separate [aci_](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/) resource. This attribute is supported in ACI versions: 1.1(1j) and later.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the EPG uSeg Block Statement object.
      - Default: `orchestrator:terraform`
  * `description` (descr) - (string) The description of the EPG uSeg Block Statement object.
  * `match` (match) - (string) The Matching Rule Type of the EPG uSeg Block Statement object.
      - Default: `any`
      - Valid Values: `all`, `any`.
  * `name` (name) - (string) The name of the EPG uSeg Block Statement object.
  * `name_alias` (nameAlias) - (string) The name alias of the EPG uSeg Block Statement object.
  * `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
  * `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
  * `precedence` (prec) - (string) The precedence of the EPG uSeg Block Statement object.
  * `scope` (scope) - (string) The scope of the EPG uSeg Block Statement object.
      - Default: `scope-bd`
      - Valid Values: `scope-bd`, `scope-vrf`.

* `relation_to_application_epg_monitoring_policy` - (list) A list of Relation To Application EPG Monitoring Policy (ACI object [fvRsAEPgMonPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsAEPgMonPol/overview)) pointing to Monitoring Policy (ACI Object [monEPGPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/monEPGPol/overview)) which can be configured using the [aci_monitoring_policy](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/monitoring_policy) resource.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Application EPG Monitoring Policy object.
      - Default: `orchestrator:terraform`
  * `monitoring_policy_name` (tnMonEPGPolName) - (string) The name of the monitoring policy.

* `relation_to_bridge_domain` - (list) A list of Relation To Bridge Domain (ACI object [fvRsBd](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsBd/overview)) pointing to Bridge Domain (ACI Object [fvBD](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvBD/overview)) which can be configured using the [aci_bridge_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/bridge_domain) resource.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Bridge Domain object.
      - Default: `orchestrator:terraform`
  * `bridge_domain_name` (tnFvBDName) - (string) The name of the bridge domain associated with this EPG.

* `relation_to_consumed_contracts` - (list) A list of Relation To Consumed Contracts (ACI object [fvRsCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCons/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource.
  
  #### Required ####
  
  * `contract_name` (tnVzBrCPName) - (string) The consumer contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Consumed Contract object.
      - Default: `orchestrator:terraform`
  * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.

* `relation_to_imported_contracts` - (list) A list of Relation To Imported Contracts (ACI object [fvRsConsIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsConsIf/overview)) pointing to Imported Contract (ACI Object [vzCPIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzCPIf/overview)) which can be configured using the [aci_imported_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/imported_contract) resource.
  
  #### Required ####
  
  * `imported_contract_name` (tnVzCPIfName) - (string) The contract interface name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/imported_contract) with `aci_imported_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/imported_contract) with `data.aci_imported_contract.example.name`.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Imported Contract object.
      - Default: `orchestrator:terraform`
  * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.

* `relation_to_custom_qos_policy` - (list) A list of Relation To Custom Qos Policy (ACI object [fvRsCustQosPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCustQosPol/overview)) pointing to Custom Qos Policy (ACI Object [qosCustomPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/qosCustomPol/overview)) which can be configured using the [aci_custom_qos_policy](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/custom_qos_policy) resource.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Custom Qos Policy object.
      - Default: `orchestrator:terraform`
  * `custom_qos_policy_name` (tnQosCustomPolName) - (string) The Custom QoS traffic policy name.

      
* `relation_to_domains` - (list) A list of Relation To Domains. This relation can point to multiple ACI objects:
    - [vmmDomP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vmmDomP/overview) which can be configured using the [aci_vmm_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/vmm_domain) resource.
    - [physDomP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/physDomP/overview) which can be configured using the [aci_physical_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/physical_domain) resource.
    - [fcDomP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fcDomP/overview). Currently there is no resource is available to configure this ACI object.
    - [l2extDomP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l2extDomP/overview). Currently there is no resource is available to configure this ACI object.
  
  #### Required ####
  
  * `target_dn` (tDn) - (string) The distinguished name of the target Domain object.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Domain object.
      - Default: `orchestrator:terraform`
  * `binding_type` (bindingType) - (string) The binding type of the Relation To Domain object.
      - Default: `none`
      - Valid Values: `dynamicBinding`, `ephemeral`, `none`, `staticBinding`.
  * `class_preference` (classPref) - (string) The class preference of the Relation To Domain object. Set 'useg' to allow microsegmentation.
      - Default: `encap`
      - Valid Values: `encap`, `useg`.
  * `custom_epg_name` (customEpgName) - (string) The display name of the user configured port-group.
  * `delimiter` (delimiter) - (string) The delimiter of the Relation To Domain object.
  * `encapsulation` (encap) - (string) The encapsulation of the Relation To Domain object. The encapsulation refers to the EPG VLAN when class preference is set to 'encap, or to the Secondary VLAN when class preference is set to 'useg'.
  * `encapsulation_mode` (encapMode) - (string) The encapsulation mode of the Relation To Domain object.
      - Default: `auto`
      - Valid Values: `auto`, `vlan`, `vxlan`.
  * `epg_cos` (epgCos) - (string) The class of service (CoS) of the Relation To Domain object.
      - Default: `Cos0`
      - Valid Values: `Cos0`, `Cos1`, `Cos2`, `Cos3`, `Cos4`, `Cos5`, `Cos6`, `Cos7`.
  * `epg_cos_pref` (epgCosPref) - (string) The class of service (CoS) preference of the Relation To Domain object.
      - Default: `disabled`
      - Valid Values: `disabled`, `enabled`.
  * `deployment_immediacy` (instrImedcy) - (string) The deployment immediacy of the Relation To Domain object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).
      - Default: `lazy`
      - Valid Values: `immediate`, `lazy`.
  * `ipam_dhcp_override` (ipamDhcpOverride) - (string) The IP address management (IPAM) DHCP override of the Relation To Domain object. Only applicable for Nutanix domains.
  * `ipam_enabled` (ipamEnabled) - (string) The IP address management (IPAM) enabled status of the Relation To Domain object. Only applicable for Nutanix domains.
      - Default: `no`
      - Valid Values: `no`, `yes`.
  * `ipam_gateway` (ipamGateway) - (string) The IP address management (IPAM) gateway of the Relation To Domain object. Only applicable for Nutanix domains.
  * `lag_policy_name` (lagPolicyName) - (string) The link aggregation group (LAG) policy name of the Relation To Domain object.
  * `netflow_direction` (netflowDir) - (string) The NetFlow monitoring direction of the Relation To Domain object.
      - Default: `both`
      - Valid Values: `both`, `egress`, `ingress`.
  * `enable_netflow` (netflowPref) - (string) The Netflow enabled status for the Relation To Domain object.
      - Default: `disabled`
      - Valid Values: `disabled`, `enabled`.
  * `number_of_ports` (numPorts) - (string) The number of ports of the Relation To Domain object.
  * `port_allocation` (portAllocation) - (string) Port allocation for ports.
      - Default: `none`
      - Valid Values: `elastic`, `fixed`, `none`.
  * `primary_encapsulation` (primaryEncap) - (string) The primary encapsulation of the Relation To Domain object. This is used when the class preference is set to 'useg'.
  * `primary_encapsulation_inner` (primaryEncapInner) - (string) The primary inner encapsulation of the Relation To Domain object. This is used for the portgroup at the VMWare Distributed Virtual Switch (DVS). This VLAN is internal to the DVS and is used for communication between the other VMs and the AVE VM at a host. Traffic is not forwarded to the fabric over the VLAN. Only applicable for Cisco ACI Virtual Edge (AVE) domains.
  * `resolution_immediacy` (resImedcy) - (string) The resolution immediacy of the Relation To Domain object. Specifies if policies are resolved immmediately or when needed.
      - Default: `lazy`
      - Valid Values: `immediate`, `lazy`, `pre-provision`.
  * `secondary_encapsulation_inner` (secondaryEncapInner) - (string) The secondary inner encapsulation of the Relation To Domain object. This is used for the portgroup at the VMWare Distributed Virtual Switch (DVS). This VLAN is internal to the DVS and is used for communication between the other VMs and the AVE VM at a host. Traffic is not forwarded to the fabric over the VLAN. Only applicable for Cisco ACI Virtual Edge (AVE) domains.
  * `switching_mode` (switchingMode) - (string) The switching mode of the Relation To Domain object.
      - Default: `native`
      - Valid Values: `AVE`, `native`.
  * `untagged` (untagged) - (string) The untagged status of the Relation To Domain object.
      - Default: `no`
      - Valid Values: `no`, `yes`.
  * `vnet_only` (vnetOnly) - (string) The VNET only status of the Relation To Domain object.
      - Default: `no`
      - Valid Values: `no`, `yes`.

* `relation_to_data_plane_policing_policy` - (list) A list of Relation To Data Plane Policing Policy (ACI object [fvRsDppPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsDppPol/overview)) pointing to Data Plane Policing Policy (ACI Object [qosDppPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/qosDppPol/overview)) which can be configured using the [aci_data_plane_policing_policy](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/data_plane_policing_policy) resource. This attribute is supported in ACI versions: 3.0(1k) and later.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Data Plane Policing Policy object.
      - Default: `orchestrator:terraform`
  * `data_plane_policing_policy_name` (tnQosDppPolName) - (string) Name.

* `relation_to_fibre_channel_paths` - (list) A list of Relation To Fibre Channel Paths (ACI object [fvRsFcPathAtt](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsFcPathAtt/overview)) pointing to  (ACI Object [fabricPathEp](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fabricPathEp/overview)). This attribute is supported in ACI versions: 2.0(1m) and later.
  
  #### Required ####
  
  * `target_dn` (tDn) - (string) The distinguished name of the target.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Fibre Channel Path object.
      - Default: `orchestrator:terraform`
  * `description` (descr) - (string) The description of the Relation To Fibre Channel Path object.
  * `vsan` (vsan) - (string) The virtual storage area network (VSAN) of the Relation To Fibre Channel Path object.
      - Default: `unknown`
  * `vsan_mode` (vsanMode) - (string) The virtual storage area network (VSAN) mode of the Relation To Fibre Channel Path object.
      - Default: `regular`
      - Valid Values: `native`, `regular`.

* `relation_to_intra_epg_contracts` - (list) A list of Relation To Intra EPG Contracts (ACI object [fvRsIntraEpg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsIntraEpg/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource. This attribute is supported in ACI versions: 3.0(1k) and later.
  
  #### Required ####
  
  * `contract_name` (tnVzBrCPName) - (string) The contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Intra EPG Contract object.
      - Default: `orchestrator:terraform`

* `relation_to_static_leafs` - (list) A list of Relation To Static Leafs (ACI object [fvRsNodeAtt](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsNodeAtt/overview)) pointing to  (ACI Object [fabricNode](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fabricNode/overview)).
  
  #### Required ####
  
  * `encapsulation` (encap) - (string) The VLAN encapsulation of the Relation To Static Leaf object.
  * `target_dn` (tDn) - (string) The distinguished name of the target of this static binding.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Static Leaf object.
      - Default: `orchestrator:terraform`
  * `description` (descr) - (string) The description of the Relation To Static Leaf object.
  * `deployment_immediacy` (instrImedcy) - (string) The deployment immediacy of the Relation To Static Leaf object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).
      - Default: `lazy`
      - Valid Values: `immediate`, `lazy`.
  * `mode` (mode) - (string) The static association mode with the path of the Relation To Static Leaf object.
      - Default: `regular`
      - Valid Values: `native`, `regular`, `untagged`.

* `relation_to_static_paths` - (list) A list of Relation To Static Paths (ACI object [fvRsPathAtt](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsPathAtt/overview)) pointing to  (ACI Object [fabricPathEp](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fabricPathEp/overview)).
  
  #### Required ####
  
  * `encapsulation` (encap) - (string) The VLAN encapsulation of the Relation To Static Path object.
  * `target_dn` (tDn) - (string) null.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Static Path object.
      - Default: `orchestrator:terraform`
  * `description` (descr) - (string) The description of the Relation To Static Path object.
  * `deployment_immediacy` (instrImedcy) - (string) The deployment immediacy of the Relation To Static Path object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).
      - Default: `lazy`
      - Valid Values: `immediate`, `lazy`.
  * `mode` (mode) - (string) The static association mode of the Relation To Static Path object.
      - Default: `regular`
      - Valid Values: `native`, `regular`, `untagged`.
  * `primary_encapsulation` (primaryEncap) - (string) The primary VLAN encapsulation of the Relation To Static Path object.

* `relation_to_taboo_contracts` - (list) A list of Relation To Taboo Contracts (ACI object [fvRsProtBy](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsProtBy/overview)) pointing to Taboo Contract (ACI Object [vzTaboo](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzTaboo/overview)) which can be configured using the [aci_taboo_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/taboo_contract) resource.
  
  #### Required ####
  
  * `taboo_contract_name` (tnVzTabooName) - (string) A contract for denying specific classes of traffic. Taboo rules are applied in the hardware before applying the rules of regular contracts. Without a contract, the default forwarding policy is to not allow any communication between EPGs. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/taboo_contract) with `aci_taboo_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/taboo_contract) with `data.aci_taboo_contract.example.name`.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Taboo Contract object.
      - Default: `orchestrator:terraform`

* `relation_to_provided_contracts` - (list) A list of Relation To Provided Contracts (ACI object [fvRsProv](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsProv/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource.
  
  #### Required ####
  
  * `contract_name` (tnVzBrCPName) - (string) The provider contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Provided Contract object.
      - Default: `orchestrator:terraform`
  * `match_criteria` (matchT) - (string) The provider label match criteria.
      - Default: `AtleastOne`
      - Valid Values: `All`, `AtleastOne`, `AtmostOne`, `None`.
  * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.

* `relation_to_contract_masters` - (list) A list of Relation To Contract Masters (ACI object [fvRsSecInherited](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsSecInherited/overview)) pointing to Application EPG (ACI Object [fvAEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAEPg/overview)) which can be configured using the [aci_application_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_epg) resource. This attribute is supported in ACI versions: 2.3(1e) and later.
  
  #### Required ####
  
  * `target_dn` (tDn) - (string) The distinguished name of the target.

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Contract Master object.
      - Default: `orchestrator:terraform`

* `relation_to_trust_control_policy` - (list) A list of Relation To Trust Control Policy (ACI object [fvRsTrustCtrl](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsTrustCtrl/overview)) pointing to Trust Control Policy (ACI Object [fhsTrustCtrlPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fhsTrustCtrlPol/overview)) which can be configured using the [aci_trust_control_policy](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/trust_control_policy) resource. This attribute is supported in ACI versions: 3.0(1k) and later.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation To Trust Control Policy object.
      - Default: `orchestrator:terraform`
  * `trust_control_policy_name` (tnFhsTrustCtrlPolName) - (string) Name.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  
  #### Required ####
  
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  
  #### Required ####
  
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

## Importing

An existing Application EPG can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_application_epg.example_application_profile uni/tn-{name}/ap-{name}/epg-{name}
```

Starting in Terraform version 1.5, an existing Application EPG can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/epg-{name}"
  to = aci_application_epg.example_application_profile
}
```