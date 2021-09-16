
data "ciscoise_anc_policy" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_anc_policy_example" {
    value = data.ciscoise_anc_policy.example.item_name
}

data "ciscoise_anc_policy" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_anc_policy_example" {
    value = data.ciscoise_anc_policy.example.item_id
}

data "ciscoise_anc_policy" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_anc_policy_example" {
    value = data.ciscoise_anc_policy.example.items
}
