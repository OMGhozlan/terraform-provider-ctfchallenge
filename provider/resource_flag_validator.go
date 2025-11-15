package provider

import (
	"context"
	"encoding/json"
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
			"proof_of_work": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Direct proof that you completed the challenge requirements",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"resource_proof": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Proof from Terraform resources with their complete configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the resource",
						},
						"resource_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name/identifier of the resource",
						},
						"attributes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Resource attributes",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"lifecycle_config": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JSON-encoded lifecycle configuration",
						},
						"meta_arguments": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Meta-arguments used (count, for_each, depends_on, etc.)",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"data_source_proof": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Proof from Terraform data sources",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the data source",
						},
						"data_source_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name/identifier of the data source",
						},
						"attributes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Data source attributes",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"lifecycle_config": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JSON-encoded lifecycle configuration",
						},
					},
				},
			},
			"module_proof": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Proof from module configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"module_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the module",
						},
						"input_validations": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JSON-encoded input validations",
						},
						"output_validations": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "JSON-encoded output validations",
						},
						"resources_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of resources in module",
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
				Description: "Source of the proof",
			},
			"validation_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Detailed validation results",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func extractProofData(d *schema.ResourceData) (*challenges.ProofData, diag.Diagnostics) {
	var diags diag.Diagnostics
	proofData := &challenges.ProofData{
		Resources:   []challenges.ResourceProof{},
		DataSources: []challenges.DataSourceProof{},
		Manual:      make(map[string]interface{}),
	}

	// Extract resource proofs
	if v, ok := d.GetOk("resource_proof"); ok {
		resourceProofList := v.([]interface{})
		for _, rp := range resourceProofList {
			resourceProof := rp.(map[string]interface{})

			proof := challenges.ResourceProof{
				ResourceType:  resourceProof["resource_type"].(string),
				ResourceName:  resourceProof["resource_name"].(string),
				Attributes:    make(map[string]interface{}),
				MetaArguments: make(map[string]interface{}),
			}

			// Extract attributes
			if attrs, ok := resourceProof["attributes"].(map[string]interface{}); ok {
				for k, v := range attrs {
					proof.Attributes[k] = v
				}
			}

			// Extract lifecycle config
			if lifecycleJSON, ok := resourceProof["lifecycle_config"].(string); ok && lifecycleJSON != "" {
				var lifecycle challenges.LifecycleConfig
				if err := json.Unmarshal([]byte(lifecycleJSON), &lifecycle); err == nil {
					proof.Lifecycle = &lifecycle
				} else {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "Failed to parse lifecycle config",
						Detail:   fmt.Sprintf("Resource %s: %v", proof.ResourceName, err),
					})
				}
			}

			// Extract meta-arguments
			if meta, ok := resourceProof["meta_arguments"].(map[string]interface{}); ok {
				for k, v := range meta {
					proof.MetaArguments[k] = v
				}
			}

			proofData.Resources = append(proofData.Resources, proof)
		}

		if len(proofData.Resources) > 0 {
			proofData.Source = fmt.Sprintf("resources:%d", len(proofData.Resources))
		}
	}

	// Extract data source proofs
	if v, ok := d.GetOk("data_source_proof"); ok {
		dataProofList := v.([]interface{})
		for _, dp := range dataProofList {
			dataProof := dp.(map[string]interface{})

			proof := challenges.DataSourceProof{
				DataSourceType: dataProof["data_source_type"].(string),
				DataSourceName: dataProof["data_source_name"].(string),
				Attributes:     make(map[string]interface{}),
			}

			// Extract attributes
			if attrs, ok := dataProof["attributes"].(map[string]interface{}); ok {
				for k, v := range attrs {
					proof.Attributes[k] = v
				}
			}

			// Extract lifecycle config
			if lifecycleJSON, ok := dataProof["lifecycle_config"].(string); ok && lifecycleJSON != "" {
				var lifecycle challenges.LifecycleConfig
				if err := json.Unmarshal([]byte(lifecycleJSON), &lifecycle); err == nil {
					proof.Lifecycle = &lifecycle
				}
			}

			proofData.DataSources = append(proofData.DataSources, proof)
		}

		if len(proofData.DataSources) > 0 {
			proofData.Source = fmt.Sprintf("data_sources:%d", len(proofData.DataSources))
		}
	}

	// Extract module proof
	if v, ok := d.GetOk("module_proof"); ok {
		moduleProofList := v.([]interface{})
		if len(moduleProofList) > 0 {
			moduleProof := moduleProofList[0].(map[string]interface{})

			proof := challenges.ModuleProof{
				ModuleName: moduleProof["module_name"].(string),
			}

			// Extract input validations
			if inputJSON, ok := moduleProof["input_validations"].(string); ok && inputJSON != "" {
				var validations []challenges.ValidationRule
				if err := json.Unmarshal([]byte(inputJSON), &validations); err == nil {
					proof.InputValidations = validations
				}
			}

			// Extract output validations
			if outputJSON, ok := moduleProof["output_validations"].(string); ok && outputJSON != "" {
				var validations []challenges.ValidationRule
				if err := json.Unmarshal([]byte(outputJSON), &validations); err == nil {
					proof.OutputValidations = validations
				}
			}

			if count, ok := moduleProof["resources_count"].(int); ok {
				proof.ResourcesCount = count
			}

			proofData.Module = &proof
			proofData.Source = "module"
		}
	}

	// Fall back to manual proof of work
	if proofData.Source == "" {
		if v, ok := d.GetOk("proof_of_work"); ok {
			proofData.Manual = v.(map[string]interface{})
			proofData.Source = "manual"
		}
	}

	if proofData.Source == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No proof provided",
			Detail:   "You must provide one of: proof_of_work, resource_proof, data_source_proof, or module_proof",
		})
		return nil, diags
	}

	return proofData, diags
}

func resourceFlagValidatorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeID := d.Get("challenge_id").(string)

	challenge, exists := challenges.Challenges[challengeID]
	if !exists {
		return diag.Errorf("Unknown challenge: %s", challengeID)
	}

	// Extract proof data
	proofData, extractDiags := extractProofData(d)
	diags = append(diags, extractDiags...)

	if proofData == nil {
		return diags
	}

	// Validate using the enhanced validator
	result := challenge.ValidateProof(proofData)

	d.Set("proof_source", proofData.Source)
	d.Set("validated", result.Success)
	d.Set("message", result.Message)
	d.Set("validation_details", result.Details)

	if result.Success {
		d.Set("points", challenge.Points)
		d.Set("flag", result.Flag)
		d.Set("timestamp", time.Now().UTC().Format(time.RFC3339))

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "🎉 Challenge Completed!",
			Detail: fmt.Sprintf("You earned %d points for completing '%s'. Check the 'flag' output for your reward!\n\nValidation details:\n%s",
				challenge.Points, challenge.Name, formatDetails(result.Details)),
		})
	} else {
		d.Set("points", 0)
		d.Set("flag", "")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Challenge validation failed",
			Detail:   fmt.Sprintf("%s\n\nValidation details:\n%s", result.Message, formatDetails(result.Details)),
		})
	}

	d.SetId(fmt.Sprintf("%s-%d", challengeID, time.Now().Unix()))
	return diags
}

func formatDetails(details []string) string {
	if len(details) == 0 {
		return "No additional details"
	}
	result := ""
	for i, detail := range details {
		result += fmt.Sprintf("  %d. %s\n", i+1, detail)
	}
	return result
}

func resourceFlagValidatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceFlagValidatorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceFlagValidatorCreate(ctx, d, m)
}

func resourceFlagValidatorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
