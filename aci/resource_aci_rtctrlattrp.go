package aci

import (
	"context"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAciActionRuleProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciActionRuleProfileCreate,
		UpdateContext: resourceAciActionRuleProfileUpdate,
		ReadContext:   resourceAciActionRuleProfileRead,
		DeleteContext: resourceAciActionRuleProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciActionRuleProfileImport,
		},

		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"tenant_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"local_pref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtctrl_set_pref_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "local-pref",
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtctrl_set_tag_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "rt-tag",
			},
			"rtctrl_set_weight_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "rt-weight",
			},
			"weight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtctrl_set_rt_metric_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "metric",
			},
			"metric_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ospf-type1",
					"ospf-type2",
				}, false),
			},
			"rtctrl_set_rt_metric_type_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "metric-type",
			},
			"half_life": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_suppress_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"suppress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtctrl_set_damp_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "dampening-pol",
			},

			"community": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"set_criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"append",
					"none",
					"replace",
				}, false),
			},
			"rtctrl_set_comm_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "community",
			},

			"additional_community": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"additional_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"additional_set_criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"append",
					"none",
					"replace",
				}, false),
			},
			"rtctrl_set_add_comm_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "community",
			},

			"addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtctrl_set_nh_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "ip-nh",
			},

			"redistribute_multipath_action_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "redist-multipath",
			},

			"nexthop_unchanged_action_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "nh-unchanged",
			},

			"set_as_path_criteria": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"prepend",
					"prepend-last-as",
				}, false),
			},
			"lastnum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"set_as_path_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"as-path",
					"community",
					"dampening-pol",
					"ip-nh",
					"local-pref",
					"metric",
					"metric-type",
					"nh-unchanged",
					"ospf-fwd-addr",
					"ospf-nssa",
					"redist-multipath",
					"rt-tag",
					"rt-weight",
				}, false),
				Default: "as-path",
			},

			"set_as_path_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"asn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"order": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		})),
	}
}
func getRemoteActionRuleProfile(client *client.Client, dn string) (*models.ActionRuleProfile, error) {
	rtctrlAttrPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	rtctrlAttrP := models.ActionRuleProfileFromContainer(rtctrlAttrPCont)

	if rtctrlAttrP.DistinguishedName == "" {
		return nil, fmt.Errorf("ActionRuleProfile %s not found", rtctrlAttrP.DistinguishedName)
	}

	return rtctrlAttrP, nil
}

func setActionRuleProfileAttributes(rtctrlAttrP *models.ActionRuleProfile, d *schema.ResourceData) (*schema.ResourceData, error) {
	dn := d.Id()
	d.SetId(rtctrlAttrP.DistinguishedName)
	d.Set("description", rtctrlAttrP.Description)
	if dn != rtctrlAttrP.DistinguishedName {
		d.Set("tenant_dn", "")
	}
	rtctrlAttrPMap, err := rtctrlAttrP.ToMap()
	if err != nil {
		return d, err
	}
	d.Set("name", rtctrlAttrPMap["name"])
	d.Set("tenant_dn", GetParentDn(dn, fmt.Sprintf("/attr-%s", rtctrlAttrPMap["name"])))
	d.Set("annotation", rtctrlAttrPMap["annotation"])
	d.Set("name_alias", rtctrlAttrPMap["nameAlias"])
	return d, nil
}

func resourceAciActionRuleProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)

	dn := d.Id()

	rtctrlAttrP, err := getRemoteActionRuleProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	rtctrlAttrPMap, err := rtctrlAttrP.ToMap()
	if err != nil {
		return nil, err
	}
	name := rtctrlAttrPMap["name"]
	pDN := GetParentDn(dn, fmt.Sprintf("/attr-%s", name))
	d.Set("tenant_dn", pDN)
	schemaFilled, err := setActionRuleProfileAttributes(rtctrlAttrP, d)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())

	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciActionRuleProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] ActionRuleProfile: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	rtctrlAttrPAttr := models.ActionRuleProfileAttributes{}

	nameAlias := ""
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		nameAlias = NameAlias.(string)
	}

	if Annotation, ok := d.GetOk("annotation"); ok {
		rtctrlAttrPAttr.Annotation = Annotation.(string)
	} else {
		rtctrlAttrPAttr.Annotation = "{}"
	}
	if Name, ok := d.GetOk("name"); ok {
		rtctrlAttrPAttr.Name = Name.(string)
	}
	rtctrlAttrP := models.NewActionRuleProfile(fmt.Sprintf(models.RnrtctrlAttrP, name), TenantDn, desc, nameAlias, rtctrlAttrPAttr)

	err := aciClient.Save(rtctrlAttrP)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rtctrlAttrP.DistinguishedName)

	// rtctrlSetPrefAttr := models.RtctrlSetPrefAttributes{}

	// if LocalPref, ok := d.GetOk("local_pref"); ok {
	// 	rtctrlSetPrefAttr.LocalPref = LocalPref.(string)
	// }

	// if RtctrlSetPref_type, ok := d.GetOk("rtctrl_set_pref_type"); ok {
	// 	rtctrlSetPrefAttr.RtctrlSetPref_type = RtctrlSetPref_type.(string)
	// }

	// rtctrlSetPref := models.NewRtctrlSetPref(fmt.Sprintf(models.RnrtctrlSetPref), d.Id(), "", "", rtctrlSetPrefAttr)

	// err1 := aciClient.Save(rtctrlSetPref)
	// if err1 != nil {
	// 	return diag.FromErr(err1)
	// }

	// rtctrlSetTagAttr := models.RtctrlSetTagAttributes{}

	// if Tag, ok := d.GetOk("tag"); ok {
	// 	rtctrlSetTagAttr.Tag = Tag.(string)
	// }

	// if RtctrlSetTag_type, ok := d.GetOk("rtctrl_set_tag_type"); ok {
	// 	rtctrlSetTagAttr.RtctrlSetTag_type = RtctrlSetTag_type.(string)
	// }
	// rtctrlSetTagAttr.Name = "test_tag"
	// rtctrlSetTag := models.NewRtctrlSetTag(fmt.Sprintf(models.RnrtctrlSetTag), d.Id(), "desc", "nameAlias", rtctrlSetTagAttr)

	// rtctrlSetWeightAttr := models.RtctrlSetWeightAttributes{}

	// if RtctrlSetWeight_type, ok := d.GetOk("rtctrl_set_weight_type"); ok {
	// 	rtctrlSetWeightAttr.RtctrlSetWeight_type = RtctrlSetWeight_type.(string)
	// }

	// if Weight, ok := d.GetOk("weight"); ok {
	// 	rtctrlSetWeightAttr.Weight = Weight.(string)
	// }
	// rtctrlSetWeightAttr.Name = "test_name"
	// rtctrlSetWeight := models.NewRtctrlSetWeight(fmt.Sprintf(models.RnrtctrlSetWeight), d.Id(), "desc", "nameAlias", rtctrlSetWeightAttr)

	// rtctrlSetRtMetricAttr := models.RtctrlSetRtMetricAttributes{}

	// if Metric, ok := d.GetOk("metric"); ok {
	// 	rtctrlSetRtMetricAttr.Metric = Metric.(string)
	// }

	// if Name, ok := d.GetOk("name"); ok {
	// 	rtctrlSetRtMetricAttr.Name = Name.(string)
	// }

	// if RtctrlSetRtMetric_type, ok := d.GetOk("rtctrl_set_rt_metric_type"); ok {
	// 	rtctrlSetRtMetricAttr.RtctrlSetRtMetric_type = RtctrlSetRtMetric_type.(string)
	// }
	// rtctrlSetRtMetric := models.NewRtctrlSetRtMetric(fmt.Sprintf(models.RnrtctrlSetRtMetric), d.Id(), "desc", "nameAlias", rtctrlSetRtMetricAttr)

	// rtctrlSetRtMetricTypeAttr := models.RtctrlSetRtMetricTypeAttributes{}

	// if MetricType, ok := d.GetOk("metric_type"); ok {
	// 	rtctrlSetRtMetricTypeAttr.MetricType = MetricType.(string)
	// }

	// if RtctrlSetRtMetricType_type, ok := d.GetOk("rtctrl_set_rt_metric_type_type"); ok {
	// 	rtctrlSetRtMetricTypeAttr.RtctrlSetRtMetricType_type = RtctrlSetRtMetricType_type.(string)
	// }
	// rtctrlSetRtMetricTypeAttr.Name = "test_metric_type"
	// rtctrlSetRtMetricType := models.NewRtctrlSetRtMetricType(fmt.Sprintf(models.RnrtctrlSetRtMetricType), d.Id(), "desc", "nameAlias", rtctrlSetRtMetricTypeAttr)

	// rtctrlSetDampAttr := models.RtctrlSetDampAttributes{}

	// if HalfLife, ok := d.GetOk("half_life"); ok {
	// 	rtctrlSetDampAttr.HalfLife = HalfLife.(string)
	// }

	// if MaxSuppressTime, ok := d.GetOk("max_suppress_time"); ok {
	// 	rtctrlSetDampAttr.MaxSuppressTime = MaxSuppressTime.(string)
	// }

	// if Reuse, ok := d.GetOk("reuse"); ok {
	// 	rtctrlSetDampAttr.Reuse = Reuse.(string)
	// }

	// if Suppress, ok := d.GetOk("suppress"); ok {
	// 	rtctrlSetDampAttr.Suppress = Suppress.(string)
	// }

	// if RtctrlSetDamp_type, ok := d.GetOk("rtctrl_set_damp_type"); ok {
	// 	rtctrlSetDampAttr.RtctrlSetDamp_type = RtctrlSetDamp_type.(string)
	// }

	// rtctrlSetDampAttr.Name = d.Get("name").(string)

	// rtctrlSetDamp := models.NewRtctrlSetDamp(fmt.Sprintf(models.RnrtctrlSetDamp), d.Id(), "desc", "nameAlias", rtctrlSetDampAttr)

	// rtctrlSetCommAttr := models.RtctrlSetCommAttributes{}

	// if Community, ok := d.GetOk("community"); ok {
	// 	rtctrlSetCommAttr.Community = Community.(string)
	// }

	// if SetCriteria, ok := d.GetOk("set_criteria"); ok {
	// 	rtctrlSetCommAttr.SetCriteria = SetCriteria.(string)
	// }

	// if RtctrlSetComm_type, ok := d.GetOk("rtctrl_set_comm_type"); ok {
	// 	rtctrlSetCommAttr.RtctrlSetComm_type = RtctrlSetComm_type.(string)
	// }

	// rtctrlSetCommAttr.Name = d.Get("name").(string)

	// rtctrlSetComm := models.NewRtctrlSetComm(fmt.Sprintf(models.RnrtctrlSetComm), d.Id(), "desc", "nameAlias", rtctrlSetCommAttr)

	// additional_community := d.Get("additional_community").(string)

	// rtctrlSetAddCommAttr := models.RtctrlSetAddCommAttributes{}

	// if Community, ok := d.GetOk("additional_community"); ok {
	// 	rtctrlSetAddCommAttr.Community = Community.(string)
	// }

	// if SetCriteria, ok := d.GetOk("additional_set_criteria"); ok {
	// 	rtctrlSetAddCommAttr.SetCriteria = SetCriteria.(string)
	// }

	// if RtctrlSetAddComm_type, ok := d.GetOk("rtctrl_set_add_comm_type"); ok {
	// 	rtctrlSetAddCommAttr.RtctrlSetAddComm_type = RtctrlSetAddComm_type.(string)
	// }

	// rtctrlSetAddCommAttr.Name = d.Get("name").(string)

	// rtctrlSetAddComm := models.NewRtctrlSetAddComm(fmt.Sprintf(models.RnrtctrlSetAddComm, additional_community), d.Id(), "desc", "nameAlias", rtctrlSetAddCommAttr)

	// rtctrlSetNhAttr := models.RtctrlSetNhAttributes{}

	// if Addr, ok := d.GetOk("addr"); ok {
	// 	rtctrlSetNhAttr.Addr = Addr.(string)
	// }

	// rtctrlSetNhAttr.Name = d.Get("name").(string)

	// if RtctrlSetNh_type, ok := d.GetOk("rtctrl_set_nh_type"); ok {
	// 	rtctrlSetNhAttr.RtctrlSetNh_type = RtctrlSetNh_type.(string)
	// }
	// rtctrlSetNh := models.NewRtctrlSetNh(fmt.Sprintf(models.RnrtctrlSetNh), d.Id(), "desc", "nameAlias", rtctrlSetNhAttr)

	// rtctrlSetNhUnchangedAttr := models.NexthopUnchangedActionAttributes{}

	// if NexthopUnchangedAction_type, ok := d.GetOk("nexthop_unchanged_action_type"); ok {
	// 	rtctrlSetNhUnchangedAttr.NexthopUnchangedAction_type = NexthopUnchangedAction_type.(string)
	// }
	// rtctrlSetNhUnchanged := models.NewNexthopUnchangedAction(fmt.Sprintf(models.RnrtctrlSetNhUnchanged), d.Id(), "desc", "nameAlias", rtctrlSetNhUnchangedAttr)

	// rtctrlSetRedistMultipathAttr := models.RedistributeMultipathActionAttributes{}

	// rtctrlSetRedistMultipathAttr.Name = d.Get("name").(string)
	// if RedistributeMultipathAction_type, ok := d.GetOk("redistribute_multipath_action_type"); ok {
	// 	rtctrlSetRedistMultipathAttr.RedistributeMultipathAction_type = RedistributeMultipathAction_type.(string)
	// }
	// rtctrlSetRedistMultipath := models.NewRedistributeMultipathAction(fmt.Sprintf(models.RnrtctrlSetRedistMultipath), d.Id(), "desc", "nameAlias", rtctrlSetRedistMultipathAttr)

	// err1 := aciClient.Save(rtctrlSetRedistMultipath)
	// if err1 != nil {
	// 	return diag.FromErr(err1)
	// }

	rtctrlSetASPathAttr := models.SetASPathAttributes{}
	set_as_path_criteria := d.Get("set_as_path_criteria").(string)

	if Criteria, ok := d.GetOk("set_as_path_criteria"); ok {
		rtctrlSetASPathAttr.Criteria = Criteria.(string)
	}

	if Lastnum, ok := d.GetOk("lastnum"); ok {
		rtctrlSetASPathAttr.Lastnum = Lastnum.(string)
	}

	if SetASPath_type, ok := d.GetOk("set_as_path_type"); ok {
		rtctrlSetASPathAttr.Type = SetASPath_type.(string)
	}

	rtctrlSetASPathAttr.Name = d.Get("name").(string)

	rtctrlSetASPath := models.NewSetASPath(fmt.Sprintf(models.RnrtctrlSetASPath, set_as_path_criteria), d.Id(), "desc", "nameAlias", rtctrlSetASPathAttr)

	err0 := aciClient.Save(rtctrlSetASPath)
	if err0 != nil {
		return diag.FromErr(err0)
	}

	SetASPathDn := rtctrlSetASPath.DistinguishedName

	rtctrlSetASPathASNAttr := models.ASNumberAttributes{}

	if Asn, ok := d.GetOk("asn"); ok {
		rtctrlSetASPathASNAttr.Asn = Asn.(string)
	}
	order := d.Get("order").(string)

	if Order, ok := d.GetOk("order"); ok {
		rtctrlSetASPathASNAttr.Order = Order.(string)
	}
	rtctrlSetASPathASN := models.NewASNumber(fmt.Sprintf(models.RnrtctrlSetASPathASN, order), SetASPathDn, "desc", "nameAlias", rtctrlSetASPathASNAttr)

	err2 := aciClient.Save(rtctrlSetASPathASN)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciActionRuleProfileRead(ctx, d, m)
}

func resourceAciActionRuleProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] ActionRuleProfile: Beginning Update")

	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	rtctrlAttrPAttr := models.ActionRuleProfileAttributes{}
	nameAlias := ""
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		nameAlias = NameAlias.(string)
	}

	if Annotation, ok := d.GetOk("annotation"); ok {
		rtctrlAttrPAttr.Annotation = Annotation.(string)
	} else {
		rtctrlAttrPAttr.Annotation = "{}"
	}

	if Name, ok := d.GetOk("name"); ok {
		rtctrlAttrPAttr.Name = Name.(string)
	}

	rtctrlAttrP := models.NewActionRuleProfile(fmt.Sprintf(models.RnrtctrlAttrP, name), TenantDn, desc, nameAlias, rtctrlAttrPAttr)

	rtctrlAttrP.Status = "modified"

	err := aciClient.Save(rtctrlAttrP)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rtctrlAttrP.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())

	// rtctrlSetPrefAttr := models.RtctrlSetPrefAttributes{}

	// if LocalPref, ok := d.GetOk("local_pref"); ok {
	// 	rtctrlSetPrefAttr.LocalPref = LocalPref.(string)
	// }

	// if RtctrlSetPref_type, ok := d.GetOk("rtctrl_set_pref_type"); ok {
	// 	rtctrlSetPrefAttr.RtctrlSetPref_type = RtctrlSetPref_type.(string)
	// }

	// rtctrlSetPref := models.NewRtctrlSetPref(fmt.Sprintf(models.RnrtctrlSetPref), d.Id(), "", "", rtctrlSetPrefAttr)

	// err1 := aciClient.Save(rtctrlSetPref)
	// if err1 != nil {
	// 	return diag.FromErr(err1)
	// }

	// rtctrlSetTagAttr := models.RtctrlSetTagAttributes{}

	// if Tag, ok := d.GetOk("tag"); ok {
	// 	rtctrlSetTagAttr.Tag = Tag.(string)
	// }

	// if RtctrlSetTag_type, ok := d.GetOk("rtctrl_set_tag_type"); ok {
	// 	rtctrlSetTagAttr.RtctrlSetTag_type = RtctrlSetTag_type.(string)
	// }
	// rtctrlSetTag := models.NewRtctrlSetTag(fmt.Sprintf(models.RnrtctrlSetTag), d.Id(), "desc", "nameAlias", rtctrlSetTagAttr)

	// err1 := aciClient.Save(rtctrlSetTag)
	// if err1 != nil {
	// 	return diag.FromErr(err1)
	// }

	return resourceAciActionRuleProfileRead(ctx, d, m)

}

func resourceAciActionRuleProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	rtctrlAttrP, err := getRemoteActionRuleProfile(aciClient, dn)

	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}
	_, err = setActionRuleProfileAttributes(rtctrlAttrP, d)
	if err != nil {
		d.SetId("")
		return nil
	}
	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciActionRuleProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "rtctrlAttrP")
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())

	d.SetId("")
	return diag.FromErr(err)
}
