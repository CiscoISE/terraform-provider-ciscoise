
data "ciscoise_tasks" "example" {
    provider = ciscoise
}

output "ciscoise_tasks_example" {
    value = data.ciscoise_tasks.example.items
}

data "ciscoise_tasks" "example" {
    provider = ciscoise
    task_id = "string"
}

output "ciscoise_tasks_example" {
    value = data.ciscoise_tasks.example.item
}
