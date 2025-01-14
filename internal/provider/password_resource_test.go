package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPasswordResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "qbee_password" "test" {
  tag = "terraform:acctest:password"
  extend = true
  users = [
    {
      username = "testuser"
      password_hash = "$5$XyfC.GiB.I8hP8cT$eyBg53DYYuWG5iAdd1Lm8T2rO/tsq0As2jbkK1lTi3D"
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("qbee_password.test", "tag", "terraform:acctest:password"),
					resource.TestCheckNoResourceAttr("qbee_password.test", "node"),
					resource.TestCheckResourceAttr("qbee_password.test", "extend", "true"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.#", "1"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.username", "testuser"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.password_hash", "$5$XyfC.GiB.I8hP8cT$eyBg53DYYuWG5iAdd1Lm8T2rO/tsq0As2jbkK1lTi3D"),
				),
			},
			// Update to a different template
			{
				Config: providerConfig + `
resource "qbee_password" "test" {
  tag = "terraform:acctest:password"
  extend = true
  users = [
    {
      username = "testuser"
      password_hash = "$5$rxiCYTuoljJlNdvd$sD00V.1VO9ePdFkogkTos6mSzQuuZLkjLXxyYAkfjSA"
    },
    {
      username = "seconduser"
      password_hash = "$5$C1XMOaIfW1niwc1n$qezUJc1c8UPVQwHkyD7BvF5JLQU8dZ0r6uQ4X4e8IbB"
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("qbee_password.test", "tag", "terraform:acctest:password"),
					resource.TestCheckNoResourceAttr("qbee_password.test", "node"),
					resource.TestCheckResourceAttr("qbee_password.test", "extend", "true"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.#", "2"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.username", "testuser"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.password_hash", "$5$rxiCYTuoljJlNdvd$sD00V.1VO9ePdFkogkTos6mSzQuuZLkjLXxyYAkfjSA"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.1.username", "seconduser"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.1.password_hash", "$5$C1XMOaIfW1niwc1n$qezUJc1c8UPVQwHkyD7BvF5JLQU8dZ0r6uQ4X4e8IbB"),
				),
			},
			// Import tag
			{
				ResourceName:                         "qbee_password.test",
				ImportState:                          true,
				ImportStateId:                        "tag:terraform:acctest:password",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "tag",
			},
			// Update to be for a node
			{
				Config: providerConfig + `
resource "qbee_password" "test" {
  node = "integrationtests"
  extend = false
  users = [
    {
      username = "testuser"
      password_hash = "$5$rxiCYTuoljJlNdvd$sD00V.1VO9ePdFkogkTos6mSzQuuZLkjLXxyYAkfjSA"
    },
    {
      username = "seconduser"
      password_hash = "$5$C1XMOaIfW1niwc1n$qezUJc1c8UPVQwHkyD7BvF5JLQU8dZ0r6uQ4X4e8IbB"
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("qbee_password.test", "tag"),
					resource.TestCheckResourceAttr("qbee_password.test", "node", "integrationtests"),
					resource.TestCheckResourceAttr("qbee_password.test", "extend", "false"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.#", "2"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.username", "testuser"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.0.password_hash", "$5$rxiCYTuoljJlNdvd$sD00V.1VO9ePdFkogkTos6mSzQuuZLkjLXxyYAkfjSA"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.1.username", "seconduser"),
					resource.TestCheckResourceAttr("qbee_password.test", "users.1.password_hash", "$5$C1XMOaIfW1niwc1n$qezUJc1c8UPVQwHkyD7BvF5JLQU8dZ0r6uQ4X4e8IbB"),
				),
			},
			// Import testing
			{
				ResourceName:                         "qbee_password.test",
				ImportState:                          true,
				ImportStateId:                        "node:integrationtests",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "node",
			},
		},
	})
}
