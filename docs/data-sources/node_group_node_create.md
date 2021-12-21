---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_node_group_node_create Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs create operation on Node Group.
  This data source action adds a node to the node group in the cluster. When a node that belongs to a node group fails,
  another node in the same node group issues a Change of Authorization (CoA) for all the URL-redirected sessions on the
  failed node.
---

# ciscoise_node_group_node_create (Data Source)

It performs create operation on Node Group.

- This data source action adds a node to the node group in the cluster. When a node that belongs to a node group fails,
another node in the same node group issues a Change of Authorization (CoA) for all the URL-redirected sessions on the
failed node.

## Example Usage

```terraform
data "ciscoise_node_group_node_create" "example" {
  provider        = ciscoise
  node_group_name = "string"
  hostname        = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **node_group_name** (String) nodeGroupName path parameter. Name of the existing node group.

### Optional

- **hostname** (String)
- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **success** (List of Object) (see [below for nested schema](#nestedobjatt--item--success))
- **version** (String)

<a id="nestedobjatt--item--success"></a>
### Nested Schema for `item.success`

Read-Only:

- **message** (String)

