
resource "ciscoise_sgt" "example" {
  provider = ciscoise
  parameters {

    description       = "string"
    generation_id     = "string"
    id                = "string"
    is_read_only      = "false"
    name              = "string"
    propogate_to_apic = "false"
    value             = 1
  }
}

output "ciscoise_sgt_example" {
  value = ciscoise_sgt.example
}