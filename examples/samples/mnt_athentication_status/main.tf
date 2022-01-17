terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_athentication_status" "example" {
  provider  = ciscoise
  mac       = "2C:54:91:88:C9:E3"
  rec_ord_s = "1"
  sec_ond_s = "1"
}

output "ciscoise_mnt_athentication_status_example" {
  value = data.ciscoise_mnt_athentication_status.example.item
}
