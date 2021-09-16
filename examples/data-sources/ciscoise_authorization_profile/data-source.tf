
data "ciscoise_authorization_profile" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_authorization_profile_example" {
  value = data.ciscoise_authorization_profile.example.item_name
}

data "ciscoise_authorization_profile" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_authorization_profile_example" {
  value = data.ciscoise_authorization_profile.example.item_id
}

data "ciscoise_authorization_profile" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_authorization_profile_example" {
  value = data.ciscoise_authorization_profile.example.items
}
