terraform {
  required_providers {
    ciscoise = {
      version = "1.0.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_tasks" "response" {
  provider = ciscoise
}
output "ciscoise__tasks_response" {
  value = data.ciscoise_tasks.response
}

# data "ciscoise_tasks" "single_response" {
#   provider = ciscoise
#   task_id       = data.ciscoise_tasks.response.items[0].id
# }

# output "ciscoise__tasks_single_response" {
#   value = data.ciscoise_tasks.single_response
# }
