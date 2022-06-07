terraform {
  required_providers {
    ciscoise = {
      version = "0.6.3-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_backup_restore" "response" {
  provider = ciscoise
  parameters {
    backup_encryption_key = "string"
    repository_name       = "string"
    restore_file          = "string"
    restore_include_adeos = "string"
  }
}

