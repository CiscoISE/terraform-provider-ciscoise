
data "ciscoise_licensing_feature_to_tier_mapping" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_feature_to_tier_mapping_example" {
  value = data.ciscoise_licensing_feature_to_tier_mapping.example.items
}
