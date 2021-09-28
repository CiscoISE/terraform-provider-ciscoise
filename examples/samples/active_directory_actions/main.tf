terraform {
  required_providers {
    ciscoise = {
      version = "0.0.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


data "ciscoise_active_directory" "response" {
  provider = ciscoise
  name = "cisco.com"
}

data "ciscoise_active_directory_add_groups" "add_groups_resp" {
  provider = ciscoise
  id       = data.ciscoise_active_directory.response.item_name[0].id
  name = data.ciscoise_active_directory.response.item_name[0].name
  description = data.ciscoise_active_directory.response.item_name[0].description
  domain = data.ciscoise_active_directory.response.item_name[0].domain
  adgroups {
    groups {
      name = "cisco.com/operators"
      sid  = "S-1-5-32-548"
    }
  }
}

output "add_groups_resp" {
    value = data.ciscoise_active_directory_add_groups.add_groups_resp
}


# data "ciscoise_active_directory_join_domain_with_all_nodes" "join_all_resp" {
#   provider = ciscoise
#   id       = data.ciscoise_active_directory.response.item_name[0].id
#   additional_data {
#     name  = "username"
#     value = "admin"
#   }
#   additional_data {
#     name  = "password"
#     value = "C1sco12345"
#   }
# }

# output "join_all_resp" {
#     value = data.ciscoise_active_directory_join_domain_with_all_nodes.join_all_resp
# }

data "ciscoise_active_directory_get_groups_by_domain_info" "get_groups_by_domain_resp" {
  provider = ciscoise
  id       = data.ciscoise_active_directory.response.item_name[0].id
  additional_data {
    name  = "name"
    value = data.ciscoise_active_directory.response.item_name[0].name
  }
  additional_data {
    name  = "domain"
    value = data.ciscoise_active_directory.response.item_name[0].domain
  }
}

output "get_groups_by_domain_resp" {
    value = data.ciscoise_active_directory_get_groups_by_domain_info.get_groups_by_domain_resp
}