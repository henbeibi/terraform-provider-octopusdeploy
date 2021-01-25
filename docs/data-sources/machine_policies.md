---
page_title: "octopusdeploy_machine_policies Data Source - terraform-provider-octopusdeploy"
subcategory: ""
description: |-
  Provides information about existing machine policies.
---

# Data Source `octopusdeploy_machine_policies`

Provides information about existing machine policies.



## Schema

### Optional

- **id** (String, Optional) The ID of this resource.
- **ids** (List of String, Optional) A filter to search by a list of IDs.
- **partial_name** (String, Optional) A filter to search by the partial match of a name.
- **skip** (Number, Optional) A filter to specify the number of items to skip in the response.
- **take** (Number, Optional) A filter to specify the number of items to take (or return) in the response.

### Read-only

- **machine_policies** (Block List) A list of machine policies that match the filter(s). (see [below for nested schema](#nestedblock--machine_policies))

<a id="nestedblock--machine_policies"></a>
### Nested Schema for `machine_policies`

Read-only:

- **connection_connect_timeout** (Number, Read-only)
- **connection_retry_count_limit** (Number, Read-only)
- **connection_retry_sleep_interval** (Number, Read-only)
- **connection_retry_time_limit** (Number, Read-only)
- **description** (String, Read-only) The description of this resource.
- **id** (String, Read-only) The unique ID for this resource.
- **is_default** (Boolean, Read-only)
- **machine_cleanup_policy** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--machine_policies--machine_cleanup_policy))
- **machine_connectivity_policy** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--machine_policies--machine_connectivity_policy))
- **machine_health_check_policy** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--machine_policies--machine_health_check_policy))
- **machine_update_policy** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--machine_policies--machine_update_policy))
- **name** (String, Read-only) The name of this resource.
- **polling_request_maximum_message_processing_timeout** (Number, Read-only)
- **polling_request_queue_timeout** (Number, Read-only)
- **space_id** (String, Read-only) The space ID associated with this resource.

<a id="nestedatt--machine_policies--machine_cleanup_policy"></a>
### Nested Schema for `machine_policies.machine_cleanup_policy`

- **delete_machines_behavior** (String)
- **delete_machines_elapsed_timespan** (Number)


<a id="nestedatt--machine_policies--machine_connectivity_policy"></a>
### Nested Schema for `machine_policies.machine_connectivity_policy`

- **machine_connectivity_behavior** (String)


<a id="nestedatt--machine_policies--machine_health_check_policy"></a>
### Nested Schema for `machine_policies.machine_health_check_policy`

- **bash_health_check_policy** (List of Object) (see [below for nested schema](#nestedobjatt--machine_policies--machine_health_check_policy--bash_health_check_policy))
- **health_check_cron** (String)
- **health_check_cron_timezone** (String)
- **health_check_interval** (Number)
- **health_check_type** (String)
- **powershell_health_check_policy** (List of Object) (see [below for nested schema](#nestedobjatt--machine_policies--machine_health_check_policy--powershell_health_check_policy))

<a id="nestedobjatt--machine_policies--machine_health_check_policy--bash_health_check_policy"></a>
### Nested Schema for `machine_policies.machine_health_check_policy.bash_health_check_policy`

- **run_type** (String)
- **script_body** (String)


<a id="nestedobjatt--machine_policies--machine_health_check_policy--powershell_health_check_policy"></a>
### Nested Schema for `machine_policies.machine_health_check_policy.powershell_health_check_policy`

- **run_type** (String)
- **script_body** (String)



<a id="nestedatt--machine_policies--machine_update_policy"></a>
### Nested Schema for `machine_policies.machine_update_policy`

- **calamari_update_behavior** (String)
- **tentacle_update_account_id** (String)
- **tentacle_update_behavior** (String)

