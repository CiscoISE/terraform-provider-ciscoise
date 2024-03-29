---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_my_device_portal Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on MyDevicePortal.
  This resource allows the client to update a my device portal by ID.This resource deletes a my device portal by ID.This resource creates a my device portal.
---

# ciscoise_my_device_portal (Resource)

It manages create, read, update and delete operations on MyDevicePortal.

- This resource allows the client to update a my device portal by ID.

- This resource deletes a my device portal by ID.

- This resource creates a my device portal.

## Example Usage

```terraform
resource "ciscoise_my_device_portal" "example" {
  provider = ciscoise
  parameters {

    customizations {

      global_customizations {

        background_image {

          data = "string"
        }
        banner_image {

          data = "string"
        }
        banner_title = "string"
        contact_text = "string"
        desktop_logo_image {

          data = "string"
        }
        footer_element = "string"
        mobile_logo_image {

          data = "string"
        }
      }
      language {

        view_language = "string"
      }
      page_customizations {

        data {

          key   = "string"
          value = "string"
        }
      }
      portal_theme {

        id         = "string"
        name       = "string"
        theme_data = "string"
      }
      portal_tweak_settings {

        banner_color              = "string"
        banner_text_color         = "string"
        page_background_color     = "string"
        page_label_and_text_color = "string"
      }
    }
    description     = "string"
    id              = "string"
    name            = "string"
    portal_test_url = "string"
    portal_type     = "string"
    settings {

      aup_settings {

        display_frequency               = "string"
        display_frequency_interval_days = 1
        include_aup                     = "false"
        require_scrolling               = "false"
      }
      employee_change_password_settings {

        allow_employee_to_change_pwd = "false"
      }
      login_page_settings {

        aup_display                           = "string"
        include_aup                           = "false"
        max_failed_attempts_before_rate_limit = 1
        require_aup_acceptance                = "false"
        require_scrolling                     = "false"
        time_between_logins_during_rate_limit = 1
      }
      portal_settings {

        allowed_interfaces      = ["string"]
        always_used_language    = "string"
        certificate_group_tag   = "string"
        display_lang            = "string"
        endpoint_identity_group = "string"
        fallback_language       = "string"
        https_port              = 1
      }
      post_access_banner_settings {

        include_post_access_banner = "false"
      }
      post_login_banner_settings {

        include_post_access_banner = "false"
      }
      support_info_settings {

        default_empty_field_value  = "string"
        empty_field_display        = "string"
        include_browser_user_agent = "false"
        include_failure_code       = "false"
        include_ip_address         = "false"
        include_mac_addr           = "false"
        include_policy_server      = "false"
        include_support_info_page  = "false"
      }
    }
  }
}

output "ciscoise_my_device_portal_example" {
  value = ciscoise_my_device_portal.example
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

- `customizations` (Block List) Defines all of the Portal Customizations available (see [below for nested schema](#nestedblock--parameters--customizations))
- `description` (String)
- `name` (String)
- `portal_test_url` (String) URL to bring up a test page for this portal
- `portal_type` (String) Allowed values:
		- BYOD,
		- HOTSPOTGUEST,
		- MYDEVICE,
		- SELFREGGUEST,
		- SPONSOR,
		- SPONSOREDGUEST
- `settings` (Block List) Defines all of the settings groups available for a Mydevice portal (see [below for nested schema](#nestedblock--parameters--settings))

Read-Only:

- `id` (String) The ID of this resource.
- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedblock--parameters--customizations"></a>
### Nested Schema for `parameters.customizations`

Optional:

- `global_customizations` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--global_customizations))
- `language` (Block List) This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported (see [below for nested schema](#nestedblock--parameters--customizations--language))
- `page_customizations` (Block List) Represent the entire page customization as a giant dictionary (see [below for nested schema](#nestedblock--parameters--customizations--page_customizations))
- `portal_theme` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--portal_theme))
- `portal_tweak_settings` (Block List) The Tweak Settings are a customization of the Portal Theme that has been selected for the portal.
		When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme.
		The Tweak Settings can subsequently be changed by the user (see [below for nested schema](#nestedblock--parameters--customizations--portal_tweak_settings))

<a id="nestedblock--parameters--customizations--global_customizations"></a>
### Nested Schema for `parameters.customizations.global_customizations`

Optional:

- `background_image` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--global_customizations--background_image))
- `banner_image` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--global_customizations--banner_image))
- `banner_title` (String)
- `contact_text` (String)
- `desktop_logo_image` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--global_customizations--desktop_logo_image))
- `footer_element` (String)
- `mobile_logo_image` (Block List) (see [below for nested schema](#nestedblock--parameters--customizations--global_customizations--mobile_logo_image))

<a id="nestedblock--parameters--customizations--global_customizations--background_image"></a>
### Nested Schema for `parameters.customizations.global_customizations.mobile_logo_image`

Optional:

- `data` (String) Represented as base 64 encoded string of the image byte array


<a id="nestedblock--parameters--customizations--global_customizations--banner_image"></a>
### Nested Schema for `parameters.customizations.global_customizations.mobile_logo_image`

Optional:

- `data` (String) Represented as base 64 encoded string of the image byte array


<a id="nestedblock--parameters--customizations--global_customizations--desktop_logo_image"></a>
### Nested Schema for `parameters.customizations.global_customizations.mobile_logo_image`

Optional:

- `data` (String) Represented as base 64 encoded string of the image byte array


<a id="nestedblock--parameters--customizations--global_customizations--mobile_logo_image"></a>
### Nested Schema for `parameters.customizations.global_customizations.mobile_logo_image`

Optional:

- `data` (String) Represented as base 64 encoded string of the image byte array



<a id="nestedblock--parameters--customizations--language"></a>
### Nested Schema for `parameters.customizations.language`

Optional:

- `view_language` (String)


<a id="nestedblock--parameters--customizations--page_customizations"></a>
### Nested Schema for `parameters.customizations.page_customizations`

Optional:

- `data` (Block List) The Dictionary will be exposed here as key value pair (see [below for nested schema](#nestedblock--parameters--customizations--page_customizations--data))

<a id="nestedblock--parameters--customizations--page_customizations--data"></a>
### Nested Schema for `parameters.customizations.page_customizations.data`

Optional:

- `key` (String)
- `value` (String)



<a id="nestedblock--parameters--customizations--portal_theme"></a>
### Nested Schema for `parameters.customizations.portal_theme`

Optional:

- `name` (String) The system- or user-assigned name of the portal theme
- `theme_data` (String) A CSS file, represented as a Base64-encoded byte array

Read-Only:

- `id` (String) The ID of this resource.


<a id="nestedblock--parameters--customizations--portal_tweak_settings"></a>
### Nested Schema for `parameters.customizations.portal_tweak_settings`

Optional:

- `banner_color` (String) Hex value of color
- `banner_text_color` (String)
- `page_background_color` (String)
- `page_label_and_text_color` (String)



<a id="nestedblock--parameters--settings"></a>
### Nested Schema for `parameters.settings`

Optional:

- `aup_settings` (Block List) Configuration of the Acceptable Use Policy (AUP) for a portal (see [below for nested schema](#nestedblock--parameters--settings--aup_settings))
- `employee_change_password_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--employee_change_password_settings))
- `login_page_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--login_page_settings))
- `portal_settings` (Block List) The port, interface, certificate, and other basic settings of a portal (see [below for nested schema](#nestedblock--parameters--settings--portal_settings))
- `post_access_banner_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--post_access_banner_settings))
- `post_login_banner_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--post_login_banner_settings))
- `support_info_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--settings--support_info_settings))

<a id="nestedblock--parameters--settings--aup_settings"></a>
### Nested Schema for `parameters.settings.aup_settings`

Optional:

- `display_frequency` (String) How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true.
		Allowed Values:
		- FIRSTLOGIN,
		- EVERYLOGIN,
		- RECURRING
- `display_frequency_interval_days` (Number) Number of days between AUP confirmations (when displayFrequency = recurring)
- `include_aup` (String) Require the portal user to read and accept an AUP
- `require_scrolling` (String) Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true


<a id="nestedblock--parameters--settings--employee_change_password_settings"></a>
### Nested Schema for `parameters.settings.employee_change_password_settings`

Optional:

- `allow_employee_to_change_pwd` (String)


<a id="nestedblock--parameters--settings--login_page_settings"></a>
### Nested Schema for `parameters.settings.login_page_settings`

Optional:

- `aup_display` (String) How the AUP should be displayed, either on page or as a link.
		Only valid if includeAup = true.
		Allowed values:
		-  ONPAGE,
		- ASLINK
- `include_aup` (String) Include an Acceptable Use Policy (AUP) that should be displayed during login
- `max_failed_attempts_before_rate_limit` (Number) Maximum failed login attempts before rate limiting
- `require_aup_acceptance` (String) Require the portal user to accept the AUP.
		Only valid if includeAup = true
- `require_scrolling` (String) Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true
- `social_configs` (List of String)
- `time_between_logins_during_rate_limit` (Number) Time between login attempts when rate limiting


<a id="nestedblock--parameters--settings--portal_settings"></a>
### Nested Schema for `parameters.settings.portal_settings`

Optional:

- `allowed_interfaces` (List of String) Interfaces that the portal will be reachable on.
		Allowed values:
		- eth0,
		- eth1,
		- eth2,
		- eth3,
		- eth4,
		- eth5,
		- bond0,
		- bond1,
		- bond2
- `always_used_language` (String)
- `certificate_group_tag` (String) Logical name of the x.509 server certificate that will be used for the portal
- `display_lang` (String) Allowed values:
		- USEBROWSERLOCALE,
		- ALWAYSUSE
- `endpoint_identity_group` (String) Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal
- `fallback_language` (String) Used when displayLang = USEBROWSERLOCALE
- `https_port` (Number) The port number that the allowed interfaces will listen on. Range from 8000 to 8999


<a id="nestedblock--parameters--settings--post_access_banner_settings"></a>
### Nested Schema for `parameters.settings.post_access_banner_settings`

Optional:

- `include_post_access_banner` (String)


<a id="nestedblock--parameters--settings--post_login_banner_settings"></a>
### Nested Schema for `parameters.settings.post_login_banner_settings`

Optional:

- `include_post_access_banner` (String) Include a Post-Login Banner page


<a id="nestedblock--parameters--settings--support_info_settings"></a>
### Nested Schema for `parameters.settings.support_info_settings`

Optional:

- `default_empty_field_value` (String) The default value displayed for an empty field.
		Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE
- `empty_field_display` (String) Specifies how empty fields are handled on the Support Information Page. Allowed values:
		- HIDE,
		- DISPLAYWITHNOVALUE,
		- DISPLAYWITHDEFAULTVALUE
- `include_browser_user_agent` (String)
- `include_failure_code` (String)
- `include_ip_address` (String)
- `include_mac_addr` (String)
- `include_policy_server` (String)
- `include_support_info_page` (String)



<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



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
- `employee_change_password_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--employee_change_password_settings))
- `login_page_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--login_page_settings))
- `portal_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--portal_settings))
- `post_access_banner_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--post_access_banner_settings))
- `post_login_banner_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--post_login_banner_settings))
- `support_info_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--settings--support_info_settings))

<a id="nestedobjatt--item--settings--aup_settings"></a>
### Nested Schema for `item.settings.aup_settings`

Read-Only:

- `display_frequency` (String)
- `display_frequency_interval_days` (Number)
- `include_aup` (String)
- `require_scrolling` (String)


<a id="nestedobjatt--item--settings--employee_change_password_settings"></a>
### Nested Schema for `item.settings.employee_change_password_settings`

Read-Only:

- `allow_employee_to_change_pwd` (String)


<a id="nestedobjatt--item--settings--login_page_settings"></a>
### Nested Schema for `item.settings.login_page_settings`

Read-Only:

- `aup_display` (String)
- `include_aup` (String)
- `max_failed_attempts_before_rate_limit` (Number)
- `require_aup_acceptance` (String)
- `require_scrolling` (String)
- `social_configs` (List of String)
- `time_between_logins_during_rate_limit` (Number)


<a id="nestedobjatt--item--settings--portal_settings"></a>
### Nested Schema for `item.settings.portal_settings`

Read-Only:

- `allowed_interfaces` (List of String)
- `always_used_language` (String)
- `certificate_group_tag` (String)
- `display_lang` (String)
- `endpoint_identity_group` (String)
- `fallback_language` (String)
- `https_port` (Number)


<a id="nestedobjatt--item--settings--post_access_banner_settings"></a>
### Nested Schema for `item.settings.post_access_banner_settings`

Read-Only:

- `include_post_access_banner` (String)


<a id="nestedobjatt--item--settings--post_login_banner_settings"></a>
### Nested Schema for `item.settings.post_login_banner_settings`

Read-Only:

- `include_post_access_banner` (String)


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

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_my_device_portal.example "id:=string"
```
