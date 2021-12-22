
resource "ciscoise_portal_theme" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    theme_data  = "string"
  }
}

output "ciscoise_portal_theme_example" {
  value = ciscoise_portal_theme.example
}