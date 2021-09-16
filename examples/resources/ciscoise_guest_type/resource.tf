
resource "ciscoise_guest_type" "example" {
    provider = ciscoise
    item {
      
      access_time {
        
        allow_access_on_specific_days_times = false
        day_time_limits {
          
          days = ["string"]
          end_time = "string"
          start_time = "string"
        }
        default_duration = 1
        duration_time_unit = "string"
        from_first_login = false
        max_account_duration = 1
      }
      description = "string"
      expiration_notification {
        
        advance_notification_duration = 1
        advance_notification_units = "string"
        email_text = "string"
        enable_notification = false
        send_email_notification = false
        send_sms_notification = false
        sms_text = "string"
      }
      id = "string"
      is_default_type = false
      login_options {
        
        allow_guest_portal_bypass = false
        failure_action = "string"
        identity_group_id = "string"
        limit_simultaneous_logins = false
        max_registered_devices = 1
        max_simultaneous_logins = 1
      }
      name = "string"
      sponsor_groups = ["string"]
    }
}

output "ciscoise_guest_type_example" {
    value = ciscoise_guest_type.example
}