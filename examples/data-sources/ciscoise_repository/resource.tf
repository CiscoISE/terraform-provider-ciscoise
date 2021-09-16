
data "ciscoise_repository" "example" {
    provider = ciscoise
}

output "ciscoise_repository_example" {
    value = data.ciscoise_repository.example.items
}

data "ciscoise_repository" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_repository_example" {
    value = data.ciscoise_repository.example.item
}
