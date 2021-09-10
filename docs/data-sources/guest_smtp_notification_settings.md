---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_guest_smtp_notification_settings Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_guest_smtp_notification_settings (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filter** (List of String)
- **filter_type** (String)
- **id** (String) The ID of this resource.
- **page** (Number)
- **size** (Number)
- **sortasc** (String)
- **sortdsc** (String)

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **connection_timeout** (String)
- **default_from_address** (String)
- **id** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- **notification_enabled** (Boolean)
- **password** (String)
- **smtp_port** (String)
- **smtp_server** (String)
- **use_default_from_address** (Boolean)
- **use_password_authentication** (Boolean)
- **use_tlsor_ssl_encryption** (Boolean)
- **user_name** (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **id** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--items--link))

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

