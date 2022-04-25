
resource "ciscoise_csr_generate" "example" {
  provider             = ciscoise
  allow_wild_card_cert = "false"
  certificate_policies = "string"
  digest_type          = "string"
  hostnames            = ["string"]
  key_length           = "string"
  key_type             = "string"
  portal_group_tag     = "string"
  san_dns              = ["string"]
  san_dir              = ["string"]
  san_ip               = ["string"]
  san_uri              = ["string"]
  subject_city         = "string"
  subject_common_name  = "string"
  subject_country      = "string"
  subject_org          = "string"
  subject_org_unit     = "string"
  subject_state        = "string"
  used_for             = "string"
}