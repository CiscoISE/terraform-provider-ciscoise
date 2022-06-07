terraform {
  required_providers {
    ciscoise = {
      version = "0.6.2-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_tasks" "response" {
  provider = ciscoise
}
output "ciscoise_tasks_response" {
  value = data.ciscoise_tasks.response
}

data "ciscoise_tasks" "single_response" {
  provider = ciscoise
  task_id  = data.ciscoise_tasks.response.items[0].id
}

output "ciscoise_tasks_single_response" {
  value = data.ciscoise_tasks.single_response
}
