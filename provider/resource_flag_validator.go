package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"proof_of_work": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Direct proof that you completed the challenge requirements",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"resource_proof": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Proof from a Terraform resource",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the resource (e.g., 'ctfchallenge_puzzle_box')",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the resource instance",
						},
						"attributes": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: "Relevant attributes from the resource",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"data_source_proof": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Proof from a Terraform data source",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the data source (e.g., 'ctfchallenge_challenge_info')",
						},
						"data_source_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the data source",
						},
						"attributes": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: "Relevant attributes from the data source",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			"flag": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The flag revealed upon successful completion",
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
			"proof_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source of the proof (manual, resource, or data_source)",
			},
		},
	}
}

func extractProofData(d *schema.ResourceData) (map[string]interface{}, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var proofMap map[string]interface{}
	var proofSource string

	// Check for resource_proof
	if v, ok := d.GetOk("resource_proof"); ok {
		resourceProofList := v.([]interface{})
		if len(resourceProofList) > 0 {
			resourceProof := resourceProofList[0].(map[string]interface{})
			proofSource = fmt.Sprintf("resource:%s", resourceProof["resource_type"].(string))

			// Extract attributes
			attributes := resourceProof["attributes"].(map[string]interface{})
			proofMap = make(map[string]interface{})

			// Add metadata
			proofMap["_source_type"] = "resource"
			proofMap["_resource_type"] = resourceProof["resource_type"].(string)
			proofMap["_resource_id"] = resourceProof["resource_id"].(string)

			// Copy all attributes
			for k, v := range attributes {
				proofMap[k] = v
			}

			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Using Resource Proof",
				Detail:   fmt.Sprintf("Validating using resource: %s (ID: %s)", resourceProof["resource_type"], resourceProof["resource_id"]),
			})

			return proofMap, proofSource, diags
		}
	}

	// Check for data_source_proof
	if v, ok := d.GetOk("data_source_proof"); ok {
		dataProofList := v.([]interface{})
		if len(dataProofList) > 0 {
			dataProof := dataProofList[0].(map[string]interface{})
			proofSource = fmt.Sprintf("data:%s", dataProof["data_source_type"].(string))

			// Extract attributes
			attributes := dataProof["attributes"].(map[string]interface{})
			proofMap = make(map[string]interface{})

			// Add metadata
			proofMap["_source_type"] = "data"
			proofMap["_data_source_type"] = dataProof["data_source_type"].(string)
			proofMap["_data_source_id"] = dataProof["data_source_id"].(string)

			// Copy all attributes
			for k, v := range attributes {
				proofMap[k] = v
			}

			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Using Data Source Proof",
				Detail:   fmt.Sprintf("Validating using data source: %s (ID: %s)", dataProof["data_source_type"], dataProof["data_source_id"]),
			})

			return proofMap, proofSource, diags
		}
	}

	// Fall back to proof_of_work
	if v, ok := d.GetOk("proof_of_work"); ok {
		proofMap = v.(map[string]interface{})
		proofSource = "manual"
		return proofMap, proofSource, diags
	}

	// No proof provided
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "No proof provided",
		Detail:   "You must provide one of: proof_of_work, resource_proof, or data_source_proof",
	})

	return nil, "", diags
}

func resourceFlagValidatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeID := d.Get("challenge_id").(string)

	challenge, exists := challenges.Challenges[challengeID]
	if !exists {
		return diag.Errorf("Unknown challenge: %s", challengeID)
	}

	// Extract proof data from resource, data source, or manual input
	proofMap, proofSource, extractDiags := extractProofData(d)
	diags = append(diags, extractDiags...)

	if proofMap == nil {
		return diags
	}

	validated, correctFlag, err := challenge.Validator(proofMap)

	d.Set("proof_source", proofSource)

	if err != nil {
		d.Set("validated", false)
		d.Set("points", 0)
		d.Set("flag", "")
		d.Set("message", fmt.Sprintf("❌ Challenge failed: %s", err.Error()))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Challenge validation failed",
			Detail:   fmt.Sprintf("Source: %s | Error: %s", proofSource, err.Error()),
		})
	} else if validated {
		d.Set("validated", true)
		d.Set("points", challenge.Points)
		d.Set("flag", correctFlag)
		d.Set("message", fmt.Sprintf("🎉 Congratulations! You solved '%s' and earned %d points!", challenge.Name, challenge.Points))
		d.Set("timestamp", time.Now().UTC().Format(time.RFC3339))

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "🎉 Challenge Completed!",
			Detail:   fmt.Sprintf("You earned %d points for completing '%s'. Check the 'flag' output for your reward! (Source: %s)", challenge.Points, challenge.Name, proofSource),
		})
	} else {
		d.Set("validated", false)
		d.Set("points", 0)
		d.Set("flag", "")
		d.Set("message", "❌ Challenge requirements not met")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Challenge not completed",
			Detail:   fmt.Sprintf("The proof from %s did not meet the challenge requirements", proofSource),
		})
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
