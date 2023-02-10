terraform {
  required_providers {
    ciscoise = {
      version = "0.6.14-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_allowed_protocols" "response" {
  provider = ciscoise
  parameters {
    name                         = var.name
    description                  = var.description
    process_host_lookup          = "false"
    allow_pap_ascii              = "false"
    allow_chap                   = "false"
    allow_ms_chap_v1             = "false"
    allow_ms_chap_v2             = "false"
    allow_eap_md5                = "false"
    allow_leap                   = "false"
    allow_eap_ttls               = "false"
    allow_eap_fast               = "false"
    allow_teap                   = "false"
    allow_preferred_eap_protocol = "false"
    eap_tls_l_bit                = "false"
    allow_weak_ciphers_for_eap   = "false"
    require_message_auth         = "false"

    # Set to true if eap_tls is uncommented
    # allow_eap_tls = "true"
    # Set to false if eap_tls is commented
    allow_eap_tls = var.allow_eap_tls
    dynamic "eap_tls" {
      for_each = var.allow_eap_tls == "true" ? [1] : []
      content {
        allow_eap_tls_auth_of_expired_certs = "false"
        # Set to false if following args are commented
        eap_tls_enable_stateless_session_resume = "false"
        # Set to true if following args are uncommented
        # eap_tls_enable_stateless_session_resume = "true"
        # Following args
        # eap_tls_session_ticket_precentage = 1
        # eap_tls_session_ticket_ttl        = 1
        # eap_tls_session_ticket_ttl_units  = "SECONDS"
      }
    }
    # Set to true if peap is uncommented
    allow_peap = "true"
    # Set to false if eap_tls is commented
    # allow_peap = "false"
    peap {
      allow_peap_eap_ms_chap_v2                    = "true"
      allow_peap_eap_ms_chap_v2_pwd_change         = "true"
      allow_peap_eap_ms_chap_v2_pwd_change_retries = 1
      allow_peap_eap_gtc                           = "true"
      allow_peap_eap_gtc_pwd_change                = "true"
      allow_peap_eap_gtc_pwd_change_retries        = 1
      allow_peap_eap_tls                           = "true"
      allow_peap_eap_tls_auth_of_expired_certs     = "false"
      require_cryptobinding                        = "false"
      allow_peap_v0                                = "false"
    }

  }
}

data "ciscoise_allowed_protocols" "example" {
  depends_on = [
    ciscoise_allowed_protocols.response
  ]
  provider = ciscoise
  name     = var.name
}
