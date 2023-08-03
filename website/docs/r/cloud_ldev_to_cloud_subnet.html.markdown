---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_ldev_to_cloud_subnet"
sidebar_current: "docs-aci-resource-cloud_ldev_to_cloud_subnet"
description: |-
  Manages ACI Relation from Cloud LDev to Cloud Subnet
---

# aci_cloud_ldev_to_cloud_subnet #

Manages ACI Relation from Cloud LDev to Cloud Subnet

## API Information ##

* `Class` - cloudRsLDevToCloudSubnet
* `Distinguished Name` - uni/tn-{name}/cld-{name}/rsLDevToCloudSubnet-[{tDn}]

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
resource "aci_cloud_ldev_to_cloud_subnet" "example" {
  cloud_l4-l7device_dn  = aci_cloud_l4-l7device.example.id
  tDn  = "example"
  annotation = "orchestrator:terraform"
  t_dn = 
}
```

## Argument Reference ##

* `cloud_l4-l7device_dn` - (Required) Distinguished name of the parent CloudL4-L7Device object.
* `tDn` - (Required) TDn of the Relation from Cloud LDev to Cloud Subnet object.
* `annotation` - (Optional) Annotation of the Relation from Cloud LDev to Cloud Subnet object.
* `name_alias` - (Optional) Name Alias of the Relation from Cloud LDev to Cloud Subnet object.
* `t_dn` - (Optional) Target-dn.The distinguished name of the target.


## Importing ##

An existing Relation from Cloud LDev to Cloud Subnet can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import aci_cloud_ldev_to_cloud_subnet.example <Dn>
```