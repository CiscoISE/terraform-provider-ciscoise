---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_licensing_connection_type Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Licensing.
  Get connection type
---

# ciscoise_licensing_connection_type (Data Source)

It performs read operation on Licensing.

- Get connection type

## Example Usage

```terraform
data "ciscoise_licensing_connection_type" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_connection_type_example" {
  value = data.ciscoise_licensing_connection_type.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **connection_type** (String)
- **state** (String)

