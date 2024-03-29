---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_selfsigned_certificate_generate Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs create operation on Certificates.
  - Generate Self-signed Certificate
  NOTE:
  The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.
  NOTE:
  Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.
  NOTE:
  Wildcard certificate and SAML certificate can be generated only on PPAN or Standalone
---

# ciscoise_selfsigned_certificate_generate (Resource)

It performs create operation on Certificates.
- Generate Self-signed Certificate
NOTE:
The certificate may have a validity period longer than 398 days. It may be untrusted by many browsers.
NOTE:
Request parameters accepting True and False as input can be replaced by 1 and 0 respectively.
NOTE:
Wildcard certificate and SAML certificate can be generated only on PPAN or Standalone


~>Warning: This resource does not represent a real-world entity in Cisco ISE, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Cisco ISE workflow. It is executed in ISE without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "ciscoise_selfsigned_certificate_generate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    admin                                      = "false"
    allow_extended_validity                    = "false"
    allow_portal_tag_transfer_for_same_subject = "false"
    allow_replacement_of_certificates          = "false"
    allow_replacement_of_portal_group_tag      = "false"
    allow_role_transfer_for_same_subject       = "false"
    allow_san_dns_bad_name                     = "false"
    allow_san_dns_non_resolvable               = "false"
    allow_wild_card_certificates               = "false"
    certificate_policies                       = "string"
    digest_type                                = "string"
    eap                                        = "false"
    expiration_ttl                             = 1
    expiration_ttl_unit                        = "string"
    host_name                                  = "string"

    key_length          = "string"
    key_type            = "string"
    name                = "string"
    portal              = "false"
    portal_group_tag    = "string"
    pxgrid              = "false"
    radius              = "false"
    saml                = "false"
    san_dns             = ["string"]
    san_ip              = ["string"]
    san_uri             = ["string"]
    subject_city        = "string"
    subject_common_name = "string"
    subject_country     = "string"
    subject_org         = "string"
    subject_org_unit    = "string"
    subject_state       = "string"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `admin` (String) Use certificate to authenticate the Cisco ISE Admin Portal
- `allow_extended_validity` (String) Allow generation of self-signed certificate with validity greater than 398 days
- `allow_portal_tag_transfer_for_same_subject` (String) Allow overwriting the portal tag from matching certificate of same subject
- `allow_replacement_of_certificates` (String) Allow Replacement of certificates
- `allow_replacement_of_portal_group_tag` (String) Allow Replacement of Portal Group Tag
- `allow_role_transfer_for_same_subject` (String) Allow transfer of roles for certificate with matching subject
- `allow_san_dns_bad_name` (String) Allow usage of SAN DNS Bad name
- `allow_san_dns_non_resolvable` (String) Allow use of non resolvable Common Name or SAN Values
- `allow_wild_card_certificates` (String) Allow Wildcard Certificates
- `certificate_policies` (String) Certificate Policies
- `digest_type` (String) Digest to sign with
- `eap` (String) Use certificate for EAP protocols that use SSL/TLS tunneling
- `expiration_ttl` (Number) Certificate expiration value
- `expiration_ttl_unit` (String) Certificate expiration unit
- `host_name` (String) Hostname of the Cisco ISE node in which self-signed certificate should be generated.
- `key_length` (String) Bit size of public key
- `key_type` (String) Algorithm to use for certificate public key creation
- `name` (String) Friendly name of the certificate.
- `portal` (String) Use for portal
- `portal_group_tag` (String) Set Group tag
- `pxgrid` (String) Use certificate for the pxGrid Controller
- `radius` (String) Use certificate for the RADSec server
- `saml` (String) Use certificate for SAML Signing
- `san_dns` (List of String) Array of SAN (Subject Alternative Name) DNS entries
- `san_ip` (List of String) Array of SAN IP entries
- `san_uri` (List of String) Array of SAN URI entries
- `subject_city` (String) Certificate city or locality (L)
- `subject_common_name` (String) Certificate common name (CN)
- `subject_country` (String) Certificate country (C)
- `subject_org` (String) Certificate organization (O)
- `subject_org_unit` (String) Certificate organizational unit (OU)
- `subject_state` (String) Certificate state (ST)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `message` (String)
- `status` (String)


