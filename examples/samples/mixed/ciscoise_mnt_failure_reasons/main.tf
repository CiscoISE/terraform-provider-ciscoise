terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_failure_reasons" "example" {
  provider = ciscoise
}

# Terraform will detect the resource as having been deleted
# each time a configuration is applied on a new machine where the file is not present and will
# generate a diff to re-create it. This may cause "noise" in diffs in environments
# where configurations are routinely applied by many different users or within automation systems.
resource "local_file" "mnt_failure_reasons" {
  content  = data.ciscoise_mnt_failure_reasons.example.item
  filename = "${path.module}/mnt_failure_reasons.xml"
}
