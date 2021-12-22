
data "ciscoise_hotpatch" "example" {
  provider = ciscoise
}

output "ciscoise_hotpatch_example" {
  value = data.ciscoise_hotpatch.example.items
}
