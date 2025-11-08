package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/omghozlan/terraform-provider-ctfchallenge/challenges"
)

func dataSourceHint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHintRead,
		Schema: map[string]*schema.Schema{
			"challenge_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The challenge ID to get a hint for",
			},
			"level": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Hint level (0-2, higher levels are more revealing)",
			},
			"hint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hint text",
			},
			"cost": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Point penalty for requesting this hint",
			},
		},
	}
}

func dataSourceHintRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeID := d.Get("challenge_id").(string)
	level := d.Get("level").(int)

	if _, exists := challenges.Challenges[challengeID]; !exists {
		return diag.Errorf("Unknown challenge: %s", challengeID)
	}

	hint := challenges.GetHint(challengeID, level)
	cost := (level + 1) * 10 // Hints cost 10, 20, or 30 points

	d.Set("hint", hint)
	d.Set("cost", cost)
	d.SetId(fmt.Sprintf("%s_hint_%d", challengeID, level))

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  fmt.Sprintf("Hint requested (-%d points)", cost),
		Detail:   fmt.Sprintf("Hint level %d for challenge '%s'", level, challengeID),
	})

	return diags
}
