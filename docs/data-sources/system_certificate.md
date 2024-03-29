---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_system_certificate Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Certificates.
  This data source supports Filtering, Sorting and Pagination.
  Filtering and Sorting supported on below mentioned attributes:
  friendlyName
  issuedTo
  issuedBy
  validFrom
  Supported Date Format: yyyy-MM-dd HH:mm:ss
  Supported Operators: EQ, NEQ, GT and LT
  expirationDate
  Supported Date Format: yyyy-MM-dd HH:mm:ss
  Supported Operators: EQ, NEQ, GT and LT
  This data source provides details of a System Certificate of a particular node based on given HostName and ID.
---

# ciscoise_system_certificate (Data Source)

It performs read operation on Certificates.

- This data source supports Filtering, Sorting and Pagination.

Filtering and Sorting supported on below mentioned attributes:


friendlyName

issuedTo

issuedBy

validFrom


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT


expirationDate


Supported Date Format: yyyy-MM-dd HH:mm:ss

Supported Operators: EQ, NEQ, GT and LT




- This data source provides details of a System Certificate of a particular node based on given HostName and ID.

## Example Usage

```terraform
data "ciscoise_system_certificate" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  host_name   = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_system_certificate_example" {
  value = data.ciscoise_system_certificate.example.items
}

data "ciscoise_system_certificate" "example" {
  provider  = ciscoise
  host_name = "string"
  id        = "string"
}

output "ciscoise_system_certificate_example" {
  value = data.ciscoise_system_certificate.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (List of String) filter query parameter. 
 
 
 
Simple filtering
 should be available through the filter query string parameter. The structure of a filter is a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the 
"filterType=or"
 query string parameter. Each resource Data model description should specify if an attribute is a filtered field. 
 
 
 
 
 
OPERATOR
 
DESCRIPTION
 
 
 
 
 
EQ
 
Equals
 
 
 
NEQ
 
Not Equals
 
 
 
GT
 
Greater Than
 
 
 
LT
 
Less Then
 
 
 
STARTSW
 
Starts With
 
 
 
NSTARTSW
 
Not Starts With
 
 
 
ENDSW
 
Ends With
 
 
 
NENDSW
 
Not Ends With
 
 
 
CONTAINS
 
Contains
 
 
 
NCONTAINS
 
Not Contains
- `filter_type` (String) filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter
- `host_name` (String) hostName path parameter. Name of the host of which system certificates should be returned
- `id` (String) id path parameter. ID of the system certificate
- `page` (Number) page query parameter. Page number
- `size` (Number) size query parameter. Number of objects returned per page
- `sort` (String) sort query parameter. sort type asc or desc
- `sort_by` (String) sortBy query parameter. sort column by which objects needs to be sorted

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `expiration_date` (String)
- `friendly_name` (String)
- `group_tag` (String)
- `id` (String)
- `issued_by` (String)
- `issued_to` (String)
- `key_size` (Number)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `portals_using_the_tag` (String)
- `self_signed` (String)
- `serial_number_decimal_format` (String)
- `sha256_fingerprint` (String)
- `signature_algorithm` (String)
- `used_by` (String)
- `valid_from` (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `expiration_date` (String)
- `friendly_name` (String)
- `group_tag` (String)
- `id` (String)
- `issued_by` (String)
- `issued_to` (String)
- `key_size` (Number)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--items--link))
- `portals_using_the_tag` (String)
- `self_signed` (String)
- `serial_number_decimal_format` (String)
- `sha256_fingerprint` (String)
- `signature_algorithm` (String)
- `used_by` (String)
- `valid_from` (String)

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


