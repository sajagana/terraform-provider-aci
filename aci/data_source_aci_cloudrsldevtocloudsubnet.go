package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciRelationFromCloudLDevToCloudSubnet() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciRelationFromCloudLDevToCloudSubnetRead,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"cloud_device_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAciRelationFromCloudLDevToCloudSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	tDn := d.Get("target_dn").(string)
	cloudDeviceDn := d.Get("cloud_device_dn").(string)
	rn := fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn)
	dn := fmt.Sprintf("%s/%s", cloudDeviceDn, rn)

	cloudRsLDevToCloudSubnet, err := getRemoteRelationFromCloudLDevToCloudSubnet(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setRelationFromCloudLDevToCloudSubnetAttributes(cloudRsLDevToCloudSubnet, d)
	if err != nil {
		return nil
	}

	return nil
}
