
data "ciscoise_filter_policy" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_filter_policy_example" {
    value = data.ciscoise_filter_policy.example.items
}

data "ciscoise_filter_policy" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_filter_policy_example" {
    value = data.ciscoise_filter_policy.example.item
}
