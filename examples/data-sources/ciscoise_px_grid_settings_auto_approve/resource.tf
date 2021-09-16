
data "ciscoise_px_grid_settings_auto_approve" "example" {
    provider = ciscoise
    allow_password_based_accounts = false
    auto_approve_cert_based_accounts = false
}