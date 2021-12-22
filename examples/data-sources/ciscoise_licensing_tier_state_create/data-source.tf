
data "ciscoise_licensing_tier_state_create" "example" {
  provider = ciscoise

  payload {

    name   = "string"
    status = "string"
  }
}