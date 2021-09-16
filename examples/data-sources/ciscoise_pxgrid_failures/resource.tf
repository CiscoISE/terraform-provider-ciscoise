
data "ciscoise_pxgrid_failures" "example" {
    provider = ciscoise
}

output "ciscoise_pxgrid_failures_example" {
    value = data.ciscoise_pxgrid_failures.example.item
}
