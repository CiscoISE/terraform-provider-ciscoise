---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_device_administration_global_exception_rules Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on Device Administration - Authorization Global
  Exception Rules.
  Device Admin Create global exception authorization rule:
  Rule must include name and condition.
  Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be
  match.
  Condition can be either reference to a stored Library condition, using model
  ConditionReference
  or dynamically built conditions which are not stored in the conditions Library, using models
  ConditionAttributes, ConditionAndBlock, ConditionOrBlock
  .
  Device Admin Update global exception authorization rule.Device Admin Delete global exception authorization rule.
---

# ciscoise_device_administration_global_exception_rules (Resource)

It manages create, read, update and delete operations on Device Administration - Authorization Global
Exception Rules.

- Device Admin Create global exception authorization rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be
match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


- Device Admin Update global exception authorization rule.

- Device Admin Delete global exception authorization rule.

## Example Usage

```terraform
resource "ciscoise_device_administration_global_exception_rules" "example" {
  provider = ciscoise
  parameters {

    commands = ["string"]
    id       = "string"

    profile = "string"
    rule {

      condition {

        attribute_name  = "string"
        attribute_value = "string"
        children {

          condition_type = "string"
          is_negate      = "false"

        }
        condition_type = "string"
        dates_range {

          end_date   = "string"
          start_date = "string"
        }
        dates_range_exception {

          end_date   = "string"
          start_date = "string"
        }
        description      = "string"
        dictionary_name  = "string"
        dictionary_value = "string"
        hours_range {

          end_time   = "string"
          start_time = "string"
        }
        hours_range_exception {

          end_time   = "string"
          start_time = "string"
        }
        id        = "string"
        is_negate = "false"

        name                = "string"
        operator            = "string"
        week_days           = ["string"]
        week_days_exception = ["string"]
      }
      default    = "false"
      hit_counts = 1
      id         = "string"
      name       = "string"
      rank       = 1
      state      = "string"
    }
  }
}

output "ciscoise_device_administration_global_exception_rules_example" {
  value = ciscoise_device_administration_global_exception_rules.example
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

Required:

- `id` (String) id path parameter. Rule id

Optional:

- `commands` (List of String) Command sets enforce the specified list of commands that can be executed by a device administrator
- `profile` (String) Device admin profiles control the initial login session of the device administrator
- `rule` (Block List) Common attributes in rule authentication/authorization (see [below for nested schema](#nestedblock--parameters--rule))

Read-Only:

- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedblock--parameters--rule"></a>
### Nested Schema for `parameters.rule`

Optional:

- `condition` (Block List) (see [below for nested schema](#nestedblock--parameters--rule--condition))
- `default` (String) Indicates if this rule is the default one
- `hit_counts` (Number) The amount of times the rule was matched
- `id` (String) The identifier of the rule
- `name` (String) Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
- `rank` (Number) The rank(priority) in relation to other rules. Lower rank is higher priority.
- `state` (String) The state that the rule is in. A disabled rule cannot be matched.

<a id="nestedblock--parameters--rule--condition"></a>
### Nested Schema for `parameters.rule.condition`

Optional:

- `children` (Block List) In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition (see [below for nested schema](#nestedblock--parameters--rule--condition--children))
- `condition_type` (String) <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
- `is_negate` (String) Indicates whereas this condition is in negate mode

Read-Only:

- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--rule--condition--link))

<a id="nestedblock--parameters--rule--condition--children"></a>
### Nested Schema for `parameters.rule.condition.link`

Optional:

- `attribute_id` (String) Dictionary attribute id (Optional), used for additional verification
- `attribute_name` (String) Dictionary attribute name
- `attribute_value` (String) <ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>
- `dictionary_name` (String) Dictionary name
- `dictionary_value` (String) Dictionary value
- `end_date` (String)
- `name` (String) Dictionary attribute name
- `operator` (String) Equality operator
- `start_date` (String) <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>

Read-Only:

- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--rule--condition--link--link))

<a id="nestedatt--parameters--rule--condition--link--link"></a>
### Nested Schema for `parameters.rule.condition.link.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--parameters--rule--condition--link"></a>
### Nested Schema for `parameters.rule.condition.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)




<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `commands` (List of String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `profile` (String)
- `rule` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule))

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)


<a id="nestedobjatt--item--rule"></a>
### Nested Schema for `item.rule`

Read-Only:

- `condition` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition))
- `default` (String)
- `hit_counts` (Number)
- `id` (String)
- `name` (String)
- `rank` (Number)
- `state` (String)

<a id="nestedobjatt--item--rule--condition"></a>
### Nested Schema for `item.rule.condition`

Read-Only:

- `attribute_name` (String)
- `attribute_value` (String)
- `children` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--children))
- `condition_type` (String)
- `dates_range` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--dates_range))
- `dates_range_exception` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--dates_range_exception))
- `description` (String)
- `dictionary_name` (String)
- `dictionary_value` (String)
- `hours_range` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--hours_range))
- `hours_range_exception` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--hours_range_exception))
- `id` (String)
- `is_negate` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--link))
- `name` (String)
- `operator` (String)
- `week_days` (List of String)
- `week_days_exception` (List of String)

<a id="nestedobjatt--item--rule--condition--children"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `condition_type` (String)
- `is_negate` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--rule--condition--week_days_exception--link))

<a id="nestedobjatt--item--rule--condition--week_days_exception--link"></a>
### Nested Schema for `item.rule.condition.week_days_exception.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedobjatt--item--rule--condition--dates_range"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `end_date` (String)
- `start_date` (String)


<a id="nestedobjatt--item--rule--condition--dates_range_exception"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `end_date` (String)
- `start_date` (String)


<a id="nestedobjatt--item--rule--condition--hours_range"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `end_time` (String)
- `start_time` (String)


<a id="nestedobjatt--item--rule--condition--hours_range_exception"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `end_time` (String)
- `start_time` (String)


<a id="nestedobjatt--item--rule--condition--link"></a>
### Nested Schema for `item.rule.condition.week_days_exception`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_device_administration_global_exception_rules.example "id:=string"
```
