
data "ciscoise_pan_ha" "example" {
  provider = ciscoise
}

output "ciscoise_pan_ha_example" {
  value = data.ciscoise_pan_ha.example.item
}
