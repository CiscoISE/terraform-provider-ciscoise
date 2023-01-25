terraform {
  required_providers {
    ciscoise = {
      version = "0.6.11-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_authorization_rules" "example" {
  provider = ciscoise
  parameters {
    profile = ["Blackhole_Wireless_Access"]
    rule {
      # condition {
      #   condition_type = "ConditionAttributes"
      #   is_negate      = "false"
      #   dictionary_name = "IdentityGroup"
      #   attribute_name = "Name"
      #   operator = "equals"
      #   attribute_value = "Endpoint Identity Groups:IAC_Lab1"
      # }
      default = "true"
      # hit_counts = 1
      # id         = "716b2dfd-a1e0-4d94-9802-eb4d3bc614d3"
      name  = "Default"
      rank  = 0
      state = "enabled"
    }
    policy_id="244253d8-cfc9-43b7-8241-38199c71c852"
    id = "00b8bea1-375e-41fe-acb4-ba5ca368f0e8"
    # security_group = "BYOD"
  }
}

# {
#     "profile": [
#         "Blackhole_Wireless_Access"
#     ],
#     "rule": {
#         "default": false,
#         "name": "TestAnsibleIssue81",
#         "rank": 0,
#         "state": "enabled",
#         "condition": {
#             "conditionType": "ConditionAttributes",
#             "isNegate": false,
#             "dictionaryName": "IdentityGroup",
#             "attributeName": "Name",
#             "operator": "equals",
#             "attributeValue": "Endpoint Identity Groups:IAC_Lab1"
#         }
#     }
# }