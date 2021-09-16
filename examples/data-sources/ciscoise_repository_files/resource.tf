
data "ciscoise_repository_files" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_repository_files_example" {
  value = data.ciscoise_repository_files.example.items
}
