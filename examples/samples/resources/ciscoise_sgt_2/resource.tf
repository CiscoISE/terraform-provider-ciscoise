terraform {
  required_providers {
    ciscoise = {
      version = "0.6.8-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}
# Configure provider with your  Cisco Identity Services Engine SDK credentials
provider "ciscoise" {
  enable_auto_import = "false"
}
resource "ciscoise_sgt" "example" {
  provider = ciscoise
  parameters {

    description       = "Terraform created with updatieee"
    #generation_id     = "string123"
    # id                = "8ed911d4-b2aa-46b8-ba97-97684d88f1d5"
    is_read_only      = "false"
    name              = "TerraformSGT00000000"
    propogate_to_apic = "true"
    value             = 111
  }
}