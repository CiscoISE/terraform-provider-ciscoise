
data "ciscoise_native_supplicant_profile" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_native_supplicant_profile_example" {
    value = data.ciscoise_native_supplicant_profile.example.items
}

data "ciscoise_native_supplicant_profile" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_native_supplicant_profile_example" {
    value = data.ciscoise_native_supplicant_profile.example.item
}
