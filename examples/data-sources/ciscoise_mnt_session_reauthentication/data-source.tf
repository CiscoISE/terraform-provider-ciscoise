
data "ciscoise_mnt_session_reauthentication" "example" {
  provider       = ciscoise
  end_poi_ntm_ac = "string"
  psn_nam_e      = "string"
  rea_uth_typ_e  = "string"
}

output "ciscoise_mnt_session_reauthentication_example" {
  value = data.ciscoise_mnt_session_reauthentication.example.item
}
