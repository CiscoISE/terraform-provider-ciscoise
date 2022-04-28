
resource "ciscoise_backup_restore" "example" {
  provider              = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    backup_encryption_key = "string"
    repository_name       = "string"
    restore_file          = "string"
    restore_include_adeos = "string"
  }
  
}