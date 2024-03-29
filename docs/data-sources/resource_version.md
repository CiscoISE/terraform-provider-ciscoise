---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_resource_version Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on VersionInfo.
  Get all VersionInfo
---

# ciscoise_resource_version (Data Source)

It performs read operation on VersionInfo.

- Get all VersionInfo

## Example Usage

```terraform
data "ciscoise_resource_version" "example" {
  provider = ciscoise
  resource = "string"
}

output "ciscoise_resource_version_example" {
  value = data.ciscoise_resource_version.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource` (String) resource path parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `current_server_version` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `supported_versions` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


