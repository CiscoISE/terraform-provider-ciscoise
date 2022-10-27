terraform {
  required_providers {
    ciscoise = {
      version = "0.6.11-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_guest_smtp_notification_settings" "example" {
  provider = ciscoise

}

output "ciscoise_guest_smtp_notification_settings_example" {
  value = data.ciscoise_guest_smtp_notification_settings.example.items
}

data "ciscoise_guest_smtp_notification_settings" "example1" {
  provider = ciscoise
  id       = data.ciscoise_guest_smtp_notification_settings.example.items[0].id
}

output "ciscoise_guest_smtp_notification_settings_example1" {
  value = data.ciscoise_guest_smtp_notification_settings.example1.item
}

resource "ciscoise_guest_smtp_notification_settings" "actual" {
  provider = ciscoise
  parameters {
    connection_timeout          = data.ciscoise_guest_smtp_notification_settings.example1.item[0].connection_timeout
    default_from_address        = data.ciscoise_guest_smtp_notification_settings.example1.item[0].default_from_address
    id                          = data.ciscoise_guest_smtp_notification_settings.example1.item[0].id
    notification_enabled        = "false"
    password                    = data.ciscoise_guest_smtp_notification_settings.example1.item[0].password
    smtp_port                   = data.ciscoise_guest_smtp_notification_settings.example1.item[0].smtp_port
    smtp_server                 = "dcloud.cisco.com"
    use_default_from_address    = data.ciscoise_guest_smtp_notification_settings.example1.item[0].use_default_from_address
    use_password_authentication = data.ciscoise_guest_smtp_notification_settings.example1.item[0].use_password_authentication
    use_tlsor_ssl_encryption    = data.ciscoise_guest_smtp_notification_settings.example1.item[0].use_tlsor_ssl_encryption
    user_name                   = data.ciscoise_guest_smtp_notification_settings.example1.item[0].user_name
  }
}

output "ciscoise_guest_smtp_notification_settings_actual" {
  value = ciscoise_guest_smtp_notification_settings.actual.item[0].notification_enabled
}
