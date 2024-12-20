---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qbee_softwaremanagement Resource - qbee"
subcategory: ""
description: |-
  SoftwareManagement controls software in the system.
---

# qbee_softwaremanagement (Resource)

SoftwareManagement controls software in the system.

## Example Usage

```terraform
resource "qbee_softwaremanagement" "example" {
  node = "root"
  # or 'tag = "tagname"'
  extend = true

  items = [
    {
      package       = "mysql",
      service_name  = "service",
      pre_condition = "/a/b/c/",
      config_files = [
        {
          template = "/testtemplate"
          location = "/testlocation"
        }
      ]
      parameters = [
        {
          key   = "param1"
          value = "value1"
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `extend` (Boolean) If the configuration should extend configuration from the parent nodes of the node the configuration is applied to. If set to false, configuration from parent nodes is ignored.
- `items` (Attributes List) The filesets that must be distributed (see [below for nested schema](#nestedatt--items))

### Optional

- `node` (String) The node for which to set the configuration. Either tag or node is required.
- `tag` (String) The tag for which to set the configuration. Either tag or node is required.

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Required:

- `package` (String) Package name (with .deb) from package in File manager or package name (without .deb ending) to install it from a apt repository configured on the device (e.g. mc for midnight commander will install from repository)

Optional:

- `config_files` (Attributes List) (see [below for nested schema](#nestedatt--items--config_files))
- `parameters` (Attributes List) (see [below for nested schema](#nestedatt--items--parameters))
- `pre_condition` (String) Script/executable that needs to return successfully before software package is installed. We expect 0 or true. We assume true if left empty. For example, call: /bin/true or finish with exit(0)
- `service_name` (String) Define a service name if it differs from the package name. If empty then service name will be assumed to be the same as the package name

<a id="nestedatt--items--config_files"></a>
### Nested Schema for `items.config_files`

Required:

- `location` (String) The destination of the file on the target device.
- `template` (String) The source of the file. Must correspond to a file in the qbee filemanager.


<a id="nestedatt--items--parameters"></a>
### Nested Schema for `items.parameters`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
# Softwaremanagement can be imported by specifying the type (tag or node), followed by a colon, and
# finally either the tag or the node id.
terraform import qbee_softwaremanagement.example tag:example_tag
terraform import qbee_softwaremanagement.example node:example_node
```
