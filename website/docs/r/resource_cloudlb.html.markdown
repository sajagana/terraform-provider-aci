---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_l4-l7load_balancer"
sidebar_current: "docs-aci-resource-cloud_l4-l7load_balancer"
description: |-
  Manages ACI Cloud L4-L7 Load Balancer
---

# aci_cloud_l4-l7load_balancer #

Manages ACI Cloud L4-L7 Load Balancer

## API Information ##

* `Class` - cloudLB
* `Distinguished Name` - uni/tn-{name}/clb-{name}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
resource "aci_cloud_l4-l7load_balancer" "example" {
  tenant_dn  = aci_tenant.example.id
  name  = "example"
  version = "1.0"
  active_active = "false"
  allow_all = "false"
  annotation = "orchestrator:terraform"
  auto_scaling = "false"
  context_aware = "single-Context"
  custom_rg = 
  devtype = "CLOUD"
  firewall_mode = 
  firewall_status = "false"
  func_type = "GoTo"
  instance_count = "2"
  is_copy = "false"
  is_instantiation = "false"
  is_static_ip = "false"
  l4l7device_application_security_group = 
  l4l7third_party_device = 
  managed = "true"
  max_instance_count = "10"
  min_instance_count = "0"
  mode = "legacy-Mode"

  name_alias = 
  native_lbname = 
  oid = "0"
  package_model = 
  prom_mode = "false"
  scheme = "internet"
  size = "medium"
  sku = "standard"
  svc_type = "NATIVELB"
  target_mode = "unspecified"
  trunking = "false"
  cloud_l4-l7load_balancer_type = "application"

  cloud_rs_ldev_to_cloud_subnet = [aci_resource.example.id]

  cloud_rs_ldev_to_compute_pol = aci_resource.example.id

  cloud_rs_ldev_to_mgmt_pol = aci_resource.example.id

  vns_rs_aldev_to_dev_mgr = aci_resource.example.id

  vns_rs_aldev_to_dom_p = {
      target_dn = aci_resource.example.id
  }

  vns_rs_aldev_to_phys_dom_p = aci_resource.example.id

  vns_rs_aldev_to_vlan_inst_p = [aci_resource.example.id]

  vns_rs_aldev_to_vxlan_inst_p = aci_resource.example.id

  vns_rs_dev_epg = aci_resource.example.id

  vns_rs_mdev_att = aci_resource.example.id

  vns_rs_svc_mgmt_epg = aci_resource.example.id
}
```

## Argument Reference ##

* `tenant_dn` - (Required) Distinguished name of the parent Tenant object.
* `name` - (Required) Name of the Cloud L4-L7 Load Balancer object.
* `annotation` - (Optional) Annotation of the Cloud L4-L7 Load Balancer object.
* `name_alias` - (Optional) Name Alias of the Cloud L4-L7 Load Balancer object.* `version` - (Optional) Version. Allowed values are and default value is "1.0".
* `active_active` - (Optional) Active-Active mode. Allowed values are "no", "yes", and default value is "false". Type: String.
* `allow_all` - (Optional) allowAll. Allowed values are "no", "yes", and default value is "false". Type: String.

* `auto_scaling` - (Optional) autoScaling. Allowed values are "no", "yes", and default value is "false". Type: String.
* `context_aware` - (Optional) Tenancy.A value to determine if the L4-L7 device cluster supports multiple contexts (VRFs). Allowed values are "multi-Context", "single-Context", and default value is "single-Context". Type: String.
* `custom_rg` - (Optional) customRG.
* `devtype` - (Optional) devtype.virtual or physical service devices in this cluster Allowed values are "CLOUD", "PHYSICAL", "VIRTUAL", and default value is "CLOUD". Type: String.
* `firewall_mode` - (Optional) firewallMode. Allowed values are "detection", "prevention",  Type: String.
* `firewall_status` - (Optional) firewallStatus. Allowed values are "no", "yes", and default value is "false". Type: String.
* `func_type` - (Optional) Function Type.The function type of the L4-L7 device cluster. Allowed values are "GoThrough", "GoTo", "L1", "L2", "None", and default value is "GoTo". Type: String.
* `instance_count` - (Optional) instanceCount.
* `is_copy` - (Optional) isCopy.if the device is a copy device Allowed values are "no", "yes", and default value is "false". Type: String.
* `is_instantiation` - (Optional) isInstantiation. Allowed values are "no", "yes", and default value is "false". Type: String.
* `is_static_ip` - (Optional) isStaticIP. Allowed values are "no", "yes", and default value is "false". Type: String.
* `l4l7device_application_security_group` - (Optional) Naming for Third Party Device ASG.
* `l4l7third_party_device` - (Optional) Naming for Third Party Device.
* `managed` - (Optional) managed.Specified if the device is a managed device Allowed values are "no", "yes", and default value is "true". Type: String.
* `max_instance_count` - (Optional) maxInstanceCount.
* `min_instance_count` - (Optional) minInstanceCount.
* `mode` - (Optional) Mode.The value for specifying if the device is legacy (classical VLAN/VXLAN) or supports service tag switching (STS). Allowed values are "legacy-Mode", and default value is "legacy-Mode". Type: String.

* `native_lbname` - (Optional) Naming for Native Load Balancer Devices.
* `oid` - (Optional) oid.Tracks the identifier of and object in the external VM management system. For example, a managed object reference of the virtual machine in the VMware vCenter. For internal use only.
* `package_model` - (Optional) Package Model.
* `prom_mode` - (Optional) Promiscuous Mode.Promiscuous Mode Allowed values are "no", "yes", and default value is "false". Type: String.
* `scheme` - (Optional) scheme. Allowed values are "internal", "internet", and default value is "internet". Type: String.
* `size` - (Optional) size.The size of the firmware image. Allowed values are "large", "medium", "small", "v2", and default value is "medium". Type: String.
* `sku` - (Optional) sku. Allowed values are "WAF", "WAF_v2", "standard", "standard_v2", and default value is "standard". Type: String.
* `svc_type` - (Optional) svcType.UI Template type Allowed values are "ADC", "COPY", "FW", "NATIVELB", "OTHERS", and default value is "NATIVELB". Type: String.
* `target_mode` - (Optional) targetMode. Allowed values are "primary", "secondary", "unspecified", and default value is "unspecified". Type: String.
* `trunking` - (Optional) trunking.For virtual devices, if a trunking port group is to be used Allowed values are "no", "yes", and default value is "false". Type: String.
* `cloud_l4-l7load_balancer_type` - (Optional) type.The specific type of the object or component. Allowed values are "application", "network", and default value is "application". Type: String.
* `relation_cloudrs_ldev_to_cloud_subnet` - (Optional) Represents the relation to a Relation from Cloud LDev to Cloud Subnet (class cloudSubnet).  Type: List.
* `relation_cloudrs_ldev_to_compute_pol` - (Optional) Represents the relation to a Relation from Cloud LDev to Compute Instance Policy (class cloudComputePol).  Type: String.
* `relation_cloudrs_ldev_to_mgmt_pol` - (Optional) Represents the relation to a Relation from Cloud LDev to Mgmt Policy (class cloudMgmtPol).  Type: String.
* `relation_vnsrs_aldev_to_dev_mgr` - (Optional) Represents the relation to a Device Manager (class vnsDevMgr).  Type: String.
* `relation_vnsrs_aldev_to_dom_p` - (Optional) A map representing the relation to a Relation from L4-L7 Devices to a Vmm Domain (class vmmDomP). A source relation to the VMM domain profile. Type: Map.
* `relation_vnsrs_aldev_to_phys_dom_p` - (Optional) Represents the relation to a Relation from L4-L7 Devices to a Vlan Namespace Instance Profile (class physDomP). A source relation to a physical domain profile. Type: String.
* `relation_vnsrs_aldev_to_vlan_inst_p` - (Optional) Represents the relation to a VLAN Pool (class fvnsVlanInstP). A source relation to the policy definition for ID ranges used for VLAN encapsulation. Type: List.
* `relation_vnsrs_aldev_to_vxlan_inst_p` - (Optional) Represents the relation to a Relation from L4-L7 Devices to a Vxlan Namespace Instance Profile (class fvnsVxlanInstP). A source relation to the VxLAN namespace policy. Note that this relation is an internal object. Type: String.
* `relation_vnsrs_dev_epg` - (Optional) Represents the relation to a Relation from a L4-L7 Devices to an EPG (class fvEPg). A source relation for storing management endpoint group related information. This object indicates that the device belongs to the management endpoint group. Type: String.
* `relation_vnsrs_mdev_att` - (Optional) Represents the relation to a Relation from a Service L4-L7 Devices to Meta Device Information (class vnsMDev). A source relation to the metadata definitions for a service device type. Note that this relation is an internal object. Type: String.
* `relation_vnsrs_svc_mgmt_epg` - (Optional) Represents the relation to a Management End Point Group (class fvEPg). A source relation to a set of endpoints. Note that this relation is an internal object. Type: String.


## Importing ##

An existing CloudL4-L7LoadBalancer can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import aci_cloud_l4-l7load_balancer.example <Dn>
```