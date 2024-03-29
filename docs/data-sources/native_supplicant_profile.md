---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_native_supplicant_profile Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on NativeSupplicantProfile.
  This data source allows the client to get a native supplicant profile by ID.This data source allows the client to get all the native supplicant profiles.
---

# ciscoise_native_supplicant_profile (Data Source)

It performs read operation on NativeSupplicantProfile.

- This data source allows the client to get a native supplicant profile by ID.

- This data source allows the client to get all the native supplicant profiles.

## Example Usage

```terraform
data "ciscoise_native_supplicant_profile" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_native_supplicant_profile_example" {
  value = data.ciscoise_native_supplicant_profile.example.items
}

data "ciscoise_native_supplicant_profile" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_native_supplicant_profile_example" {
  value = data.ciscoise_native_supplicant_profile.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter.
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `wireless_profiles` (List of Object) (see [below for nested schema](#nestedobjatt--item--wireless_profiles))

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


<a id="nestedobjatt--item--wireless_profiles"></a>
### Nested Schema for `item.wireless_profiles`

Read-Only:

- `action_type` (String)
- `allowed_protocol` (String)
- `certificate_template_id` (String)
- `previous_ssid` (String)
- `ssid` (String)



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


