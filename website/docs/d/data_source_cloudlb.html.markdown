---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_l4-l7load_balancer"
sidebar_current: "docs-aci-data-source-cloud_l4-l7load_balancer"
description: |-
  Data source for ACI Cloud L4-L7 Load Balancer
---

# aci_cloud_l4-l7load_balancer #

Data source for ACI Cloud L4-L7 Load Balancer
Note: This resource is supported in Cloud APIC only.

## API Information ##

* `Class` - cloudLB
* `Distinguished Name` - uni/tn-{name}/clb-{name}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
data "aci_cloud_l4-l7load_balancer" "example" {
  tenant_dn  = aci_tenant.example.id
  name  = "example"
}
```

## Argument Reference ##

* `tenant_dn` - (Required) Distinguished name of the parent Tenant object.
* `name` - (Required) Name of the Cloud L4-L7 Load Balancer object.

## Attribute Reference ##
* `id` - Attribute id set to the Dn of the Cloud L4-L7 Load Balancer.
* `annotation` - (Optional) Annotation of the Cloud L4-L7 Load Balancer object.
* `name_alias` - (Optional) Name Alias of the Cloud L4-L7 Load Balancer object.
* `version` - (Optional) Version. 
* `active_active` - (Optional) Active-Active mode. 
* `allow_all` - (Optional) allowAll. 
* `auto_scaling` - (Optional) autoScaling. 
* `context_aware` - (Optional) Tenancy. A value to determine if the L4-L7 device cluster supports multiple contexts (VRFs).
* `custom_rg` - (Optional) customRG. 
* `devtype` - (Optional) devtype. virtual or physical service devices in this cluster
* `firewall_mode` - (Optional) firewallMode. 
* `firewall_status` - (Optional) firewallStatus. 
* `func_type` - (Optional) Function Type. The function type of the L4-L7 device cluster.
* `instance_count` - (Optional) instanceCount. 
* `is_copy` - (Optional) isCopy. if the device is a copy device
* `is_instantiation` - (Optional) isInstantiation. 
* `is_static_ip` - (Optional) isStaticIP. 
* `l4l7device_application_security_group` - (Optional) Naming for Third Party Device ASG. 
* `l4l7third_party_device` - (Optional) Naming for Third Party Device. 
* `managed` - (Optional) managed. Specified if the device is a managed device
* `max_instance_count` - (Optional) maxInstanceCount. 
* `min_instance_count` - (Optional) minInstanceCount. 
* `mode` - (Optional) Mode. The value for specifying if the device is legacy (classical VLAN/VXLAN) or supports service tag switching (STS).
* `native_lbname` - (Optional) Naming for Native Load Balancer Devices. 
* `oid` - (Optional) oid. Tracks the identifier of and object in the external VM management system. For example, a managed object reference of the virtual machine in the VMware vCenter. For internal use only.
* `package_model` - (Optional) Package Model. 
* `prom_mode` - (Optional) Promiscuous Mode. Promiscuous Mode
* `scheme` - (Optional) scheme. 
* `size` - (Optional) size. The size of the firmware image.
* `sku` - (Optional) sku. 
* `svc_type` - (Optional) svcType. UI Template type
* `target_mode` - (Optional) targetMode. 
* `trunking` - (Optional) trunking. For virtual devices, if a trunking port group is to be used
* `cloud_l4-l7load_balancer_type` - (Optional) type. The specific type of the object or component.

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