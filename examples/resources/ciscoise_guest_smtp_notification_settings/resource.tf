
resource "ciscoise_guest_smtp_notification_settings" "example" {
  provider = ciscoise
  parameters {

    connection_timeout          = "string"
    default_from_address        = "string"
    id                          = "string"
    notification_enabled        = "false"
    password                    = "******"
    smtp_port                   = "string"
    smtp_server                 = "string"
    use_default_from_address    = "false"
    use_password_authentication = "false"
    use_tlsor_ssl_encryption    = "false"
    user_name                   = "string"
  }
}

output "ciscoise_guest_smtp_notification_settings_example" {
  value = ciscoise_guest_smtp_notification_settings.example
}