terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.1"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


# Run it once for each backup
data "ciscoise_backup_config" "backup_config" {
  provider              = ciscoise
  backup_encryption_key = "My3ncryptionkey"
  backup_name           = "myBackup"
  repository_name       = "Temp"
}

output "backup_config" {
  value = data.ciscoise_backup_config.backup_config
}

# data "ciscoise_tasks" "task_status" {
#   task_id = "fd4f4ef0-2092-11ec-92b5-f645229364e5"
#   # data.ciscoise_backup_config.backup_config.item[0].id
# }

# output "task_status" {
#   value = data.ciscoise_tasks.task_status.item
# }


data "ciscoise_backup_last_status" "backup_status" {
  provider = ciscoise
}

output "ciscoise_backup_last_status_backup_status" {
  value = data.ciscoise_backup_last_status.backup_status.item
}

# data "ciscoise_backup_cancel" "example" {
#   provider = ciscoise
# }


# data "ciscoise_backup_restore" "backup_restore" {
#   provider              = ciscoise
#   backup_encryption_key = "My3ncryptionkey"
#   repository_name       = "Temp"
#   restore_file          = "myBackup-CFG10-210928-1933.tar.gpg"
#   restore_include_adeos = "true"
# }

# output "backup_restore" {
#   value = data.ciscoise_backup_restore.backup_restore
# }