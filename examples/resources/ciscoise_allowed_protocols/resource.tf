
resource "ciscoise_allowed_protocols" "example" {
  provider = ciscoise
  parameters {

    allow_chap                   = "false"
    allow_eap_fast               = "false"
    allow_eap_md5                = "false"
    allow_eap_tls                = "false"
    allow_eap_ttls               = "false"
    allow_leap                   = "false"
    allow_ms_chap_v1             = "false"
    allow_ms_chap_v2             = "false"
    allow_pap_ascii              = "false"
    allow_peap                   = "false"
    allow_preferred_eap_protocol = "false"
    allow_teap                   = "false"
    allow_weak_ciphers_for_eap   = "false"
    description                  = "string"
    eap_fast {

      allow_eap_fast_eap_gtc                                                  = "false"
      allow_eap_fast_eap_gtc_pwd_change                                       = "false"
      allow_eap_fast_eap_gtc_pwd_change_retries                               = 1
      allow_eap_fast_eap_ms_chap_v2                                           = "false"
      allow_eap_fast_eap_ms_chap_v2_pwd_change                                = "false"
      allow_eap_fast_eap_ms_chap_v2_pwd_change_retries                        = 1
      allow_eap_fast_eap_tls                                                  = "false"
      allow_eap_fast_eap_tls_auth_of_expired_certs                            = "false"
      eap_fast_dont_use_pacs_accept_client_cert                               = "false"
      eap_fast_dont_use_pacs_allow_machine_authentication                     = "false"
      eap_fast_enable_eap_chaining                                            = "false"
      eap_fast_use_pacs                                                       = "false"
      eap_fast_use_pacs_accept_client_cert                                    = "false"
      eap_fast_use_pacs_allow_anonym_provisioning                             = "false"
      eap_fast_use_pacs_allow_authen_provisioning                             = "false"
      eap_fast_use_pacs_allow_machine_authentication                          = "false"
      eap_fast_use_pacs_authorization_pac_ttl                                 = 1
      eap_fast_use_pacs_authorization_pac_ttl_units                           = "string"
      eap_fast_use_pacs_machine_pac_ttl                                       = 1
      eap_fast_use_pacs_machine_pac_ttl_units                                 = "string"
      eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning = "false"
      eap_fast_use_pacs_stateless_session_resume                              = "false"
      eap_fast_use_pacs_tunnel_pac_ttl                                        = 1
      eap_fast_use_pacs_tunnel_pac_ttl_units                                  = "string"
      eap_fast_use_pacs_use_proactive_pac_update_precentage                   = 1
    }
    eap_tls {

      allow_eap_tls_auth_of_expired_certs     = "false"
      eap_tls_enable_stateless_session_resume = "false"
      eap_tls_session_ticket_precentage       = 1
      eap_tls_session_ticket_ttl              = 1
      eap_tls_session_ticket_ttl_units        = "string"
    }
    eap_tls_l_bit = "false"
    eap_ttls {

      eap_ttls_chap                              = "false"
      eap_ttls_eap_md5                           = "false"
      eap_ttls_eap_ms_chap_v2                    = "false"
      eap_ttls_eap_ms_chap_v2_pwd_change         = "false"
      eap_ttls_eap_ms_chap_v2_pwd_change_retries = 1
      eap_ttls_ms_chap_v1                        = "false"
      eap_ttls_ms_chap_v2                        = "false"
      eap_ttls_pap_ascii                         = "false"
    }
    id   = "string"
    name = "string"
    peap {

      allow_peap_eap_gtc                           = "false"
      allow_peap_eap_gtc_pwd_change                = "false"
      allow_peap_eap_gtc_pwd_change_retries        = 1
      allow_peap_eap_ms_chap_v2                    = "false"
      allow_peap_eap_ms_chap_v2_pwd_change         = "false"
      allow_peap_eap_ms_chap_v2_pwd_change_retries = 1
      allow_peap_eap_tls                           = "false"
      allow_peap_eap_tls_auth_of_expired_certs     = "false"
      allow_peap_v0                                = "false"
      require_cryptobinding                        = "false"
    }
    preferred_eap_protocol = "string"
    process_host_lookup    = "false"
    require_message_auth   = "false"
    teap {

      accept_client_cert_during_tunnel_est         = "false"
      allow_downgrade_msk                          = "false"
      allow_teap_eap_ms_chap_v2                    = "false"
      allow_teap_eap_ms_chap_v2_pwd_change         = "false"
      allow_teap_eap_ms_chap_v2_pwd_change_retries = 1
      allow_teap_eap_tls                           = "false"
      allow_teap_eap_tls_auth_of_expired_certs     = "false"
      enable_eap_chaining                          = "false"
      request_basic_pwd_auth                       = "false"
    }
  }
}

output "ciscoise_allowed_protocols_example" {
  value = ciscoise_allowed_protocols.example
}