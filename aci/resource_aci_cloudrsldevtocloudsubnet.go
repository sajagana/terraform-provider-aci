package aci

import (
	"context"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAciRelationFromCloudLDevToCloudSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciRelationFromCloudLDevToCloudSubnetCreate,
		ReadContext:   resourceAciRelationFromCloudLDevToCloudSubnetRead,
		DeleteContext: resourceAciRelationFromCloudLDevToCloudSubnetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciRelationFromCloudLDevToCloudSubnetImport,
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"cloud_device_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func getRemoteRelationFromCloudLDevToCloudSubnet(client *client.Client, dn string) (*models.RelationFromCloudLDevToCloudSubnet, error) {
	cloudRsLDevToCloudSubnetCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudRsLDevToCloudSubnet := models.RelationFromCloudLDevToCloudSubnetFromContainer(cloudRsLDevToCloudSubnetCont)
	if cloudRsLDevToCloudSubnet.DistinguishedName == "" {
		return nil, fmt.Errorf("Relation from Cloud LDev to Cloud Subnet %s not found", dn)
	}
	return cloudRsLDevToCloudSubnet, nil
}

func setRelationFromCloudLDevToCloudSubnetAttributes(cloudRsLDevToCloudSubnet *models.RelationFromCloudLDevToCloudSubnet, d *schema.ResourceData) (*schema.ResourceData, error) {
	d.SetId(cloudRsLDevToCloudSubnet.DistinguishedName)
	cloudRsLDevToCloudSubnetMap, err := cloudRsLDevToCloudSubnet.ToMap()
	if err != nil {
		return d, err
	}
	dn := d.Id()
	if dn != cloudRsLDevToCloudSubnet.DistinguishedName {
		d.Set("cloud_device_dn", "")
	} else {
		d.Set("cloud_device_dn", GetParentDn(cloudRsLDevToCloudSubnet.DistinguishedName, fmt.Sprintf("/"+models.RnCloudRsLDevToCloudSubnet, cloudRsLDevToCloudSubnetMap["tDn"])))
	}
	d.Set("annotation", cloudRsLDevToCloudSubnetMap["annotation"])
	d.Set("target_dn", cloudRsLDevToCloudSubnetMap["tDn"])
	return d, nil
}

func resourceAciRelationFromCloudLDevToCloudSubnetImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	cloudRsLDevToCloudSubnet, err := getRemoteRelationFromCloudLDevToCloudSubnet(aciClient, dn)
	if err != nil {
		return nil, err
	}
	schemaFilled, err := setRelationFromCloudLDevToCloudSubnetAttributes(cloudRsLDevToCloudSubnet, d)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciRelationFromCloudLDevToCloudSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Relation from Cloud LDev to Cloud Subnet: Beginning Creation")
	aciClient := m.(*client.Client)
	tDn := d.Get("target_dn").(string)
	cloudDeviceDn := d.Get("cloud_device_dn").(string)

	cloudRsLDevToCloudSubnetAttr := models.RelationFromCloudLDevToCloudSubnetAttributes{}

	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudRsLDevToCloudSubnetAttr.Annotation = Annotation.(string)
	} else {
		cloudRsLDevToCloudSubnetAttr.Annotation = "{}"
	}

	cloudRsLDevToCloudSubnetAttr.TDn = tDn
	cloudRsLDevToCloudSubnet := models.NewRelationFromCloudLDevToCloudSubnet(fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn), cloudDeviceDn, cloudRsLDevToCloudSubnetAttr)

	err := aciClient.Save(cloudRsLDevToCloudSubnet)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cloudRsLDevToCloudSubnet.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciRelationFromCloudLDevToCloudSubnetRead(ctx, d, m)
}

func resourceAciRelationFromCloudLDevToCloudSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	cloudRsLDevToCloudSubnet, err := getRemoteRelationFromCloudLDevToCloudSubnet(aciClient, dn)
	if err != nil {

		return errorForObjectNotFound(err, dn, d)
	}

	_, err = setRelationFromCloudLDevToCloudSubnetAttributes(cloudRsLDevToCloudSubnet, d)
	if err != nil {
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciRelationFromCloudLDevToCloudSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	err := aciClient.DeleteByDn(dn, models.CloudRsLDevToCloudSubnetClassName)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return diag.FromErr(err)
}
