---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_certificate_template Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on CertificateTemplate.
  This data source allows the client to get a certificate template by name.This data source allows the client to get a certificate template by ID.This data source allows the client to get aall the certificate templates.
---

# ciscoise_certificate_template (Data Source)

It performs read operation on CertificateTemplate.

- This data source allows the client to get a certificate template by name.

- This data source allows the client to get a certificate template by ID.

- This data source allows the client to get aall the certificate templates.

## Example Usage

```terraform
data "ciscoise_certificate_template" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_certificate_template_example" {
  value = data.ciscoise_certificate_template.example.item_name
}

data "ciscoise_certificate_template" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_certificate_template_example" {
  value = data.ciscoise_certificate_template.example.item_id
}

data "ciscoise_certificate_template" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_certificate_template_example" {
  value = data.ciscoise_certificate_template.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter.
- `name` (String) name path parameter.
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page

### Read-Only

- `item_id` (List of Object) (see [below for nested schema](#nestedatt--item_id))
- `item_name` (List of Object) (see [below for nested schema](#nestedatt--item_name))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item_id"></a>
### Nested Schema for `item_id`

Read-Only:

- `description` (String)
- `id` (String)
- `key_size` (Number)
- `name` (String)
- `raprofile` (String)
- `validity_period` (Number)


<a id="nestedatt--item_name"></a>
### Nested Schema for `item_name`

Read-Only:

- `description` (String)
- `id` (String)
- `key_size` (Number)
- `name` (String)
- `raprofile` (String)
- `validity_period` (Number)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String)
- `id` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--items--link))
- `name` (String)

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


