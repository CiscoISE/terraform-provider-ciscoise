---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_sxp_local_bindings Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on SXPLocalBindings.
  This data source allows the client to get a SXP local binding by ID.This data source allows the client to get all the SXP local bindings.
  Filter:
  [name, description]
  Sorting:
  [name, description]
---

# ciscoise_sxp_local_bindings (Data Source)

It performs read operation on SXPLocalBindings.

- This data source allows the client to get a SXP local binding by ID.

- This data source allows the client to get all the SXP local bindings.

Filter:

[name, description]

Sorting:

[name, description]

## Example Usage

```terraform
data "ciscoise_sxp_local_bindings" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sxp_local_bindings_example" {
  value = data.ciscoise_sxp_local_bindings.example.items
}

data "ciscoise_sxp_local_bindings" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_sxp_local_bindings_example" {
  value = data.ciscoise_sxp_local_bindings.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (List of String) filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains
- `filter_type` (String) filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
- `id` (String) id path parameter.
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page
- `sortasc` (String) sortasc query parameter. sort asc
- `sortdsc` (String) sortdsc query parameter. sort desc

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `binding_name` (String)
- `description` (String)
- `id` (String)
- `ip_address_or_host` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `sgt` (String)
- `sxp_vpn` (String)
- `vns` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--items--link))

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


