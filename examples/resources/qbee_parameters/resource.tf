resource "qbee_parameters" "example_tag" {
  tag    = "example_tag"
  extend = true
  parameters = [
    {
      key   = "parameter-key-1"
      value = "parameter-value-1"
    }
  ]
}

resource "qbee_parameters" "example_with_secrets" {
  tag    = "example_tag"
  extend = true
  secrets = [
    {
      key   = "secret-key"
      value = "secret-value"
    }
  ]
}

resource "qbee_filedistribution" "example_node" {
  node   = "example_node"
  extend = true
  parameters = [
    {
      key   = "parameter-key-1"
      value = "parameter-value-1"
    },
    {
      key   = "parameter-key-2"
      value = "parameter-value-2"
    }
  ]
}
