package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/omghozlan/terraform-provider-ctfchallenge/challenges"
)

func resourceFlagValidator() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlagValidatorCreate,
		ReadContext:   resourceFlagValidatorRead,
		UpdateContext: resourceFlagValidatorUpdate,
		DeleteContext: resourceFlagValidatorDelete,
		Schema: map[string]*schema.Schema{
			"challenge_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the challenge to validate",
			},
			"flag": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The flag you discovered for this challenge",
			},
			"proof_of_work": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Proof that you completed the challenge requirements",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"validated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the challenge was successfully validated",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Validation result message",
			},
			"points": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Points awarded for this challenge",
			},
			"timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the challenge was completed",
			},
		},
	}
}

func resourceFlagValidatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeID := d.Get("challenge_id").(string)
	flag := d.Get("flag").(string)

	challenge, exists := challenges.Challenges[challengeID]
	if !exists {
		return diag.Errorf("Unknown challenge: %s", challengeID)
	}

	proofMap := d.Get("proof_of_work").(map[string]interface{})
	validated, correctFlag, err := challenge.Validator(proofMap)

	if err != nil {
		d.Set("validated", false)
		d.Set("points", 0)
		d.Set("message", fmt.Sprintf("❌ Challenge failed: %s", err.Error()))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Challenge validation failed",
			Detail:   err.Error(),
		})
	} else if validated && flag == correctFlag {
		d.Set("validated", true)
		d.Set("points", challenge.Points)
		d.Set("message", fmt.Sprintf("🎉 Congratulations! You solved '%s' and earned %d points!", challenge.Name, challenge.Points))
		d.Set("timestamp", time.Now().UTC().Format(time.RFC3339))

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Challenge Completed!",
			Detail:   fmt.Sprintf("You earned %d points for completing '%s'", challenge.Points, challenge.Name),
		})
	} else if validated && flag != correctFlag {
		d.Set("validated", false)
		d.Set("points", 0)
		d.Set("message", "✓ Challenge requirements met, but flag is incorrect")
	} else {
		d.Set("validated", false)
		d.Set("points", 0)
		d.Set("message", "❌ Challenge requirements not met and/or incorrect flag")
	}

	d.SetId(fmt.Sprintf("%s-%d", challengeID, time.Now().Unix()))
	return diags
}

func resourceFlagValidatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// This is an ephemeral resource, so read is a no-op
	return nil
}

func resourceFlagValidatorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Re-run validation on update
	return resourceFlagValidatorCreate(ctx, d, m)
}

func resourceFlagValidatorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Remove from state
	d.SetId("")
	return nil
}
