
resource "ciscoise_trustsec_vn" "example" {
  provider = ciscoise
  item {





  }
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