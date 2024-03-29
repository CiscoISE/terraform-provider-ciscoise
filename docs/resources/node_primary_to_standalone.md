---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_node_primary_to_standalone Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs create operation on Node Deployment.
  - This resource changes the primary PAN in a single node cluster on which the API is invoked, to a standalone
  node.
---

# ciscoise_node_primary_to_standalone (Resource)

It performs create operation on Node Deployment.
- This resource changes the primary PAN in a single node cluster on which the API is invoked, to a standalone
node.


~>Warning: This resource does not represent a real-world entity in Cisco ISE, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Cisco ISE workflow. It is executed in ISE without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "ciscoise_node_primary_to_standalone" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `success` (List of Object) (see [below for nested schema](#nestedobjatt--item--success))
- `version` (String)

<a id="nestedobjatt--item--success"></a>
### Nested Schema for `item.success`

Read-Only:

- `message` (String)


