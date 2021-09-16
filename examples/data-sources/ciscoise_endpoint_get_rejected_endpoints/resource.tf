
data "ciscoise_endpoint_get_rejected_endpoints" "example" {
    provider = ciscoise
}

output "ciscoise_endpoint_get_rejected_endpoints_example" {
    value = data.ciscoise_endpoint_get_rejected_endpoints.example.item
}
