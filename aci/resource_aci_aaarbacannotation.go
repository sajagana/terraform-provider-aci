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

func resourceAciAnnotationToCaptureRbacInfo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciAnnotationToCaptureRbacInfoCreate,
		UpdateContext: resourceAciAnnotationToCaptureRbacInfoUpdate,
		ReadContext:   resourceAciAnnotationToCaptureRbacInfoRead,
		DeleteContext: resourceAciAnnotationToCaptureRbacInfoDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciAnnotationToCaptureRbacInfoImport,
		},

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"parent_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"child_regex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aaa_domain_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func getRemoteAnnotationToCaptureRbacInfo(client *client.Client, dn string) (*models.AnnotationToCaptureRbacInfo, error) {
	aaaRbacAnnotationCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}
	aaaRbacAnnotation := models.AnnotationToCaptureRbacInfoFromContainer(aaaRbacAnnotationCont)
	if aaaRbacAnnotation.DistinguishedName == "" {
		return nil, fmt.Errorf("AAA Domain Annotation %s not found", dn)
	}
	return aaaRbacAnnotation, nil
}

func setAnnotationToCaptureRbacInfoAttributes(aaaRbacAnnotation *models.AnnotationToCaptureRbacInfo, d *schema.ResourceData) (*schema.ResourceData, error) {
	d.SetId(aaaRbacAnnotation.DistinguishedName)
	aaaRbacAnnotationMap, err := aaaRbacAnnotation.ToMap()
	if err != nil {
		return d, err
	}
	dn := d.Id()
	if dn != aaaRbacAnnotation.DistinguishedName {
		d.Set("parent_dn", "")
	} else {
		d.Set("parent_dn", GetParentDn(aaaRbacAnnotation.DistinguishedName, fmt.Sprintf("/"+models.RnAaaRbacAnnotation, aaaRbacAnnotationMap["name"])))
	}
	d.Set("child_regex", aaaRbacAnnotationMap["childRegex"])
	d.Set("aaa_domain_dn", aaaRbacAnnotationMap["domain"])
	return d, nil
}

func resourceAciAnnotationToCaptureRbacInfoImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	aaaRbacAnnotation, err := getRemoteAnnotationToCaptureRbacInfo(aciClient, dn)
	if err != nil {
		return nil, err
	}
	schemaFilled, err := setAnnotationToCaptureRbacInfoAttributes(aaaRbacAnnotation, d)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciAnnotationToCaptureRbacInfoCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] AAA Domain Annotation Info: Beginning Creation")
	aciClient := m.(*client.Client)
	parentDn := d.Get("parent_dn").(string)

	aaaRbacAnnotationAttr := models.AnnotationToCaptureRbacInfoAttributes{}

	if ChildRegex, ok := d.GetOk("child_regex"); ok {
		aaaRbacAnnotationAttr.ChildRegex = ChildRegex.(string)
	}

	if Domain, ok := d.GetOk("aaa_domain_dn"); ok {
		aaaRbacAnnotationAttr.Domain = GetMOName(Domain.(string))
	}
	aaaRbacAnnotation := models.NewAnnotationToCaptureRbacInfo(fmt.Sprintf(models.RnAaaRbacAnnotation, aaaRbacAnnotationAttr.Domain), parentDn, aaaRbacAnnotationAttr)

	err := aciClient.Save(aaaRbacAnnotation)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaaRbacAnnotation.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciAnnotationToCaptureRbacInfoRead(ctx, d, m)
}
func resourceAciAnnotationToCaptureRbacInfoUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] AAA Domain Annotation Info: Beginning Update")
	aciClient := m.(*client.Client)
	parentDn := d.Get("parent_dn").(string)

	aaaRbacAnnotationAttr := models.AnnotationToCaptureRbacInfoAttributes{}

	if ChildRegex, ok := d.GetOk("child_regex"); ok {
		aaaRbacAnnotationAttr.ChildRegex = ChildRegex.(string)
	}

	if Domain, ok := d.GetOk("aaa_domain_dn"); ok {
		aaaRbacAnnotationAttr.Domain = GetMOName(Domain.(string))
	}
	aaaRbacAnnotation := models.NewAnnotationToCaptureRbacInfo(fmt.Sprintf(models.RnAaaRbacAnnotation, aaaRbacAnnotationAttr.Domain), parentDn, aaaRbacAnnotationAttr)

	aaaRbacAnnotation.Status = "modified"

	err := aciClient.Save(aaaRbacAnnotation)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaaRbacAnnotation.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciAnnotationToCaptureRbacInfoRead(ctx, d, m)
}

func resourceAciAnnotationToCaptureRbacInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	aaaRbacAnnotation, err := getRemoteAnnotationToCaptureRbacInfo(aciClient, dn)
	if err != nil {
		return errorForObjectNotFound(err, dn, d)
	}

	_, err = setAnnotationToCaptureRbacInfoAttributes(aaaRbacAnnotation, d)
	if err != nil {
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciAnnotationToCaptureRbacInfoDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	err := aciClient.DeleteByDn(dn, models.AaaRbacAnnotationClassName)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return diag.FromErr(err)
}
