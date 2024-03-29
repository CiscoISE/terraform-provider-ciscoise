---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_endpoint_register Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs update operation on endpoint.
  - This resource allows the client to register an endpoint.
---

# ciscoise_endpoint_register (Resource)

It performs update operation on endpoint.
- This resource allows the client to register an endpoint.


~>Warning: This resource does not represent a real-world entity in Cisco ISE, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Cisco ISE workflow. It is executed in ISE without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "ciscoise_endpoint_register" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    custom_attributes {

    }
    description       = "string"
    group_id          = "string"
    id                = "string"
    identity_store    = "string"
    identity_store_id = "string"

    mac = "string"
    mdm_attributes {

      mdm_compliance_status = "false"
      mdm_encrypted         = "false"
      mdm_enrolled          = "false"
      mdm_ime_i             = "string"
      mdm_jail_broken       = "false"
      mdm_manufacturer      = "string"
      mdm_model             = "string"
      mdm_os                = "string"
      mdm_phone_number      = "string"
      mdm_pinlock           = "false"
      mdm_reachable         = "false"
      mdm_serial            = "string"
      mdm_server_name       = "string"
    }
    name                      = "string"
    portal_user               = "string"
    profile_id                = "string"
    static_group_assignment   = "false"
    static_profile_assignment = "false"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `custom_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--custom_attributes))
- `description` (String)
- `group_id` (String)
- `identity_store` (String)
- `identity_store_id` (String)
- `mac` (String)
- `mdm_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--mdm_attributes))
- `name` (String)
- `portal_user` (String)
- `profile_id` (String)
- `static_group_assignment` (String)
- `static_profile_assignment` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--parameters--custom_attributes"></a>
### Nested Schema for `parameters.custom_attributes`

Optional:

- `custom_attributes` (Map of String) Key value map


<a id="nestedblock--parameters--mdm_attributes"></a>
### Nested Schema for `parameters.mdm_attributes`

Optional:

- `mdm_compliance_status` (String)
- `mdm_encrypted` (String)
- `mdm_enrolled` (String)
- `mdm_ime_i` (String)
- `mdm_jail_broken` (String)
- `mdm_manufacturer` (String)
- `mdm_model` (String)
- `mdm_os` (String)
- `mdm_phone_number` (String)
- `mdm_pinlock` (String)
- `mdm_reachable` (String)
- `mdm_serial` (String)
- `mdm_server_name` (String)


