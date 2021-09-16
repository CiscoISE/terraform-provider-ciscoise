
data "ciscoise_profiler_profile" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_profiler_profile_example" {
    value = data.ciscoise_profiler_profile.example.items
}

data "ciscoise_profiler_profile" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_profiler_profile_example" {
    value = data.ciscoise_profiler_profile.example.item
}
