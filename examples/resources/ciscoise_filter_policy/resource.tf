
resource "ciscoise_filter_policy" "example" {
    provider = ciscoise
    item {
      
      domains = "string"
      id = "string"
      sgt = "string"
      subnet = "string"
      vn = "string"
    }
}

output "ciscoise_filter_policy_example" {
    value = ciscoise_filter_policy.example
}