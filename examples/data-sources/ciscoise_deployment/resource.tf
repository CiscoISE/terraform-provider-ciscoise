
data "ciscoise_deployment" "example" {
    provider = ciscoise
}

output "ciscoise_deployment_example" {
    value = data.ciscoise_deployment.example.item
}
