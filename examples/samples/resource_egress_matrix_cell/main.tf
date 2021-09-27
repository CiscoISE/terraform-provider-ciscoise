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

# data "ciscoise_egress_matrix_cell" "example" {
#     provider = ciscoise
#     id = "92c1a900-8c01-11e6-996c-525400b48521"
# }

# output "ciscoise_egress_matrix_cell_example" {
#     value = data.ciscoise_egress_matrix_cell.example.item
# }

data "ciscoise_sgt" "sgt_src" {
  provider = ciscoise
  filter   = ["name.EQ.Quarantined_Systems"]
}
output "ciscoise_sgt_sgt_src_id" {
  value = data.ciscoise_sgt.sgt_src.items[0].id
}

data "ciscoise_sgt" "sgt_dest" {
  provider = ciscoise
  filter   = ["name.EQ.Guests"]
}
output "ciscoise_sgt_sgt_dest_id" {
  value = data.ciscoise_sgt.sgt_dest.items[0].id
}

resource "ciscoise_egress_matrix_cell" "name" {
  provider = ciscoise
  item {
    default_rule       = "DENY_IP"
    matrix_cell_status = "ENABLED"
    description        = "Updated by import utility (3)."
    sgacls             = ["92919850-8c01-11e6-996c-525400b48521"]
    name               = "${data.ciscoise_sgt.sgt_src.items[0].name}-${data.ciscoise_sgt.sgt_dest.items[0].name}"
    destination_sgt_id = data.ciscoise_sgt.sgt_dest.items[0].id
    source_sgt_id      = data.ciscoise_sgt.sgt_src.items[0].id
  }
}