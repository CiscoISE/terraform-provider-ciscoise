---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_external_radius_server Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on ExternalRADIUSServer.
  This resource allows the client to update an external RADIUS server.This resource deletes an external RADIUS server.This resource creates an external RADIUS server.
---

# ciscoise_external_radius_server (Resource)

It manages create, read, update and delete operations on ExternalRADIUSServer.

- This resource allows the client to update an external RADIUS server.

- This resource deletes an external RADIUS server.

- This resource creates an external RADIUS server.

## Example Usage

```terraform
resource "ciscoise_external_radius_server" "example" {
  provider = ciscoise
  parameters {

    accounting_port     = 1
    authentication_port = 1
    authenticator_key   = "string"
    description         = "string"
    enable_key_wrap     = "false"
    encryption_key      = "string"
    host_ip             = "string"
    id                  = "string"
    key_input_format    = "string"
    name                = "string"
    proxy_timeout       = 1
    retries             = 1
    shared_secret       = "string"
    timeout             = 1
  }
}

output "ciscoise_external_radius_server_example" {
  value = ciscoise_external_radius_server.example
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

- `accounting_port` (Number) Valid Range 1 to 65535
- `authentication_port` (Number) Valid Range 1 to 65535
- `authenticator_key` (String) The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
		The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
- `description` (String)
- `enable_key_wrap` (String) KeyWrap may only be enabled if it is supported on the device.
		When running in FIPS mode this option should be enabled for such devices
- `encryption_key` (String) The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
		The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')
- `host_ip` (String) The IP of the host - must be a valid IPV4 address
- `key_input_format` (String) Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'.
		Allowed Values:
		- ASCII
		- HEXADECIMAL
- `name` (String) Resource Name. Allowed charactera are alphanumeric and _ (underscore).
- `proxy_timeout` (Number) Valid Range 1 to 600
- `retries` (Number) Valid Range 1 to 9
- `shared_secret` (String) Shared secret maximum length is 128 characters
- `timeout` (Number) Valid Range 1 to 120

Read-Only:

- `id` (String) The ID of this resource.
- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `accounting_port` (Number)
- `authentication_port` (Number)
- `authenticator_key` (String)
- `description` (String)
- `enable_key_wrap` (String)
- `encryption_key` (String)
- `host_ip` (String)
- `id` (String)
- `key_input_format` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `proxy_timeout` (Number)
- `retries` (Number)
- `shared_secret` (String)
- `timeout` (Number)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_external_radius_server.example "id:=string\name:=string"
```
