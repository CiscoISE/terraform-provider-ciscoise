output "ciscoise_allowed_protocols_response_item" {
  depends_on = [
    ciscoise_allowed_protocols.response
  ]
  value = ciscoise_allowed_protocols.response.item
}

output "ciscoise_allowed_protocols_response_id" {
  depends_on = [
    ciscoise_allowed_protocols.response
  ]
  value = ciscoise_allowed_protocols.response.id
}

output "ciscoise_allowed_protocols_response_item_datasource" {
  depends_on = [
    data.ciscoise_allowed_protocols.example
  ]
  value = data.ciscoise_allowed_protocols.example.item_name
}

output "ciscoise_allowed_protocols_response_item_description" {
  depends_on = [
    ciscoise_allowed_protocols.response
  ]
  value = length(ciscoise_allowed_protocols.response.item) > 0 ? ciscoise_allowed_protocols.response.item.0.description : ""
}

output "ciscoise_allowed_protocols_response_item_description_datasource" {
  depends_on = [
    data.ciscoise_allowed_protocols.example
  ]
  value = length(data.ciscoise_allowed_protocols.example.item_name) > 0 ? data.ciscoise_allowed_protocols.example.item_name.0.description : ""
}
