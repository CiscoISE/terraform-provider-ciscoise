---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise Provider"
subcategory: ""
description: |-
  
---

# ciscoise Provider



## Example Usage

```terraform
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

  # Timeout (in seconds) for the RESTful HTTP requests
  single_request_timeout = 60
  # it can be set using the environment variable ISE_SINGLE_REQUEST_TIMEOUT

  # Boolean to enable or disable autoimport on resources
  enable_auto_import = "false"
  # it can be set using the environment variable ISE_ENABLE_AUTO_IMPORT
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `base_url` (String) Identity Services Engine base URL, FQDN or IP. If not set, it uses the ISE_BASE_URL environment variable.
- `debug` (String) Flag for Identity Services Engine to enable debugging. If not set, it uses the ISE_DEBUG environment variable; defaults to `false`.
- `enable_auto_import` (String) Flag to enable or disable terraform automatic import (Automatic import means that when Terraform attempts to create the resource, it will perform a get operation if it founds a matching resource, it will perform an import of the resource it found, this is a similar operation to the terraform import command.) in resources, this is a configuration added to the provider, it uses the ISE_ENABLE_AUTO_IMPORT environment varible; `true` to enable it, defaults to `false`.
- `password` (String, Sensitive) Identity Services Engine password to authenticate. If not set, it uses the ISE_PASSWORD environment variable.
- `single_request_timeout` (Number) Timeout (in seconds) for the RESTful HTTP requests. If not set, it uses the ISE_SINGLE_REQUEST_TIMEOUT environment varible; defaults to 60.
- `ssl_verify` (String, Sensitive) Flag to enable or disable SSL certificate verification. If not set, it uses the ISE_SSL_VERIFY environment variable; defaults to `true`.
- `use_api_gateway` (String) Flag to enable or disable the usage of the ISE's API Gateway. If not set, it uses the ISE_USE_API_GATEWAY environment variable; defaults to `false`.
- `use_csrf_token` (String) Flag to enable or disable the usage of the X-CSRF-Token header. If not set, it uses the ISE_USE_CSRF_TOKEN environment varible; defaults to `false`.
- `username` (String, Sensitive) Identity Services Engine username to authenticate. If not set, it uses the ISE_USERNAME environment variable.
