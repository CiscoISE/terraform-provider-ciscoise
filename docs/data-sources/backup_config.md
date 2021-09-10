---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_backup_config Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_backup_config (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **backup_encryption_key** (String)
- **backup_name** (String)
- **id** (String) The ID of this resource.
- **repository_name** (String)

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **id** (String)
- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- **message** (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

