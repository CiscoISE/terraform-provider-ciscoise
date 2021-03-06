
resource "ciscoise_repository" "example" {
  provider = ciscoise
  parameters {

    enable_pki      = "false"
    name            = "string"
    password        = "******"
    path            = "string"
    protocol        = "string"
    repository_name = "string"
    server_name     = "string"
    user_name       = "string"
  }
}

output "ciscoise_repository_example" {
  value = ciscoise_repository.example
}