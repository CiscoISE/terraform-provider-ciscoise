
resource "ciscoise_sponsor_group" "example" {
    provider = ciscoise
    item {
      
      auto_notification = false
      create_permissions {
        
        can_create_random_accounts = false
        can_import_multiple_accounts = false
        can_set_future_start_date = false
        can_specify_username_prefix = false
        default_username_prefix = "string"
        import_batch_size_limit = 1
        random_batch_size_limit = 1
        start_date_future_limit_days = 1
      }
      description = "string"
      guest_types = ["string"]
      id = "string"
      is_default_group = false
      is_enabled = false
      locations = ["string"]
      manage_permission = "string"
      member_groups = ["string"]
      name = "string"
      other_permissions {
        
        can_access_via_rest = false
        can_approve_selfreg_guests = false
        can_delete_guest_accounts = false
        can_extend_guest_accounts = false
        can_reinstate_suspended_accounts = false
        can_reset_guest_passwords = false
        can_send_sms_notifications = false
        can_suspend_guest_accounts = false
        can_update_guest_contact_info = false
        can_view_guest_passwords = false
        limit_approval_to_sponsors_guests = false
        require_suspension_reason = false
      }
    }
}

output "ciscoise_sponsor_group_example" {
    value = ciscoise_sponsor_group.example
}