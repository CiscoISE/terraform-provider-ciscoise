
data "ciscoise_sg_to_vn_to_vlan" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
}

output "ciscoise_sg_to_vn_to_vlan_example" {
    value = data.ciscoise_sg_to_vn_to_vlan.example.items
}

data "ciscoise_sg_to_vn_to_vlan" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_sg_to_vn_to_vlan_example" {
    value = data.ciscoise_sg_to_vn_to_vlan.example.item
}
