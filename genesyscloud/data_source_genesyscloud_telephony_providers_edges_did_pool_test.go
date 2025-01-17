package genesyscloud

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)


func TestAccDataSourceDidPoolBasic(t *testing.T) {
	var (
		didPoolStartPhoneNumber = "+45465550001"
		didPoolEndPhoneNumber = "+45465550333"
		didPoolRes = "didPool"
		didPoolDataRes = "didPoolData"
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// Create
				Config: generateDidPoolResource(&didPoolStruct{
					didPoolRes,
					didPoolStartPhoneNumber,
					didPoolEndPhoneNumber,
					nullValue, // No description
					nullValue, // No comments
					nullValue, // No provider
				}) + generateDidPoolDataSource(didPoolDataRes,
					didPoolStartPhoneNumber,
					didPoolEndPhoneNumber,
					"genesyscloud_telephony_providers_edges_did_pool." + didPoolRes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.genesyscloud_telephony_providers_edges_did_pool." + didPoolDataRes, "id", "genesyscloud_telephony_providers_edges_did_pool." + didPoolRes, "id"),
				),
			},
		},
	})
}

func generateDidPoolDataSource(
	resourceID string,
	startPhoneNumber string,
	endPhoneNumber string,
// Must explicitly use depends_on in terraform v0.13 when a data source references a resource
// Fixed in v0.14 https://github.com/hashicorp/terraform/pull/26284
	dependsOnResource string) string {
	return fmt.Sprintf(`data "genesyscloud_telephony_providers_edges_did_pool" "%s" {
		start_phone_number = "%s"
		end_phone_number = "%s"
		depends_on=[%s]
	}
	`, resourceID, startPhoneNumber, endPhoneNumber, dependsOnResource)
}