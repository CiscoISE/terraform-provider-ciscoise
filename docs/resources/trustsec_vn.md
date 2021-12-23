---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_trustsec_vn Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on virtualNetwork.
  Create Virtual NetworkUpdate Virtual NetworkDelete Virtual Network
---

# ciscoise_trustsec_vn (Resource)

It manages create, read, update and delete operations on virtualNetwork.

- Create Virtual Network

- Update Virtual Network

- Delete Virtual Network

## Example Usage

```terraform
resource "ciscoise_trustsec_vn" "example" {
  provider = ciscoise
  parameters {

    additional_attributes = "string"
    id                    = "string"
    last_update           = "string"
    name                  = "string"
  }
}

output "ciscoise_trustsec_vn_example" {
  value = ciscoise_trustsec_vn.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **parameters** (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **additional_attributes** (String) JSON String of additional attributes for the Virtual Network
- **id** (String) Identifier of the Virtual Network
- **last_update** (String) Timestamp for the last update of the Virtual Network
- **name** (String) Name of the Virtual Network


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **additional_attributes** (String)
- **id** (String)
- **last_update** (String)
- **name** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_trustsec_vn.example "id:=string"
```