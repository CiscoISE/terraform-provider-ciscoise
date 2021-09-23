
resource "ciscoise_id_store_sequence" "example" {
  provider = ciscoise
  item {

    break_on_store_fail                = "false"
    certificate_authentication_profile = "string"
    description                        = "string"
    id                                 = "string"
    id_seq_item {

      idstore = "string"
      order   = 1
    }
    name   = "string"
    parent = "string"
  }
}

output "ciscoise_id_store_sequence_example" {
  value = ciscoise_id_store_sequence.example
}