
data "ciscoise_aci_settings" "example" {
    provider = ciscoise
}

output "ciscoise_aci_settings_example" {
    value = data.ciscoise_aci_settings.example.item
}
