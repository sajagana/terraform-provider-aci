package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciCloudL4L7LogicalInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciCloudL4L7LogicalInterfaceRead,
		SchemaVersion: 1,
		Schema: AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"cloud_l4_l7_device_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allow_all": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"annotation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_point_selectors": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"annotation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_expression": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_alias": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		}),
	}
}

func dataSourceAciCloudL4L7LogicalInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	name := d.Get("name").(string)
	CloudL4L7DeviceDn := d.Get("cloud_l4_l7_device_dn").(string)
	rn := fmt.Sprintf(models.RnCloudLIf, name)
	dn := fmt.Sprintf("%s/%s", CloudL4L7DeviceDn, rn)

	_, err := getAndSetRemoteCloudL4L7LogicalInterface(aciClient, dn, d)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
