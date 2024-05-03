
data "ciscoise_px_grid_direct_dictionary_info" "example" {
  provider = ciscoise
}

output "ciscoise_px_grid_direct_dictionary_info_example" {
  value = data.ciscoise_px_grid_direct_dictionary_info.example.item
}
