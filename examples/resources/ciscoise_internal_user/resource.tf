
resource "ciscoise_internal_user" "example" {
  provider = ciscoise
  item {

    change_password = "false"
    custom_attributes {}
    description         = "string"
    email               = "string"
    enable_password     = "string"
    enabled             = "false"
    expiry_date         = "string"
    expiry_date_enabled = "false"
    first_name          = "string"
    id                  = "string"
    identity_groups     = "string"
    last_name           = "string"
    name                = "string"
    password            = "******"
    password_idstore    = "******"
  }
}

output "ciscoise_internal_user_example" {
  value = ciscoise_internal_user.example
}