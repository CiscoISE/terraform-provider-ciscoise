
resource "ciscoise_tacacs_command_sets" "example" {
  provider = ciscoise
  parameters {

    commands {

      command_list {

        arguments = "string"
        command   = "string"
        grant     = "string"
      }
    }
    description      = "string"
    id               = "string"
    name             = "string"
    permit_unmatched = "false"
  }
}

output "ciscoise_tacacs_command_sets_example" {
  value = ciscoise_tacacs_command_sets.example
}