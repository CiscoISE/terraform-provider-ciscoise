package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAllowedProtocols() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AllowedProtocols.

- This data source allows the client to get an allowed protocol by name.

- This data source allows the client to get an allowed protocol by ID.

- This data source allows the client to get all the allowed protocols.
`,

		ReadContext: dataSourceAllowedProtocolsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_chap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_fast": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_md5": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_tls": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_ttls": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_leap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_ms_chap_v1": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_ms_chap_v2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_pap_ascii": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_peap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_preferred_eap_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_teap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_weak_ciphers_for_eap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"eap_fast": &schema.Schema{
							Description: `The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_fast_eap_gtc": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_gtc_pwd_change": &schema.Schema{
										Description: `The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_eap_fast_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"allow_eap_fast_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_dont_use_pacs_accept_client_cert": &schema.Schema{
										Description: `The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_dont_use_pacs_allow_machine_authentication": &schema.Schema{
										Description: `The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_accept_client_cert": &schema.Schema{
										Description: `The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_anonym_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_authen_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_machine_authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true,
otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true,
otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning
is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_stateless_session_resume": &schema.Schema{
										Description: `The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_use_proactive_pac_update_precentage": &schema.Schema{
										Description: `The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"eap_tls": &schema.Schema{
							Description: `The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_tls_enable_stateless_session_resume": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_tls_session_ticket_precentage": &schema.Schema{
										Description: `The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_tls_session_ticket_ttl": &schema.Schema{
										Description: `Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"eap_tls_session_ticket_ttl_units": &schema.Schema{
										Description: `Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true,
otherwise it must be ignored. Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"eap_tls_l_bit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"eap_ttls": &schema.Schema{
							Description: `The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored.
The object eapTtls contains the settings for EAP TTLS protocol`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eap_ttls_chap": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_md5": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_ttls_ms_chap_v1": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_pap_ascii": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Resource UUID, Mandatory for update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Resource Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"peap": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_peap_eap_gtc": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_gtc_pwd_change": &schema.Schema{
										Description: `The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_peap_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_peap_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_peap_v0": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"require_cryptobinding": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"preferred_eap_protocol": &schema.Schema{
							Description: `The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored.
Allowed Values: 
- EAP_FAST,
- PEAP,
- LEAP,
- EAP_MD5,
- EAP_TLS,
- EAP_TTLS,
- TEAP`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"process_host_lookup": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"require_message_auth": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"teap": &schema.Schema{
							Description: `The teap is required only if allowTeap is true, otherwise it must be ignored.
The object teap contains the settings for TEAP protocol`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accept_client_cert_during_tunnel_est": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_downgrade_msk": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true,
otherwise it must be ignored.
Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_teap_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_chap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_fast": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_md5": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_tls": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_eap_ttls": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_leap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_ms_chap_v1": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_ms_chap_v2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_pap_ascii": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_peap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_preferred_eap_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_teap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_weak_ciphers_for_eap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"eap_fast": &schema.Schema{
							Description: `The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_fast_eap_gtc": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_gtc_pwd_change": &schema.Schema{
										Description: `The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_eap_fast_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"allow_eap_fast_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_eap_fast_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_dont_use_pacs_accept_client_cert": &schema.Schema{
										Description: `The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_dont_use_pacs_allow_machine_authentication": &schema.Schema{
										Description: `The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_accept_client_cert": &schema.Schema{
										Description: `The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_anonym_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_authen_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_machine_authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsAuthorizationPacTtlUnits is required only if eapFastUsePacsStatelessSessionResume is true,
otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsMachinePacTtlUnits is required only if eapFastUsePacsAllowMachineAuthentication is true,
otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning
is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_stateless_session_resume": &schema.Schema{
										Description: `The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored.
Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_fast_use_pacs_use_proactive_pac_update_precentage": &schema.Schema{
										Description: `The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"eap_tls": &schema.Schema{
							Description: `The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_tls_enable_stateless_session_resume": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_tls_session_ticket_precentage": &schema.Schema{
										Description: `The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true,
otherwise it must be ignored`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_tls_session_ticket_ttl": &schema.Schema{
										Description: `Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"eap_tls_session_ticket_ttl_units": &schema.Schema{
										Description: `Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true,
otherwise it must be ignored. Allowed Values:
- SECONDS,
- MINUTES,
- HOURS,
- DAYS,
- WEEKS`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"eap_tls_l_bit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"eap_ttls": &schema.Schema{
							Description: `The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored.
The object eapTtls contains the settings for EAP TTLS protocol`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eap_ttls_chap": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_md5": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"eap_ttls_ms_chap_v1": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_ttls_pap_ascii": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Resource UUID, Mandatory for update`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Resource Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"peap": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_peap_eap_gtc": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_gtc_pwd_change": &schema.Schema{
										Description: `The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_peap_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true,
otherwise it must be ignored`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true,
otherwise it must be ignored. Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_peap_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_peap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_peap_v0": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"require_cryptobinding": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"preferred_eap_protocol": &schema.Schema{
							Description: `The preferredEapProtocol is required only if allowPreferredEapProtocol is true, otherwise it must be ignored.
Allowed Values: 
- EAP_FAST,
- PEAP,
- LEAP,
- EAP_MD5,
- EAP_TLS,
- EAP_TTLS,
- TEAP`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"process_host_lookup": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"require_message_auth": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"teap": &schema.Schema{
							Description: `The teap is required only if allowTeap is true, otherwise it must be ignored.
The object teap contains the settings for TEAP protocol`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accept_client_cert_during_tunnel_est": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_downgrade_msk": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true,
otherwise it must be ignored.
Valid range is 0-3`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allow_teap_eap_tls": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"allow_teap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description: `The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAllowedProtocolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAllowedProtocols")
		queryParams1 := isegosdk.GetAllowedProtocolsQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.AllowedProtocols.GetAllowedProtocols(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocols", err,
				"Failure at GetAllowedProtocols, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.AllowedProtocols.GetAllowedProtocols(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenAllowedProtocolsGetAllowedProtocolsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocols response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetAllowedProtocolByName")
		vvName := vName.(string)

		response2, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByName", err,
				"Failure at GetAllowedProtocolByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenAllowedProtocolsGetAllowedProtocolByNameItemName(response2.AllowedProtocols)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetAllowedProtocolByID")
		vvID := vID.(string)

		response3, _, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByID", err,
				"Failure at GetAllowedProtocolByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenAllowedProtocolsGetAllowedProtocolByIDItemID(response3.AllowedProtocols)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAllowedProtocolsGetAllowedProtocolsItems(items *[]isegosdk.ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenAllowedProtocolsGetAllowedProtocolsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAllowedProtocolsGetAllowedProtocolsItemsLink(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolsSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemName(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocols) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["eap_tls"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapTls(item.EapTls)
	respItem["peap"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNamePeap(item.Peap)
	respItem["eap_fast"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapFast(item.EapFast)
	respItem["eap_ttls"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapTtls(item.EapTtls)
	respItem["teap"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNameTeap(item.Teap)
	respItem["process_host_lookup"] = boolPtrToString(item.ProcessHostLookup)
	respItem["allow_pap_ascii"] = boolPtrToString(item.AllowPapAscii)
	respItem["allow_chap"] = boolPtrToString(item.AllowChap)
	respItem["allow_ms_chap_v1"] = boolPtrToString(item.AllowMsChapV1)
	respItem["allow_ms_chap_v2"] = boolPtrToString(item.AllowMsChapV2)
	respItem["allow_eap_md5"] = boolPtrToString(item.AllowEapMd5)
	respItem["allow_leap"] = boolPtrToString(item.AllowLeap)
	respItem["allow_eap_tls"] = boolPtrToString(item.AllowEapTls)
	respItem["allow_eap_ttls"] = boolPtrToString(item.AllowEapTtls)
	respItem["allow_eap_fast"] = boolPtrToString(item.AllowEapFast)
	respItem["allow_peap"] = boolPtrToString(item.AllowPeap)
	respItem["allow_teap"] = boolPtrToString(item.AllowTeap)
	respItem["allow_preferred_eap_protocol"] = boolPtrToString(item.AllowPreferredEapProtocol)
	respItem["preferred_eap_protocol"] = item.PreferredEapProtocol
	respItem["eap_tls_l_bit"] = boolPtrToString(item.EapTlsLBit)
	respItem["allow_weak_ciphers_for_eap"] = boolPtrToString(item.AllowWeakCiphersForEap)
	respItem["require_message_auth"] = boolPtrToString(item.RequireMessageAuth)
	respItem["link"] = flattenAllowedProtocolsGetAllowedProtocolByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapTls(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTls) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowEapTlsAuthOfExpiredCerts)
	respItem["eap_tls_enable_stateless_session_resume"] = boolPtrToString(item.EapTlsEnableStatelessSessionResume)
	respItem["eap_tls_session_ticket_ttl"] = item.EapTlsSessionTicketTtl
	respItem["eap_tls_session_ticket_ttl_units"] = item.EapTlsSessionTicketTtlUnits
	respItem["eap_tls_session_ticket_precentage"] = item.EapTlsSessionTicketPrecentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNamePeap(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsPeap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_peap_eap_ms_chap_v2"] = boolPtrToString(item.AllowPeapEapMsChapV2)
	respItem["allow_peap_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowPeapEapMsChapV2PwdChange)
	respItem["allow_peap_eap_ms_chap_v2_pwd_change_retries"] = item.AllowPeapEapMsChapV2PwdChangeRetries
	respItem["allow_peap_eap_gtc"] = boolPtrToString(item.AllowPeapEapGtc)
	respItem["allow_peap_eap_gtc_pwd_change"] = boolPtrToString(item.AllowPeapEapGtcPwdChange)
	respItem["allow_peap_eap_gtc_pwd_change_retries"] = item.AllowPeapEapGtcPwdChangeRetries
	respItem["allow_peap_eap_tls"] = boolPtrToString(item.AllowPeapEapTls)
	respItem["allow_peap_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowPeapEapTlsAuthOfExpiredCerts)
	respItem["require_cryptobinding"] = boolPtrToString(item.RequireCryptobinding)
	respItem["allow_peap_v0"] = boolPtrToString(item.AllowPeapV0)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapFast(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapFast) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_eap_fast_eap_ms_chap_v2"] = boolPtrToString(item.AllowEapFastEapMsChapV2)
	respItem["allow_eap_fast_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowEapFastEapMsChapV2PwdChange)
	respItem["allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"] = item.AllowEapFastEapMsChapV2PwdChangeRetries
	respItem["allow_eap_fast_eap_gtc"] = boolPtrToString(item.AllowEapFastEapGtc)
	respItem["allow_eap_fast_eap_gtc_pwd_change"] = boolPtrToString(item.AllowEapFastEapGtcPwdChange)
	respItem["allow_eap_fast_eap_gtc_pwd_change_retries"] = item.AllowEapFastEapGtcPwdChangeRetries
	respItem["allow_eap_fast_eap_tls"] = boolPtrToString(item.AllowEapFastEapTls)
	respItem["allow_eap_fast_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowEapFastEapTlsAuthOfExpiredCerts)
	respItem["eap_fast_use_pacs"] = boolPtrToString(item.EapFastUsePacs)
	respItem["eap_fast_use_pacs_tunnel_pac_ttl"] = item.EapFastUsePacsTunnelPacTtl
	respItem["eap_fast_use_pacs_tunnel_pac_ttl_units"] = item.EapFastUsePacsTunnelPacTtlUnits
	respItem["eap_fast_use_pacs_use_proactive_pac_update_precentage"] = item.EapFastUsePacsUseProactivePacUpdatePrecentage
	respItem["eap_fast_use_pacs_allow_anonym_provisioning"] = boolPtrToString(item.EapFastUsePacsAllowAnonymProvisioning)
	respItem["eap_fast_use_pacs_allow_authen_provisioning"] = boolPtrToString(item.EapFastUsePacsAllowAuthenProvisioning)
	respItem["eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning"] = boolPtrToString(item.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning)
	respItem["eap_fast_use_pacs_accept_client_cert"] = boolPtrToString(item.EapFastUsePacsAcceptClientCert)
	respItem["eap_fast_use_pacs_machine_pac_ttl"] = item.EapFastUsePacsMachinePacTtl
	respItem["eap_fast_use_pacs_machine_pac_ttl_units"] = item.EapFastUsePacsMachinePacTtlUnits
	respItem["eap_fast_use_pacs_allow_machine_authentication"] = boolPtrToString(item.EapFastUsePacsAllowMachineAuthentication)
	respItem["eap_fast_use_pacs_stateless_session_resume"] = boolPtrToString(item.EapFastUsePacsStatelessSessionResume)
	respItem["eap_fast_use_pacs_authorization_pac_ttl"] = item.EapFastUsePacsAuthorizationPacTtl
	respItem["eap_fast_use_pacs_authorization_pac_ttl_units"] = item.EapFastUsePacsAuthorizationPacTtlUnits
	respItem["eap_fast_dont_use_pacs_accept_client_cert"] = boolPtrToString(item.EapFastDontUsePacsAcceptClientCert)
	respItem["eap_fast_dont_use_pacs_allow_machine_authentication"] = boolPtrToString(item.EapFastDontUsePacsAllowMachineAuthentication)
	respItem["eap_fast_enable_eap_chaining"] = boolPtrToString(item.EapFastEnableEApChaining)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNameEapTtls(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsEapTtls) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["eap_ttls_pap_ascii"] = boolPtrToString(item.EapTtlsPapAscii)
	respItem["eap_ttls_chap"] = boolPtrToString(item.EapTtlsChap)
	respItem["eap_ttls_ms_chap_v1"] = boolPtrToString(item.EapTtlsMsChapV1)
	respItem["eap_ttls_ms_chap_v2"] = boolPtrToString(item.EapTtlsMsChapV2)
	respItem["eap_ttls_eap_md5"] = boolPtrToString(item.EapTtlsEapMd5)
	respItem["eap_ttls_eap_ms_chap_v2"] = boolPtrToString(item.EapTtlsEapMsChapV2)
	respItem["eap_ttls_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.EapTtlsEapMsChapV2PwdChange)
	respItem["eap_ttls_eap_ms_chap_v2_pwd_change_retries"] = item.EapTtlsEapMsChapV2PwdChangeRetries

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNameTeap(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsTeap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_teap_eap_ms_chap_v2"] = boolPtrToString(item.AllowTeapEapMsChapV2)
	respItem["allow_teap_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowTeapEapMsChapV2PwdChange)
	respItem["allow_teap_eap_ms_chap_v2_pwd_change_retries"] = item.AllowTeapEapMsChapV2PwdChangeRetries
	respItem["allow_teap_eap_tls"] = boolPtrToString(item.AllowTeapEapTls)
	respItem["allow_teap_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowTeapEapTlsAuthOfExpiredCerts)
	respItem["accept_client_cert_during_tunnel_est"] = boolPtrToString(item.AcceptClientCertDuringTunnelEst)
	respItem["enable_eap_chaining"] = boolPtrToString(item.EnableEapChaining)
	respItem["allow_downgrade_msk"] = boolPtrToString(item.AllowDowngradeMsk)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByNameItemNameLink(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByNameAllowedProtocolsLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemID(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocols) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["eap_tls"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapTls(item.EapTls)
	respItem["peap"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDPeap(item.Peap)
	respItem["eap_fast"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapFast(item.EapFast)
	respItem["eap_ttls"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapTtls(item.EapTtls)
	respItem["teap"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDTeap(item.Teap)
	respItem["process_host_lookup"] = boolPtrToString(item.ProcessHostLookup)
	respItem["allow_pap_ascii"] = boolPtrToString(item.AllowPapAscii)
	respItem["allow_chap"] = boolPtrToString(item.AllowChap)
	respItem["allow_ms_chap_v1"] = boolPtrToString(item.AllowMsChapV1)
	respItem["allow_ms_chap_v2"] = boolPtrToString(item.AllowMsChapV2)
	respItem["allow_eap_md5"] = boolPtrToString(item.AllowEapMd5)
	respItem["allow_leap"] = boolPtrToString(item.AllowLeap)
	respItem["allow_eap_tls"] = boolPtrToString(item.AllowEapTls)
	respItem["allow_eap_ttls"] = boolPtrToString(item.AllowEapTtls)
	respItem["allow_eap_fast"] = boolPtrToString(item.AllowEapFast)
	respItem["allow_peap"] = boolPtrToString(item.AllowPeap)
	respItem["allow_teap"] = boolPtrToString(item.AllowTeap)
	respItem["allow_preferred_eap_protocol"] = boolPtrToString(item.AllowPreferredEapProtocol)
	respItem["preferred_eap_protocol"] = item.PreferredEapProtocol
	respItem["eap_tls_l_bit"] = boolPtrToString(item.EapTlsLBit)
	respItem["allow_weak_ciphers_for_eap"] = boolPtrToString(item.AllowWeakCiphersForEap)
	respItem["require_message_auth"] = boolPtrToString(item.RequireMessageAuth)
	respItem["link"] = flattenAllowedProtocolsGetAllowedProtocolByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapTls(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTls) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowEapTlsAuthOfExpiredCerts)
	respItem["eap_tls_enable_stateless_session_resume"] = boolPtrToString(item.EapTlsEnableStatelessSessionResume)
	respItem["eap_tls_session_ticket_ttl"] = item.EapTlsSessionTicketTtl
	respItem["eap_tls_session_ticket_ttl_units"] = item.EapTlsSessionTicketTtlUnits
	respItem["eap_tls_session_ticket_precentage"] = item.EapTlsSessionTicketPrecentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDPeap(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsPeap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_peap_eap_ms_chap_v2"] = boolPtrToString(item.AllowPeapEapMsChapV2)
	respItem["allow_peap_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowPeapEapMsChapV2PwdChange)
	respItem["allow_peap_eap_ms_chap_v2_pwd_change_retries"] = item.AllowPeapEapMsChapV2PwdChangeRetries
	respItem["allow_peap_eap_gtc"] = boolPtrToString(item.AllowPeapEapGtc)
	respItem["allow_peap_eap_gtc_pwd_change"] = boolPtrToString(item.AllowPeapEapGtcPwdChange)
	respItem["allow_peap_eap_gtc_pwd_change_retries"] = item.AllowPeapEapGtcPwdChangeRetries
	respItem["allow_peap_eap_tls"] = boolPtrToString(item.AllowPeapEapTls)
	respItem["allow_peap_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowPeapEapTlsAuthOfExpiredCerts)
	respItem["require_cryptobinding"] = boolPtrToString(item.RequireCryptobinding)
	respItem["allow_peap_v0"] = boolPtrToString(item.AllowPeapV0)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapFast(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapFast) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_eap_fast_eap_ms_chap_v2"] = boolPtrToString(item.AllowEapFastEapMsChapV2)
	respItem["allow_eap_fast_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowEapFastEapMsChapV2PwdChange)
	respItem["allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"] = item.AllowEapFastEapMsChapV2PwdChangeRetries
	respItem["allow_eap_fast_eap_gtc"] = boolPtrToString(item.AllowEapFastEapGtc)
	respItem["allow_eap_fast_eap_gtc_pwd_change"] = boolPtrToString(item.AllowEapFastEapGtcPwdChange)
	respItem["allow_eap_fast_eap_gtc_pwd_change_retries"] = item.AllowEapFastEapGtcPwdChangeRetries
	respItem["allow_eap_fast_eap_tls"] = boolPtrToString(item.AllowEapFastEapTls)
	respItem["allow_eap_fast_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowEapFastEapTlsAuthOfExpiredCerts)
	respItem["eap_fast_use_pacs"] = boolPtrToString(item.EapFastUsePacs)
	respItem["eap_fast_use_pacs_tunnel_pac_ttl"] = item.EapFastUsePacsTunnelPacTtl
	respItem["eap_fast_use_pacs_tunnel_pac_ttl_units"] = item.EapFastUsePacsTunnelPacTtlUnits
	respItem["eap_fast_use_pacs_use_proactive_pac_update_precentage"] = item.EapFastUsePacsUseProactivePacUpdatePrecentage
	respItem["eap_fast_use_pacs_allow_anonym_provisioning"] = boolPtrToString(item.EapFastUsePacsAllowAnonymProvisioning)
	respItem["eap_fast_use_pacs_allow_authen_provisioning"] = boolPtrToString(item.EapFastUsePacsAllowAuthenProvisioning)
	respItem["eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning"] = boolPtrToString(item.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning)
	respItem["eap_fast_use_pacs_accept_client_cert"] = boolPtrToString(item.EapFastUsePacsAcceptClientCert)
	respItem["eap_fast_use_pacs_machine_pac_ttl"] = item.EapFastUsePacsMachinePacTtl
	respItem["eap_fast_use_pacs_machine_pac_ttl_units"] = item.EapFastUsePacsMachinePacTtlUnits
	respItem["eap_fast_use_pacs_allow_machine_authentication"] = boolPtrToString(item.EapFastUsePacsAllowMachineAuthentication)
	respItem["eap_fast_use_pacs_stateless_session_resume"] = boolPtrToString(item.EapFastUsePacsStatelessSessionResume)
	respItem["eap_fast_use_pacs_authorization_pac_ttl"] = item.EapFastUsePacsAuthorizationPacTtl
	respItem["eap_fast_use_pacs_authorization_pac_ttl_units"] = item.EapFastUsePacsAuthorizationPacTtlUnits
	respItem["eap_fast_dont_use_pacs_accept_client_cert"] = boolPtrToString(item.EapFastDontUsePacsAcceptClientCert)
	respItem["eap_fast_dont_use_pacs_allow_machine_authentication"] = boolPtrToString(item.EapFastDontUsePacsAllowMachineAuthentication)
	respItem["eap_fast_enable_eap_chaining"] = boolPtrToString(item.EapFastEnableEApChaining)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDEapTtls(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsEapTtls) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["eap_ttls_pap_ascii"] = boolPtrToString(item.EapTtlsPapAscii)
	respItem["eap_ttls_chap"] = boolPtrToString(item.EapTtlsChap)
	respItem["eap_ttls_ms_chap_v1"] = boolPtrToString(item.EapTtlsMsChapV1)
	respItem["eap_ttls_ms_chap_v2"] = boolPtrToString(item.EapTtlsMsChapV2)
	respItem["eap_ttls_eap_md5"] = boolPtrToString(item.EapTtlsEapMd5)
	respItem["eap_ttls_eap_ms_chap_v2"] = boolPtrToString(item.EapTtlsEapMsChapV2)
	respItem["eap_ttls_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.EapTtlsEapMsChapV2PwdChange)
	respItem["eap_ttls_eap_ms_chap_v2_pwd_change_retries"] = item.EapTtlsEapMsChapV2PwdChangeRetries

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDTeap(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsTeap) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_teap_eap_ms_chap_v2"] = boolPtrToString(item.AllowTeapEapMsChapV2)
	respItem["allow_teap_eap_ms_chap_v2_pwd_change"] = boolPtrToString(item.AllowTeapEapMsChapV2PwdChange)
	respItem["allow_teap_eap_ms_chap_v2_pwd_change_retries"] = item.AllowTeapEapMsChapV2PwdChangeRetries
	respItem["allow_teap_eap_tls"] = boolPtrToString(item.AllowTeapEapTls)
	respItem["allow_teap_eap_tls_auth_of_expired_certs"] = boolPtrToString(item.AllowTeapEapTlsAuthOfExpiredCerts)
	respItem["accept_client_cert_during_tunnel_est"] = boolPtrToString(item.AcceptClientCertDuringTunnelEst)
	respItem["enable_eap_chaining"] = boolPtrToString(item.EnableEapChaining)
	respItem["allow_downgrade_msk"] = boolPtrToString(item.AllowDowngradeMsk)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAllowedProtocolsGetAllowedProtocolByIDItemIDLink(item *isegosdk.ResponseAllowedProtocolsGetAllowedProtocolByIDAllowedProtocolsLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
