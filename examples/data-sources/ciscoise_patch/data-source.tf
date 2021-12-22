
data "ciscoise_patch" "example" {
  provider = ciscoise
}

output "ciscoise_patch_example" {
  value = data.ciscoise_patch.example.item
}
