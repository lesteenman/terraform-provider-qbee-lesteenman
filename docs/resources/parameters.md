---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qbee_parameters Resource - qbee"
subcategory: ""
description: |-
  
---

# qbee_parameters (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `extend` (Boolean) If the configuration should extend configuration from the parent nodes of the node the configuration is applied to. If set to false, configuration from parent nodes is ignored.

### Optional

- `node` (String) The node for which to set the configuration. Either tag or node is required.
- `parameters` (Attributes List) (see [below for nested schema](#nestedatt--parameters))
- `tag` (String) The tag for which to set the configuration. Either tag or node is required.

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Required:

- `key` (String)
- `value` (String)