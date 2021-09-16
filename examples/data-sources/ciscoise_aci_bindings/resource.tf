
data "ciscoise_aci_bindings" "example" {
    provider = ciscoise
    filter_by = ["string"]
    filter_value = ["string"]
    page = 1
    size = 1
    sort = "string"
    sort_by = "string"
}

output "ciscoise_aci_bindings_example" {
    value = data.ciscoise_aci_bindings.example.item
}
