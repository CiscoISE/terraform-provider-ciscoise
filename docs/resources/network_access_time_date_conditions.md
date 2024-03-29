---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_network_access_time_date_conditions Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on Network Access - Time/Date Conditions.
  Network Access Creates time/date conditionNetwork Access Update network conditionNetwork Access Delete Time/Date condition.
---

# ciscoise_network_access_time_date_conditions (Resource)

It manages create, read, update and delete operations on Network Access - Time/Date Conditions.

- Network Access Creates time/date condition

- Network Access Update network condition

- Network Access Delete Time/Date condition.

## Example Usage

```terraform
resource "ciscoise_network_access_time_date_conditions" "example" {
  provider = ciscoise
  parameters {

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
}

output "ciscoise_network_access_time_date_conditions_example" {
  value = ciscoise_network_access_time_date_conditions.example
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

Optional:

- `children` (Block List) In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition (see [below for nested schema](#nestedblock--parameters--children))
- `condition_type` (String) <ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>
- `dates_range` (Block List) <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p> (see [below for nested schema](#nestedblock--parameters--dates_range))
- `dates_range_exception` (Block List) <p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p> (see [below for nested schema](#nestedblock--parameters--dates_range_exception))
- `description` (String) Condition description
- `hours_range` (Block List) <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p> (see [below for nested schema](#nestedblock--parameters--hours_range))
- `hours_range_exception` (Block List) <p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p> (see [below for nested schema](#nestedblock--parameters--hours_range_exception))
- `is_negate` (String) Indicates whereas this condition is in negate mode
- `name` (String) Condition name
- `week_days` (List of String) <p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>
- `week_days_exception` (List of String) <p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>

Read-Only:

- `id` (String) The ID of this resource.
- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--link))

<a id="nestedblock--parameters--children"></a>
### Nested Schema for `parameters.children`

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

- `link` (List of Object) (see [below for nested schema](#nestedatt--parameters--children--link))

<a id="nestedatt--parameters--children--link"></a>
### Nested Schema for `parameters.children.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedblock--parameters--dates_range"></a>
### Nested Schema for `parameters.dates_range`

Optional:

- `end_date` (String)
- `start_date` (String)


<a id="nestedblock--parameters--dates_range_exception"></a>
### Nested Schema for `parameters.dates_range_exception`

Optional:

- `end_date` (String)
- `start_date` (String)


<a id="nestedblock--parameters--hours_range"></a>
### Nested Schema for `parameters.hours_range`

Optional:

- `end_time` (String)
- `start_time` (String)


<a id="nestedblock--parameters--hours_range_exception"></a>
### Nested Schema for `parameters.hours_range_exception`

Optional:

- `end_time` (String)
- `start_time` (String)


<a id="nestedatt--parameters--link"></a>
### Nested Schema for `parameters.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `attribute_name` (String)
- `attribute_value` (String)
- `children` (List of Object) (see [below for nested schema](#nestedobjatt--item--children))
- `condition_type` (String)
- `dates_range` (List of Object) (see [below for nested schema](#nestedobjatt--item--dates_range))
- `dates_range_exception` (List of Object) (see [below for nested schema](#nestedobjatt--item--dates_range_exception))
- `description` (String)
- `dictionary_name` (String)
- `dictionary_value` (String)
- `hours_range` (List of Object) (see [below for nested schema](#nestedobjatt--item--hours_range))
- `hours_range_exception` (List of Object) (see [below for nested schema](#nestedobjatt--item--hours_range_exception))
- `id` (String)
- `is_negate` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- `name` (String)
- `operator` (String)
- `week_days` (List of String)
- `week_days_exception` (List of String)

<a id="nestedobjatt--item--children"></a>
### Nested Schema for `item.children`

Read-Only:

- `condition_type` (String)
- `is_negate` (String)
- `link` (List of Object) (see [below for nested schema](#nestedobjatt--item--children--link))

<a id="nestedobjatt--item--children--link"></a>
### Nested Schema for `item.children.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)



<a id="nestedobjatt--item--dates_range"></a>
### Nested Schema for `item.dates_range`

Read-Only:

- `end_date` (String)
- `start_date` (String)


<a id="nestedobjatt--item--dates_range_exception"></a>
### Nested Schema for `item.dates_range_exception`

Read-Only:

- `end_date` (String)
- `start_date` (String)


<a id="nestedobjatt--item--hours_range"></a>
### Nested Schema for `item.hours_range`

Read-Only:

- `end_time` (String)
- `start_time` (String)


<a id="nestedobjatt--item--hours_range_exception"></a>
### Nested Schema for `item.hours_range_exception`

Read-Only:

- `end_time` (String)
- `start_time` (String)


<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- `href` (String)
- `rel` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_network_access_time_date_conditions.example "id:=string"
```
