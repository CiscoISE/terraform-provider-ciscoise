
resource "ciscoise_radius_server_sequence" "example" {
  provider = ciscoise
  item {

    before_accept_attr_manipulators_list {

      action          = "string"
      attribute_name  = "string"
      changed_val     = "string"
      dictionary_name = "string"
      value           = "string"
    }
    on_request_attr_manipulator_list {

      action          = "string"
      attribute_name  = "string"
      changed_val     = "string"
      dictionary_name = "string"
      value           = "string"
    }
    radius_server_list      = ["string"]
    continue_authorz_policy = "false"
    description             = "string"
    id                      = "string"
    local_accounting        = "false"
    name                    = "string"
    prefix_separator        = "string"
    remote_accounting       = "false"
    strip_prefix            = "false"
    strip_suffix            = "false"
    suffix_separator        = "string"
    use_attr_set_before_acc = "false"
    use_attr_set_on_request = "false"
  }
}

output "ciscoise_radius_server_sequence_example" {
  value = ciscoise_radius_server_sequence.example
}