
resource "ciscoise_node_services_profiler_probe_config" "example" {
  provider = ciscoise
  parameters {

    days_before_rescan = 1
  }
}

output "ciscoise_node_services_profiler_probe_config_example" {
  value = ciscoise_node_services_profiler_probe_config.example
}