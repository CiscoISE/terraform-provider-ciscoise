
resource "ciscoise_trustsec_nbar_app" "example" {
  provider = ciscoise
  item {




    network_identities {



    }
  }
  parameters {


    id = "string"

    network_identities {



    }
    ports    = "string"
    protocol = "string"
  }
}

output "ciscoise_trustsec_nbar_app_example" {
  value = ciscoise_trustsec_nbar_app.example
}