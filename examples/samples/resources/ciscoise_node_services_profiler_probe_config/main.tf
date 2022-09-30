terraform {
  required_providers {
    ciscoise = {
      version = "0.6.7-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_node_services_profiler_probe_config" "example" {
  provider = ciscoise
  parameters {
    hostname = "ise"
    active_directory {
      days_before_rescan = 2
    }
    http {
      interfaces {
        interface = "GigabitEthernet 0"
      }
    }
    dhcp {
      interfaces {
        interface = "GigabitEthernet 0"
      }
      port = 67
    }
    pxgrid = "false"
    radius = "true"
    nmap   = "false"
    dns {
      timeout = 2
    }
    netflow {
      interfaces {
        interface = "GigabitEthernet 0"
      }
      port = 9996
    }
    dhcp_span {
      interfaces {
        interface = "GigabitEthernet 0"
      }
    }
    snmp_query {
      retries       = 2
      timeout       = 1000
      event_timeout = 30
    }
    snmp_trap {
      interfaces {
        interface = "GigabitEthernet 0"
      }
      link_trap_query = "true"
      mac_trap_query  = "true"
      port            = 162
    }
  }
}

# output "ciscoise_node_services_profiler_probe_config_example" {
#   value = ciscoise_node_services_profiler_probe_config.example
# }
