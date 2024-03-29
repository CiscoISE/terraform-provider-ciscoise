---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_active_directory Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update* and delete operations on ActiveDirectory.
  This resource deletes an AD join point from Cisco ISE.This resource creates an AD join point in Cisco ISE.*This resource action loads domain groups configuration from Active Directory into Cisco ISE.
---

# ciscoise_active_directory (Resource)

It manages create, read, update* and delete operations on ActiveDirectory.

- This resource deletes an AD join point from Cisco ISE.

- This resource creates an AD join point in Cisco ISE.

- *This resource action loads domain groups configuration from Active Directory into Cisco ISE.

## Example Usage

```terraform
resource "ciscoise_active_directory" "example" {
  provider = ciscoise
  parameters {

    ad_attributes {

      attributes {

        default_value = "string"
        internal_name = "string"
        name          = "string"
        type          = "string"
      }
    }
    ad_scopes_names = "string"
    adgroups {

      groups {

        name = "string"
        sid  = "string"
        type = "string"
      }
    }
    advanced_settings {

      aging_time                        = 1
      auth_protection_type              = "string"
      country                           = "string"
      department                        = "string"
      email                             = "string"
      enable_callback_for_dialin_client = "false"
      enable_dialin_permission_check    = "false"
      enable_failed_auth_protection     = "false"
      enable_machine_access             = "false"
      enable_machine_auth               = "false"
      enable_pass_change                = "false"
      enable_rewrites                   = "false"
      failed_auth_threshold             = 1
      first_name                        = "string"
      identity_not_in_ad_behaviour      = "string"
      job_title                         = "string"
      last_name                         = "string"
      locality                          = "string"
      organizational_unit               = "string"
      plaintext_auth                    = "false"
      rewrite_rules {

        rewrite_match  = "string"
        rewrite_result = "string"
        row_id         = 1
      }
      schema                        = "string"
      state_or_province             = "string"
      street_address                = "string"
      telephone                     = "string"
      unreachable_domains_behaviour = "string"
    }
    description              = "string"
    domain                   = "string"
    enable_domain_white_list = "false"
    id                       = "string"
    name                     = "string"
  }
}

output "ciscoise_active_directory_example" {
  value = ciscoise_active_directory.example
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

- `ad_attributes` (Block List) Holds list of AD Attributes (see [below for nested schema](#nestedblock--parameters--ad_attributes))
- `ad_scopes_names` (String) String that contains the names of the scopes that the active directory belongs to. Names are separated by comma. Alphanumeric, underscore (_) characters are allowed
- `adgroups` (Block List) Holds list of AD Groups (see [below for nested schema](#nestedblock--parameters--adgroups))
- `advanced_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--advanced_settings))
- `description` (String) No character restriction
- `domain` (String) The AD domain. Alphanumeric, hyphen (-) and dot (.) characters are allowed
- `enable_domain_white_list` (String)
- `id` (String) Resource UUID value
- `name` (String) Resource Name. Maximum 32 characters allowed. Allowed characters are alphanumeric and .-_/\\ characters

Read-Only:

- `enable_domain_allowed_list` (String)
- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedblock--parameters--ad_attributes"></a>
### Nested Schema for `parameters.ad_attributes`

Optional:

- `attributes` (Block List) List of Attributes (see [below for nested schema](#nestedblock--parameters--ad_attributes--attributes))

<a id="nestedblock--parameters--ad_attributes--attributes"></a>
### Nested Schema for `parameters.ad_attributes.attributes`

Optional:

- `default_value` (String) Required for each attribute in the attribute list. Can contain an empty string. All characters are allowed except <%"
- `internal_name` (String) Required for each attribute in the attribute list. All characters are allowed except <%"
- `name` (String) Required for each attribute in the attribute list with no duplication between attributes. All characters are allowed except <%"
- `type` (String) Required for each group in the group list. Allowed values: STRING, IP, BOOLEAN, INT, OCTET_STRING



<a id="nestedblock--parameters--adgroups"></a>
### Nested Schema for `parameters.adgroups`

Optional:

- `groups` (Block List) List of Groups (see [below for nested schema](#nestedblock--parameters--adgroups--groups))

<a id="nestedblock--parameters--adgroups--groups"></a>
### Nested Schema for `parameters.adgroups.groups`

Optional:

- `name` (String) Required for each group in the group list with no duplication between groups. All characters are allowed except %
- `sid` (String) Cisco ISE uses security identifiers (SIDs) for optimization of group membership evaluation. SIDs are useful for efficiency (speed) when the groups are evaluated. All characters are allowed except %
- `type` (String) No character restriction



<a id="nestedblock--parameters--advanced_settings"></a>
### Nested Schema for `parameters.advanced_settings`

Optional:

- `aging_time` (Number) Range 1-8760 hours
- `auth_protection_type` (String) Enable prevent AD account lockout. Allowed values:
		- WIRELESS,
		- WIRED,
		- BOTH
- `country` (String) User info attribute. All characters are allowed except %
- `department` (String) User info attribute. All characters are allowed except %
- `email` (String) User info attribute. All characters are allowed except %
- `enable_callback_for_dialin_client` (String)
- `enable_dialin_permission_check` (String)
- `enable_failed_auth_protection` (String) Enable prevent AD account lockout due to too many bad password attempts
- `enable_machine_access` (String)
- `enable_machine_auth` (String)
- `enable_pass_change` (String)
- `enable_rewrites` (String)
- `failed_auth_threshold` (Number) Number of bad password attempts
- `first_name` (String) User info attribute. All characters are allowed except %
- `identity_not_in_ad_behaviour` (String) Allowed values: REJECT, SEARCH_JOINED_FOREST, SEARCH_ALL
- `job_title` (String) User info attribute. All characters are allowed except %
- `last_name` (String) User info attribute. All characters are allowed except %
- `locality` (String) User info attribute. All characters are allowed except %
- `organizational_unit` (String) User info attribute. All characters are allowed except %
- `plaintext_auth` (String)
- `rewrite_rules` (Block List) Identity rewrite is an advanced feature that directs Cisco ISE to manipulate the identity
		before it is passed to the external Active Directory system. You can create rules to change
		the identity to a desired format that includes or excludes a domain prefix and/or suffix or
		other additional markup of your choice (see [below for nested schema](#nestedblock--parameters--advanced_settings--rewrite_rules))
- `schema` (String) Allowed values: ACTIVE_DIRECTORY, CUSTOM.
		Choose ACTIVE_DIRECTORY schema when the AD attributes defined in AD can be copied to relevant attributes
		in Cisco ISE. If customization is needed, choose CUSTOM schema. All User info attributes are always set to
		default value if schema is ACTIVE_DIRECTORY. Values can be changed only for CUSTOM schema
- `state_or_province` (String) User info attribute. All characters are allowed except %
- `street_address` (String) User info attribute. All characters are allowed except %
- `telephone` (String) User info attribute. All characters are allowed except %
- `unreachable_domains_behaviour` (String) Allowed values: PROCEED, DROP

<a id="nestedblock--parameters--advanced_settings--rewrite_rules"></a>
### Nested Schema for `parameters.advanced_settings.rewrite_rules`

Optional:

- `rewrite_match` (String) Required for each rule in the list with no duplication between rules. All characters are allowed except %"
- `rewrite_result` (String) Required for each rule in the list. All characters are allowed except %"
- `row_id` (Number) Required for each rule in the list in serial order



<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `ad_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--item--ad_attributes))
- `ad_scopes_names` (String)
- `adgroups` (List of Object) (see [below for nested schema](#nestedobjatt--item--adgroups))
- `advanced_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--advanced_settings))
- `description` (String)
- `domain` (String)
- `enable_domain_allowed_list` (String)
- `enable_domain_white_list` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)

<a id="nestedobjatt--item--ad_attributes"></a>
### Nested Schema for `item.ad_attributes`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--item--ad_attributes--attributes))

<a id="nestedobjatt--item--ad_attributes--attributes"></a>
### Nested Schema for `item.ad_attributes.attributes`

Read-Only:

- `default_value` (String)
- `internal_name` (String)
- `name` (String)
- `type` (String)



<a id="nestedobjatt--item--adgroups"></a>
### Nested Schema for `item.adgroups`

Read-Only:

- `groups` (List of Object) (see [below for nested schema](#nestedobjatt--item--adgroups--groups))

<a id="nestedobjatt--item--adgroups--groups"></a>
### Nested Schema for `item.adgroups.groups`

Read-Only:

- `name` (String)
- `sid` (String)
- `type` (String)



<a id="nestedobjatt--item--advanced_settings"></a>
### Nested Schema for `item.advanced_settings`

Read-Only:

- `aging_time` (Number)
- `auth_protection_type` (String)
- `country` (String)
- `department` (String)
- `email` (String)
- `enable_callback_for_dialin_client` (String)
- `enable_dialin_permission_check` (String)
- `enable_failed_auth_protection` (String)
- `enable_machine_access` (String)
- `enable_machine_auth` (String)
- `enable_pass_change` (String)
- `enable_rewrites` (String)
- `failed_auth_threshold` (Number)
- `first_name` (String)
- `identity_not_in_ad_behaviour` (String)
- `job_title` (String)
- `last_name` (String)
- `locality` (String)
- `organizational_unit` (String)
- `plaintext_auth` (String)
- `rewrite_rules` (List of Object) (see [below for nested schema](#nestedobjatt--item--advanced_settings--rewrite_rules))
- `schema` (String)
- `state_or_province` (String)
- `street_address` (String)
- `telephone` (String)
- `unreachable_domains_behaviour` (String)

<a id="nestedobjatt--item--advanced_settings--rewrite_rules"></a>
### Nested Schema for `item.advanced_settings.rewrite_rules`

Read-Only:

- `rewrite_match` (String)
- `rewrite_result` (String)
- `row_id` (Number)



<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_active_directory.example "id:=string\name:=string"
```
