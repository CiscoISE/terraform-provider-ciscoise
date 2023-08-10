## 0.6.22-beta (August 09, 2023)
BUG FIXES:
* Resource 'ciscoise_network_access_policy_set' does not support nested children blocks #101
* No request_basic_pwd_auth parameter available in ciscoise_allowed_protocols resource in teap block #99 

## 0.6.21-beta (June 27, 2023)
BUG FIXES:
* ciscoise_internal_user import results in Error #88
* Ciscoise_network_device_group missing othername #94
* No parent_id parameter available in cisco_endpoint_group resource #95

## 0.6.20-beta (June 19, 2023)
BUG FIXES:
* Resource ciscoise_network_access_policy_set wrong state handling when using children blocks #92
* Ciscoise_network_device_group missing othername #94
* No parent_id parameter available in cisco_endpoint_group resource #95
* Getting all network access policy sets throwing error #90

## 0.6.19-beta (Jun 14, 2023)
BUG FIXES:
*  Resource 'ciscoise_network_access_conditions' does not support children blocks #91 
*  Getting all network access policy sets throwing error #90 
*   Resource ciscoise_network_access_policy_set wrong state handling when using children blocks #92

## 0.6.18-beta (Mar 03, 2023)
UPGRADE NOTES:
* The go version of the provider was updated to 1.20, this due to the new prerequisites of terraform gorealeser, in which it is detailed that you must have a GO version of 1.18 or higher. Here are [gorealeaser docs](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-release-publish).

## 0.6.17-beta (Feb 22, 2023)
BUG FIXES:
* Errors handled in new resources:
  - `resource_device_administration_authorization_rules_update`
  - `resource_network_access_authorization_rules_update`

## 0.6.16-beta (Feb 21, 2023)
FEATURES:
* **New Resource** `resource_device_administration_authorization_rules_update` issue #71.
* **New Resource** `resource_network_access_authorization_rules_update` issue #71.

BUG FIXES:
* Issue #83 =>  Update README as requested.

## 0.6.15-beta (Feb 15, 2023)
BUG FIXES:
* Issue #77 =>  flattenNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDItemRuleConditionChildren parameter `is_negate` parsed to string.

## 0.6.14-beta (Feb 10, 2023)
BUG FIXES:
* Issue #77 =>  ConditionReference support in children section of resource ciscoise_network_access_authorization_rules, ciscoise_network_access_policy_set, ciscoise_network_access_authentication_rules (idempotency problem) #77 [Fixed]

## 0.6.13-beta (Feb 01, 2023)
BUG FIXES:
* Issue #77 =>  ConditionReference support in children section of resource ciscoise_network_access_authorization_rules, ciscoise_network_access_policy_set, ciscoise_network_access_authentication_rules #77 [Fixed]

## 0.6.12-beta (Jan 30, 2023)
BUG FIXES:
* Issue #76 =>Resource ciscoise_network_access_authorization_rules wrong state handling [Fixed]
* Issue #71 =>  ciscoise_network_access_authorization_rules fixing issues [Fixed]
* Issue #77 =>  ConditionReference support in children section of resource ciscoise_network_access_authorization_rules [Fixed]

## 0.6.11-beta (Oct 27, 2022)
BUG FIXES:
* Issue #68 => Resource 'ciscoise_network_access_authentication_rules' does not support nested children blocks [Fixed]
* Issue #69 =>  Resource 'ciscoise_device_administration_policy_set' with single condition match returns multiple unsupported argument errors [Fixed]

IMPROVEMENTS:
* Documentation updated.
* Samples added.

## 0.6.10-beta (Oct 21, 2022)
BUG FIXES:
*  Issue #62 group_name => name change parameter name
*  Issue #65, bad parameters on resource
*  Issue #64 changing "othername" to "ndgtype"

## 0.6.9-beta (Oct 12, 2022)

BUG FIXES:
* Function `compareBoolean` fixed for no omit false in some times.
* `ciscoise_trusted_certificate` re-designed for use with `ciscoise_trusted_certificate_import`.(Problems with idempotency in API should be fixed after)

IMPROVEMENTS:
* Documentation updated.
* Samples added.

## 0.6.8-beta (Oct 03, 2022)

BUG FIXES:
* Issue #53, network_access_policy_set resource update doesn't work, Fixed
* Resource ciscoise_network_access_authentication_rules fails #54, Fixed
* Resource ciscoise_network_access_authorization_rules fails, Fixed

IMPROVEMENTS:
* Documentation updated.
* Samples added.

## 0.6.8-beta (Sep 30, 2022)

BUG FIXES:
* Parameters `name`, `id`, `description` added to resource `network_access_policy_set`

IMPROVEMENTS:
* Documentation updated.
* Samples added.

## 0.6.8-beta (Sep 21, 2022)

FEATURES:
* New provider configuration variable added:
  - **enable_auto_import** (String) Flag to enable or disable terraform automatic import (Automatic import means that when Terraform attempts to create the resource, it will perform a get operation if it founds a matching resource, it will perform an import of the resource it found, this is a similar operation to the terraform import command.) in resources, this is a configuration added to the provider, it uses the ISE_ENABLE_AUTO_IMPORT environment variable; `true` to enable it, defaults to `false`.

BUG FIXES:
* The provider overwrites the configuration to the one in the `tf file` even if it is updated outside of the provider.
* On `resource_sgt` `value` parameter adds new validation, it only allows numbers greatest or equal than `2` and lower or equal than `65519`.

## 0.6.5-beta (Jun 07, 2022)

FEATURES:
* New util file funcs for personas resources `personas_utils`
* **New Resource** `resource_personas_check_standalone`
* **New Resource** `resource_personas_export_certs`
* **New Resource** `resource_personas_promote_primary`
* **New Resource** `resource_personas_register_node`
* **New Resource** `resource_personas_update_roles_services`
## 0.6.4-beta (Jun 07, 2022)

BUG FIXES:

*  `data_source_network_access_policy_set.go` fixed.
*  `resource_network_access_policy_set.go` fixed.
*  `go.mod` updated.
*  `go.sum` updated.
## 0.6.3-beta (Jun 07, 2022)

BUG FIXES:
* `go.mod` and `go.sum` files updated.

## 0.6.2-beta (Jun 07, 2022)

BUG FIXES:
* Fixed ISE version 3.1.1 to 3.1_Patch_1 which is the correct version name.
## 0.6.1-beta (May 18, 2022)

IMPROVEMENTS:
* Removed the `remove_parameters` function call from resources
* The error 400 is now skipped when it defaults at the time of deleting `ciscoise_device_administration_authentication_rules`

## 0.6.0-beta (May 12, 2022)

IMPROVEMENTS:
* provider: Update ciscoisesdk from v1.1.3 to 1.1.4
* Fix func `remove_parameters` from `utils`

## 0.5.0-beta (May 06, 2022)
FEATURES:

* **New Resource:** `ciscoise_system_certificate_import`
* **New Resource:** `ciscoise_trusted_certificate_import`

## 0.4.0-beta (April 28, 2022)

FEATURES:

* **New Resource:** `ciscoise_backup_restore`
* **New Resource:** `ciscoise_backup_schedule_config_update`
* **New Resource:** `ciscoise_backup_schedule_config`
* **New Resource:** `ciscoise_ise_root_ca_regenerate`
* **New Resource:** `ciscoise_renew_certificate`
* **New Resource:** `ciscoise_bind_signed_certificate`
* **New Resource:** `ciscoise_selfsigned_certificate_generate`
* **New Resource:** `ciscoise_node_standalone_to_primary`
* **New Resource:** `ciscoise_node_secondary_to_primary`
* **New Resource:** `ciscoise_node_primary_to_standalone`
* **New Resource:** `ciscoise_node_deployment_sync`
* **New Resource:** `ciscoise_trustsec_sg_vn_mapping_bulk_create`
* **New Resource:** `ciscoise_trustsec_sg_vn_mapping_bulk_delete`
* **New Resource:** `ciscoise_trustsec_sg_vn_mapping_bulk_update`
* **New Resource:** `ciscoise_trustsec_vn_bulk_create`
* **New Resource:** `ciscoise_trustsec_vn_bulk_delete`
* **New Resource:** `ciscoise_trustsec_vn_bulk_update`
* **New Resource:** `ciscoise_trustsec_vn_vlan_mapping_bulk_delete`
* **New Resource:** `ciscoise_trustsec_vn_vlan_mapping_bulk_update`
* **New Resource:** `ciscoise_pxgrid_access_secret`
* **New Resource:** `ciscoise_pxgrid_account_activate`
* **New Resource:** `ciscoise_pxgrid_account_create`
* **New Resource:** `ciscoise_pxgrid_authorization`
* **New Resource:** `ciscoise_pxgrid_service_lookup`
* **New Resource:** `ciscoise_pxgrid_service_register`
* **New Resource:** `ciscoise_pxgrid_service_reregister`
* **New Resource:** `ciscoise_pxgrid_service_unregister`
* **New Resource:** `ciscoise_device_administration_authentication_reset_hitcount`
* **New Resource:** `ciscoise_device_administration_authorization_reset_hitcount`
* **New Resource:** `ciscoise_device_administration_local_exception_rules_reset_hitcount`
* **New Resource:** `ciscoise_device_administration_global_exception_rules_reset_hitcount`
* **New Resource:** `ciscoise_device_administration_policy_set_reset_hitcount`
* **New Resource:** `ciscoise_active_directory_join_domain`
* **New Resource:** `ciscoise_active_directory_join_domain_with_all_nodes`
* **New Resource:** `ciscoise_active_directory_leave_domain`
* **New Resource:** `ciscoise_active_directory_leave_domain_with_all_nodes`
* **New Resource:** `ciscoise_anc_endpoint_bulk_request`
* **New Resource:** `ciscoise_anc_policy_bulk_request`
* **New Resource:** `ciscoise_egress_matrix_cell_bulk_request`
* **New Resource:** `ciscoise_egress_matrix_cell_clear_all`
* **New Resource:** `ciscoise_egress_matrix_cell_clone`
* **New Resource:** `ciscoise_egress_matrix_cell_set_all_status`
* **New Resource:** `ciscoise_endpoint_deregister`
* **New Resource:** `ciscoise_endpoint_bulk_request`
* **New Resource:** `ciscoise_endpoint_register`
* **New Resource:** `ciscoise_endpoint_certificate`
* **New Resource:** `ciscoise_guest_user_approve`
* **New Resource:** `ciscoise_guest_user_bulk_request`
* **New Resource:** `ciscoise_guest_user_deny`
* **New Resource:** `ciscoise_guest_user_reinstate`
* **New Resource:** `ciscoise_guest_user_reinstate`
* **New Resource:** `ciscoise_guest_user_reset_password`
* **New Resource:** `ciscoise_guest_user_suspend`
* **New Resource:** `ciscoise_guest_user_suspend`
* **New Resource:** `ciscoise_network_device_bulk_request`
* **New Resource:** `ciscoise_px_grid_settings_auto_approve`
* **New Resource:** `ciscoise_sg_acl_bulk_request`
* **New Resource:** `ciscoise_sg_mapping_deploy`
* **New Resource:** `ciscoise_sg_mapping_bulk_request`
* **New Resource:** `ciscoise_sg_mapping_deploy_all`
* **New Resource:** `ciscoise_sg_mapping_group_deploy`
* **New Resource:** `ciscoise_sg_mapping_group_bulk_request`
* **New Resource:** `ciscoise_sg_mapping_group_deploy_all`
* **New Resource:** `ciscoise_sgt_bulk_request`
* **New Resource:** `ciscoise_sg_to_vn_to_vlan_bulk_request`
* **New Resource:** `ciscoise_sxp_connections_bulk_request`
* **New Resource:** `ciscoise_sxp_local_bindings_bulk_request`
* **New Resource:** `ciscoise_sxp_vpns_bulk_request`
* **New Resource:** `ciscoise_threat_vulnerabilities_clear`
* **New Resource:** `ciscoise_network_access_authentication_rules_reset_hitcount`
* **New Resource:** `ciscoise_network_access_authorization_rules_reset_hitcount`
* **New Resource:** `ciscoise_network_access_local_exception_rules_reset_hitcounts`
* **New Resource:** `ciscoise_network_access_global_exception_rules_reset_hitcount`
* **New Resource:** `ciscoise_network_access_policy_set_reset_hitcount`
* **New Resource:** `ciscoise_mnt_session_delete_all`


## 0.3.0-beta (March 17, 2022)

IMPROVEMENTS:

* provider: Update ciscoisesdk from v1.1.2 to 1.1.3
* provider: Update terraform-plugin-sdk/v2 from v2.7.1 to v2.10.1
* resource/ciscoise_guest_user: Add `UpdateGuestUserEmail`.
* resource/ciscoise_active_directory: Add `AddGroupsLoadGroupsFromDomain` to update context
* resource/ciscoise_guest_user: Update examples, samples and documentation
* resource/ciscoise_active_directory: Update examples, samples and documentation

FEATURES:

* **New Resource:** `ciscoise_pan_ha`
* **New Resource:** `ciscoise_licensing_registration`
* **New Resource:** `ciscoise_licensing_tier_state`
* **New Resource:** `ciscoise_px_grid_node`
* **New Resource:** `ciscoise_anc_endpoint`
* **New Resource:** `ciscoise_node_group_node`


## 0.2.0 (February 08, 2022)

NOTES:

* Data Sources of type 'action' have been removed. Removed data Sources of type 'action' have been classified as unsafe by the team. [GH-23]
* Next 0.3.0 version will transform some of them to resources. [GH-23]

BREAKING CHANGES:

* Data Source of type 'action' `ciscoise_backup_config` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_backup_cancel` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_backup_restore` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_backup_schedule_config_update` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_backup_schedule_config` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_csr_generate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_csr_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_csr_generate_intermediate_ca` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_ise_root_ca_regenerate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_renew_certificate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_bind_signed_certificate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_selfsigned_certificate_generate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_system_certificate_import` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trusted_certificate_import` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_group_node_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_group_node_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pan_ha_update` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_standalone_to_primary` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_secondary_to_primary` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_primary_to_standalone` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_node_deployment_sync` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_hotpatch_install` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_hotpatch_rollback` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_licensing_registration_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_licensing_smart_state_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_licensing_tier_state_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_patch_install` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_patch_rollback` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_update` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_update` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_update` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_access_secret` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_account_activate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_account_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_authorization` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_service_lookup` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_service_register` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_service_reregister` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_pxgrid_service_unregister` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_device_administration_authentication_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_device_administration_authorization_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_device_administration_local_exception_rules_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_device_administration_global_exception_rules_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_device_administration_policy_set_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_active_directory_add_groups` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_active_directory_join_domain` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_active_directory_join_domain_with_all_nodes` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_active_directory_leave_domain` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_active_directory_leave_domain_with_all_nodes` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_anc_endpoint_apply` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_anc_endpoint_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_anc_endpoint_clear` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_anc_policy_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_egress_matrix_cell_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_egress_matrix_cell_clear_all` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_egress_matrix_cell_clone` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_egress_matrix_cell_set_all_status` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_px_grid_node_delete` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_px_grid_node_approve` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_endpoint_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_endpoint_deregister` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_endpoint_certificate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_endpoint_release_rejected_endpoint` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_endpoint_register` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_type_email` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_type_sms` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_approve` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_change_sponsor_password` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_deny` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_email` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_reinstate` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_reset_password` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_device_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_sms` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_guest_user_suspend` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_px_grid_settings_auto_approve` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_acl_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_deploy` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_deploy_all` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_group_deploy` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_group_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_mapping_group_deploy_all` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sgt_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sg_to_vn_to_vlan_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_support_bundle` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sxp_connections_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sxp_local_bindings_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_sxp_vpns_bulk_request` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_system_certificate_create` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_threat_vulnerabilities_clear` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_access_authentication_rules_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_access_authorization_rules_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_access_local_exception_rules_reset_hitcounts` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_access_global_exception_rules_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_network_access_policy_set_reset_hitcount` has been removed [GH-23]
* Data Source of type 'action' `ciscoise_mnt_session_delete_all` has been removed [GH-23]

## 0.1.0 (February 02, 2022)

BUG FIXES:

* provider: Add single_request_timeout as a provider configuration argument. Changest at ciscoise/config. [GH-22]
* resource: Change "parameters" behavior from `Optional` to `Required`. [GH-22]

## 0.1.0-rc.4 (January 21, 2022)

BUG FIXES:

* data_source/ciscoise_pan_ha_update: Fix issue 18. Add missing Tf schema. [GH-19]
* data_source/ciscoise_pan_ha_update: Update documentation. [GH-19]

## 0.1.0-rc.3 (January 17, 2022)

BUG FIXES:

* provider: Fix resty Logger to be Tf compatible and always visible. Changes at ciscoise/logger. [GH-17]
* resource/ciscoise_node_services_profiler_probe_config: Change parameters (nmap, pxgrid, radius) type from TypeList to TypeString["", "true", "false". [GH-17]

IMPROVEMENTS:

* resource/ciscoise_node_services_profiler_probe_config: Update examples, samples and documentation [GH-17]
* resource/ciscoise_aci_settings:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_native_supplicant_profile:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_node_services_profiler_probe_config:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_node_services_sxp_interfaces:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_portal_global_setting:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_proxy_connection_settings:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_system_certificate:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_transport_gateway_settings:  Add 'update' call to CreateContext that are 'empty'. [GH-17]
* resource/ciscoise_trusted_certificate:  Add 'update' call to CreateContext that are 'empty'. [GH-17]

## 0.1.0-rc.2 (January 11, 2022)

IMPROVEMENTS:

* provider: Update examples and documentation. [GH-14]
* provider: Change resty Logger to be Tf compatible, changes at ciscoise/config [GH-14]
* resource/ciscoise_aci_settings: Change behaviour of "id" from Optional to Required. [GH-14]
* resource/ciscoise_system_certificate:  Add "id" to optional parameters. [GH-14]
* resource/ciscoise_system_certificate:  Add "host_name" to Tf id. [GH-14]
* resource/ciscoise_portal_global_setting: Change behaviour of "id" from Optional to Required. [GH-14]
* resource/ciscoise_transport_gateway_settings:  Change "url" to be the Tf id. [GH-14]
* resource/ciscoise_node_deployment: Add "fqdn" to Tf id. [GH-14]
* resource/ciscoise_proxy_connection_settings:  Add "user_name" to Tf id. [GH-14]
* resource/ciscoise_repository: Fix update func selection method. [GH-14]
* resource/ciscoise_node_group: Fix update func selection method. [GH-14]
* resource/ciscoise_node_services_profiler_probe_config:  Change "hostname" to be the Tf id. [GH-14]
* resource/ciscoise_byod_portal: Update documented import id format. [GH-14]
* resource/ciscoise_device_administration_network_conditions: documented import id format.Update  [GH-14]
* resource/ciscoise_device_administration_policy_set: documented import id format.Update  [GH-14]
* resource/ciscoise_device_administration_time_date_conditions: documented import id format.Update  [GH-14]
* resource/ciscoise_downloadable_acl: Update documented import id format. [GH-14]
* resource/ciscoise_egress_matrix_cell: Update documented import id format. [GH-14]
* resource/ciscoise_filter_policy: Update documented import id format. [GH-14]
* resource/ciscoise_guest_ssid: Update documented import id format. [GH-14]
* resource/ciscoise_guest_type: Update documented import id format. [GH-14]
* resource/ciscoise_hotspot_portal: Update documented import id format. [GH-14]
* resource/ciscoise_node_deployment: Update documented import id format. [GH-14]
* resource/ciscoise_proxy_connection_settings: Update documented import id format. [GH-14]
* resource/ciscoise_repository: Update documented import id format. [GH-14]
* resource/ciscoise_self_registered_portal: Update documented import id format. [GH-14]
* resource/ciscoise_sg_acl: Update documented import id format. [GH-14]
* resource/ciscoise_sg_mapping: Update documented import id format. [GH-14]
* resource/ciscoise_sg_mapping_group: Update documented import id format. [GH-14]
* resource/ciscoise_sg_to_vn_to_vlan: Update documented import id format. [GH-14]
* resource/ciscoise_sgt: Update documented import id format. [GH-14]
* resource/ciscoise_sponsor_group: Update documented import id format. [GH-14]
* resource/ciscoise_sponsor_portal: Update documented import id format. [GH-14]
* resource/ciscoise_sponsored_guest_portal: Update documented import id format. [GH-14]
* resource/ciscoise_sxp_vpns: Update documented import id format. [GH-14]
* resource/ciscoise_system_certificate: Update documented import id format. [GH-14]
* resource/ciscoise_transport_gateway_settings: Update documented import id format. [GH-14]
* resource/ciscoise_trusted_certificate: Update documented import id format. [GH-14]
* resource/ciscoise_trustsec_nbar_app: Update documented import id format. [GH-14]
* resource/ciscoise_trustsec_sg_vn_mapping: Update documented import id format. [GH-14]
* resource/ciscoise_trustsec_vn: Update documented import id format. [GH-14]
* resource/ciscoise_trustsec_vn_vlan_mapping: Update documented import id format. [GH-14]

BUG FIXES:

* data_source/ciscoise_deployment: Fix schemas (resolves expected string got []interface{}). [GH-14]
* data_source/ciscoise_node_services_profiler_probe_config: Fix schemas (resolves expected string got []interface{}). [GH-14]
* data_source/ciscoise_sg_acl: Fix schemas (resolves expected string got []interface{}). [GH-14]
* data_source/ciscoise_tasks: Fix schemas (resolves expected string got []interface{}). [GH-14]
* resource/ciscoise_node_services_profiler_probe_config: Fix schemas (resolves expected string got []interface{}). [GH-14]

## 0.1.0-rc.1 (December 24, 2021)

IMPROVEMENTS:

* resource: Change Read not found behaviour to setId(""). [GH-11]
* resource: Add log to indicate resource context execution. [GH-11]

## 0.1.0-rc (December 23, 2021)

BUG FIXES:

* data_source/ciscoise_system_certificate: Fix read only schema below `items` parameter. [GH-9]
* resource/ciscoise_self_registered_portal: Fix resource Id value for create operation. [GH-9]
* data_source/ciscoise_device_administration_authentication_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_device_administration_authorization_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_device_administration_local_exception_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_network_access_authentication_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_network_access_authorization_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_network_access_dictionary_attribute: Fix method selection logic. [GH-9]
* data_source/ciscoise_network_access_local_exception_rules: Fix method selection logic. [GH-9]
* data_source/ciscoise_system_certificate: Fix method selection logic. [GH-9]
* resource/ciscoise_device_administration_authentication_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_device_administration_authorization_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_device_administration_local_exception_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_network_access_authentication_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_network_access_authorization_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_network_access_local_exception_rules: Fix method selection logic. [GH-9]
* resource/ciscoise_system_certificate: Fix method selection logic. [GH-9]

IMPROVEMENTS:

* provider: Update examples and documentation [GH-9]
* resource: Remove number from some logs to avoid confusion. [GH-9]
* resource: Updated logs to use %v instead of %q. [GH-9]
* resource: Removed _ param in `for range` code. [GH-9]
* data_source: Remove number from some logs to avoid confusion. [GH-9]
* data_source: Updated logs to use %v instead of %q. [GH-9]
* data_source: Removed _ param in `for range` code. [GH-9]

## 0.0.3-beta (December 22, 2021)

BREAKING CHANGES:

* Data Source `ciscoise_mnt_athentication_status` has been removed [GH-8]
* Data Source `ciscoise_node_promotion` has been removed [GH-8]
* Data Source `ciscoise_node_replication_status` has been removed [GH-8]
* Data Source `ciscoise_node_sync` has been removed [GH-8]
* Resource `resource_ciscoise_pan_ha` has been removed [GH-8]

FEATURES:

* **New Data Source:** `ciscoise_hotpatch` [GH-8]
* **New Data Source:** `ciscoise_hotpatch_install` [GH-8]
* **New Data Source:** `ciscoise_hotpatch_rollback` [GH-8]
* **New Data Source:** `ciscoise_licensing_connection_type` [GH-8]
* **New Data Source:** `ciscoise_licensing_eval_license` [GH-8]
* **New Data Source:** `ciscoise_licensing_feature_to_tier_mapping` [GH-8]
* **New Data Source:** `ciscoise_licensing_registration` [GH-8]
* **New Data Source:** `ciscoise_licensing_registration_create` [GH-8]
* **New Data Source:** `ciscoise_licensing_smart_state` [GH-8]
* **New Data Source:** `ciscoise_licensing_smart_state_create` [GH-8]
* **New Data Source:** `ciscoise_licensing_tier_state` [GH-8]
* **New Data Source:** `ciscoise_licensing_tier_state_create` [GH-8]
* **New Data Source:** `ciscoise_mnt_authentication_status` [GH-8]
* **New Data Source:** `ciscoise_node_deployment_sync` [GH-8]
* **New Data Source:** `ciscoise_node_group_node` [GH-8]
* **New Data Source:** `ciscoise_node_group_node_create` [GH-8]
* **New Data Source:** `ciscoise_node_group_node_delete` [GH-8]
* **New Data Source:** `ciscoise_node_primary_to_standalone` [GH-8]
* **New Data Source:** `ciscoise_node_secondary_to_primary` [GH-8]
* **New Data Source:** `ciscoise_node_services_interfaces` [GH-8]
* **New Data Source:** `ciscoise_node_services_profiler_probe_config` [GH-8]
* **New Data Source:** `ciscoise_node_services_sxp_interfaces` [GH-8]
* **New Data Source:** `ciscoise_node_standalone_to_primary` [GH-8]
* **New Data Source:** `ciscoise_pan_ha_update` [GH-8]
* **New Data Source:** `ciscoise_patch` [GH-8]
* **New Data Source:** `ciscoise_patch_install` [GH-8]
* **New Data Source:** `ciscoise_patch_rollback` [GH-8]
* **New Data Source:** `ciscoise_proxy_connection_settings` [GH-8]
* **New Data Source:** `ciscoise_selfsigned_certificate_generate` [GH-8]
* **New Data Source:** `ciscoise_transport_gateway_settings` [GH-8]
* **New Data Source:** `ciscoise_trustsec_nbar_app` [GH-8]
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping` [GH-8]
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_create` [GH-8]
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_delete` [GH-8]
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_update` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_bulk_create` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_bulk_delete` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_bulk_update` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_create` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_delete` [GH-8]
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_update` [GH-8]
* **New Resource:** `ciscoise_node_services_profiler_probe_config` [GH-8]
* **New Resource:** `ciscoise_node_services_sxp_interfaces` [GH-8]
* **New Resource:** `ciscoise_proxy_connection_settings` [GH-8]
* **New Resource:** `ciscoise_transport_gateway_settings` [GH-8]
* **New Resource:** `ciscoise_trustsec_nbar_app` [GH-8]
* **New Resource:** `ciscoise_trustsec_sg_vn_mapping` [GH-8]
* **New Resource:** `ciscoise_trustsec_vn` [GH-8]
* **New Resource:** `ciscoise_trustsec_vn_vlan_mapping` [GH-8]

IMPROVEMENTS:

* provider: Update examples and documentation [GH-8]
* resource: Separated `Computed` and `Optional`/`Required` parameters. `Computed` parameters still reside inside `item` schema, while `Optional`/`Required` now reside inside `parameters` schema. [GH-8]


## 0.0.2-beta (September 28, 2021)

IMPROVEMENTS:

* provider: Add provider description at docs/index.md

## 0.0.1-beta (September 28, 2021)

* Initial Release [GH-6]