
data "ciscoise_sg_mapping_group" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_sg_mapping_group_example" {
    value = data.ciscoise_sg_mapping_group.example.items
}

data "ciscoise_sg_mapping_group" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_sg_mapping_group_example" {
    value = data.ciscoise_sg_mapping_group.example.item
}
