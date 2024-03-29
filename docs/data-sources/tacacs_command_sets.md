---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_tacacs_command_sets Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on TACACSCommandSets.
  This data source allows the client to get TACACS command sets by name.This data source allows the client to get TACACS command sets by ID.This data source allows the client to get all the TACACS command sets.
---

# ciscoise_tacacs_command_sets (Data Source)

It performs read operation on TACACSCommandSets.

- This data source allows the client to get TACACS command sets by name.

- This data source allows the client to get TACACS command sets by ID.

- This data source allows the client to get all the TACACS command sets.

## Example Usage

```terraform
data "ciscoise_tacacs_command_sets" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_tacacs_command_sets_example" {
  value = data.ciscoise_tacacs_command_sets.example.item_name
}

data "ciscoise_tacacs_command_sets" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_tacacs_command_sets_example" {
  value = data.ciscoise_tacacs_command_sets.example.item_id
}

data "ciscoise_tacacs_command_sets" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_tacacs_command_sets_example" {
  value = data.ciscoise_tacacs_command_sets.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter.
- `name` (String) name path parameter.
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page

### Read-Only

- `item_id` (List of Object) (see [below for nested schema](#nestedatt--item_id))
- `item_name` (List of Object) (see [below for nested schema](#nestedatt--item_name))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item_id"></a>
### Nested Schema for `item_id`

Read-Only:

- `commands` (List of Object) (see [below for nested schema](#nestedobjatt--item_id--commands))
- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item_id--link))
- `name` (String)
- `permit_unmatched` (String)

<a id="nestedobjatt--item_id--commands"></a>
### Nested Schema for `item_id.commands`

Read-Only:

- `command_list` (List of Object) (see [below for nested schema](#nestedobjatt--item_id--commands--command_list))

<a id="nestedobjatt--item_id--commands--command_list"></a>
### Nested Schema for `item_id.commands.command_list`

Read-Only:

- `arguments` (String)
- `command` (String)
- `grant` (String)



<a id="nestedobjatt--item_id--link"></a>
### Nested Schema for `item_id.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item_name"></a>
### Nested Schema for `item_name`

Read-Only:

- `commands` (List of Object) (see [below for nested schema](#nestedobjatt--item_name--commands))
- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item_name--link))
- `name` (String)
- `permit_unmatched` (String)

<a id="nestedobjatt--item_name--commands"></a>
### Nested Schema for `item_name.commands`

Read-Only:

- `command_list` (List of Object) (see [below for nested schema](#nestedobjatt--item_name--commands--command_list))

<a id="nestedobjatt--item_name--commands--command_list"></a>
### Nested Schema for `item_name.commands.command_list`

Read-Only:

- `arguments` (String)
- `command` (String)
- `grant` (String)



<a id="nestedobjatt--item_name--link"></a>
### Nested Schema for `item_name.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--items--link))
- `name` (String)

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


