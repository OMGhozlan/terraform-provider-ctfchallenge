package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/omghozlan/terraform-provider-ctfchallenge/challenges"
)

func resourcePuzzleBox() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePuzzleBoxCreate,
		ReadContext:   resourcePuzzleBoxRead,
		UpdateContext: resourcePuzzleBoxUpdate,
		DeleteContext: resourcePuzzleBoxDelete,
		Schema: map[string]*schema.Schema{
			"inputs": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Puzzle inputs to validate",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"solved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the puzzle was solved",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Result message",
			},
			"secret_output": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Secret revealed when puzzle is solved",
			},
		},
	}
}

func resourcePuzzleBoxCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	inputs := d.Get("inputs").(map[string]interface{})
	solved, message := challenges.ValidatePuzzleInput(inputs)

	d.Set("solved", solved)
	d.Set("message", message)

	if solved {
		d.Set("secret_output", "flag{xor_puzzl3_s0lv3d}")
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Puzzle Solved!",
			Detail:   "Check the secret_output for your reward",
		})
	}

	d.SetId(fmt.Sprintf("puzzle-%d", time.Now().Unix()))
	return diags
}

func resourcePuzzleBoxRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourcePuzzleBoxUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePuzzleBoxCreate(ctx, d, m)
}

func resourcePuzzleBoxDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
