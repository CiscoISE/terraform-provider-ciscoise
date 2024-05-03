
data "ciscoise_reservation" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_reservation_example" {
  value = data.ciscoise_reservation.example.items
}

data "ciscoise_reservation" "example" {
  provider  = ciscoise
  client_id = "string"
}

output "ciscoise_reservation_example" {
  value = data.ciscoise_reservation.example.item
}
