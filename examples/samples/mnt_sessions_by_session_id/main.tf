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

data "ciscoise_mnt_sessions_by_session_id" "example" {
  provider   = ciscoise
  session_id = "null"
}

output "ciscoise_mnt_sessions_by_session_id_example" {
  value = data.ciscoise_mnt_sessions_by_session_id.example.item
}
