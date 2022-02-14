## 0.3.0-beta (Unreleased)

IMPROVEMENTS:

* provider: Update ciscoisesdk from v1.1.2 to 1.1.3
* provider: Update terraform-plugin-sdk/v2 from v2.7.1 to v2.10.1
* resource/ciscoise_guest_user: Add `ChangeSponsorPassword`, `UpdateGuestUserEmail` and `UpdateGuestUserSms`.
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
* **New Resource:** `ciscoise_hotpatch`
* **New Resource:** `ciscoise_patch`

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