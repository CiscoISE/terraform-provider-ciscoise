terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_session_profiler_count" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_profiler_count_example" {
  value = data.ciscoise_mnt_session_profiler_count.example.item
}
