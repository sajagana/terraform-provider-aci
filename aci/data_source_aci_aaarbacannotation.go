package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciAnnotationToCaptureRbacInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciAnnotationToCaptureRbacInfoRead,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"parent_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"child_regex": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aaa_domain_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAciAnnotationToCaptureRbacInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	aaaDomainDn := GetMOName(d.Get("aaa_domain_dn").(string))
	parentDn := d.Get("parent_dn").(string)
	rn := fmt.Sprintf(models.RnAaaRbacAnnotation, aaaDomainDn)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	aaaRbacAnnotation, err := getRemoteAnnotationToCaptureRbacInfo(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setAnnotationToCaptureRbacInfoAttributes(aaaRbacAnnotation, d)
	if err != nil {
		return nil
	}

	return nil
}
