---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qbee_password Resource - qbee"
subcategory: ""
description: |-
  Password bundle sets passwords for existing users.
---

# qbee_password (Resource)

Password bundle sets passwords for existing users.

## Example Usage

```terraform
resource "qbee_password" "example_tag" {
  tag    = "example-tag"
  extend = true
  users = [
    {
      username      = "testuser"
      password_hash = "$5$XyfC.GiB.I8hP8cT$eyBg53DYYuWG5iAdd1Lm8T2rO/tsq0As2jbkK1lTi3D"
    }
  ]
}

resource "qbee_password" "example_node" {
  node   = "example-node-id"
  extend = true
  users = [
    {
      username      = "testuser"
      password_hash = "$5$rxiCYTuoljJlNdvd$sD00V.1VO9ePdFkogkTos6mSzQuuZLkjLXxyYAkfjSA"
    },
    {
      username      = "seconduser"
      password_hash = "$5$C1XMOaIfW1niwc1n$qezUJc1c8UPVQwHkyD7BvF5JLQU8dZ0r6uQ4X4e8IbB"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `extend` (Boolean) If the configuration should extend configuration from the parent nodes of the node the configuration is applied to. If set to false, configuration from parent nodes is ignored.
- `users` (Attributes List) A list of users and their password hashes. (see [below for nested schema](#nestedatt--users))

### Optional

- `node` (String) The node for which to set the configuration. Either tag or node is required.
- `tag` (String) The tag for which to set the configuration. Either tag or node is required.

<a id="nestedatt--users"></a>
### Nested Schema for `users`

Required:

- `password_hash` (String) The password hash for the user. See https://qbee.io/docs/qbee-password.html for more information.
- `username` (String) The username of the user for which the password hash is set.

## Import

Import is supported using the following syntax:

```shell
# qbee_password can be imported by specifying the type (tag or node), followed by a colon,
# and finally either the tag or the node id.

terraform import qbee_password.example_tag tag:example-tag
terraform import qbee_password.example_node node:example-node-id
```