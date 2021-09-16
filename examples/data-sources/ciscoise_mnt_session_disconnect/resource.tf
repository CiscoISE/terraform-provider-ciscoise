
data "ciscoise_mnt_session_disconnect" "example" {
    provider = ciscoise
    dis_con_nec_tty_pe = "string"
    end_poi_nti_p = "string"
    mac = "string"
    nas_ipv4 = "string"
    psn_nam_e = "string"
}

output "ciscoise_mnt_session_disconnect_example" {
    value = data.ciscoise_mnt_session_disconnect.example.item
}
