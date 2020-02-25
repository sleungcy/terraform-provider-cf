package cloudfoundry

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/satori/go.uuid"
	"github.com/terraform-providers/terraform-provider-cloudfoundry/cloudfoundry/managers"
	"log"
)

func dataSourceUserProvidedServiceInstance() *schema.Resource {

	return &schema.Resource{

		Read: dataSourceUserProvidedServiceInstanceRead,

		Schema: map[string]*schema.Schema{

			"name_or_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"space": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceUserProvidedServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	session := meta.(*managers.Session)

	var (
		name_or_id                  string
		space                       string
		UserProvidedServiceInstance ccv2.UserProvidedServiceInstance
	)

	name_or_id = d.Get("name_or_id").(string)
	space = d.Get("space").(string)
	isUUID := uuid.FromStringOrNil(name_or_id)
	if uuid.Equal(isUUID, uuid.Nil) {

		log.Printf("[AAAA] Finding by NAME")

		UserProvidedServiceInstances, _, err := session.ClientV2.GetUserProvServiceInstances(ccv2.FilterByName(name_or_id), ccv2.FilterEqual(constant.SpaceGUIDFilter, space))
		if err != nil {
			return err
		}
		if len(UserProvidedServiceInstances) == 0 {
			return NotFound
		}
		UserProvidedServiceInstance = UserProvidedServiceInstances[0]
	} else {
		var err error
		UserProvidedServiceInstance, _, err = session.ClientV2.GetUserProvidedServiceInstance(name_or_id)
		if err != nil {
			return err
		}
	}

	d.SetId(UserProvidedServiceInstance.GUID)
	d.Set("name", UserProvidedServiceInstance.Name)
	d.Set("tags", UserProvidedServiceInstance.Tags)

	return nil
}
