---
subcategory: -
layout: "aci"
page_title: "ACI: aci_role_based_access_control"
sidebar_current: "docs-aci-data-source-role_based_access_control"
description: |-
  Data source for ACI Annotation to capture rbac info
---

# aci_role_based_access_control #

Data source for ACI Annotation to capture rbac info


## API Information ##

* `Class` - aaaRbacAnnotation
* `Distinguished Name` - uni/monitor/telemetrypol-{name}/ftiflt-[{subnetIp}]/rbacDom-{domain}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
data "aci_role_based_access_control" "example" {
  parent_dn  = aci_coopgroup_policy.example.id
  domain  = "example"
}
```

## Argument Reference ##

* `parent_dn` - (Required) Distinguished name of the parent COOPGroupPolicy object.
* `domain` - (Required) Domain of the Annotation to capture rbac info object.

## Attribute Reference ##
* `id` - Attribute id set to the Dn of the Annotation to capture rbac info.
* `child_regex` - (Optional) childregx. 
* `domain` - (Optional) Domain. The domain of the counts object.
