
data "ciscoise_repository_files" "example" {
  provider        = ciscoise
  repository_name = "string"
}

output "ciscoise_repository_files_example" {
  value = data.ciscoise_repository_files.example.items
}
