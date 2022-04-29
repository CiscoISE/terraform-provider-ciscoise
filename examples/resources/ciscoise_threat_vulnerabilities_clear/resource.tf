resource "ciscoise_threat_vulnerabilities_clear" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    mac_addresses = "string"
  }
}