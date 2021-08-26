# Configure provider with your {{config.client}} {{config.api}} SDK credentials
provider "ciscoise" {
  # {{config.client}} {{config.api}} user name
  username = "admin"
  # it can be set using the environment variable ISE_BASE_URL

  # {{config.client}} {{config.api}} password
  password = "admin123"
  # it can be set using the environment variable ISE_USERNAME

  # {{config.client}} {{config.api}} base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable ISE_PASSWORD

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