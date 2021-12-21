
data "ciscoise_node_services_profiler_probe_config" "example" {
  provider = ciscoise
  hostname = "string"
}

output "ciscoise_node_services_profiler_probe_config_example" {
  value = data.ciscoise_node_services_profiler_probe_config.example.item
}
