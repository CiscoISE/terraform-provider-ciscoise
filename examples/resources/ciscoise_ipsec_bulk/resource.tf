
resource "ciscoise_ipsec_bulk" "example" {
  provider = ciscoise
  parameters {

    auth_type                 = "string"
    cert_id                   = "string"
    configure_vti             = "false"
    esp_ah_protocol           = "string"
    host_name                 = "string"
    iface                     = "string"
    ike_re_auth_time          = 1
    ike_version               = "string"
    local_internal_ip         = "string"
    mode_option               = "string"
    nad_ip                    = "string"
    phase_one_dhgroup         = "string"
    phase_one_encryption_algo = "string"
    phase_one_hash_algo       = "string"
    phase_one_life_time       = 1
    phase_two_dhgroup         = "string"
    phase_two_encryption_algo = "string"
    phase_two_hash_algo       = "string"
    phase_two_life_time       = 1
    psk                       = "string"
    remote_peer_internal_ip   = "string"
  }
}

output "ciscoise_ipsec_bulk_example" {
  value = ciscoise_ipsec_bulk.example
}