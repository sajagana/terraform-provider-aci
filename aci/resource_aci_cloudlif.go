package aci

import (
	"context"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAciCloudL4L7LogicalInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciCloudL4L7LogicalInterfaceCreate,
		UpdateContext: resourceAciCloudL4L7LogicalInterfaceUpdate,
		ReadContext:   resourceAciCloudL4L7LogicalInterfaceRead,
		DeleteContext: resourceAciCloudL4L7LogicalInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudL4L7LogicalInterfaceImport,
		},

		SchemaVersion: 1,
		Schema: AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"cloud_l4_l7_device_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allow_all": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"annotation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_point_selectors": {
				Type:     schema.TypeSet,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"annotation": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"match_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name_alias": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		}),
	}
}

func getAndSetRemoteCloudL4L7LogicalInterface(client *client.Client, dn string, d *schema.ResourceData) (*schema.ResourceData, error) {
	dnUrl := fmt.Sprintf("%s/%s.json?rsp-subtree=full", client.MOURL, dn)
	cloudLIfCont, err := client.GetViaURL(dnUrl)
	if err != nil {
		return nil, err
	}

	cloudLIfAttrs := cloudLIfCont.S("imdata").Index(0).S("cloudLIf").S("attributes")
	cloudLIfDistinguishedName := models.StripQuotes(cloudLIfAttrs.S("dn").String())
	if cloudLIfDistinguishedName == "" {
		return nil, fmt.Errorf("Cloud L4L7 Logical Interface %s not found", dn)
	}

	cloudLifName := models.StripQuotes(cloudLIfAttrs.S("name").String())
	d.Set("cloud_l4_l7_device_dn", GetParentDn(cloudLIfDistinguishedName, "/"+fmt.Sprintf(models.RnCloudLIf, cloudLifName)))
	d.Set("allow_all", models.StripQuotes(cloudLIfAttrs.S("allowAll").String()))
	d.Set("annotation", models.StripQuotes(cloudLIfAttrs.S("annotation").String()))
	d.Set("name", cloudLifName)
	d.Set("name_alias", models.StripQuotes(cloudLIfAttrs.S("nameAlias").String()))

	cloudLIfChildCount, err := cloudLIfCont.S("imdata").Index(0).S("cloudLIf").ArrayCount("children")
	if err == nil {
		cloudLIfChild := cloudLIfCont.S("imdata").Index(0).S("cloudLIf").S("children")
		cloudLIfChildList := make([]map[string]string, cloudLIfChildCount)
		for i := 0; i < cloudLIfChildCount; i++ {
			endPointSelectorAttrs := cloudLIfChild.Index(i).S("cloudEPSelector").S("attributes")
			endPointSelectorsMap := make(map[string]string)
			endPointSelectorsMap["annotation"] = models.StripQuotes(endPointSelectorAttrs.S("annotation").String())
			endPointSelectorsMap["name"] = models.StripQuotes(endPointSelectorAttrs.S("name").String())
			endPointSelectorsMap["match_expression"] = models.StripQuotes(endPointSelectorAttrs.S("matchExpression").String())
			endPointSelectorsMap["description"] = models.StripQuotes(endPointSelectorAttrs.S("descr").String())
			endPointSelectorsMap["name_alias"] = models.StripQuotes(endPointSelectorAttrs.S("nameAlias").String())
			cloudLIfChildList[i] = endPointSelectorsMap
		}
		d.Set("end_point_selectors", cloudLIfChildList)
	}
	d.SetId(cloudLIfDistinguishedName)
	return d, nil
}

func resourceAciCloudL4L7LogicalInterfaceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())

	aciClient := m.(*client.Client)
	schemaFilled, err := getAndSetRemoteCloudL4L7LogicalInterface(aciClient, d.Id(), d)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func getCloudL4L7LogicalInterfaceObject(d *schema.ResourceData, aciClient *client.Client) error {
	name := d.Get("name").(string)
	cloudL4L7DeviceDn := d.Get("cloud_l4_l7_device_dn").(string)
	var nameAlias, allowAll, annotation string
	if AllowAll, ok := d.GetOk("allow_all"); ok {
		allowAll = AllowAll.(string)
	}

	if NameAlias, ok := d.GetOk("name_alias"); ok {
		nameAlias = NameAlias.(string)
	}

	annotation = "{}"
	if Annotation, ok := d.GetOk("annotation"); ok {
		annotation = Annotation.(string)
	}

	cloudLifMap := map[string]string{
		"annotation": annotation,
		"name":       name,
		"allowAll":   allowAll,
		"nameAlias":  nameAlias,
	}

	cloudEPSelectorsList := make([]interface{}, 0)
	if endPointSelectors, ok := d.GetOk("end_point_selectors"); ok {
		endPointSelectorsMap := endPointSelectors.(*schema.Set).List()
		for _, endPointSelector := range endPointSelectorsMap {
			endPointSelectorMap := endPointSelector.(map[string]interface{})
			cloudEndPointSelectorMap := map[string]interface{}{
				"class_name": models.CloudepselectorClassName,
				"content": map[string]string{
					"annotation":      endPointSelectorMap["annotation"].(string),
					"name":            endPointSelectorMap["name"].(string),
					"matchExpression": endPointSelectorMap["match_expression"].(string),
					"nameAlias":       endPointSelectorMap["name_alias"].(string),
					"descr":           endPointSelectorMap["description"].(string),
				},
			}
			cloudEPSelectorsList = append(cloudEPSelectorsList, cloudEndPointSelectorMap)
		}
	}

	cloudLifCont, err := preparePayload(models.CloudLIfClassName, cloudLifMap, cloudEPSelectorsList)
	if err != nil {
		return err
	}

	httpRequestPayload, err := aciClient.MakeRestRequest("POST", fmt.Sprintf("%s/%s.json", client.DefaultMOURL, cloudL4L7DeviceDn), cloudLifCont, true)
	if err != nil {
		return err
	}

	respCont, _, err := aciClient.Do(httpRequestPayload)
	if err != nil {
		return err
	}

	err = client.CheckForErrors(respCont, "POST", false)
	if err != nil {
		return err
	}

	d.SetId(cloudL4L7DeviceDn + "/" + fmt.Sprintf(models.RnCloudLIf, name))
	return nil
}

func resourceAciCloudL4L7LogicalInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] CloudL4L7LogicalInterface: Beginning Creation")

	aciClient := m.(*client.Client)
	err := getCloudL4L7LogicalInterfaceObject(d, aciClient)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciCloudL4L7LogicalInterfaceRead(ctx, d, m)
}
func resourceAciCloudL4L7LogicalInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] CloudL4L7LogicalInterface: Beginning Update")

	aciClient := m.(*client.Client)
	err := getCloudL4L7LogicalInterfaceObject(d, aciClient)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciCloudL4L7LogicalInterfaceRead(ctx, d, m)
}

func resourceAciCloudL4L7LogicalInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	_, err := getAndSetRemoteCloudL4L7LogicalInterface(aciClient, dn, d)
	if err != nil {
		return errorForObjectNotFound(err, dn, d)
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciCloudL4L7LogicalInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	err := aciClient.DeleteByDn(dn, models.CloudLIfClassName)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return diag.FromErr(err)
}
