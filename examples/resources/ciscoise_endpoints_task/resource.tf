
resource "ciscoise_endpoints_task" "example" {
  provider = ciscoise
}

output "ciscoise_endpoints_task_example" {
  value = ciscoise_endpoints_task.example
}