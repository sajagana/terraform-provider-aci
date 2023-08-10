---
subcategory: -
layout: "aci"
page_title: "ACI: aci_cloud_l4-l7logical_interface"
sidebar_current: "docs-aci-data-source-cloud_l4-l7logical_interface"
description: |-
  Data source for ACI Cloud L4-L7 Logical Interface
---

# aci_cloud_l4-l7logical_interface #

Data source for ACI Cloud L4-L7 Logical Interface
Note: This resource is supported in Cloud APIC only.

## API Information ##

* `Class` - cloudLIf
* `Distinguished Name` - uni/tn-{name}/cld-{name}/clif-{name}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
data "aci_cloud_l4-l7logical_interface" "example" {
  cloud_l4-l7device_dn  = aci_cloud_l4-l7device.example.id
  name  = "example"
}
```

## Argument Reference ##

* `cloud_l4-l7device_dn` - (Required) Distinguished name of the parent CloudL4-L7Device object.
* `name` - (Required) Name of the Cloud L4-L7 Logical Interface object.

## Attribute Reference ##
* `id` - Attribute id set to the Dn of the Cloud L4-L7 Logical Interface.
* `annotation` - (Optional) Annotation of the Cloud L4-L7 Logical Interface object.
* `name_alias` - (Optional) Name Alias of the Cloud L4-L7 Logical Interface object.
* `allow_all` - (Optional) allowAll. 

* `relation_cloudrs_lif_to_cloud_subnet` - (Optional) Represents the relation to a Relation from Cloud LDev to Cloud Subnet (class cloudSubnet).  Type: String.