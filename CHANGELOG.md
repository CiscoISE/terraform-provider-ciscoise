## 0.2.0 (February 08, 2022)

NOTES:
* Data Sources of type 'action' have been removed. Removed data Sources of type 'action' have been classified as unsafe by the team.
* Next 0.3.0 version will transform some of them to resources.

BREAKING CHANGES:

* Data Source of type 'action' `ciscoise_backup_config` has been removed
* Data Source of type 'action' `ciscoise_backup_cancel` has been removed
* Data Source of type 'action' `ciscoise_backup_restore` has been removed
* Data Source of type 'action' `ciscoise_backup_schedule_config_update` has been removed
* Data Source of type 'action' `ciscoise_backup_schedule_config` has been removed
* Data Source of type 'action' `ciscoise_csr_generate` has been removed
* Data Source of type 'action' `ciscoise_csr_delete` has been removed
* Data Source of type 'action' `ciscoise_csr_generate_intermediate_ca` has been removed
* Data Source of type 'action' `ciscoise_ise_root_ca_regenerate` has been removed
* Data Source of type 'action' `ciscoise_renew_certificate` has been removed
* Data Source of type 'action' `ciscoise_bind_signed_certificate` has been removed
* Data Source of type 'action' `ciscoise_selfsigned_certificate_generate` has been removed
* Data Source of type 'action' `ciscoise_system_certificate_import` has been removed
* Data Source of type 'action' `ciscoise_trusted_certificate_import` has been removed
* Data Source of type 'action' `ciscoise_node_group_node_create` has been removed
* Data Source of type 'action' `ciscoise_node_group_node_delete` has been removed
* Data Source of type 'action' `ciscoise_pan_ha_update` has been removed
* Data Source of type 'action' `ciscoise_node_standalone_to_primary` has been removed
* Data Source of type 'action' `ciscoise_node_secondary_to_primary` has been removed
* Data Source of type 'action' `ciscoise_node_primary_to_standalone` has been removed
* Data Source of type 'action' `ciscoise_node_deployment_sync` has been removed
* Data Source of type 'action' `ciscoise_hotpatch_install` has been removed
* Data Source of type 'action' `ciscoise_hotpatch_rollback` has been removed
* Data Source of type 'action' `ciscoise_licensing_registration_create` has been removed
* Data Source of type 'action' `ciscoise_licensing_smart_state_create` has been removed
* Data Source of type 'action' `ciscoise_licensing_tier_state_create` has been removed
* Data Source of type 'action' `ciscoise_patch_install` has been removed
* Data Source of type 'action' `ciscoise_patch_rollback` has been removed
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_create` has been removed
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_delete` has been removed
* Data Source of type 'action' `ciscoise_trustsec_sg_vn_mapping_bulk_update` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_create` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_delete` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_bulk_update` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_create` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_delete` has been removed
* Data Source of type 'action' `ciscoise_trustsec_vn_vlan_mapping_bulk_update` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_access_secret` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_account_activate` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_account_create` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_authorization` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_service_lookup` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_service_register` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_service_reregister` has been removed
* Data Source of type 'action' `ciscoise_pxgrid_service_unregister` has been removed
* Data Source of type 'action' `ciscoise_device_administration_authentication_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_device_administration_authorization_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_device_administration_local_exception_rules_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_device_administration_global_exception_rules_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_device_administration_policy_set_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_active_directory_add_groups` has been removed
* Data Source of type 'action' `ciscoise_active_directory_join_domain` has been removed
* Data Source of type 'action' `ciscoise_active_directory_join_domain_with_all_nodes` has been removed
* Data Source of type 'action' `ciscoise_active_directory_leave_domain` has been removed
* Data Source of type 'action' `ciscoise_active_directory_leave_domain_with_all_nodes` has been removed
* Data Source of type 'action' `ciscoise_anc_endpoint_apply` has been removed
* Data Source of type 'action' `ciscoise_anc_endpoint_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_anc_endpoint_clear` has been removed
* Data Source of type 'action' `ciscoise_anc_policy_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_egress_matrix_cell_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_egress_matrix_cell_clear_all` has been removed
* Data Source of type 'action' `ciscoise_egress_matrix_cell_clone` has been removed
* Data Source of type 'action' `ciscoise_egress_matrix_cell_set_all_status` has been removed
* Data Source of type 'action' `ciscoise_px_grid_node_delete` has been removed
* Data Source of type 'action' `ciscoise_px_grid_node_approve` has been removed
* Data Source of type 'action' `ciscoise_endpoint_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_endpoint_deregister` has been removed
* Data Source of type 'action' `ciscoise_endpoint_certificate` has been removed
* Data Source of type 'action' `ciscoise_endpoint_release_rejected_endpoint` has been removed
* Data Source of type 'action' `ciscoise_endpoint_register` has been removed
* Data Source of type 'action' `ciscoise_guest_type_email` has been removed
* Data Source of type 'action' `ciscoise_guest_user_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_guest_type_sms` has been removed
* Data Source of type 'action' `ciscoise_guest_user_approve` has been removed
* Data Source of type 'action' `ciscoise_guest_user_change_sponsor_password` has been removed
* Data Source of type 'action' `ciscoise_guest_user_deny` has been removed
* Data Source of type 'action' `ciscoise_guest_user_email` has been removed
* Data Source of type 'action' `ciscoise_guest_user_reinstate` has been removed
* Data Source of type 'action' `ciscoise_guest_user_reset_password` has been removed
* Data Source of type 'action' `ciscoise_network_device_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_guest_user_sms` has been removed
* Data Source of type 'action' `ciscoise_guest_user_suspend` has been removed
* Data Source of type 'action' `ciscoise_px_grid_settings_auto_approve` has been removed
* Data Source of type 'action' `ciscoise_sg_acl_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_deploy` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_deploy_all` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_group_deploy` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_group_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sg_mapping_group_deploy_all` has been removed
* Data Source of type 'action' `ciscoise_sgt_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sg_to_vn_to_vlan_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_support_bundle` has been removed
* Data Source of type 'action' `ciscoise_sxp_connections_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sxp_local_bindings_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_sxp_vpns_bulk_request` has been removed
* Data Source of type 'action' `ciscoise_system_certificate_create` has been removed
* Data Source of type 'action' `ciscoise_threat_vulnerabilities_clear` has been removed
* Data Source of type 'action' `ciscoise_network_access_authentication_rules_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_network_access_authorization_rules_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_network_access_local_exception_rules_reset_hitcounts` has been removed
* Data Source of type 'action' `ciscoise_network_access_global_exception_rules_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_network_access_policy_set_reset_hitcount` has been removed
* Data Source of type 'action' `ciscoise_mnt_session_delete_all` has been removed

## 0.1.0 (February 02, 2022)

BUG FIXES:
* ciscoise/resource*: Change "parameters" behavior from `Optional` to `Required`.
* ciscoise/config*: Add single_request_timeout as a provider configuration argument.

## 0.1.0-rc.4 (January 21, 2022)

BUG FIXES:
* ciscoise/data_source_pan_ha_update: Fix issue 18. Add missing Tf schema.
* ciscoise/data_source_pan_ha_update: Update documentation.

## 0.1.0-rc.3 (January 17, 2022)

BUG FIXES:
* ciscoise/logger: Fix resty Logger to be Tf compatible and always visible.
* ciscoise/resource_node_services_profiler_probe_config: Change parameters (nmap, pxgrid, radius) type from TypeList to TypeString["", "true", "false".

IMPROVEMENTS:
* ciscoise/resource_node_services_profiler_probe_config: Update examples, samples and documentation
* ciscoise/resource_aci_settings:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_native_supplicant_profile:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_node_services_profiler_probe_config:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_node_services_sxp_interfaces:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_portal_global_setting:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_proxy_connection_settings:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_system_certificate:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_transport_gateway_settings:  Add 'update' call to CreateContext that are 'empty'.
* ciscoise/resource_trusted_certificate:  Add 'update' call to CreateContext that are 'empty'.

## 0.1.0-rc.2 (January 11, 2022)

IMPROVEMENTS:
* ciscoise/config: Change resty Logger to be Tf compatible.
* ciscoise/resource_aci_settings: Change behaviour of "id" from Optional to Required.
* ciscoise/resource_system_certificate:  Add "id" to optional parameters.
* ciscoise/resource_system_certificate:  Add "host_name" to Tf id.
* ciscoise/resource_portal_global_setting: Change behaviour of "id" from Optional to Required.
* ciscoise/resource_transport_gateway_settings:  Change "url" to be the Tf id.
* ciscoise/resource_node_deployment: Add "fqdn" to Tf id.
* ciscoise/resource_proxy_connection_settings:  Add "user_name" to Tf id.
* ciscoise/resource_repository: Fix update func selection method.
* ciscoise/resource_node_group: Fix update func selection method.
* ciscoise/resource_node_services_profiler_probe_config:  Change "hostname" to be the Tf id.
* resource/ciscoise_byod_portal: Update documented import id format.
* resource/ciscoise_device_administration_network_conditions: documented import id format.Update 
* resource/ciscoise_device_administration_policy_set: documented import id format.Update 
* resource/ciscoise_device_administration_time_date_conditions: documented import id format.Update 
* resource/ciscoise_downloadable_acl: Update documented import id format.
* resource/ciscoise_egress_matrix_cell: Update documented import id format.
* resource/ciscoise_filter_policy: Update documented import id format.
* resource/ciscoise_guest_ssid: Update documented import id format.
* resource/ciscoise_guest_type: Update documented import id format.
* resource/ciscoise_hotspot_portal: Update documented import id format.
* resource/ciscoise_node_deployment: Update documented import id format.
* resource/ciscoise_proxy_connection_settings: Update documented import id format.
* resource/ciscoise_repository: Update documented import id format.
* resource/ciscoise_self_registered_portal: Update documented import id format.
* resource/ciscoise_sg_acl: Update documented import id format.
* resource/ciscoise_sg_mapping: Update documented import id format.
* resource/ciscoise_sg_mapping_group: Update documented import id format.
* resource/ciscoise_sg_to_vn_to_vlan: Update documented import id format.
* resource/ciscoise_sgt: Update documented import id format.
* resource/ciscoise_sponsor_group: Update documented import id format.
* resource/ciscoise_sponsor_portal: Update documented import id format.
* resource/ciscoise_sponsored_guest_portal: Update documented import id format.
* resource/ciscoise_sxp_vpns: Update documented import id format.
* resource/ciscoise_system_certificate: Update documented import id format.
* resource/ciscoise_transport_gateway_settings: Update documented import id format.
* resource/ciscoise_trusted_certificate: Update documented import id format.
* resource/ciscoise_trustsec_nbar_app: Update documented import id format.
* resource/ciscoise_trustsec_sg_vn_mapping: Update documented import id format.
* resource/ciscoise_trustsec_vn: Update documented import id format.
* resource/ciscoise_trustsec_vn_vlan_mapping: Update documented import id format.
* Update examples and documentation.

BUG FIXES:
* ciscoise/data_source_deployment: Fix schemas (resolves expected string got []interface{}).
ciscoise/* data_source_node_services_profiler_probe_config: Fix schemas (resolves expected string got []interface{}).
* ciscoise/data_source_sg_acl: Fix schemas (resolves expected string got []interface{}).
* ciscoise/data_source_tasks: Fix schemas (resolves expected string got []interface{}).
* ciscoise/resource_node_services_profiler_probe_config: Fix schemas (resolves expected string got []interface{}).

## 0.1.0-rc.1 (December 24, 2021)

IMPROVEMENTS:
* ciscoise/resource_*: Change Read not found behaviour to setId("").
* ciscoise/resource_*: Add log to indicate resource context execution.

## 0.1.0-rc (December 23, 2021)

BUG FIXES:
* data_source/system_certificate: Fix read only schema below `items` parameter.
* resource/self_registered_portal: Fix resource Id value for create operation.
* data_source/device_administration_authentication_rules: Fix method selection logic.
* data_source/device_administration_authorization_rules: Fix method selection logic.
* data_source/device_administration_local_exception_rules: Fix method selection logic.
* data_source/network_access_authentication_rules: Fix method selection logic.
* data_source/network_access_authorization_rules: Fix method selection logic.
* data_source/network_access_dictionary_attribute: Fix method selection logic.
* data_source/network_access_local_exception_rules: Fix method selection logic.
* data_source/system_certificate: Fix method selection logic.
* resource/device_administration_authentication_rules: Fix method selection logic.
* resource/device_administration_authorization_rules: Fix method selection logic.
* resource/device_administration_local_exception_rules: Fix method selection logic.
* resource/network_access_authentication_rules: Fix method selection logic.
* resource/network_access_authorization_rules: Fix method selection logic.
* resource/network_access_local_exception_rules: Fix method selection logic.
* resource/system_certificate: Fix method selection logic.

IMPROVEMENTS:
* ciscoise/resource_*: Remove number from some logs to avoid confusion.
* ciscoise/resource_*: Updated logs to use %v instead of %q.
* ciscoise/resource_*: Removed _ param in `for range` code.
* ciscoise/data_source_*: Remove number from some logs to avoid confusion.
* ciscoise/data_source_*: Updated logs to use %v instead of %q.
* ciscoise/data_source_*: Removed _ param in `for range` code.
* Update examples and documentation

## 0.0.3-beta (December 22, 2021)

BREAKING CHANGES:

* Data Source `ciscoise_mnt_athentication_status` has been removed
* Data Source `ciscoise_node_promotion` has been removed
* Data Source `ciscoise_node_replication_status` has been removed
* Data Source `ciscoise_node_sync` has been removed
* Resource `resource_ciscoise_pan_ha` has been removed

FEATURES:

* **New Data Source:** `ciscoise_hotpatch`
* **New Data Source:** `ciscoise_hotpatch_install`
* **New Data Source:** `ciscoise_hotpatch_rollback`
* **New Data Source:** `ciscoise_licensing_connection_type`
* **New Data Source:** `ciscoise_licensing_eval_license`
* **New Data Source:** `ciscoise_licensing_feature_to_tier_mapping`
* **New Data Source:** `ciscoise_licensing_registration`
* **New Data Source:** `ciscoise_licensing_registration_create`
* **New Data Source:** `ciscoise_licensing_smart_state`
* **New Data Source:** `ciscoise_licensing_smart_state_create`
* **New Data Source:** `ciscoise_licensing_tier_state`
* **New Data Source:** `ciscoise_licensing_tier_state_create`
* **New Data Source:** `ciscoise_mnt_authentication_status`
* **New Data Source:** `ciscoise_node_deployment_sync`
* **New Data Source:** `ciscoise_node_group_node`
* **New Data Source:** `ciscoise_node_group_node_create`
* **New Data Source:** `ciscoise_node_group_node_delete`
* **New Data Source:** `ciscoise_node_primary_to_standalone`
* **New Data Source:** `ciscoise_node_secondary_to_primary`
* **New Data Source:** `ciscoise_node_services_interfaces`
* **New Data Source:** `ciscoise_node_services_profiler_probe_config`
* **New Data Source:** `ciscoise_node_services_sxp_interfaces`
* **New Data Source:** `ciscoise_node_standalone_to_primary`
* **New Data Source:** `ciscoise_pan_ha_update`
* **New Data Source:** `ciscoise_patch`
* **New Data Source:** `ciscoise_patch_install`
* **New Data Source:** `ciscoise_patch_rollback`
* **New Data Source:** `ciscoise_proxy_connection_settings`
* **New Data Source:** `ciscoise_selfsigned_certificate_generate`
* **New Data Source:** `ciscoise_transport_gateway_settings`
* **New Data Source:** `ciscoise_trustsec_nbar_app`
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping`
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_create`
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_delete`
* **New Data Source:** `ciscoise_trustsec_sg_vn_mapping_bulk_update`
* **New Data Source:** `ciscoise_trustsec_vn`
* **New Data Source:** `ciscoise_trustsec_vn_bulk_create`
* **New Data Source:** `ciscoise_trustsec_vn_bulk_delete`
* **New Data Source:** `ciscoise_trustsec_vn_bulk_update`
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping`
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_create`
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_delete`
* **New Data Source:** `ciscoise_trustsec_vn_vlan_mapping_bulk_update`
* **New Resource:** `ciscoise_node_services_profiler_probe_config`
* **New Resource:** `ciscoise_node_services_sxp_interfaces`
* **New Resource:** `ciscoise_proxy_connection_settings`
* **New Resource:** `ciscoise_transport_gateway_settings`
* **New Resource:** `ciscoise_trustsec_nbar_app`
* **New Resource:** `ciscoise_trustsec_sg_vn_mapping`
* **New Resource:** `ciscoise_trustsec_vn`
* **New Resource:** `ciscoise_trustsec_vn_vlan_mapping`

IMPROVEMENTS:
* ciscoise/resource_*: Separated `Computed` and `Optional`/`Required` parameters. `Computed` parameters still reside inside `item` schema, while `Optional`/`Required` now reside inside `parameters` schema.
* Update examples and documentation


## 0.0.2-beta (September 28, 2021)

IMPROVEMENTS:

* docs/index.md: Add provider description

## 0.0.1-beta (September 28, 2021)

* Initial Release