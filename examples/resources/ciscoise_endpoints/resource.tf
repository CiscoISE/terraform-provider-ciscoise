
resource "ciscoise_endpoints" "example" {
  provider = ciscoise
  item {























  }
  parameters {






















    value = "string"

  }
}

output "ciscoise_endpoints_example" {
  value = ciscoise_endpoints.example
}