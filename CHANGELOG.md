## 0.1.0-rc.1 (December 24, 2021)

IMPROVEMENTS:
* ciscoise/resource_*: Change Read not found behaviour to setId("")
* ciscoise/resource_*: Add log to indicate resource context execution

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