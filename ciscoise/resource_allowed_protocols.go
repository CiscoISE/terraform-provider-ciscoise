package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAllowedProtocols() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on AllowedProtocols.

- This resource allows the client to update an allowed protocol.

- This resource deletes an allowed protocol.

- This resource creates an allowed protocol.
`,

		CreateContext: resourceAllowedProtocolsCreate,
		ReadContext:   resourceAllowedProtocolsRead,
		UpdateContext: resourceAllowedProtocolsUpdate,
		DeleteContext: resourceAllowedProtocolsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_chap": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_eap_fast": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_eap_md5": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_eap_tls": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_eap_ttls": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_leap": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_ms_chap_v1": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_ms_chap_v2": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_pap_ascii": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_peap": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_preferred_eap_protocol": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_teap": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"allow_weak_ciphers_for_eap": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"eap_fast": &schema.Schema{
							Description:      `The eapFast is required only if allowEapFast is true, otherwise it must be ignored. The object eapFast contains the settings for EAP FAST protocol`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_fast_eap_gtc": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_eap_fast_eap_gtc_pwd_change": &schema.Schema{
										Description:      `The allowEapFastEapGtcPwdChange is required only if allowEapFastEapGtc is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_eap_fast_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowEapFastEapGtcPwdChangeRetries is required only if allowEapFastEapGtc is true,
		otherwise it must be ignored. Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"allow_eap_fast_eap_ms_chap_v2": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description:      `The allowEapFastEapMsChapV2PwdChange is required only if allowEapFastEapMsChapV2 is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description:      `The allowEapFastEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored. Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"allow_eap_fast_eap_tls": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_eap_fast_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description:      `The allowEapFastEapTlsAuthOfExpiredCerts is required only if allowEapFastEapTls is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_dont_use_pacs_accept_client_cert": &schema.Schema{
										Description:      `The eapFastDontUsePacsAcceptClientCert is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_dont_use_pacs_allow_machine_authentication": &schema.Schema{
										Description:      `The eapFastDontUsePacsAllowMachineAuthentication is required only if eapFastUsePacs is FALSE, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_enable_eap_chaining": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_accept_client_cert": &schema.Schema{
										Description: `The eapFastUsePacsAcceptClientCert is required only if eapFastUsePacsAllowAuthenProvisioning is true,
		otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_allow_anonym_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAnonymProvisioning is required only if eapFastUsePacs is true,
		otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_allow_authen_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsAllowAuthenProvisioning is required only if eapFastUsePacs is true,
		otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_allow_machine_authentication": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsAuthorizationPacTtl is required only if eapFastUsePacsStatelessSessionResume is true,
		otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
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
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_fast_use_pacs_machine_pac_ttl": &schema.Schema{
										Description: `The eapFastUsePacsMachinePacTtl is required only if eapFastUsePacsAllowMachineAuthentication is true,
		otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
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
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning": &schema.Schema{
										Description: `The eapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning
		is required only if eapFastUsePacsAllowAuthenProvisioning is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_stateless_session_resume": &schema.Schema{
										Description:      `The eapFastUsePacsStatelessSessionResume is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl": &schema.Schema{
										Description:      `The eapFastUsePacsTunnelPacTtl is required only if eapFastUsePacs is true, otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl_units": &schema.Schema{
										Description: `The eapFastUsePacsTunnelPacTtlUnits is required only if eapFastUsePacs is true, otherwise it must be ignored.
		Allowed Values:
		- SECONDS,
		- MINUTES,
		- HOURS,
		- DAYS,
		- WEEKS`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_fast_use_pacs_use_proactive_pac_update_precentage": &schema.Schema{
										Description: `The eapFastUsePacsUseProactivePacUpdatePrecentage is required only if eapFastUsePacs is true,
		otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"eap_tls": &schema.Schema{
							Description:      `The eapTls is required only if allowEapTls is true, otherwise it must be ignored. The object eapTls contains the settings for EAP TLS protocol`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_tls_enable_stateless_session_resume": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_tls_session_ticket_precentage": &schema.Schema{
										Description: `The eapTlsSessionTicketPrecentage is required only if eapTlsEnableStatelessSessionResume is true,
		otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_tls_session_ticket_ttl": &schema.Schema{
										Description:      `Time to live. The eapTlsSessionTicketTtl is required only if eapTlsEnableStatelessSessionResume is true, otherwise it must be ignored`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_tls_session_ticket_ttl_units": &schema.Schema{
										Description: `Time to live time units. The eapTlsSessionTicketTtlUnits is required only if eapTlsEnableStatelessSessionResume is true,
		otherwise it must be ignored. Allowed Values:
		- SECONDS,
		- MINUTES,
		- HOURS,
		- DAYS,
		- WEEKS`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"eap_tls_l_bit": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"eap_ttls": &schema.Schema{
							Description: `The eapTtls is required only if allowEapTtls is true, otherwise it must be ignored.
		The object eapTtls contains the settings for EAP TTLS protocol`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eap_ttls_chap": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_eap_md5": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_eap_ms_chap_v2": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description:      `The eapTtlsEapMsChapV2PwdChange is required only if eapTtlsEapMsChapV2 is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The eapTtlsEapMsChapV2PwdChangeRetries is required only if eapTtlsEapMsChapV2 is true,
		otherwise it must be ignored. Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"eap_ttls_ms_chap_v1": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_ms_chap_v2": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"eap_ttls_pap_ascii": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description:      `Resource UUID, Mandatory for update`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
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
							Description:      `Resource Name`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"peap": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_peap_eap_gtc": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_eap_gtc_pwd_change": &schema.Schema{
										Description:      `The allowPeapEapGtcPwdChange is required only if allowPeapEapGtc is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_eap_gtc_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapGtcPwdChangeRetries is required only if allowPeapEapGtc is true,
		otherwise it must be ignored. Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"allow_peap_eap_ms_chap_v2": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChange is required only if allowPeapEapMsChapV2 is true,
		otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowPeapEapMsChapV2PwdChangeRetries is required only if allowPeapEapMsChapV2 is true,
		otherwise it must be ignored. Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"allow_peap_eap_tls": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description:      `The allowPeapEapTlsAuthOfExpiredCerts is required only if allowPeapEapTls is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_peap_v0": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"require_cryptobinding": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
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
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"process_host_lookup": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"require_message_auth": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"teap": &schema.Schema{
							Description: `The teap is required only if allowTeap is true, otherwise it must be ignored.
		The object teap contains the settings for TEAP protocol`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accept_client_cert_during_tunnel_est": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_downgrade_msk": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_teap_eap_ms_chap_v2": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Description:      `The allowTeapEapMsChapV2PwdChange is required only if allowTeapEapMsChapV2 is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Description: `The allowTeapEapMsChapV2PwdChangeRetries is required only if allowTeapEapMsChapV2 is true,
		otherwise it must be ignored.
		Valid range is 0-3`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"allow_teap_eap_tls": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"allow_teap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Description:      `The allowTeapEapTlsAuthOfExpiredCerts is required only if allowTeapEapTls is true, otherwise it must be ignored`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"enable_eap_chaining": &schema.Schema{
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceAllowedProtocolsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AllowedProtocols create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAllowedProtocolsCreateAllowedProtocol(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse1, _, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)
			if err == nil && getResponse1 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceAllowedProtocolsRead(ctx, d, m)
			}
		}
		if okName && vvName != "" {
			getResponse2, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = getResponse2.AllowedProtocols.ID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceAllowedProtocolsRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.AllowedProtocols.CreateAllowedProtocol(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAllowedProtocol", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAllowedProtocol", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceAllowedProtocolsRead(ctx, d, m)
}

func resourceAllowedProtocolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AllowedProtocols read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetAllowedProtocolByName")
		vvName := vName

		response1, restyResp1, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenAllowedProtocolsGetAllowedProtocolByNameItemName(response1.AllowedProtocols)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllowedProtocolByID")
		vvID := vID

		response2, restyResp2, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenAllowedProtocolsGetAllowedProtocolByIDItemID(response2.AllowedProtocols)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceAllowedProtocolsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AllowedProtocols update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByName", err,
				"Failure at GetAllowedProtocolByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.AllowedProtocols != nil {
			vvID = getResp.AllowedProtocols.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestAllowedProtocolsUpdateAllowedProtocolByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.AllowedProtocols.UpdateAllowedProtocolByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAllowedProtocolByID", err, restyResp1.String(),
					"Failure at UpdateAllowedProtocolByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAllowedProtocolByID", err,
				"Failure at UpdateAllowedProtocolByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceAllowedProtocolsRead(ctx, d, m)
}

func resourceAllowedProtocolsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AllowedProtocols delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.AllowedProtocols != nil {
			vvID = getResp.AllowedProtocols.ID
		}
	}
	restyResp1, err := client.AllowedProtocols.DeleteAllowedProtocolByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAllowedProtocolByID", err, restyResp1.String(),
				"Failure at DeleteAllowedProtocolByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAllowedProtocolByID", err,
			"Failure at DeleteAllowedProtocolByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAllowedProtocolsCreateAllowedProtocol(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocol {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocol{}
	request.AllowedProtocols = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls")))) {
		request.EapTls = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls(ctx, key+".eap_tls.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peap")))) {
		request.Peap = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap(ctx, key+".peap.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast")))) {
		request.EapFast = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast(ctx, key+".eap_fast.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls")))) {
		request.EapTtls = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls(ctx, key+".eap_ttls.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".teap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".teap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".teap")))) {
		request.Teap = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap(ctx, key+".teap.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".process_host_lookup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".process_host_lookup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".process_host_lookup")))) {
		request.ProcessHostLookup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_pap_ascii")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_pap_ascii")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_pap_ascii")))) {
		request.AllowPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_chap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_chap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_chap")))) {
		request.AllowChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_ms_chap_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_ms_chap_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_ms_chap_v1")))) {
		request.AllowMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_ms_chap_v2")))) {
		request.AllowMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_md5")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_md5")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_md5")))) {
		request.AllowEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_leap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_leap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_leap")))) {
		request.AllowLeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_tls")))) {
		request.AllowEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_ttls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_ttls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_ttls")))) {
		request.AllowEapTtls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast")))) {
		request.AllowEapFast = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap")))) {
		request.AllowPeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap")))) {
		request.AllowTeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_preferred_eap_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_preferred_eap_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_preferred_eap_protocol")))) {
		request.AllowPreferredEapProtocol = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_eap_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_eap_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_eap_protocol")))) {
		request.PreferredEapProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_l_bit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_l_bit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_l_bit")))) {
		request.EapTlsLBit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_weak_ciphers_for_eap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_weak_ciphers_for_eap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_weak_ciphers_for_eap")))) {
		request.AllowWeakCiphersForEap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_message_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_message_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_message_auth")))) {
		request.RequireMessageAuth = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_tls_auth_of_expired_certs")))) {
		request.AllowEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_enable_stateless_session_resume")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_enable_stateless_session_resume")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_enable_stateless_session_resume")))) {
		request.EapTlsEnableStatelessSessionResume = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl")))) {
		request.EapTlsSessionTicketTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl_units")))) {
		request.EapTlsSessionTicketTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_precentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_precentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_precentage")))) {
		request.EapTlsSessionTicketPrecentage = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2")))) {
		request.AllowPeapEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_ms_chap_v2_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2_pwd_change")))) {
		request.AllowPeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_ms_chap_v2_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_ms_chap_v2_pwd_change_retries")))) {
		request.AllowPeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc")))) {
		request.AllowPeapEapGtc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change")))) {
		request.AllowPeapEapGtcPwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change_retries")))) {
		request.AllowPeapEapGtcPwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_tls")))) {
		request.AllowPeapEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_tls_auth_of_expired_certs")))) {
		request.AllowPeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_cryptobinding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_cryptobinding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_cryptobinding")))) {
		request.RequireCryptobinding = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_v0")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_v0")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_v0")))) {
		request.AllowPeapV0 = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2")))) {
		request.AllowEapFastEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change")))) {
		request.AllowEapFastEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries")))) {
		request.AllowEapFastEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc")))) {
		request.AllowEapFastEapGtc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change")))) {
		request.AllowEapFastEapGtcPwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change_retries")))) {
		request.AllowEapFastEapGtcPwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls")))) {
		request.AllowEapFastEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls_auth_of_expired_certs")))) {
		request.AllowEapFastEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs")))) {
		request.EapFastUsePacs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_tunnel_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl")))) {
		request.EapFastUsePacsTunnelPacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_tunnel_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl_units")))) {
		request.EapFastUsePacsTunnelPacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_use_proactive_pac_update_precentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_use_proactive_pac_update_precentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_use_proactive_pac_update_precentage")))) {
		request.EapFastUsePacsUseProactivePacUpdatePrecentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_anonym_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_anonym_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_anonym_provisioning")))) {
		request.EapFastUsePacsAllowAnonymProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_authen_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_authen_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_authen_provisioning")))) {
		request.EapFastUsePacsAllowAuthenProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")))) {
		request.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_accept_client_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_accept_client_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_accept_client_cert")))) {
		request.EapFastUsePacsAcceptClientCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_machine_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl")))) {
		request.EapFastUsePacsMachinePacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_machine_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl_units")))) {
		request.EapFastUsePacsMachinePacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_machine_authentication")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_machine_authentication")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_machine_authentication")))) {
		request.EapFastUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_stateless_session_resume")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_stateless_session_resume")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_stateless_session_resume")))) {
		request.EapFastUsePacsStatelessSessionResume = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_authorization_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl")))) {
		request.EapFastUsePacsAuthorizationPacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_authorization_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl_units")))) {
		request.EapFastUsePacsAuthorizationPacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_dont_use_pacs_accept_client_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_accept_client_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_accept_client_cert")))) {
		request.EapFastDontUsePacsAcceptClientCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_dont_use_pacs_allow_machine_authentication")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_allow_machine_authentication")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_allow_machine_authentication")))) {
		request.EapFastDontUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_enable_eap_chaining")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_enable_eap_chaining")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_enable_eap_chaining")))) {
		request.EapFastEnableEApChaining = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_pap_ascii")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_pap_ascii")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_pap_ascii")))) {
		request.EapTtlsPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_chap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_chap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_chap")))) {
		request.EapTtlsChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_ms_chap_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_ms_chap_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_ms_chap_v1")))) {
		request.EapTtlsMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_ms_chap_v2")))) {
		request.EapTtlsMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_eap_md5")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_eap_md5")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_eap_md5")))) {
		request.EapTtlsEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_eap_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2")))) {
		request.EapTtlsEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_eap_ms_chap_v2_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2_pwd_change")))) {
		request.EapTtlsEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_eap_ms_chap_v2_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_eap_ms_chap_v2_pwd_change_retries")))) {
		request.EapTtlsEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_ms_chap_v2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2")))) {
		request.AllowTeapEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_ms_chap_v2_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2_pwd_change")))) {
		request.AllowTeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_ms_chap_v2_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_ms_chap_v2_pwd_change_retries")))) {
		request.AllowTeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_tls")))) {
		request.AllowTeapEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_tls_auth_of_expired_certs")))) {
		request.AllowTeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accept_client_cert_during_tunnel_est")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accept_client_cert_during_tunnel_est")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accept_client_cert_during_tunnel_est")))) {
		request.AcceptClientCertDuringTunnelEst = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_eap_chaining")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_eap_chaining")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_eap_chaining")))) {
		request.EnableEapChaining = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_downgrade_msk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_downgrade_msk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_downgrade_msk")))) {
		request.AllowDowngradeMsk = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByID {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByID{}
	request.AllowedProtocols = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}

	vAllowEapTls, okAllowEapTls := d.GetOk(fixKeyAccess(key + ".allow_eap_tls"))
	vvAllowEapTls := interfaceToBoolPtr(vAllowEapTls)
	if okAllowEapTls && vvAllowEapTls != nil && *vvAllowEapTls {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls")))) {
			request.EapTls = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls(ctx, key+".eap_tls.0", d)
		}
	}
	vAllowPeap, okAllowPeap := d.GetOk(fixKeyAccess(key + ".allow_peap"))
	vvAllowPeap := interfaceToBoolPtr(vAllowPeap)
	if okAllowPeap && vvAllowPeap != nil && *vvAllowPeap {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".peap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peap")))) {
			request.Peap = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap(ctx, key+".peap.0", d)
		}
	}
	vAllowEapFast, okAllowEapFast := d.GetOk(fixKeyAccess(key + ".allow_eap_fast"))
	vvAllowEapFast := interfaceToBoolPtr(vAllowEapFast)
	if okAllowEapFast && vvAllowEapFast != nil && *vvAllowEapFast {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast")))) {
			request.EapFast = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast(ctx, key+".eap_fast.0", d)
		}
	}
	vAllowEapTtls, okAllowEapTtls := d.GetOk(fixKeyAccess(key + ".allow_eap_ttls"))
	vvAllowEapTtls := interfaceToBoolPtr(vAllowEapTtls)
	if okAllowEapTtls && vvAllowEapTtls != nil && *vvAllowEapTtls {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls")))) {
			request.EapTtls = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls(ctx, key+".eap_ttls.0", d)
		}
	}
	vAllowTeap, okAllowTeap := d.GetOk(fixKeyAccess(key + ".allow_teap"))
	vvAllowTeap := interfaceToBoolPtr(vAllowTeap)
	if okAllowTeap && vvAllowTeap != nil && *vvAllowTeap {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".teap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".teap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".teap")))) {
			request.Teap = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap(ctx, key+".teap.0", d)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".process_host_lookup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".process_host_lookup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".process_host_lookup")))) {
		request.ProcessHostLookup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_pap_ascii")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_pap_ascii")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_pap_ascii")))) {
		request.AllowPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_chap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_chap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_chap")))) {
		request.AllowChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_ms_chap_v1"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_ms_chap_v1"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_ms_chap_v1"))) {
		request.AllowMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_ms_chap_v2"))) {
		request.AllowMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_md5"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_eap_md5"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_eap_md5"))) {
		request.AllowEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_leap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_leap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_leap")))) {
		request.AllowLeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_tls")))) {
		request.AllowEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_ttls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_ttls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_ttls")))) {
		request.AllowEapTtls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast")))) {
		request.AllowEapFast = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap")))) {
		request.AllowPeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap")))) {
		request.AllowTeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_preferred_eap_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_preferred_eap_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_preferred_eap_protocol")))) {
		request.AllowPreferredEapProtocol = interfaceToBoolPtr(v)
	}
	vAllowPreferredEapProtocol, okAllowPreferredEapProtocol := d.GetOk(fixKeyAccess(key + ".allow_preferred_eap_protocol"))
	vvAllowPreferredEapProtocol := interfaceToBoolPtr(vAllowPreferredEapProtocol)
	if okAllowPreferredEapProtocol && vvAllowPreferredEapProtocol != nil && *vvAllowPreferredEapProtocol {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_eap_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_eap_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_eap_protocol")))) {
			request.PreferredEapProtocol = interfaceToString(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_l_bit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_l_bit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_l_bit")))) {
		request.EapTlsLBit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_weak_ciphers_for_eap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_weak_ciphers_for_eap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_weak_ciphers_for_eap")))) {
		request.AllowWeakCiphersForEap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_message_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_message_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_message_auth")))) {
		request.RequireMessageAuth = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls{}
	vEapTlsEnableStatelessSessionResume, okEapTlsEnableStatelessSessionResume := d.GetOk(fixKeyAccess(key + ".eap_tls_enable_stateless_session_resume"))
	vvEapTlsEnableStatelessSessionResume := interfaceToBoolPtr(vEapTlsEnableStatelessSessionResume)
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_tls_auth_of_expired_certs")))) {
		request.AllowEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_enable_stateless_session_resume")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_enable_stateless_session_resume")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_enable_stateless_session_resume")))) {
		request.EapTlsEnableStatelessSessionResume = interfaceToBoolPtr(v)
	}
	if okEapTlsEnableStatelessSessionResume && vvEapTlsEnableStatelessSessionResume != nil && *vvEapTlsEnableStatelessSessionResume {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl")))) {
			request.EapTlsSessionTicketTtl = interfaceToIntPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_ttl_units")))) {
			request.EapTlsSessionTicketTtlUnits = interfaceToString(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_tls_session_ticket_precentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_tls_session_ticket_precentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_tls_session_ticket_precentage")))) {
			request.EapTlsSessionTicketPrecentage = interfaceToIntPtr(v)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap{}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_peap_eap_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_peap_eap_ms_chap_v2"))) {
		request.AllowPeapEapMsChapV2 = interfaceToBoolPtr(v)
	}

	vAllowPeapEapMsChapV2, okAllowPeapEapMsChapV2 := d.GetOk(key + ".allow_peap_eap_ms_chap_v2")
	vvAllowPeapEapMsChapV2 := interfaceToBoolPtr(vAllowPeapEapMsChapV2)
	if okAllowPeapEapMsChapV2 && vvAllowPeapEapMsChapV2 != nil && *vvAllowPeapEapMsChapV2 {
		if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2_pwd_change"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_peap_eap_ms_chap_v2_pwd_change"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_peap_eap_ms_chap_v2_pwd_change"))) {
			request.AllowPeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2_pwd_change_retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_peap_eap_ms_chap_v2_pwd_change_retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_peap_eap_ms_chap_v2_pwd_change_retries"))) {
			request.AllowPeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc")))) {
		request.AllowPeapEapGtc = interfaceToBoolPtr(v)
	}

	vAllowPeapEapGtc, okAllowPeapEapGtc := d.GetOk(fixKeyAccess(key + ".allow_peap_eap_gtc"))
	vvAllowPeapEapGtc := interfaceToBoolPtr(vAllowPeapEapGtc)
	if okAllowPeapEapGtc && vvAllowPeapEapGtc != nil && *vvAllowPeapEapGtc {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change")))) {
			request.AllowPeapEapGtcPwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_gtc_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_gtc_pwd_change_retries")))) {
			request.AllowPeapEapGtcPwdChangeRetries = interfaceToIntPtr(v)
		}
	}

	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_tls")))) {
		request.AllowPeapEapTls = interfaceToBoolPtr(v)
	}
	vAllowPeapEapTls, okAllowPeapEapTls := d.GetOk(fixKeyAccess(key + ".allow_peap_eap_tls"))
	vvAllowPeapEapTls := interfaceToBoolPtr(vAllowPeapEapTls)
	if okAllowPeapEapTls && vvAllowPeapEapTls != nil && *vvAllowPeapEapTls {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_peap_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_peap_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_peap_eap_tls_auth_of_expired_certs")))) {
			request.AllowPeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_cryptobinding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_cryptobinding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_cryptobinding")))) {
		request.RequireCryptobinding = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_v0"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_peap_v0"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_peap_v0"))) {
		request.AllowPeapV0 = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast{}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_eap_fast_eap_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_eap_fast_eap_ms_chap_v2"))) {
		request.AllowEapFastEapMsChapV2 = interfaceToBoolPtr(v)
	}
	vAllowEapFastEapMsChapV2, okAllowEapFastEapMsChapV2 := d.GetOk(key + ".allow_eap_fast_eap_ms_chap_v2")
	vvAllowEapFastEapMsChapV2 := interfaceToBoolPtr(vAllowEapFastEapMsChapV2)
	if okAllowEapFastEapMsChapV2 && vvAllowEapFastEapMsChapV2 != nil && *vvAllowEapFastEapMsChapV2 {
		if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change"))) {
			request.AllowEapFastEapMsChapV2PwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"))) {
			request.AllowEapFastEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc")))) {
		request.AllowEapFastEapGtc = interfaceToBoolPtr(v)
	}
	vAllowEapFastEapGtc, okAllowEapFastEapGtc := d.GetOk(fixKeyAccess(key + ".allow_eap_fast_eap_gtc"))
	vvAllowEapFastEapGtc := interfaceToBoolPtr(vAllowEapFastEapGtc)
	if okAllowEapFastEapGtc && vvAllowEapFastEapGtc != nil && *vvAllowEapFastEapGtc {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc_pwd_change")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change")))) {
			request.AllowEapFastEapGtcPwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_gtc_pwd_change_retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change_retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_gtc_pwd_change_retries")))) {
			request.AllowEapFastEapGtcPwdChangeRetries = interfaceToIntPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls")))) {
		request.AllowEapFastEapTls = interfaceToBoolPtr(v)
	}
	vAllowEapFastEapTls, okAllowEapFastEapTls := d.GetOk(fixKeyAccess(key + ".allow_eap_fast_eap_tls"))
	vvAllowEapFastEapTls := interfaceToBoolPtr(vAllowEapFastEapTls)
	if okAllowEapFastEapTls && vvAllowEapFastEapTls != nil && *vvAllowEapFastEapTls {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_eap_fast_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_eap_fast_eap_tls_auth_of_expired_certs")))) {
			request.AllowEapFastEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs")))) {
		request.EapFastUsePacs = interfaceToBoolPtr(v)
	}
	vEapFastUsePacs, okEapFastUsePacs := d.GetOk(fixKeyAccess(key + ".eap_fast_use_pacs"))
	vvEapFastUsePacs := interfaceToBoolPtr(vEapFastUsePacs)
	if okEapFastUsePacs && vvEapFastUsePacs != nil && *vvEapFastUsePacs {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_tunnel_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl")))) {
			request.EapFastUsePacsTunnelPacTtl = interfaceToIntPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_tunnel_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_tunnel_pac_ttl_units")))) {
			request.EapFastUsePacsTunnelPacTtlUnits = interfaceToString(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_use_proactive_pac_update_precentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_use_proactive_pac_update_precentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_use_proactive_pac_update_precentage")))) {
			request.EapFastUsePacsUseProactivePacUpdatePrecentage = interfaceToIntPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_anonym_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_anonym_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_anonym_provisioning")))) {
			request.EapFastUsePacsAllowAnonymProvisioning = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_authen_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_authen_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_authen_provisioning")))) {
			request.EapFastUsePacsAllowAuthenProvisioning = interfaceToBoolPtr(v)
		}
		vEapFastUsePacsAllowAuthenProvisioning, okEapFastUsePacsAllowAuthenProvisioning := d.GetOk(fixKeyAccess(key + ".eap_fast_use_pacs_allow_authen_provisioning"))
		vvEapFastUsePacsAllowAuthenProvisioning := interfaceToBoolPtr(vEapFastUsePacsAllowAuthenProvisioning)
		if okEapFastUsePacsAllowAuthenProvisioning && vvEapFastUsePacsAllowAuthenProvisioning != nil && *vvEapFastUsePacsAllowAuthenProvisioning {
			if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning")))) {
				request.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning = interfaceToBoolPtr(v)
			}
			if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_accept_client_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_accept_client_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_accept_client_cert")))) {
				request.EapFastUsePacsAcceptClientCert = interfaceToBoolPtr(v)
			}
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_stateless_session_resume")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_stateless_session_resume")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_stateless_session_resume")))) {
			request.EapFastUsePacsStatelessSessionResume = interfaceToBoolPtr(v)
		}
		vEapFastUsePacsStatelessSessionResume, okEapFastUsePacsStatelessSessionResume := d.GetOk(fixKeyAccess(key + ".eap_fast_use_pacs_stateless_session_resume"))
		vvEapFastUsePacsStatelessSessionResume := interfaceToBoolPtr(vEapFastUsePacsStatelessSessionResume)
		if okEapFastUsePacsStatelessSessionResume && vvEapFastUsePacsStatelessSessionResume != nil && *vvEapFastUsePacsStatelessSessionResume {
			if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_authorization_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl")))) {
				request.EapFastUsePacsAuthorizationPacTtl = interfaceToIntPtr(v)
			}
			if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_authorization_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_authorization_pac_ttl_units")))) {
				request.EapFastUsePacsAuthorizationPacTtlUnits = interfaceToString(v)
			}
		}
	}

	vEapFastUsePacsAllowMachineAuthentication, okEapFastUsePacsAllowMachineAuthentication := d.GetOk(fixKeyAccess(key + ".eap_fast_use_pacs_allow_machine_authentication"))
	vvEapFastUsePacsAllowMachineAuthentication := interfaceToBoolPtr(vEapFastUsePacsAllowMachineAuthentication)
	if okEapFastUsePacsAllowMachineAuthentication && vvEapFastUsePacsAllowMachineAuthentication != nil && *vvEapFastUsePacsAllowMachineAuthentication {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_machine_pac_ttl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl")))) {
			request.EapFastUsePacsMachinePacTtl = interfaceToIntPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_machine_pac_ttl_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_machine_pac_ttl_units")))) {
			request.EapFastUsePacsMachinePacTtlUnits = interfaceToString(v)
		}
	}

	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_use_pacs_allow_machine_authentication")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_machine_authentication")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_use_pacs_allow_machine_authentication")))) {
		request.EapFastUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
	}
	if !(okEapFastUsePacs && vvEapFastUsePacs != nil && *vvEapFastUsePacs) {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_dont_use_pacs_accept_client_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_accept_client_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_accept_client_cert")))) {
			request.EapFastDontUsePacsAcceptClientCert = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_dont_use_pacs_allow_machine_authentication")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_allow_machine_authentication")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_dont_use_pacs_allow_machine_authentication")))) {
			request.EapFastDontUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_fast_enable_eap_chaining")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_fast_enable_eap_chaining")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_fast_enable_eap_chaining")))) {
		request.EapFastEnableEApChaining = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_pap_ascii")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_pap_ascii")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_pap_ascii")))) {
		request.EapTtlsPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_ttls_chap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_ttls_chap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_ttls_chap")))) {
		request.EapTtlsChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_ms_chap_v1"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_ms_chap_v1"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_ms_chap_v1"))) {
		request.EapTtlsMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_ms_chap_v2"))) {
		request.EapTtlsMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_md5"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_eap_md5"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_eap_md5"))) {
		request.EapTtlsEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_eap_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_eap_ms_chap_v2"))) {
		request.EapTtlsEapMsChapV2 = interfaceToBoolPtr(v)
	}
	vEapTtlsEapMsChapV2, okEapTtlsEapMsChapV2 := d.GetOk(key + ".eap_ttls_eap_ms_chap_v2")
	vvEapTtlsEapMsChapV2 := interfaceToBoolPtr(vEapTtlsEapMsChapV2)
	if okEapTtlsEapMsChapV2 && vvEapTtlsEapMsChapV2 != nil && *vvEapTtlsEapMsChapV2 {
		if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2_pwd_change"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_eap_ms_chap_v2_pwd_change"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_eap_ms_chap_v2_pwd_change"))) {
			request.EapTtlsEapMsChapV2PwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2_pwd_change_retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap_ttls_eap_ms_chap_v2_pwd_change_retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap_ttls_eap_ms_chap_v2_pwd_change_retries"))) {
			request.EapTtlsEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap{}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_teap_eap_ms_chap_v2"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_teap_eap_ms_chap_v2"))) {
		request.AllowTeapEapMsChapV2 = interfaceToBoolPtr(v)
	}
	vAllowTeapEapMsChapV2, okAllowTeapEapMsChapV2 := d.GetOk(key + ".allow_teap_eap_ms_chap_v2")
	vvAllowTeapEapMsChapV2 := interfaceToBoolPtr(vAllowTeapEapMsChapV2)
	if okAllowTeapEapMsChapV2 && vvAllowTeapEapMsChapV2 != nil && *vvAllowTeapEapMsChapV2 {
		if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2_pwd_change"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_teap_eap_ms_chap_v2_pwd_change"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_teap_eap_ms_chap_v2_pwd_change"))) {
			request.AllowTeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
		}
		if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2_pwd_change_retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_teap_eap_ms_chap_v2_pwd_change_retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_teap_eap_ms_chap_v2_pwd_change_retries"))) {
			request.AllowTeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_tls")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_tls")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_tls")))) {
		request.AllowTeapEapTls = interfaceToBoolPtr(v)
	}

	vAllowTeapEapTls, okAllowTeapEapTls := d.GetOk(fixKeyAccess(key + ".allow_teap_eap_tls"))
	vvAllowTeapEapTls := interfaceToBoolPtr(vAllowTeapEapTls)
	if okAllowTeapEapTls && vvAllowTeapEapTls != nil && *vvAllowTeapEapTls {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_teap_eap_tls_auth_of_expired_certs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_teap_eap_tls_auth_of_expired_certs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_teap_eap_tls_auth_of_expired_certs")))) {
			request.AllowTeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accept_client_cert_during_tunnel_est")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accept_client_cert_during_tunnel_est")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accept_client_cert_during_tunnel_est")))) {
		request.AcceptClientCertDuringTunnelEst = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_eap_chaining")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_eap_chaining")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_eap_chaining")))) {
		request.EnableEapChaining = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_downgrade_msk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_downgrade_msk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_downgrade_msk")))) {
		request.AllowDowngradeMsk = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
