
data "ciscoise_backup_restore" "example" {
    provider = ciscoise
    backup_encryption_key = "string"
    repository_name = "string"
    restore_file = "string"
    restore_include_adeos = "string"
}