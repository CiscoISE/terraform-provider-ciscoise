terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.4"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_guest_user" "example" {
  provider = ciscoise
  parameters {

    custom_fields = {
      Authorization = "Internet"
      Owner         = "wilhelm"
    }
    description = "Guest user test"
    guest_access_info {

      from_date = "12/22/2021 17:40"
      # group_tag  = "string"
      location = "San Jose"
      # ssid       = "string"
      to_date    = "12/23/2021 17:40"
      valid_days = 1
    }
    guest_info {

      # company               = "string"
      # creation_time         = "string"
      email_address = "user@cisco.com"
      enabled       = "true"
      first_name    = "user"
      last_name     = "example"
      # notification_language = "string"
      password = "C1sco12345"
      # phone_number          = "string"
      # sms_service_provider  = "string"
      user_name = "1user"
    }
    guest_type = "Daily (default)"
    # id                = "string"
    name             = "string"
    portal_id        = "bd48c1a1-9477-4746-8e40-e43d20c9f429"
    reason_for_visit = "ISE Guest Services"
    # sponsor_user_id   = "string"
    # sponsor_user_name = "string"
    # status            = "string"
    # status_reason     = "string"
  }
}

output "ciscoise_guest_user_example" {
  sensitive = true
  value     = ciscoise_guest_user.example
}