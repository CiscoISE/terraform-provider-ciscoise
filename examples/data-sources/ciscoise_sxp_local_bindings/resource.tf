
data "ciscoise_sxp_local_bindings" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_sxp_local_bindings_example" {
    value = data.ciscoise_sxp_local_bindings.example.items
}

data "ciscoise_sxp_local_bindings" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_sxp_local_bindings_example" {
    value = data.ciscoise_sxp_local_bindings.example.item
}
