---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_trusted_certificate_export Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Certificates.
  The response of this API carries a trusted certificate file mapped to the requested ID
---

# ciscoise_trusted_certificate_export (Data Source)

It performs read operation on Certificates.

- The response of this API carries a trusted certificate file mapped to the requested ID

## Example Usage

```terraform
data "ciscoise_trusted_certificate_export" "example" {
  provider = ciscoise
  dirpath  = "string"
  id       = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dirpath` (String) Directory absolute path in which to save the file.
- `id` (String) id path parameter. ID of the Trusted Certificate to be exported.


