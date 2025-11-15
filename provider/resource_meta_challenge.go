package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMetaChallenge() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMetaChallengeCreate,
		ReadContext:   resourceMetaChallengeRead,
		UpdateContext: resourceMetaChallengeUpdate,
		DeleteContext: resourceMetaChallengeDelete,
		Schema: map[string]*schema.Schema{
			"challenge_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of meta-argument challenge (count, for_each, depends_on, lifecycle, etc.)",
			},
			"configuration": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Your challenge solution configuration",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Additional metadata about your solution",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_arguments_used": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of meta-arguments used in this configuration",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"resource_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of resources created",
						},
						"complexity_score": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Self-assessed complexity (1-10)",
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Notes about your implementation",
						},
					},
				},
			},
			"validation_result": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Result of the meta-argument validation",
			},
			"hints_used": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Number of hints used for this challenge",
			},
			"success": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the challenge was completed successfully",
			},
		},
	}
}

func resourceMetaChallengeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	challengeType := d.Get("challenge_type").(string)
	config := d.Get("configuration").(map[string]interface{})

	// Validate based on challenge type
	success := false
	var resultMessage string

	switch challengeType {
	case "count":
		success = validateCountConfig(config)
		resultMessage = "Count meta-argument validated"
	case "for_each":
		success = validateForEachConfig(config)
		resultMessage = "For_each meta-argument validated"
	case "depends_on":
		success = validateDependsOnConfig(config)
		resultMessage = "Depends_on meta-argument validated"
	case "lifecycle":
		success = validateLifecycleConfig(config)
		resultMessage = "Lifecycle meta-argument validated"
	default:
		resultMessage = fmt.Sprintf("Unknown challenge type: %s", challengeType)
	}

	d.Set("success", success)
	d.Set("validation_result", resultMessage)

	if success {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Meta-Argument Challenge Complete!",
			Detail:   fmt.Sprintf("You successfully demonstrated %s meta-argument usage!", challengeType),
		})
	}

	d.SetId(fmt.Sprintf("meta-%s-%d", challengeType, time.Now().Unix()))
	return diags
}

func validateCountConfig(config map[string]interface{}) bool {
	// Simple validation - check for count-related keys
	_, hasCount := config["count_value"]
	_, hasResourceIds := config["resource_ids"]
	return hasCount && hasResourceIds
}

func validateForEachConfig(config map[string]interface{}) bool {
	// Check for for_each related keys
	_, hasForeach := config["foreach_type"]
	_, hasDifficulties := config["difficulties"]
	return hasForeach && hasDifficulties
}

func validateDependsOnConfig(config map[string]interface{}) bool {
	// Check for dependency chain
	_, hasChain := config["dependency_chain_length"]
	_, hasResources := config["resource_chain"]
	return hasChain && hasResources
}

func validateLifecycleConfig(config map[string]interface{}) bool {
	// Check for lifecycle rules
	_, hasCreate := config["uses_create_before_destroy"]
	_, hasIgnore := config["ignore_changes"]
	return hasCreate && hasIgnore
}

func resourceMetaChallengeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceMetaChallengeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceMetaChallengeCreate(ctx, d, m)
}

func resourceMetaChallengeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
