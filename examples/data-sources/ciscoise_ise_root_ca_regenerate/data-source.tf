
data "ciscoise_ise_root_ca_regenerate" "example" {
  provider                             = ciscoise
  remove_existing_ise_intermediate_csr = "false"
}