---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_trustsec_vn_bulk_delete Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs create operation on virtualNetwork.
  Delete Virtual Network in bulk
---

# ciscoise_trustsec_vn_bulk_delete (Data Source)

It performs create operation on virtualNetwork.

- Delete Virtual Network in bulk

## Example Usage

```terraform
data "ciscoise_trustsec_vn_bulk_delete" "example" {
  provider = ciscoise
  payload  = ["string"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **payload** (List of String) Array of RequestVirtualNetworkBulkDeleteVirtualNetworks

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **id** (String)

