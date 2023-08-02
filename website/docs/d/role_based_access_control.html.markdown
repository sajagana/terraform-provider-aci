---
subcategory: -
layout: "aci"
page_title: "ACI: aci_role_based_access_control"
sidebar_current: "docs-aci-resource-role_based_access_control"
description: |-
  Manages ACI Annotation to capture rbac info
---

# aci_role_based_access_control #

Manages ACI Annotation to capture Role-based access control info

## API Information ##

* `Class` - aaaRbacAnnotation
* `Distinguished Name` - uni/monitor/telemetrypol-{name}/ftiflt-[{subnetIp}]/rbacDom-{domain}

## GUI Information ##

* `Location` - 


## Example Usage ##

```hcl
resource "aci_role_based_access_control" "example" {
  parent_dn  = aci_coopgroup_policy.example.id
  domain  = "example"
  child_regex = "test_child_regex"
}
```

## Argument Reference ##

* `parent_dn` - (Required) Distinguished name of the parent COOPGroupPolicy object.
* `domain` - (Required) Domain of the Annotation to capture rbac info object.
* `annotation` - (Optional) Annotation of the Annotation to capture rbac info object.
* `name_alias` - (Optional) Name Alias of the Annotation to capture rbac info object.* `child_regex` - (Optional) childregx.
* `domain` - (Optional) Domain.The domain of the counts object.


## Importing ##

An existing Annotation To Capture RBAC Info can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import aci_role_based_access_control.example <Dn>
```