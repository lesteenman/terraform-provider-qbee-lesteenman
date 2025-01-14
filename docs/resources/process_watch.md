---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qbee_process_watch Resource - qbee"
subcategory: ""
description: |-
  ProcessWatch ensures running process are running (or not).
---

# qbee_process_watch (Resource)

ProcessWatch ensures running process are running (or not).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `extend` (Boolean) If the configuration should extend configuration from the parent nodes of the node the configuration is applied to. If set to false, configuration from parent nodes is ignored.
- `processes` (Attributes List) Processes to watch. (see [below for nested schema](#nestedatt--processes))

### Optional

- `node` (String) The node for which to set the configuration. Either tag or node is required.
- `tag` (String) The tag for which to set the configuration. Either tag or node is required.

<a id="nestedatt--processes"></a>
### Nested Schema for `processes`

Required:

- `command` (String) Command to use to get the process in the expected state. For ProcessPresent it should be a start command, for ProcessAbsent it should be a stop command.
- `name` (String) Name of the process to watch.
- `policy` (String) Policy for the process.
