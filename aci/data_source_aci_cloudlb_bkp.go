package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciCloudL4L7Devices() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciCloudL4L7DevicesRead,
		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			},
			"allow_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_scaling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"context_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			// "firewall_mode": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			// "firewall_status": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"function_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"is_instantiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_static_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"native_lb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "oid": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"package_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sku": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svc_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trunking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_l4_l7_lb_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "relation_cloudrs_ldev_to_cloud_subnet": &schema.Schema{
			// 	Type:        schema.TypeSet,
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// 	Optional:    true,
			// 	Description: "Query cloud:Subnet relationship object",
			// 	Set:         schema.HashString,
			// },
			// "relation_cloudrs_ldev_to_compute_pol": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query cloud:ComputePol relationship object",
			// },
			// "relation_cloudrs_ldev_to_mgmt_pol": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query cloud:MgmtPol relationship object",
			// },
			// "relation_vnsrs_aldev_to_dev_mgr": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query vns:DevMgr relationship object",
			// },
			// "relation_vnsrs_aldev_to_dom_p": &schema.Schema{
			// 	Type:        schema.TypeMap,
			// 	Optional:    true,
			// 	Description: "Query vmmDomP relationship object",
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// },
			// "relation_vnsrs_aldev_to_phys_dom_p": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query phys:DomP relationship object",
			// },
			// "relation_vnsrs_aldev_to_vlan_inst_p": &schema.Schema{
			// 	Type:        schema.TypeSet,
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// 	Optional:    true,
			// 	Description: "Query fvns:VlanInstP relationship object",
			// 	Set:         schema.HashString,
			// },
			// "relation_vnsrs_aldev_to_vxlan_inst_p": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query fvns:VxlanInstP relationship object",
			// },
			// "relation_vnsrs_dev_epg": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query fv:EPg relationship object",
			// },
			// "relation_vnsrs_mdev_att": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query vns:MDev relationship object",
			// },
			// "relation_vnsrs_svc_mgmt_epg": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Query fv:EPg relationship object",
			// },
		})),
	}
}

func dataSourceAciCloudL4L7DevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	name := d.Get("name").(string)
	TenantDn := d.Get("tenant_dn").(string)
	rn := fmt.Sprintf(models.RnCloudLB, name)
	dn := fmt.Sprintf("%s/%s", TenantDn, rn)

	cloudLB, err := getRemoteCloudL4L7LB(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setCloudL4L7LBAttributes(cloudLB, d)
	if err != nil {
		return nil
	}

	// // Get and Set Relational Attributes
	// getAndSetCloudL4L7LBRelationalAttributes(aciClient, dn, d)
	return nil
}
