# Configure provider with your  Cisco Identity Services Engine SDK credentials
provider "ciscoise" {
  #  Cisco Identity Services Engine user name
  username = "admin"
  # it can be set using the environment variable ISE_USERNAME

  #  Cisco Identity Services Engine password
  password = "admin123"
  # it can be set using the environment variable ISE_PASSWORD

  #  Cisco Identity Services Engine base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable ISE_BASE_URL

  # Boolean to enable debugging
  debug = "false"
  # it can be set using the environment variable ISE_DEBUG

  # Boolean to enable or disable SSL certificate verification
  ssl_verify = "false"
  # it can be set using the environment variable ISE_SSL_VERIFY

  # Boolean to enable or disable the usage of the ISE's API Gateway
  use_api_gateway = "false"
  # it can be set using the environment variable ISE_USE_API_GATEWAY

  # Boolean to enable or disable the usage of the X-CSRF-Token header
  use_csrf_token = "false"
  # it can be set using the environment variable ISE_USE_CSRF_TOKEN
}