
resource "ciscoise_mnt_session_delete_all" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}