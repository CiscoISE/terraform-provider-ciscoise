---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_authorization_profile Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_authorization_profile (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **name** (String)
- **page** (Number)
- **size** (Number)

### Read-Only

- **item_id** (List of Object) (see [below for nested schema](#nestedatt--item_id))
- **item_name** (List of Object) (see [below for nested schema](#nestedatt--item_name))
- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item_id"></a>
### Nested Schema for `item_id`

Read-Only:

- **access_type** (String)
- **acl** (String)
- **advanced_attributes** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--advanced_attributes))
- **agentless_posture** (Boolean)
- **airespace_acl** (String)
- **airespace_ipv6_acl** (String)
- **asa_vpn** (String)
- **authz_profile_type** (String)
- **auto_smart_port** (String)
- **avc_profile** (String)
- **dacl_name** (String)
- **description** (String)
- **easywired_session_candidate** (Boolean)
- **id** (String)
- **interface_template** (String)
- **ipv6_acl_filter** (String)
- **ipv6_dacl_name** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--link))
- **mac_sec_policy** (String)
- **name** (String)
- **neat** (Boolean)
- **profile_name** (String)
- **reauth** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--reauth))
- **service_template** (Boolean)
- **track_movement** (Boolean)
- **vlan** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--vlan))
- **voice_domain_permission** (Boolean)
- **web_auth** (Boolean)
- **web_redirection** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--web_redirection))

<a id="nestedobjatt--item_id--advanced_attributes"></a>
### Nested Schema for `item_id.advanced_attributes`

Read-Only:

- **left_hand_side_dictionary_attribue** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--advanced_attributes--left_hand_side_dictionary_attribue))
- **right_hand_side_attribue_value** (List of Object) (see [below for nested schema](#nestedobjatt--item_id--advanced_attributes--right_hand_side_attribue_value))

<a id="nestedobjatt--item_id--advanced_attributes--left_hand_side_dictionary_attribue"></a>
### Nested Schema for `item_id.advanced_attributes.left_hand_side_dictionary_attribue`

Read-Only:

- **advanced_attribute_value_type** (String)
- **attribute_name** (String)
- **dictionary_name** (String)
- **value** (String)


<a id="nestedobjatt--item_id--advanced_attributes--right_hand_side_attribue_value"></a>
### Nested Schema for `item_id.advanced_attributes.right_hand_side_attribue_value`

Read-Only:

- **advanced_attribute_value_type** (String)
- **attribute_name** (String)
- **dictionary_name** (String)
- **value** (String)



<a id="nestedobjatt--item_id--link"></a>
### Nested Schema for `item_id.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)


<a id="nestedobjatt--item_id--reauth"></a>
### Nested Schema for `item_id.reauth`

Read-Only:

- **connectivity** (String)
- **timer** (Number)


<a id="nestedobjatt--item_id--vlan"></a>
### Nested Schema for `item_id.vlan`

Read-Only:

- **name_id** (String)
- **tag_id** (Number)


<a id="nestedobjatt--item_id--web_redirection"></a>
### Nested Schema for `item_id.web_redirection`

Read-Only:

- **acl** (String)
- **display_certificates_renewal_messages** (Boolean)
- **portal_name** (String)
- **static_iphost_name_fqd_n** (String)
- **web_redirection_type** (String)



<a id="nestedatt--item_name"></a>
### Nested Schema for `item_name`

Read-Only:

- **access_type** (String)
- **acl** (String)
- **advanced_attributes** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--advanced_attributes))
- **agentless_posture** (Boolean)
- **airespace_acl** (String)
- **airespace_ipv6_acl** (String)
- **asa_vpn** (String)
- **authz_profile_type** (String)
- **auto_smart_port** (String)
- **avc_profile** (String)
- **dacl_name** (String)
- **description** (String)
- **easywired_session_candidate** (Boolean)
- **id** (String)
- **interface_template** (String)
- **ipv6_acl_filter** (String)
- **ipv6_dacl_name** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--link))
- **mac_sec_policy** (String)
- **name** (String)
- **neat** (Boolean)
- **profile_name** (String)
- **reauth** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--reauth))
- **service_template** (Boolean)
- **track_movement** (Boolean)
- **vlan** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--vlan))
- **voice_domain_permission** (Boolean)
- **web_auth** (Boolean)
- **web_redirection** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--web_redirection))

<a id="nestedobjatt--item_name--advanced_attributes"></a>
### Nested Schema for `item_name.advanced_attributes`

Read-Only:

- **left_hand_side_dictionary_attribue** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--advanced_attributes--left_hand_side_dictionary_attribue))
- **right_hand_side_attribue_value** (List of Object) (see [below for nested schema](#nestedobjatt--item_name--advanced_attributes--right_hand_side_attribue_value))

<a id="nestedobjatt--item_name--advanced_attributes--left_hand_side_dictionary_attribue"></a>
### Nested Schema for `item_name.advanced_attributes.left_hand_side_dictionary_attribue`

Read-Only:

- **advanced_attribute_value_type** (String)
- **attribute_name** (String)
- **dictionary_name** (String)
- **value** (String)


<a id="nestedobjatt--item_name--advanced_attributes--right_hand_side_attribue_value"></a>
### Nested Schema for `item_name.advanced_attributes.right_hand_side_attribue_value`

Read-Only:

- **advanced_attribute_value_type** (String)
- **attribute_name** (String)
- **dictionary_name** (String)
- **value** (String)



<a id="nestedobjatt--item_name--link"></a>
### Nested Schema for `item_name.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)


<a id="nestedobjatt--item_name--reauth"></a>
### Nested Schema for `item_name.reauth`

Read-Only:

- **connectivity** (String)
- **timer** (Number)


<a id="nestedobjatt--item_name--vlan"></a>
### Nested Schema for `item_name.vlan`

Read-Only:

- **name_id** (String)
- **tag_id** (Number)


<a id="nestedobjatt--item_name--web_redirection"></a>
### Nested Schema for `item_name.web_redirection`

Read-Only:

- **acl** (String)
- **display_certificates_renewal_messages** (Boolean)
- **portal_name** (String)
- **static_iphost_name_fqd_n** (String)
- **web_redirection_type** (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **description** (String)
- **id** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--items--link))
- **name** (String)

<a id="nestedobjatt--items--link"></a>
### Nested Schema for `items.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

