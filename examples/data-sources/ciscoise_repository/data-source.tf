
data "ciscoise_repository" "example" {
  provider = ciscoise
}

output "ciscoise_repository_example" {
  value = data.ciscoise_repository.example.items
}

data "ciscoise_repository" "example" {
  provider        = ciscoise
  repository_name = "string"
}

output "ciscoise_repository_example" {
  value = data.ciscoise_repository.example.item
}
