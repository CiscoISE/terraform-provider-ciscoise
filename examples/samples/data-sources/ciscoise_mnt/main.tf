terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_version" "response" {
  provider = ciscoise
}
output "ciscoise__mnt_version_response" {
  value = data.ciscoise_mnt_version.response
}


data "ciscoise_mnt_session_active_list" "response" {
  provider = ciscoise
}
output "ciscoise__mnt_session_active_list_response" {
  value = data.ciscoise_mnt_session_active_list.response
}

data "ciscoise_mnt_session_active_count" "response" {
  provider = ciscoise
}
output "ciscoise__mnt_session_active_count_response" {
  value = data.ciscoise_mnt_session_active_count.response
}


data "ciscoise_mnt_sessions_by_session_id" "response" {
  provider   = ciscoise
  session_id = "1"
}
output "ciscoise__mnt_sessions_by_session_id_response" {
  value = data.ciscoise_mnt_sessions_by_session_id.response
}


data "ciscoise_mnt_session_posture_count" "response" {
  provider = ciscoise
}
output "ciscoise__mnt_session_posture_count_response" {
  value = data.ciscoise_mnt_session_posture_count.response
}

data "ciscoise_mnt_session_auth_list" "response" {
  provider = ciscoise
}
output "ciscoise__mnt_session_auth_list_response" {
  value = data.ciscoise_mnt_session_auth_list.response
}
