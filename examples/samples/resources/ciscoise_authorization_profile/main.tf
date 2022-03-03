terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_authorization_profile" "example" {
  provider = ciscoise
  parameters {
    access_type                 = "ACCESS_ACCEPT"
    authz_profile_type          = "SWITCH"
    dacl_name                   = "PERMIT_ALL_IPV4_TRAFFIC"
    description                 = "Onboard the device with Cisco temporal agent"
    easywired_session_candidate = "false"
    name                        = "Cisco_Temporal_Onboard"
    profile_name                = "Cisco"
    service_template            = "false"
    track_movement              = "false"
    web_redirection {
      web_redirection_type = "ClientProvisioning"
      acl                  = "ACL_WEBAUTH_REDIRECT"
      portal_name          = "Client Provisioning Portal (default)"
    }
  }
}

output "ciscoise_authorization_profile_example" {
  value = ciscoise_authorization_profile.example
}
