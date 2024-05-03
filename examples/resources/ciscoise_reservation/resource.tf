
resource "ciscoise_reservation" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    client_id      = "string"
    client_name    = "string"
    end_index      = 1
    number_of_tags = 1

    start_index = 1

  }
}

output "ciscoise_reservation_example" {
  value = ciscoise_reservation.example
}