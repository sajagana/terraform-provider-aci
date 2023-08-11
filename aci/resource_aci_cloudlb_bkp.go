package aci

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAciCloudL4L7Devices() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciCloudL4L7DevicesCreate,
		UpdateContext: resourceAciCloudL4L7DevicesUpdate,
		ReadContext:   resourceAciCloudL4L7DevicesRead,
		DeleteContext: resourceAciCloudL4L7DevicesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudL4L7DevicesImport,
		},

		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active_active": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"allow_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"auto_scaling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"context_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"multi-Context",
					"single-Context",
				}, false),
			},
			"custom_rg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"CLOUD",
					"PHYSICAL",
					"VIRTUAL",
				}, false),
			},
			"function_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GoThrough",
					"GoTo",
					"L1",
					"L2",
					"None",
				}, false),
			},
			"instance_count": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"is_instantiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"is_static_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"l4_l7_device_application_security_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l4_l7_third_party_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"managed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"max_instance_count": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_instance_count": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"legacy-Mode",
				}, false),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"native_lb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"package_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"internal",
					"internet",
				}, false),
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"large",
					"medium",
					"small",
					"v2",
				}, false),
			},
			"sku": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"standard",
					"standard_v2",
				}, false),
			},
			"svc_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ADC",
					"COPY",
					"FW",
					"NATIVELB",
					"OTHERS",
				}, false),
			},
			"target_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"primary",
					"secondary",
					"unspecified",
				}, false),
			},
			"trunking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
			"cloud_l4_l7_lb_type": &schema.Schema{ // maps with type
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"application",
					"network",
				}, false),
			},

			"vrf_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"relation_to_cloud_subnet": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"logical_interfaces": {
				Type: schema.TypeSet,
				// MinItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"name_alias": {
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
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"match_expression": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name_alias": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		})),
	}
}

func getRemoteCloudL4L7LB(client *client.Client, dn string) (*models.CloudL4L7LB, error) {
	cloudLBCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudLB := models.CloudL4L7LBFromContainer(cloudLBCont)
	if cloudLB.DistinguishedName == "" {
		return nil, fmt.Errorf("Cloud L4-L7 Device %s not found", dn)
	}
	return cloudLB, nil
}

func setCloudL4L7LBAttributes(cloudLB *models.CloudL4L7LB, d *schema.ResourceData) (*schema.ResourceData, error) {
	d.SetId(cloudLB.DistinguishedName)
	d.Set("description", cloudLB.Description)
	cloudLBMap, err := cloudLB.ToMap()
	if err != nil {
		return d, err
	}
	dn := d.Id()
	if dn != cloudLB.DistinguishedName {
		d.Set("tenant_dn", "")
	} else {
		d.Set("tenant_dn", GetParentDn(cloudLB.DistinguishedName, fmt.Sprintf("/"+models.RnCloudLB, cloudLBMap["name"])))
	}
	d.Set("version", cloudLBMap["Version"])
	d.Set("active_active", cloudLBMap["activeActive"])
	d.Set("allow_all", cloudLBMap["allowAll"])
	d.Set("annotation", cloudLBMap["annotation"])
	d.Set("auto_scaling", cloudLBMap["autoScaling"])
	d.Set("context_aware", cloudLBMap["contextAware"])
	d.Set("custom_rg", cloudLBMap["customRG"])
	d.Set("device_type", cloudLBMap["devtype"])
	d.Set("function_type", cloudLBMap["funcType"])
	d.Set("instance_count", cloudLBMap["instanceCount"])
	d.Set("is_copy", cloudLBMap["isCopy"])
	d.Set("is_instantiation", cloudLBMap["isInstantiation"])
	d.Set("is_static_ip", cloudLBMap["isStaticIP"])
	d.Set("l4_l7_device_application_security_group", cloudLBMap["l4L7DeviceApplicationSecurityGroup"])
	d.Set("l4_l7_third_party_device", cloudLBMap["l4L7ThirdPartyDevice"])
	d.Set("managed", cloudLBMap["managed"])
	d.Set("max_instance_count", cloudLBMap["maxInstanceCount"])
	d.Set("min_instance_count", cloudLBMap["minInstanceCount"])
	d.Set("mode", cloudLBMap["mode"])
	d.Set("name", cloudLBMap["name"])
	d.Set("name_alias", cloudLBMap["nameAlias"])
	d.Set("native_lb_name", cloudLBMap["nativeLBName"])
	d.Set("package_model", cloudLBMap["packageModel"])
	d.Set("prom_mode", cloudLBMap["promMode"])
	d.Set("scheme", cloudLBMap["scheme"])
	d.Set("size", cloudLBMap["size"])
	d.Set("sku", cloudLBMap["sku"])
	d.Set("svc_type", cloudLBMap["svcType"])
	d.Set("target_mode", cloudLBMap["targetMode"])
	d.Set("trunking", cloudLBMap["trunking"])
	d.Set("cloud_l4_l7_lb_type", cloudLBMap["type"])
	return d, nil
}

// func getAndSetCloudL4L7LBRelationalAttributes(client *client.Client, dn string, d *schema.ResourceData) {

// 	log.Printf("[DEBUG] cloudRsLDevToCloudSubnet: Beginning Read")
// 	cloudRsLDevToCloudSubnetData, err := client.ReadRelationcloudRsLDevToCloudSubnet(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation cloudRsLDevToCloudSubnet %v", err)
// 		d.Set("relation_cloudrs_ldev_to_cloud_subnet", make([]string, 0, 1))
// 	} else {
// 		cloudRsLDevToCloudSubnetResultData := make([]string, 0, 1)
// 		for _, obj := range cloudRsLDevToCloudSubnetData.([]map[string]string) {
// 			cloudRsLDevToCloudSubnetResultData = append(cloudRsLDevToCloudSubnetResultData, obj["tDn"])
// 		}
// 		sort.Strings(cloudRsLDevToCloudSubnetResultData)
// 		d.Set("relation_cloudrs_ldev_to_cloud_subnet", cloudRsLDevToCloudSubnetResultData)
// 		log.Printf("[DEBUG]: cloudRsLDevToCloudSubnet: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] cloudRsLDevToComputePol: Beginning Read")

// 	cloudRsLDevToComputePolData, err := client.ReadRelationcloudRsLDevToComputePol(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation cloudRsLDevToComputePol %v", err)
// 		d.Set("relation_cloudrs_ldev_to_compute_pol", "")
// 	} else {
// 		d.Set("relation_cloudrs_ldev_to_compute_pol", cloudRsLDevToComputePolData["tDn"])
// 		log.Printf("[DEBUG]: cloudRsLDevToComputePol: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] cloudRsLDevToMgmtPol: Beginning Read")

// 	cloudRsLDevToMgmtPolData, err := client.ReadRelationcloudRsLDevToMgmtPol(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation cloudRsLDevToMgmtPol %v", err)
// 		d.Set("relation_cloudrs_ldev_to_mgmt_pol", "")
// 	} else {
// 		d.Set("relation_cloudrs_ldev_to_mgmt_pol", cloudRsLDevToMgmtPolData["tDn"])
// 		log.Printf("[DEBUG]: cloudRsLDevToMgmtPol: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsALDevToDevMgr: Beginning Read")

// 	vnsRsALDevToDevMgrData, err := client.ReadRelationvnsRsALDevToDevMgr(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsALDevToDevMgr %v", err)
// 		d.Set("relation_vnsrs_aldev_to_dev_mgr", "")
// 	} else {
// 		d.Set("relation_vnsrs_aldev_to_dev_mgr", vnsRsALDevToDevMgrData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsALDevToDevMgr: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsALDevToDomP: Beginning Read")

// 	vnsRsALDevToDomPData, err := client.ReadRelationvnsRsALDevToDomP(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsALDevToDomP %v", err)
// 		d.Set("relation_vnsrs_aldev_to_dom_p", make([]interface{}, 0, 1))
// 	} else {
// 		vnsRsALDevToDomPResultData := make([]map[string]string, 0, 1)
// 		for _, obj := range vnsRsALDevToDomPData.([]map[string]string) {
// 			vnsRsALDevToDomPResultData = append(vnsRsALDevToDomPResultData, map[string]string{
// 				"switching_mode": obj["switchingMode"],
// 				"target_dn":      obj["tDn"],
// 			})
// 		}
// 		d.Set("relation_vnsrs_aldev_to_dom_p", vnsRsALDevToDomPResultData)
// 		log.Printf("[DEBUG]: vnsRsALDevToDomP: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsALDevToPhysDomP: Beginning Read")

// 	vnsRsALDevToPhysDomPData, err := client.ReadRelationvnsRsALDevToPhysDomP(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsALDevToPhysDomP %v", err)
// 		d.Set("relation_vnsrs_aldev_to_phys_dom_p", "")
// 	} else {
// 		d.Set("relation_vnsrs_aldev_to_phys_dom_p", vnsRsALDevToPhysDomPData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsALDevToPhysDomP: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsALDevToVlanInstP: Beginning Read")
// 	vnsRsALDevToVlanInstPData, err := client.ReadRelationvnsRsALDevToVlanInstP(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsALDevToVlanInstP %v", err)
// 		d.Set("relation_vnsrs_aldev_to_vlan_inst_p", make([]string, 0, 1))
// 	} else {
// 		vnsRsALDevToVlanInstPResultData := make([]string, 0, 1)
// 		for _, obj := range vnsRsALDevToVlanInstPData.([]map[string]string) {
// 			vnsRsALDevToVlanInstPResultData = append(vnsRsALDevToVlanInstPResultData, obj["tDn"])
// 		}
// 		sort.Strings(vnsRsALDevToVlanInstPResultData)
// 		d.Set("relation_vnsrs_aldev_to_vlan_inst_p", vnsRsALDevToVlanInstPResultData)
// 		log.Printf("[DEBUG]: vnsRsALDevToVlanInstP: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsALDevToVxlanInstP: Beginning Read")

// 	vnsRsALDevToVxlanInstPData, err := client.ReadRelationvnsRsALDevToVxlanInstP(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsALDevToVxlanInstP %v", err)
// 		d.Set("relation_vnsrs_aldev_to_vxlan_inst_p", "")
// 	} else {
// 		d.Set("relation_vnsrs_aldev_to_vxlan_inst_p", vnsRsALDevToVxlanInstPData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsALDevToVxlanInstP: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsDevEpg: Beginning Read")

// 	vnsRsDevEpgData, err := client.ReadRelationvnsRsDevEpg(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsDevEpg %v", err)
// 		d.Set("relation_vnsrs_dev_epg", "")
// 	} else {
// 		d.Set("relation_vnsrs_dev_epg", vnsRsDevEpgData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsDevEpg: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsMDevAtt: Beginning Read")

// 	vnsRsMDevAttData, err := client.ReadRelationvnsRsMDevAtt(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsMDevAtt %v", err)
// 		d.Set("relation_vnsrs_mdev_att", "")
// 	} else {
// 		d.Set("relation_vnsrs_mdev_att", vnsRsMDevAttData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsMDevAtt: Reading finished successfully")
// 	}

// 	log.Printf("[DEBUG] vnsRsSvcMgmtEpg: Beginning Read")

// 	vnsRsSvcMgmtEpgData, err := client.ReadRelationvnsRsSvcMgmtEpg(dn)
// 	if err != nil {
// 		log.Printf("[DEBUG] Error while reading relation vnsRsSvcMgmtEpg %v", err)
// 		d.Set("relation_vnsrs_svc_mgmt_epg", "")
// 	} else {
// 		d.Set("relation_vnsrs_svc_mgmt_epg", vnsRsSvcMgmtEpgData["tDn"])
// 		log.Printf("[DEBUG]: vnsRsSvcMgmtEpg: Reading finished successfully")
// 	}
// }

func resourceAciCloudL4L7DevicesImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	cloudLB, err := getRemoteCloudL4L7LB(aciClient, dn)
	if err != nil {
		return nil, err
	}
	schemaFilled, err := setCloudL4L7LBAttributes(cloudLB, d)
	if err != nil {
		return nil, err
	}

	// // Get and Set Relational Attributes
	// getAndSetCloudL4L7LBRelationalAttributes(aciClient, dn, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudL4L7DevicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Cloud L4-L7 Device: Beginning Creation")

	aciClient := m.(*client.Client)

	getCloudL4L7LogicalInterfaceContainer(d, aciClient)

	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	TenantDn := d.Get("tenant_dn").(string)

	cloudLBAttr := models.CloudL4L7LBAttributes{}

	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudLBAttr.Annotation = Annotation.(string)
	} else {
		cloudLBAttr.Annotation = "{}"
	}

	if Version, ok := d.GetOk("version"); ok {
		cloudLBAttr.Version = Version.(string)
	}

	if ActiveActive, ok := d.GetOk("active_active"); ok {
		cloudLBAttr.ActiveActive = ActiveActive.(string)
	}

	if AllowAll, ok := d.GetOk("allow_all"); ok {
		cloudLBAttr.AllowAll = AllowAll.(string)
	}

	if AutoScaling, ok := d.GetOk("auto_scaling"); ok {
		cloudLBAttr.AutoScaling = AutoScaling.(string)
	}

	if ContextAware, ok := d.GetOk("context_aware"); ok {
		cloudLBAttr.ContextAware = ContextAware.(string)
	}

	if CustomRG, ok := d.GetOk("custom_rg"); ok {
		cloudLBAttr.CustomRG = CustomRG.(string)
	}

	if DeviceType, ok := d.GetOk("device_type"); ok {
		cloudLBAttr.DevType = DeviceType.(string)
	}

	if FunctionType, ok := d.GetOk("function_type"); ok {
		cloudLBAttr.FuncType = FunctionType.(string)
	}

	if InstanceCount, ok := d.GetOk("instance_count"); ok {
		cloudLBAttr.InstanceCount = InstanceCount.(string)
	}

	if IsCopy, ok := d.GetOk("is_copy"); ok {
		cloudLBAttr.IsCopy = IsCopy.(string)
	}

	if IsInstantiation, ok := d.GetOk("is_instantiation"); ok {
		cloudLBAttr.IsInstantiation = IsInstantiation.(string)
	}

	if IsStaticIP, ok := d.GetOk("is_static_ip"); ok {
		cloudLBAttr.IsStaticIP = IsStaticIP.(string)
	}

	if L4L7DeviceApplicationSecurityGroup, ok := d.GetOk("l4_l7_device_application_security_group"); ok {
		cloudLBAttr.L4L7DeviceApplicationSecurityGroup = L4L7DeviceApplicationSecurityGroup.(string)
	}

	if L4L7ThirdPartyDevice, ok := d.GetOk("l4_l7_third_party_device"); ok {
		cloudLBAttr.L4L7ThirdPartyDevice = L4L7ThirdPartyDevice.(string)
	}

	if Managed, ok := d.GetOk("managed"); ok {
		cloudLBAttr.Managed = Managed.(string)
	}

	if MaxInstanceCount, ok := d.GetOk("max_instance_count"); ok {
		cloudLBAttr.MaxInstanceCount = MaxInstanceCount.(string)
	}

	if MinInstanceCount, ok := d.GetOk("min_instance_count"); ok {
		cloudLBAttr.MinInstanceCount = MinInstanceCount.(string)
	}

	if Mode, ok := d.GetOk("mode"); ok {
		cloudLBAttr.Mode = Mode.(string)
	}

	if Name, ok := d.GetOk("name"); ok {
		cloudLBAttr.Name = Name.(string)
	}

	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudLBAttr.NameAlias = NameAlias.(string)
	}

	if NativeLBName, ok := d.GetOk("native_lb_name"); ok {
		cloudLBAttr.NativeLBName = NativeLBName.(string)
	}

	if PackageModel, ok := d.GetOk("package_model"); ok {
		cloudLBAttr.PackageModel = PackageModel.(string)
	}

	if PromMode, ok := d.GetOk("prom_mode"); ok {
		cloudLBAttr.PromMode = PromMode.(string)
	}

	if Scheme, ok := d.GetOk("scheme"); ok {
		cloudLBAttr.Scheme = Scheme.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		cloudLBAttr.Size = Size.(string)
	}

	if Sku, ok := d.GetOk("sku"); ok {
		cloudLBAttr.Sku = Sku.(string)
	}

	if SvcType, ok := d.GetOk("svc_type"); ok {
		cloudLBAttr.SvcType = SvcType.(string)
	}

	if TargetMode, ok := d.GetOk("target_mode"); ok {
		cloudLBAttr.TargetMode = TargetMode.(string)
	}

	if Trunking, ok := d.GetOk("trunking"); ok {
		cloudLBAttr.Trunking = Trunking.(string)
	}

	if CloudL4L7LB_type, ok := d.GetOk("cloud_l4_l7_lb_type"); ok {
		cloudLBAttr.CloudL4L7LB_type = CloudL4L7LB_type.(string)
	}
	cloudLB := models.NewCloudL4L7LB(fmt.Sprintf(models.RnCloudLB, name), TenantDn, desc, cloudLBAttr)

	err := aciClient.Save(cloudLB)
	if err != nil {
		return diag.FromErr(err)
	}
	// checkDns := make([]string, 0, 1)

	// if relationTocloudRsLDevToCloudSubnet, ok := d.GetOk("relation_cloudrs_ldev_to_cloud_subnet"); ok {
	// 	relationParamList := toStringList(relationTocloudRsLDevToCloudSubnet.(*schema.Set).List())
	// 	for _, relationParam := range relationParamList {
	// 		checkDns = append(checkDns, relationParam)
	// 	}
	// }

	// if relationTocloudRsLDevToComputePol, ok := d.GetOk("relation_cloudrs_ldev_to_compute_pol"); ok {
	// 	relationParam := relationTocloudRsLDevToComputePol.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTocloudRsLDevToMgmtPol, ok := d.GetOk("relation_cloudrs_ldev_to_mgmt_pol"); ok {
	// 	relationParam := relationTocloudRsLDevToMgmtPol.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsALDevToDevMgr, ok := d.GetOk("relation_vnsrs_aldev_to_dev_mgr"); ok {
	// 	relationParam := relationTovnsRsALDevToDevMgr.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsALDevToDomP, ok := d.GetOk("relation_vnsrs_aldev_to_dom_p"); ok {
	// 	relationParamMap := toStringList(relationTovnsRsALDevToDomP.(map[string]interface{}))
	// 	for _, relationParam := range relationParamMap {
	// 		checkDns = append(checkDns, relationParam)
	// 	}
	// }

	// if relationTovnsRsALDevToPhysDomP, ok := d.GetOk("relation_vnsrs_aldev_to_phys_dom_p"); ok {
	// 	relationParam := relationTovnsRsALDevToPhysDomP.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsALDevToVlanInstP, ok := d.GetOk("relation_vnsrs_aldev_to_vlan_inst_p"); ok {
	// 	relationParamList := toStringList(relationTovnsRsALDevToVlanInstP.(*schema.Set).List())
	// 	for _, relationParam := range relationParamList {
	// 		checkDns = append(checkDns, relationParam)
	// 	}
	// }

	// if relationTovnsRsALDevToVxlanInstP, ok := d.GetOk("relation_vnsrs_aldev_to_vxlan_inst_p"); ok {
	// 	relationParam := relationTovnsRsALDevToVxlanInstP.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsDevEpg, ok := d.GetOk("relation_vnsrs_dev_epg"); ok {
	// 	relationParam := relationTovnsRsDevEpg.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsMDevAtt, ok := d.GetOk("relation_vnsrs_mdev_att"); ok {
	// 	relationParam := relationTovnsRsMDevAtt.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// if relationTovnsRsSvcMgmtEpg, ok := d.GetOk("relation_vnsrs_svc_mgmt_epg"); ok {
	// 	relationParam := relationTovnsRsSvcMgmtEpg.(string)
	// 	checkDns = append(checkDns, relationParam)

	// }

	// d.Partial(true)
	// err = checkTDn(aciClient, checkDns)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }
	// d.Partial(false)

	// if relationTocloudRsLDevToCloudSubnet, ok := d.GetOk("relation_cloudrs_ldev_to_cloud_subnet"); ok {
	// 	relationParamList := toStringList(relationTocloudRsLDevToCloudSubnet.(*schema.Set).List())
	// 	for _, relationParam := range relationParamList {
	// 		err = aciClient.CreateRelationcloudRsLDevToCloudSubnet(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// }

	// if relationTocloudRsLDevToComputePol, ok := d.GetOk("relation_cloudrs_ldev_to_compute_pol"); ok {
	// 	relationParam := relationTocloudRsLDevToComputePol.(string)
	// 	err = aciClient.CreateRelationcloudRsLDevToComputePol(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTocloudRsLDevToMgmtPol, ok := d.GetOk("relation_cloudrs_ldev_to_mgmt_pol"); ok {
	// 	relationParam := relationTocloudRsLDevToMgmtPol.(string)
	// 	err = aciClient.CreateRelationcloudRsLDevToMgmtPol(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsALDevToDevMgr, ok := d.GetOk("relation_vnsrs_aldev_to_dev_mgr"); ok {
	// 	relationParam := relationTovnsRsALDevToDevMgr.(string)
	// 	err = aciClient.CreateRelationvnsRsALDevToDevMgr(cloudLB.DistinguishedName, cloudLBAttr.Annotation, GetMOName(relationParam))
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsALDevToDomP, ok := d.GetOk("relation_vnsrs_aldev_to_dom_p"); ok {
	// 	relationParamMap := relationTovnsRsALDevToDomP.(map[string]interface{})
	// 	err = aciClient.CreateRelationvnsRsALDevToDomP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParamMap["switching_mode"].(string), relationParamMap["target_dn"].(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// }

	// if relationTovnsRsALDevToPhysDomP, ok := d.GetOk("relation_vnsrs_aldev_to_phys_dom_p"); ok {
	// 	relationParam := relationTovnsRsALDevToPhysDomP.(string)
	// 	err = aciClient.CreateRelationvnsRsALDevToPhysDomP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsALDevToVlanInstP, ok := d.GetOk("relation_vnsrs_aldev_to_vlan_inst_p"); ok {
	// 	relationParamList := toStringList(relationTovnsRsALDevToVlanInstP.(*schema.Set).List())
	// 	for _, relationParam := range relationParamList {
	// 		err = aciClient.CreateRelationvnsRsALDevToVlanInstP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// }

	// if relationTovnsRsALDevToVxlanInstP, ok := d.GetOk("relation_vnsrs_aldev_to_vxlan_inst_p"); ok {
	// 	relationParam := relationTovnsRsALDevToVxlanInstP.(string)
	// 	err = aciClient.CreateRelationvnsRsALDevToVxlanInstP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsDevEpg, ok := d.GetOk("relation_vnsrs_dev_epg"); ok {
	// 	relationParam := relationTovnsRsDevEpg.(string)
	// 	err = aciClient.CreateRelationvnsRsDevEpg(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsMDevAtt, ok := d.GetOk("relation_vnsrs_mdev_att"); ok {
	// 	relationParam := relationTovnsRsMDevAtt.(string)
	// 	err = aciClient.CreateRelationvnsRsMDevAtt(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	// if relationTovnsRsSvcMgmtEpg, ok := d.GetOk("relation_vnsrs_svc_mgmt_epg"); ok {
	// 	relationParam := relationTovnsRsSvcMgmtEpg.(string)
	// 	err = aciClient.CreateRelationvnsRsSvcMgmtEpg(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relationParam)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	d.SetId(cloudLB.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciCloudL4L7DevicesRead(ctx, d, m)
}
func resourceAciCloudL4L7DevicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Cloud L4-L7 Device: Beginning Update")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	TenantDn := d.Get("tenant_dn").(string)

	cloudLBAttr := models.CloudL4L7LBAttributes{}

	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudLBAttr.Annotation = Annotation.(string)
	} else {
		cloudLBAttr.Annotation = "{}"
	}

	if Version, ok := d.GetOk("version"); ok {
		cloudLBAttr.Version = Version.(string)
	}

	if ActiveActive, ok := d.GetOk("active_active"); ok {
		cloudLBAttr.ActiveActive = ActiveActive.(string)
	}

	if AllowAll, ok := d.GetOk("allow_all"); ok {
		cloudLBAttr.AllowAll = AllowAll.(string)
	}

	if AutoScaling, ok := d.GetOk("auto_scaling"); ok {
		cloudLBAttr.AutoScaling = AutoScaling.(string)
	}

	if ContextAware, ok := d.GetOk("context_aware"); ok {
		cloudLBAttr.ContextAware = ContextAware.(string)
	}

	if CustomRG, ok := d.GetOk("custom_rg"); ok {
		cloudLBAttr.CustomRG = CustomRG.(string)
	}

	if DeviceType, ok := d.GetOk("device_type"); ok {
		cloudLBAttr.DevType = DeviceType.(string)
	}

	if FunctionType, ok := d.GetOk("function_type"); ok {
		cloudLBAttr.FuncType = FunctionType.(string)
	}

	if InstanceCount, ok := d.GetOk("instance_count"); ok {
		cloudLBAttr.InstanceCount = InstanceCount.(string)
	}

	if IsCopy, ok := d.GetOk("is_copy"); ok {
		cloudLBAttr.IsCopy = IsCopy.(string)
	}

	if IsInstantiation, ok := d.GetOk("is_instantiation"); ok {
		cloudLBAttr.IsInstantiation = IsInstantiation.(string)
	}

	if IsStaticIP, ok := d.GetOk("is_static_ip"); ok {
		cloudLBAttr.IsStaticIP = IsStaticIP.(string)
	}

	if L4L7DeviceApplicationSecurityGroup, ok := d.GetOk("l4_l7_device_application_security_group"); ok {
		cloudLBAttr.L4L7DeviceApplicationSecurityGroup = L4L7DeviceApplicationSecurityGroup.(string)
	}

	if L4L7ThirdPartyDevice, ok := d.GetOk("l4_l7_third_party_device"); ok {
		cloudLBAttr.L4L7ThirdPartyDevice = L4L7ThirdPartyDevice.(string)
	}

	if Managed, ok := d.GetOk("managed"); ok {
		cloudLBAttr.Managed = Managed.(string)
	}

	if MaxInstanceCount, ok := d.GetOk("max_instance_count"); ok {
		cloudLBAttr.MaxInstanceCount = MaxInstanceCount.(string)
	}

	if MinInstanceCount, ok := d.GetOk("min_instance_count"); ok {
		cloudLBAttr.MinInstanceCount = MinInstanceCount.(string)
	}

	if Mode, ok := d.GetOk("mode"); ok {
		cloudLBAttr.Mode = Mode.(string)
	}

	if Name, ok := d.GetOk("name"); ok {
		cloudLBAttr.Name = Name.(string)
	}

	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudLBAttr.NameAlias = NameAlias.(string)
	}

	if NativeLBName, ok := d.GetOk("native_lb_name"); ok {
		cloudLBAttr.NativeLBName = NativeLBName.(string)
	}

	if PackageModel, ok := d.GetOk("package_model"); ok {
		cloudLBAttr.PackageModel = PackageModel.(string)
	}

	if PromMode, ok := d.GetOk("prom_mode"); ok {
		cloudLBAttr.PromMode = PromMode.(string)
	}

	if Scheme, ok := d.GetOk("scheme"); ok {
		cloudLBAttr.Scheme = Scheme.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		cloudLBAttr.Size = Size.(string)
	}

	if Sku, ok := d.GetOk("sku"); ok {
		cloudLBAttr.Sku = Sku.(string)
	}

	if SvcType, ok := d.GetOk("svc_type"); ok {
		cloudLBAttr.SvcType = SvcType.(string)
	}

	if TargetMode, ok := d.GetOk("target_mode"); ok {
		cloudLBAttr.TargetMode = TargetMode.(string)
	}

	if Trunking, ok := d.GetOk("trunking"); ok {
		cloudLBAttr.Trunking = Trunking.(string)
	}

	if CloudL4L7LB_type, ok := d.GetOk("cloud_l4_l7_lb_type"); ok {
		cloudLBAttr.CloudL4L7LB_type = CloudL4L7LB_type.(string)
	}
	cloudLB := models.NewCloudL4L7LB(fmt.Sprintf(models.RnCloudLB, name), TenantDn, desc, cloudLBAttr)

	cloudLB.Status = "modified"

	err := aciClient.Save(cloudLB)
	if err != nil {
		return diag.FromErr(err)
	}

	// checkDns := make([]string, 0, 1)

	// if d.HasChange("relation_cloudrs_ldev_to_cloud_subnet") || d.HasChange("annotation") {
	// 	oldRel, newRel := d.GetChange("relation_cloudrs_ldev_to_cloud_subnet")
	// 	oldRelSet := oldRel.(*schema.Set)
	// 	newRelSet := newRel.(*schema.Set)
	// 	relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())
	// 	for _, relDn := range relToCreate {
	// 		checkDns = append(checkDns, relDn)
	// 	}
	// }

	// if d.HasChange("relation_cloudrs_ldev_to_compute_pol") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_cloudrs_ldev_to_compute_pol")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_cloudrs_ldev_to_mgmt_pol") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_cloudrs_ldev_to_mgmt_pol")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_aldev_to_dev_mgr") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_dev_mgr")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_aldev_to_dom_p") || d.HasChange("annotation") {
	// 	oldRel, newRel := d.GetChange("relation_vnsrs_aldev_to_dom_p")
	// 	oldRelSet := oldRel.(*schema.Set)
	// 	newRelSet := newRel.(*schema.Set)
	// 	relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())
	// 	for _, relDn := range relToCreate {
	// 		checkDns = append(checkDns, relDn)
	// 	}
	// }

	// if d.HasChange("relation_vnsrs_aldev_to_phys_dom_p") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_phys_dom_p")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_aldev_to_vlan_inst_p") || d.HasChange("annotation") {
	// 	oldRel, newRel := d.GetChange("relation_vnsrs_aldev_to_vlan_inst_p")
	// 	oldRelSet := oldRel.(*schema.Set)
	// 	newRelSet := newRel.(*schema.Set)
	// 	relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())
	// 	for _, relDn := range relToCreate {
	// 		checkDns = append(checkDns, relDn)
	// 	}
	// }

	// if d.HasChange("relation_vnsrs_aldev_to_vxlan_inst_p") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_vxlan_inst_p")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_dev_epg") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_dev_epg")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_mdev_att") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_mdev_att")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// if d.HasChange("relation_vnsrs_svc_mgmt_epg") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_svc_mgmt_epg")
	// 	checkDns = append(checkDns, newRelParam.(string))

	// }

	// d.Partial(true)
	// err = checkTDn(aciClient, checkDns)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }
	// d.Partial(false)

	// if d.HasChange("relation_cloudrs_ldev_to_cloud_subnet") || d.HasChange("annotation") {
	// 	oldRel, newRel := d.GetChange("relation_cloudrs_ldev_to_cloud_subnet")
	// 	oldRelSet := oldRel.(*schema.Set)
	// 	newRelSet := newRel.(*schema.Set)
	// 	relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
	// 	relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

	// 	for _, relDn := range relToDelete {
	// 		err = aciClient.DeleteRelationcloudRsLDevToCloudSubnet(cloudLB.DistinguishedName, relDn)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// 	for _, relDn := range relToCreate {
	// 		err = aciClient.CreateRelationcloudRsLDevToCloudSubnet(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relDn)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// }
	// if d.HasChange("relation_cloudrs_ldev_to_compute_pol") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_cloudrs_ldev_to_compute_pol")
	// 	err = aciClient.DeleteRelationcloudRsLDevToComputePol(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationcloudRsLDevToComputePol(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_cloudrs_ldev_to_mgmt_pol") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_cloudrs_ldev_to_mgmt_pol")
	// 	err = aciClient.DeleteRelationcloudRsLDevToMgmtPol(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationcloudRsLDevToMgmtPol(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_aldev_to_dev_mgr") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_dev_mgr")
	// 	err = aciClient.DeleteRelationvnsRsALDevToDevMgr(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsALDevToDevMgr(cloudLB.DistinguishedName, cloudLBAttr.Annotation, GetMOName(newRelParam.(string)))
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_aldev_to_dom_p") || d.HasChange("annotation") {
	// 	oldRelParam, newRelParam := d.GetChange("relation_vnsrs_aldev_to_dom_p")
	// 	oldRelParamMap := oldRelParam.(map[string]interface{})
	// 	newRelParamMap := newRelParam.(map[string]interface{})
	// 	err = aciClient.DeleteRelationvnsRsALDevToDomP(cloudLB.DistinguishedName)

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsALDevToDomP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParamMap["switching_mode"].(string), newRelParamMap["target_dn"].(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// }
	// if d.HasChange("relation_vnsrs_aldev_to_phys_dom_p") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_phys_dom_p")
	// 	err = aciClient.DeleteRelationvnsRsALDevToPhysDomP(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsALDevToPhysDomP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_aldev_to_vlan_inst_p") || d.HasChange("annotation") {
	// 	oldRel, newRel := d.GetChange("relation_vnsrs_aldev_to_vlan_inst_p")
	// 	oldRelSet := oldRel.(*schema.Set)
	// 	newRelSet := newRel.(*schema.Set)
	// 	relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
	// 	relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

	// 	for _, relDn := range relToDelete {
	// 		err = aciClient.DeleteRelationvnsRsALDevToVlanInstP(cloudLB.DistinguishedName, relDn)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// 	for _, relDn := range relToCreate {
	// 		err = aciClient.CreateRelationvnsRsALDevToVlanInstP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, relDn)

	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// }
	// if d.HasChange("relation_vnsrs_aldev_to_vxlan_inst_p") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_aldev_to_vxlan_inst_p")
	// 	err = aciClient.DeleteRelationvnsRsALDevToVxlanInstP(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsALDevToVxlanInstP(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_dev_epg") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_dev_epg")
	// 	err = aciClient.DeleteRelationvnsRsDevEpg(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsDevEpg(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_mdev_att") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_mdev_att")
	// 	err = aciClient.DeleteRelationvnsRsMDevAtt(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsMDevAtt(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }
	// if d.HasChange("relation_vnsrs_svc_mgmt_epg") || d.HasChange("annotation") {
	// 	_, newRelParam := d.GetChange("relation_vnsrs_svc_mgmt_epg")
	// 	err = aciClient.DeleteRelationvnsRsSvcMgmtEpg(cloudLB.DistinguishedName)
	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// 	err = aciClient.CreateRelationvnsRsSvcMgmtEpg(cloudLB.DistinguishedName, cloudLBAttr.Annotation, newRelParam.(string))

	// 	if err != nil {
	// 		return diag.FromErr(err)
	// 	}

	// }

	d.SetId(cloudLB.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciCloudL4L7DevicesRead(ctx, d, m)
}

func resourceAciCloudL4L7DevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	getCloudL4L7LogicalInterfaceContainer(d, aciClient)

	cloudLB, err := getRemoteCloudL4L7LB(aciClient, dn)
	if err != nil {
		d.SetId("")
		return nil
	}

	_, err = setCloudL4L7LBAttributes(cloudLB, d)
	if err != nil {
		d.SetId("")
		return nil
	}

	// // Get and Set Relational Attributes
	// getAndSetCloudL4L7LBRelationalAttributes(aciClient, dn, d)

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciCloudL4L7DevicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()

	err := aciClient.DeleteByDn(dn, models.CloudLBClassName)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return diag.FromErr(err)
}

// New Code part

func getCloudL4L7LogicalInterfaceContainer(d *schema.ResourceData, aciClient *client.Client) error {
	if logicalInterfaces, ok := d.GetOk("logical_interfaces"); ok {
		logicalInterfacesList := logicalInterfaces.(*schema.Set).List()
		parsedLogicalInterfaces := make([]map[string]interface{}, 0)
		for _, logicalInterface := range logicalInterfacesList {
			logicalInterfaceMap := logicalInterface.(map[string]interface{})
			cloudLifAttrsMap := map[string]string{
				"annotation": logicalInterfaceMap["annotation"].(string),
				"name":       logicalInterfaceMap["name"].(string),
				"allowAll":   logicalInterfaceMap["allow_all"].(string),
				"nameAlias":  logicalInterfaceMap["name_alias"].(string),
			}
			endPointSelectorsList := make([]interface{}, 0)
			for _, endPointSelector := range logicalInterfaceMap["end_point_selectors"].(*schema.Set).List() {
				endPointSelectorInterface := endPointSelector.(map[string]interface{})
				endPointSelectorMap := map[string]interface{}{
					models.CloudepselectorClassName: map[string]interface{}{
						"attributes": map[string]string{
							"annotation":      endPointSelectorInterface["annotation"].(string),
							"name":            endPointSelectorInterface["name"].(string),
							"matchExpression": endPointSelectorInterface["match_expression"].(string),
							"nameAlias":       endPointSelectorInterface["name_alias"].(string),
							"descr":           endPointSelectorInterface["description"].(string),
						},
					},
				}
				endPointSelectorsList = append(endPointSelectorsList, endPointSelectorMap)
			}

			cloudLifMap := map[string]interface{}{
				models.CloudLIfClassName: map[string]interface{}{
					"attributes": cloudLifAttrsMap,
					"children":   endPointSelectorsList,
				},
			}
			parsedLogicalInterfaces = append(parsedLogicalInterfaces, cloudLifMap)
		}
		parsedLogicalInterfacesJsonBytes, err := json.Marshal(parsedLogicalInterfaces)

		BCont, _ := container.ParseJSON(parsedLogicalInterfacesJsonBytes)

		if err != nil {
			log.Printf("[DEBUG] --------------- parsedLogicalInterfacesJsonBytes err: %v", err)
		}
		log.Printf("[DEBUG] --------------- parsedLogicalInterfacesJsonBytes: %v", string(parsedLogicalInterfacesJsonBytes))
	}
	return nil
}

func getCloudL4L7LogicalInterfaceMap(d *schema.ResourceData, aciClient *client.Client) error {
	log.Printf("[DEBUG] --------------- inside getCloudL4L7LogicalInterfaceMap")
	// name := d.Get("name").(string)
	// cloudL4L7DeviceDn := d.Get("cloud_l4_l7_device_dn").(string)

	// leakRoutesCont, leakRoutesErr := container.ParseJSON(leakRoutesJson)
	// if leakRoutesErr != nil {
	// 	// return nil, fmt.Errorf("leakRoutes ParseJSON Error: %v", leakRoutesErr)
	// 	log.Printf("[DEBUG] --------------- leakRoutesErr: %v", leakRoutesErr)
	// } else {
	// 	log.Printf("[DEBUG] --------------- leakRoutesCont: %v", leakRoutesCont)
	// }

	// // A := []byte(`
	// // {
	// // 	"A": {
	// // 		"attributes": {
	// // 			"name": "AName"
	// // 		}
	// // 	}
	// // }`)

	// // ACont, _ := container.ParseJSON(A)

	// // B := []byte(`
	// // {
	// // 	"B": {
	// // 		"attributes": {
	// // 			"name": "BName"
	// // 		}
	// // 	}
	// // }`)

	// // BCont, _ := container.ParseJSON(B)

	// // C := []byte(`
	// // {
	// // 	"C": {
	// // 		"attributes": {
	// // 			"name": "CName"
	// // 		}
	// // 	}
	// // }`)

	// // CCont, _ := container.ParseJSON(C)

	// // BCont.ArrayAppend(CCont.Data(), "B", "children")

	// // AContErr := ACont.ArrayAppend(BCont.Data(), "A", "children")
	// // if AContErr != nil {
	// // 	return nil, fmt.Errorf("AContErr Error:: %v", AContErr)
	// // }

	// log.Printf("[DEBUG] ---------------- vnsAbsGraphCont final value : %v", ACont)

	// return nil

	if logicalInterfaces, ok := d.GetOk("logical_interfaces"); ok {

		logicalInterfacesList := logicalInterfaces.(*schema.Set).List()
		for _, logicalInterface := range logicalInterfacesList {
			logicalInterfaceMap := logicalInterface.(map[string]interface{})
			log.Printf("[DEBUG] --------------- logicalInterfaceMap: %v", logicalInterfaceMap)

			// var nameAlias, allowAll, annotation string
			// if AllowAll, ok := d.GetOk("allow_all"); ok {
			// 	allowAll = AllowAll.(string)
			// }

			// if NameAlias, ok := d.GetOk("name_alias"); ok {
			// 	nameAlias = NameAlias.(string)
			// }

			// annotation = "{}"
			// if Annotation, ok := d.GetOk("annotation"); ok {
			// 	annotation = Annotation.(string)
			// }

			// cloudLifMap := map[string]interface{}{
			// 	"annotation": logicalInterfaceMap["annotation"].(string),
			// 	"name":       logicalInterfaceMap["name"].(string),
			// 	"allowAll":   logicalInterfaceMap["allow_all"].(string),
			// 	"nameAlias":  logicalInterfaceMap["name_alias"].(string),
			// }

			cloudLifMap := map[string]interface{}{
				"class_name": models.CloudLIfClassName,
				"content": map[string]string{
					"annotation": logicalInterfaceMap["annotation"].(string),
					"name":       logicalInterfaceMap["name"].(string),
					"allowAll":   logicalInterfaceMap["allow_all"].(string),
					"nameAlias":  logicalInterfaceMap["name_alias"].(string),
				},
			}

			cloudEPSelectorsList := make([]interface{}, 0)
			// if endPointSelectors, ok := d.GetOk("end_point_selectors"); ok {
			// 	endPointSelectorsMap := endPointSelectors.(*schema.Set).List()
			// 	for _, endPointSelector := range endPointSelectorsMap {
			// 		endPointSelectorMap := endPointSelector.(map[string]interface{})
			// 		cloudEndPointSelectorMap := map[string]interface{}{
			// 			"class_name": models.CloudepselectorClassName,
			// 			"content": map[string]string{
			// 				"annotation":      endPointSelectorMap["annotation"].(string),
			// 				"name":            endPointSelectorMap["name"].(string),
			// 				"matchExpression": endPointSelectorMap["match_expression"].(string),
			// 				"nameAlias":       endPointSelectorMap["name_alias"].(string),
			// 				"descr":           endPointSelectorMap["description"].(string),
			// 			},
			// 		}
			// 		cloudEPSelectorsList = append(cloudEPSelectorsList, cloudEndPointSelectorMap)
			// 	}
			// }

			// endPointSelectorsMap := endPointSelectors.(*schema.Set).List()
			for _, endPointSelector := range logicalInterfaceMap["end_point_selectors"].(*schema.Set).List() {
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

			// cloudLifMap["children"] = make([]interface{}, 0)
			cloudLifMap["children"] = cloudEPSelectorsList
			log.Printf("[DEBUG] --------------- cloudLifMap: %v", cloudLifMap)

			// cloudLifCont, err := preparePayload(models.CloudLIfClassName, cloudLifMap, cloudEPSelectorsList)
			// if err != nil {
			// 	return err
			// }

			// httpRequestPayload, err := aciClient.MakeRestRequest("POST", fmt.Sprintf("%s/%s.json", client.DefaultMOURL, cloudL4L7DeviceDn), cloudLifCont, true)
			// if err != nil {
			// 	return err
			// }

			// respCont, _, err := aciClient.Do(httpRequestPayload)
			// if err != nil {
			// 	return err
			// }

			// err = client.CheckForErrors(respCont, "POST", false)
			// if err != nil {
			// 	return err
			// }

			// d.SetId(cloudL4L7DeviceDn + "/" + fmt.Sprintf(models.RnCloudLIf, name))
			return nil
		}
	}

	// leakInternalSubnetCont, leakInternalSubnetErr := preparePayload(leakInternalSubnetMap["class_name"].(string), leakInternalSubnetMap["content"].(map[string]string), leakToList)
	return nil
}
