
resource "ciscoise_rest_id_store" "example" {
  provider = ciscoise
  item {

    description = "string"
    ers_rest_idstore_attributes {

      headers {

        key   = "string"
        value = "string"
      }
      predefined      = "string"
      root_url        = "string"
      username_suffix = "string"
    }
    id   = "string"
    name = "string"
  }
}

output "ciscoise_rest_id_store_example" {
  value = ciscoise_rest_id_store.example
}