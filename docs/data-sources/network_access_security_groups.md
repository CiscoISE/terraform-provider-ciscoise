---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_network_access_security_groups Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Network Access - Security Groups.
  Network Access Return list of available security groups for authorization policy definition.
  (Other CRUD APIs available throught ERS)
---

# ciscoise_network_access_security_groups (Data Source)

It performs read operation on Network Access - Security Groups.

- Network Access Return list of available security groups for authorization policy definition.
 (Other CRUD APIs available throught ERS)

## Example Usage

```terraform
data "ciscoise_network_access_security_groups" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_security_groups_example" {
  value = data.ciscoise_network_access_security_groups.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `name` (String)


