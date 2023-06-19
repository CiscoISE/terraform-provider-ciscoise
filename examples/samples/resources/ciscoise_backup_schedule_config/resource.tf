terraform {
  required_providers {
    ciscoise = {
      version = "0.6.20-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_backup_schedule_config" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    backup_description    = "string"
    backup_encryption_key = "string"
    backup_name           = "string"
    end_date              = "string"
    frequency             = "string"
    month_day             = "string"
    repository_name       = "string"
    start_date            = "string"
    status                = "string"
    time                  = "string"
    week_day              = "string"
  }
}