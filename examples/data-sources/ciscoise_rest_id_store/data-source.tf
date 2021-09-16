
data "ciscoise_rest_id_store" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_rest_id_store_example" {
  value = data.ciscoise_rest_id_store.example.item_name
}

data "ciscoise_rest_id_store" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_rest_id_store_example" {
  value = data.ciscoise_rest_id_store.example.item_id
}

data "ciscoise_rest_id_store" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_rest_id_store_example" {
  value = data.ciscoise_rest_id_store.example.items
}
