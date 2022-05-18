terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_device_administration_time_date_conditions" "example" {
  provider = ciscoise
  parameters {
    condition_type = "TimeAndDateCondition"
    dates_range {
      end_date   = "2021-09-30"
      start_date = "2021-09-25"
    }
    dates_range_exception {
      end_date   = "2021-09-28"
      start_date = "2021-09-27"
    }
    description = "Test T&D"
    hours_range {
      start_time = "23:02"
      end_time   = "13:02"
    }
    hours_range_exception {
      end_time   = "23:02"
      start_time = "23:50"
    }
    is_negate           = "false"
    name                = "Test2"
    week_days           = ["Thursday", "Friday", "Saturday", "Sunday"]
    week_days_exception = ["Sunday"]
  }
}

output "ciscoise_device_administration_time_date_conditions_example" {
  value = ciscoise_device_administration_time_date_conditions.example
}
