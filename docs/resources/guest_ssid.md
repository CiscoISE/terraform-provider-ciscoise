---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_guest_ssid Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on GuestSSID.
  This resource allows the client to update a guest SSID by ID.This resource deletes a guest SSID by ID.This resource creates a guest SSID.
---

# ciscoise_guest_ssid (Resource)

It manages create, read, update and delete operations on GuestSSID.

- This resource allows the client to update a guest SSID by ID.

- This resource deletes a guest SSID by ID.

- This resource creates a guest SSID.

## Example Usage

```terraform
resource "ciscoise_guest_ssid" "example" {
  provider = ciscoise
  parameters {

    id   = "string"
    name = "string"
  }
}

output "ciscoise_guest_ssid_example" {
  value = ciscoise_guest_ssid.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Resource Name. Name may contain alphanumeric or any of the following characters [_.-]


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **id** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- **name** (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_guest_ssid.example "id:=string"
terraform import ciscoise_guest_ssid.example "name:=string"
```
