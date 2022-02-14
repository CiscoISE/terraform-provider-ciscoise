
resource "ciscoise_patch" "example" {
  provider = ciscoise
  parameters {
    repository_name = "Workstation"
    patch_number    = 99
    patch_name      = "ise-patchbundle-3.1.0.518-Patch99-21093017.SPA.x86_64.tar.gz"
  }
}

output "ciscoise_patch_example" {
  value = ciscoise_patch.example
}