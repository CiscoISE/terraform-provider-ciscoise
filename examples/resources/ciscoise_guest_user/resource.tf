
resource "ciscoise_guest_user" "example" {
    provider = ciscoise
    item {
      
      custom_fields {}
      description = "string"
      guest_access_info {
        
        from_date = "string"
        group_tag = "string"
        location = "string"
        ssid = "string"
        to_date = "string"
        valid_days = 1
      }
      guest_info {
        
        company = "string"
        creation_time = "string"
        email_address = "string"
        enabled = false
        first_name = "string"
        last_name = "string"
        notification_language = "string"
        password = "******"
        phone_number = "string"
        sms_service_provider = "string"
        user_name = "string"
      }
      guest_type = "string"
      id = "string"
      name = "string"
      portal_id = "string"
      reason_for_visit = "string"
      sponsor_user_id = "string"
      sponsor_user_name = "string"
      status = "string"
      status_reason = "string"
    }
}

output "ciscoise_guest_user_example" {
    value = ciscoise_guest_user.example
}