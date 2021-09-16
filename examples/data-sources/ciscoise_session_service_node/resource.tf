
data "ciscoise_session_service_node" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_session_service_node_example" {
    value = data.ciscoise_session_service_node.example.item_name
}

data "ciscoise_session_service_node" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_session_service_node_example" {
    value = data.ciscoise_session_service_node.example.item_id
}

data "ciscoise_session_service_node" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_session_service_node_example" {
    value = data.ciscoise_session_service_node.example.items
}
