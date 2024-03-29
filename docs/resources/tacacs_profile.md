---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_tacacs_profile Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on TACACSProfile.
  This resource allows the client to update a TACACS profile.This resource deletes a TACACS profile.This resource creates a TACACS profile.
---

# ciscoise_tacacs_profile (Resource)

It manages create, read, update and delete operations on TACACSProfile.

- This resource allows the client to update a TACACS profile.

- This resource deletes a TACACS profile.

- This resource creates a TACACS profile.

## Example Usage

```terraform
resource "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    session_attributes {

      session_attribute_list {

        name  = "string"
        type  = "string"
        value = "string"
      }
    }
  }
}

output "ciscoise_tacacs_profile_example" {
  value = ciscoise_tacacs_profile.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `description` (String)
- `name` (String)
- `session_attributes` (Block List) Holds list of session attributes. View type for GUI is Shell by default (see [below for nested schema](#nestedblock--parameters--session_attributes))

Read-Only:

- `id` (String) The ID of this resource.
- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedblock--parameters--session_attributes"></a>
### Nested Schema for `parameters.session_attributes`

Optional:

- `session_attribute_list` (Block List) (see [below for nested schema](#nestedblock--parameters--session_attributes--session_attribute_list))

<a id="nestedblock--parameters--session_attributes--session_attribute_list"></a>
### Nested Schema for `parameters.session_attributes.session_attribute_list`

Optional:

- `name` (String)
- `type` (String) Allowed values: MANDATORY, OPTIONAL
- `value` (String)



<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `session_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--item--session_attributes))

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


<a id="nestedobjatt--item--session_attributes"></a>
### Nested Schema for `item.session_attributes`

Read-Only:

- `session_attribute_list` (List of Object) (see [below for nested schema](#nestedobjatt--item--session_attributes--session_attribute_list))

<a id="nestedobjatt--item--session_attributes--session_attribute_list"></a>
### Nested Schema for `item.session_attributes.session_attribute_list`

Read-Only:

- `name` (String)
- `type` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_tacacs_profile.example "id:=string\name:=string"
```
