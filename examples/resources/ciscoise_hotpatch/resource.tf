
resource "ciscoise_hotpatch" "example" {
  provider = ciscoise
  parameters {
    repository_name = "Workstation"
    hotpatch_name   = "ise-apply-CSCvz53724_3.2.x_patchall-SPA.tar.gz"
  }
}

output "ciscoise_hotpatch_example" {
  value = ciscoise_hotpatch.example
}