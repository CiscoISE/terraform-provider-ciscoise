
resource "ciscoise_anc_policy" "example" {
  provider = ciscoise
  item {

    actions = ["string"]
    id      = "string"
    name    = "string"
  }
}

output "ciscoise_anc_policy_example" {
  value = ciscoise_anc_policy.example
}