---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_sponsor_portal Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on SponsorPortal.
  This data source allows the client to get a sponsor portal by ID.This data source allows the client to get all the sponsor portals.
  Filter:
  [name, description]
  Sorting:
  [name, description]
---

# ciscoise_sponsor_portal (Data Source)

It performs read operation on SponsorPortal.

- This data source allows the client to get a sponsor portal by ID.

- This data source allows the client to get all the sponsor portals.

Filter:

[name, description]

Sorting:

[name, description]

## Example Usage

```terraform
data "ciscoise_sponsor_portal" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sponsor_portal_example" {
  value = data.ciscoise_sponsor_portal.example.items
}

data "ciscoise_sponsor_portal" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_sponsor_portal_example" {
  value = data.ciscoise_sponsor_portal.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (List of String) filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains
- `filter_type` (String) filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
- `id` (String) id path parameter.
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page
- `sortasc` (String) sortasc query parameter. sort asc
- `sortdsc` (String) sortdsc query parameter. sort desc

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `customizations` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations))
- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `portal_test_url` (String)
- `portal_type` (String)
- `settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings))

<a id="nestedobjatt--item--customizations"></a>
### Nested Schema for `item.customizations`

Read-Only:

- `global_customizations` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--global_customizations))
- `language` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--language))
- `page_customizations` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--page_customizations))
- `portal_theme` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--portal_theme))
- `portal_tweak_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--portal_tweak_settings))

<a id="nestedobjatt--item--customizations--global_customizations"></a>
### Nested Schema for `item.customizations.global_customizations`

Read-Only:

- `background_image` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--global_customizations--background_image))
- `banner_image` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--global_customizations--banner_image))
- `banner_title` (String)
- `contact_text` (String)
- `desktop_logo_image` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--global_customizations--desktop_logo_image))
- `footer_element` (String)
- `mobile_logo_image` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--global_customizations--mobile_logo_image))

<a id="nestedobjatt--item--customizations--global_customizations--background_image"></a>
### Nested Schema for `item.customizations.global_customizations.mobile_logo_image`

Read-Only:

- `data` (String)


<a id="nestedobjatt--item--customizations--global_customizations--banner_image"></a>
### Nested Schema for `item.customizations.global_customizations.mobile_logo_image`

Read-Only:

- `data` (String)


<a id="nestedobjatt--item--customizations--global_customizations--desktop_logo_image"></a>
### Nested Schema for `item.customizations.global_customizations.mobile_logo_image`

Read-Only:

- `data` (String)


<a id="nestedobjatt--item--customizations--global_customizations--mobile_logo_image"></a>
### Nested Schema for `item.customizations.global_customizations.mobile_logo_image`

Read-Only:

- `data` (String)



<a id="nestedobjatt--item--customizations--language"></a>
### Nested Schema for `item.customizations.language`

Read-Only:

- `view_language` (String)


<a id="nestedobjatt--item--customizations--page_customizations"></a>
### Nested Schema for `item.customizations.page_customizations`

Read-Only:

- `data` (List of Object) (see [below for nested schema](#nestedobjatt--item--customizations--page_customizations--data))

<a id="nestedobjatt--item--customizations--page_customizations--data"></a>
### Nested Schema for `item.customizations.page_customizations.data`

Read-Only:

- `key` (String)
- `value` (String)



<a id="nestedobjatt--item--customizations--portal_theme"></a>
### Nested Schema for `item.customizations.portal_theme`

Read-Only:

- `id` (String)
- `name` (String)
- `theme_data` (String)


<a id="nestedobjatt--item--customizations--portal_tweak_settings"></a>
### Nested Schema for `item.customizations.portal_tweak_settings`

Read-Only:

- `banner_color` (String)
- `banner_text_color` (String)
- `page_background_color` (String)
- `page_label_and_text_color` (String)



<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


<a id="nestedobjatt--item--settings"></a>
### Nested Schema for `item.settings`

Read-Only:

- `aup_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--aup_settings))
- `login_page_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--login_page_settings))
- `portal_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--portal_settings))
- `post_access_banner_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--post_access_banner_settings))
- `post_login_banner_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--post_login_banner_settings))
- `sponsor_change_password_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--sponsor_change_password_settings))
- `support_info_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--support_info_settings))

<a id="nestedobjatt--item--settings--aup_settings"></a>
### Nested Schema for `item.settings.aup_settings`

Read-Only:

- `display_frequency` (String)
- `display_frequency_interval_days` (Number)
- `include_aup` (String)
- `require_scrolling` (String)


<a id="nestedobjatt--item--settings--login_page_settings"></a>
### Nested Schema for `item.settings.login_page_settings`

Read-Only:

- `aup_display` (String)
- `include_aup` (String)
- `max_failed_attempts_before_rate_limit` (Number)
- `require_aup_acceptance` (String)
- `require_aup_scrolling` (String)
- `social_configs` (List of String)
- `time_between_logins_during_rate_limit` (Number)


<a id="nestedobjatt--item--settings--portal_settings"></a>
### Nested Schema for `item.settings.portal_settings`

Read-Only:

- `allowed_interfaces` (List of String)
- `authentication_method` (String)
- `available_ssids` (List of String)
- `certificate_group_tag` (String)
- `display_lang` (String)
- `fallback_language` (String)
- `fqdn` (String)
- `https_port` (Number)
- `idle_timeout` (Number)


<a id="nestedobjatt--item--settings--post_access_banner_settings"></a>
### Nested Schema for `item.settings.post_access_banner_settings`

Read-Only:

- `include_post_access_banner` (String)


<a id="nestedobjatt--item--settings--post_login_banner_settings"></a>
### Nested Schema for `item.settings.post_login_banner_settings`

Read-Only:

- `include_post_access_banner` (String)


<a id="nestedobjatt--item--settings--sponsor_change_password_settings"></a>
### Nested Schema for `item.settings.sponsor_change_password_settings`

Read-Only:

- `allow_sponsor_to_change_pwd` (String)


<a id="nestedobjatt--item--settings--support_info_settings"></a>
### Nested Schema for `item.settings.support_info_settings`

Read-Only:

- `default_empty_field_value` (String)
- `empty_field_display` (String)
- `include_browser_user_agent` (String)
- `include_failure_code` (String)
- `include_ip_address` (String)
- `include_mac_addr` (String)
- `include_policy_server` (String)
- `include_support_info_page` (String)




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


