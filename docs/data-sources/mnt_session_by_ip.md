---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_mnt_session_by_ip Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Misc.
  Sessions by Endpoint IP
---

# ciscoise_mnt_session_by_ip (Data Source)

It performs read operation on Misc.

- Sessions by Endpoint IP

## Example Usage

```terraform
data "ciscoise_mnt_session_by_ip" "example" {
  provider      = ciscoise
  endpoint_ipv4 = "string"
}

output "ciscoise_mnt_session_by_ip_example" {
  value = data.ciscoise_mnt_session_by_ip.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint_ipv4` (String) endpoint_ipv4 path parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)


