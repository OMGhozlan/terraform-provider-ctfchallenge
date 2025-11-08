package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/omghozlan/terraform-provider-ctfchallenge/challenges"
)

func dataSourceChallengeInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChallengeInfoRead,
		Schema: map[string]*schema.Schema{
			"challenge_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The challenge ID to get info for",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Challenge name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Challenge description",
			},
			"points": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Points for completing this challenge",
			},
			"difficulty": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Difficulty level",
			},
			"category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Challenge category",
			},
		},
	}
}

func dataSourceChallengeInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeID := d.Get("challenge_id").(string)

	challenge, exists := challenges.Challenges[challengeID]
	if !exists {
		return diag.Errorf("Unknown challenge: %s", challengeID)
	}

	d.Set("name", challenge.Name)
	d.Set("description", challenge.Description)
	d.Set("points", challenge.Points)
	d.Set("difficulty", challenge.Difficulty)
	d.Set("category", challenge.Category)
	d.SetId(challengeID)

	return diags
}
