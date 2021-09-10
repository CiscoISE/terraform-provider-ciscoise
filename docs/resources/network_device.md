---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_network_device Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_network_device (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **item** (Block List) (see [below for nested schema](#nestedblock--item))

### Read-Only

- **last_updated** (String)

<a id="nestedblock--item"></a>
### Nested Schema for `item`

Optional:

- **authentication_settings** (Block List) (see [below for nested schema](#nestedblock--item--authentication_settings))
- **coa_port** (Number)
- **description** (String)
- **dtls_dns_name** (String)
- **id** (String) The ID of this resource.
- **model_name** (String)
- **name** (String)
- **network_device_group_list** (List of String)
- **network_device_iplist** (Block List) (see [below for nested schema](#nestedblock--item--network_device_iplist))
- **profile_name** (String)
- **snmpsettings** (Block List) (see [below for nested schema](#nestedblock--item--snmpsettings))
- **software_version** (String)
- **tacacs_settings** (Block List) (see [below for nested schema](#nestedblock--item--tacacs_settings))
- **trustsecsettings** (Block List) (see [below for nested schema](#nestedblock--item--trustsecsettings))

Read-Only:

- **link** (List of Object) (see [below for nested schema](#nestedatt--item--link))

<a id="nestedblock--item--authentication_settings"></a>
### Nested Schema for `item.authentication_settings`

Optional:

- **dtls_required** (Boolean)
- **enable_key_wrap** (Boolean)
- **enable_multi_secret** (String)
- **enabled** (Boolean)
- **key_encryption_key** (String)
- **key_input_format** (String)
- **message_authenticator_code_key** (String)
- **network_protocol** (String)
- **radius_shared_secret** (String)
- **second_radius_shared_secret** (String)


<a id="nestedblock--item--network_device_iplist"></a>
### Nested Schema for `item.network_device_iplist`

Optional:

- **get_ipaddress_exclude** (String)
- **ipaddress** (String)
- **mask** (Number)


<a id="nestedblock--item--snmpsettings"></a>
### Nested Schema for `item.snmpsettings`

Optional:

- **link_trap_query** (Boolean)
- **mac_trap_query** (Boolean)
- **originating_policy_services_node** (String)
- **polling_interval** (Number)
- **ro_community** (String)
- **version** (String)


<a id="nestedblock--item--tacacs_settings"></a>
### Nested Schema for `item.tacacs_settings`

Optional:

- **connect_mode_options** (String)
- **shared_secret** (String)


<a id="nestedblock--item--trustsecsettings"></a>
### Nested Schema for `item.trustsecsettings`

Optional:

- **device_authentication_settings** (Block List) (see [below for nested schema](#nestedblock--item--trustsecsettings--device_authentication_settings))
- **device_configuration_deployment** (Block List) (see [below for nested schema](#nestedblock--item--trustsecsettings--device_configuration_deployment))
- **push_id_support** (Boolean)
- **sga_notification_and_updates** (Block List) (see [below for nested schema](#nestedblock--item--trustsecsettings--sga_notification_and_updates))

<a id="nestedblock--item--trustsecsettings--device_authentication_settings"></a>
### Nested Schema for `item.trustsecsettings.device_authentication_settings`

Optional:

- **sga_device_id** (String)
- **sga_device_password** (String)


<a id="nestedblock--item--trustsecsettings--device_configuration_deployment"></a>
### Nested Schema for `item.trustsecsettings.device_configuration_deployment`

Optional:

- **enable_mode_password** (String)
- **exec_mode_password** (String)
- **exec_mode_username** (String)
- **include_when_deploying_sgt_updates** (Boolean)


<a id="nestedblock--item--trustsecsettings--sga_notification_and_updates"></a>
### Nested Schema for `item.trustsecsettings.sga_notification_and_updates`

Optional:

- **coa_source_host** (String)
- **downlaod_environment_data_every_x_seconds** (Number)
- **downlaod_peer_authorization_policy_every_x_seconds** (Number)
- **download_sga_cllists_every_x_seconds** (Number)
- **other_sga_devices_to_trust_this_device** (Boolean)
- **re_authentication_every_x_seconds** (Number)
- **send_configuration_to_device** (Boolean)
- **send_configuration_to_device_using** (String)



<a id="nestedatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

