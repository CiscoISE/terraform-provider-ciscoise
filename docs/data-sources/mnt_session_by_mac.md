---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_mnt_session_by_mac Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Misc.
  Sessions by MAC
---

# ciscoise_mnt_session_by_mac (Data Source)

It performs read operation on Misc.

- Sessions by MAC

## Example Usage

```terraform
data "ciscoise_mnt_session_by_mac" "example" {
  provider = ciscoise
  mac      = "string"
}

output "ciscoise_mnt_session_by_mac_example" {
  value = data.ciscoise_mnt_session_by_mac.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mac` (String) mac path parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)


