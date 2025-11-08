package provider

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/omghozlan/terraform-provider-ctfchallenge/challenges"
)

func dataSourceChallengeList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChallengeListRead,
		Schema: map[string]*schema.Schema{
			"difficulty": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by difficulty (beginner, intermediate, advanced)",
			},
			"category": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by category",
			},
			"challenges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of challenges",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"points": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"difficulty": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"total_points": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total points available from listed challenges",
			},
		},
	}
}

func dataSourceChallengeListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	difficultyFilter := d.Get("difficulty").(string)
	categoryFilter := d.Get("category").(string)

	var challengeList []interface{}
	totalPoints := 0

	// Get sorted challenge IDs for consistent output
	ids := challenges.GetAllChallengeIDs()
	sort.Strings(ids)

	for _, id := range ids {
		challenge := challenges.Challenges[id]

		// Apply filters
		if difficultyFilter != "" && challenge.Difficulty != difficultyFilter {
			continue
		}
		if categoryFilter != "" && challenge.Category != categoryFilter {
			continue
		}

		challengeMap := map[string]interface{}{
			"id":          challenge.ID,
			"name":        challenge.Name,
			"description": challenge.Description,
			"points":      challenge.Points,
			"difficulty":  challenge.Difficulty,
			"category":    challenge.Category,
		}
		challengeList = append(challengeList, challengeMap)
		totalPoints += challenge.Points
	}

	d.Set("challenges", challengeList)
	d.Set("total_points", totalPoints)
	d.SetId(fmt.Sprintf("challenges-%s-%s", difficultyFilter, categoryFilter))

	return diags
}