---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_ldev_to_cloud_subnet"
sidebar_current: "docs-aci-data-source-cloud_ldev_to_cloud_subnet"
description: |-
  Data source for ACI Relation from Cloud LDev to Cloud Subnet
---

# aci_cloud_ldev_to_cloud_subnet #

Data source for ACI Relation from Cloud LDev to Cloud Subnet
Note: This resource is supported in Cloud APIC only.

## API Information ##

* `Class` - cloudRsLDevToCloudSubnet
* `Distinguished Name` - uni/tn-{name}/cld-{name}/rsLDevToCloudSubnet-[{tDn}]

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
data "aci_cloud_ldev_to_cloud_subnet" "example" {
  cloud_l4-l7device_dn  = aci_cloud_l4-l7device.example.id
  tDn  = "example"
}
```

## Argument Reference ##

* `cloud_l4-l7device_dn` - (Required) Distinguished name of the parent CloudL4-L7Device object.
* `tDn` - (Required) TDn of the Relation from Cloud LDev to Cloud Subnet object.

## Attribute Reference ##
* `id` - Attribute id set to the Dn of the Relation from Cloud LDev to Cloud Subnet.
* `annotation` - (Optional) Annotation of the Relation from Cloud LDev to Cloud Subnet object.
* `name_alias` - (Optional) Name Alias of the Relation from Cloud LDev to Cloud Subnet object.
* `t_dn` - (Optional) Target-dn. The distinguished name of the target.
