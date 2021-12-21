---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_patch Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Patching.
  List all the installed patches in the system, with the patch number for rollback.
---

# ciscoise_patch (Data Source)

It performs read operation on Patching.

- List all the installed patches in the system, with the patch number for rollback.

## Example Usage

```terraform
data "ciscoise_patch" "example" {
  provider = ciscoise
}

output "ciscoise_patch_example" {
  value = data.ciscoise_patch.example.item
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

- **ise_version** (String)
- **patch_version** (List of Object) (see [below for nested schema](#nestedobjatt--item--patch_version))

<a id="nestedobjatt--item--patch_version"></a>
### Nested Schema for `item.patch_version`

Read-Only:

- **install_date** (String)
- **patch_number** (Number)

