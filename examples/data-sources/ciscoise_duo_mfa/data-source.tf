
data "ciscoise_duo_mfa" "example" {
  provider = ciscoise
}

output "ciscoise_duo_mfa_example" {
  value = data.ciscoise_duo_mfa.example.items
}

data "ciscoise_duo_mfa" "example" {
  provider        = ciscoise
  connection_name = "string"
}

output "ciscoise_duo_mfa_example" {
  value = data.ciscoise_duo_mfa.example.item
}
