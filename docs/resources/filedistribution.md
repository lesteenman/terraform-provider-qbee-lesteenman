---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qbee_filedistribution Resource - qbee"
subcategory: ""
description: |-
  
---

# qbee_filedistribution (Resource)



## Example Usage

```terraform
resource "qbee_filedistribution" "example_tag" {
  tag = "example_tag"

  extend = true
  files = [
    {
      pre_condition = "/bin/true"
      command       = "date -u > /tmp/last-updated.txt"
      templates = [
        {
          source      = "/example/file.txt.template"
          destination = "/target/path.txt"
          is_template = true
        },
        {
          source      = "/example/file2.json"
          destination = "/target/path2.json"
          is_template = false
        }
      ]
      parameters = [
        {
          key   = "param-key-1"
          value = "param-value-1"
        },
        {
          key   = "param-key-2"
          value = "param-value-2"
        }
      ]
    }
  ]
}

resource "qbee_filedistribution" "example_node" {
  node = "example_node"

  extend = true
  files = [
    {
      templates = [
        {
          source      = "/example/file2.json"
          destination = "/target/path2.json"
          is_template = false
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
- `files` (Attributes List) The filesets to distribute. (see [below for nested schema](#nestedatt--files))

### Optional

- `node` (String) The node for which to set the configuration. Either tag or node is required.
- `tag` (String) The tag for which to set the configuration. Either tag or node is required.

<a id="nestedatt--files"></a>
### Nested Schema for `files`

Required:

- `templates` (Attributes List) (see [below for nested schema](#nestedatt--files--templates))

Optional:

- `command` (String) A command that will be run on the device after this fileset is distributed. Example: `/bin/true`.
- `label` (String) An optional label for the fileset.
- `parameters` (Attributes List) (see [below for nested schema](#nestedatt--files--parameters))
- `pre_condition` (String) A command that must successfully execute on the device (return a non-zero exit code) before this fileset can be distributed. Example: `/bin/true`.

<a id="nestedatt--files--templates"></a>
### Nested Schema for `files.templates`

Required:

- `destination` (String) The destination of the file on the target device.
- `is_template` (Boolean) If this file is a template. If set to true, template substitution of '\{\{ pattern \}\}' will be performed in the file contents, using the parameters defined in this filedistribution config.
- `source` (String) The source of the file. Must correspond to a file in the qbee filemanager.


<a id="nestedatt--files--parameters"></a>
### Nested Schema for `files.parameters`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
# Filedistribution can be imported by specifying the type (tag or node), followed by a colon, and
# finally either the tag or the node id.
terraform import qbee_filedistribution.example tag:example_tag
terraform import qbee_filedistribution.example node:example_node
```