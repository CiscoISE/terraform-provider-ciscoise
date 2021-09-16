
data "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_tacacs_profile_example" {
  value = data.ciscoise_tacacs_profile.example.item_name
}

data "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_tacacs_profile_example" {
  value = data.ciscoise_tacacs_profile.example.item_id
}

data "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_tacacs_profile_example" {
  value = data.ciscoise_tacacs_profile.example.items
}
