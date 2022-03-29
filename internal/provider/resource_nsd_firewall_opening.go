package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"math/rand"
	"strconv"
)

func resourceNsdFirewallOpening() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a firewall opening in NSD.",

		CreateContext: resourceNsdFirewallOpeningCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: resourceNsdFirewallOpeningDelete,

		Schema: map[string]*schema.Schema{
			"ninety_nine_number": {
				Description: "The 99-number from STEP",
				Type:        schema.TypeString,
				Required:    true,
			},
			"application_id": {
				Description: "The application ID from STEP (confusingly also called STEP ID)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPv4Address,
			},
			"target_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPv4Address,
			},
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IsPortNumber,
			},
			"proto": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
			},
		},
	}
}

func resourceNsdFirewallOpeningCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(strconv.FormatUint(rand.Uint64(), 10))

	var diags diag.Diagnostics
	return diags
}

func resourceNsdFirewallOpeningDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	data.SetId("")

	var diags diag.Diagnostics
	return diags
}
