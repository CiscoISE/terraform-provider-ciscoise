---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_aci_test_connectivity Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs update operation on ACISettings.
  This data source action allows the client to test ACI Domain Manager connection.
---

# ciscoise_aci_test_connectivity (Data Source)

It performs update operation on ACISettings.

- This data source action allows the client to test ACI Domain Manager connection.

## Example Usage

```terraform
data "ciscoise_aci_test_connectivity" "example" {
  provider = ciscoise

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `result` (String)


