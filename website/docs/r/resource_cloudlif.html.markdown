---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_l4-l7logical_interface"
sidebar_current: "docs-aci-resource-cloud_l4-l7logical_interface"
description: |-
  Manages ACI Cloud L4-L7 Logical Interface
---

# aci_cloud_l4-l7logical_interface #

Manages ACI Cloud L4-L7 Logical Interface

## API Information ##

* `Class` - cloudLIf
* `Distinguished Name` - uni/tn-{name}/cld-{name}/clif-{name}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
resource "aci_cloud_l4-l7logical_interface" "example" {
  cloud_l4-l7device_dn  = aci_cloud_l4-l7device.example.id
  name  = "example"
  allow_all = "false"
  annotation = "orchestrator:terraform"

  name_alias = 

  cloud_rs_lif_to_cloud_subnet = aci_resource.example.id
}
```

## Argument Reference ##

* `cloud_l4-l7device_dn` - (Required) Distinguished name of the parent CloudL4-L7Device object.
* `name` - (Required) Name of the Cloud L4-L7 Logical Interface object.
* `annotation` - (Optional) Annotation of the Cloud L4-L7 Logical Interface object.
* `name_alias` - (Optional) Name Alias of the Cloud L4-L7 Logical Interface object.* `allow_all` - (Optional) allowAll. Allowed values are "no", "yes", and default value is "false". Type: String.


* `relation_cloudrs_lif_to_cloud_subnet` - (Optional) Represents the relation to a Relation from Cloud LDev to Cloud Subnet (class cloudSubnet).  Type: String.


## Importing ##

An existing CloudL4-L7LogicalInterface can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import aci_cloud_l4-l7logical_interface.example <Dn>
```