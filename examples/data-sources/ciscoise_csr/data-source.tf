
data "ciscoise_csr" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_csr_example" {
  value = data.ciscoise_csr.example.items
}

data "ciscoise_csr" "example" {
  provider  = ciscoise
  host_name = "string"
  id        = "string"
}

output "ciscoise_csr_example" {
  value = data.ciscoise_csr.example.item
}
