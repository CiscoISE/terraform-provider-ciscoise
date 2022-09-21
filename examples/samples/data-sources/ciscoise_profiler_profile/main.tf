terraform {
  required_providers {
    ciscoise = {
      version = "0.6.6-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_profiler_profile" "example" {
  provider = ciscoise

}

output "ciscoise_profiler_profile_example" {
  value = data.ciscoise_profiler_profile.example.items
}

data "ciscoise_profiler_profile" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_profiler_profile.example.items[0].id
}

output "ciscoise_profiler_profile_single_response" {
  value = data.ciscoise_profiler_profile.single_response.item
}
