package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	// Register meta-argument challenges
	registerMetaArgumentChallenges()
}

func registerMetaArgumentChallenges() {
	// Count Meta-Argument Challenge
	Challenges["count_master"] = &Challenge{
		ID:          "count_master",
		Name:        "Count Master",
		Description: "Master the 'count' meta-argument by creating exactly 3 puzzle boxes with sequential keys",
		Points:      150,
		Difficulty:  "intermediate",
		Category:    "meta-arguments",
		Validator:   validateCountChallenge,
	}

	// For_each Meta-Argument Challenge
	Challenges["foreach_wizard"] = &Challenge{
		ID:          "foreach_wizard",
		Name:        "For Each Wizard",
		Description: "Use 'for_each' to create puzzle boxes for all difficulty levels: beginner, intermediate, advanced",
		Points:      200,
		Difficulty:  "intermediate",
		Category:    "meta-arguments",
		Validator:   validateForEachChallenge,
	}

	// Depends_on Meta-Argument Challenge
	Challenges["dependency_chain"] = &Challenge{
		ID:          "dependency_chain",
		Name:        "Dependency Chain Master",
		Description: "Create a dependency chain using 'depends_on' with at least 3 resources in sequence",
		Points:      175,
		Difficulty:  "intermediate",
		Category:    "meta-arguments",
		Validator:   validateDependsOnChallenge,
	}

	// Lifecycle Meta-Argument Challenge
	Challenges["lifecycle_expert"] = &Challenge{
		ID:          "lifecycle_expert",
		Name:        "Lifecycle Expert",
		Description: "Use lifecycle rules to demonstrate create_before_destroy and ignore_changes",
		Points:      225,
		Difficulty:  "advanced",
		Category:    "meta-arguments",
		Validator:   validateLifecycleChallenge,
	}

	// Combined Meta-Arguments Challenge
	Challenges["meta_grandmaster"] = &Challenge{
		ID:          "meta_grandmaster",
		Name:        "Meta-Argument Grandmaster",
		Description: "Combine count, for_each, depends_on, and lifecycle in a single configuration",
		Points:      300,
		Difficulty:  "advanced",
		Category:    "meta-arguments",
		Validator:   validateMetaGrandmasterChallenge,
	}

	// Dynamic Blocks Challenge
	Challenges["dynamic_block_architect"] = &Challenge{
		ID:          "dynamic_blocks",
		Name:        "Dynamic Block Architect",
		Description: "Use dynamic blocks to generate configuration based on variable inputs",
		Points:      180,
		Difficulty:  "intermediate",
		Category:    "meta-arguments",
		Validator:   validateDynamicBlocksChallenge,
	}

	// Locals and Count Challenge
	Challenges["locals_count_combo"] = &Challenge{
		ID:          "locals_count_combo",
		Name:        "Locals + Count Combo",
		Description: "Use locals with count.index to create resources with computed names",
		Points:      160,
		Difficulty:  "intermediate",
		Category:    "meta-arguments",
		Validator:   validateLocalsCountChallenge,
	}

	// Conditional Creation Challenge
	Challenges["conditional_resources"] = &Challenge{
		ID:          "conditional_resources",
		Name:        "Conditional Creation Master",
		Description: "Use count = var.condition ? 1 : 0 pattern to conditionally create resources",
		Points:      140,
		Difficulty:  "beginner",
		Category:    "meta-arguments",
		Validator:   validateConditionalChallenge,
	}
}

func validateCountChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check if count was used
	countStr, hasCount := proof["count_value"].(string)
	if !hasCount {
		return false, "", fmt.Errorf("missing 'count_value' in proof - use count meta-argument")
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count != 3 {
		return false, "", fmt.Errorf("count must be exactly 3, got: %s", countStr)
	}

	// Check for sequential resource IDs
	resourceIDs, hasIDs := proof["resource_ids"].(string)
	if !hasIDs {
		return false, "", fmt.Errorf("missing 'resource_ids' - provide comma-separated list of created resource IDs")
	}

	ids := strings.Split(resourceIDs, ",")
	if len(ids) != 3 {
		return false, "", fmt.Errorf("expected 3 resource IDs, got %d", len(ids))
	}

	// Check for count.index usage
	hasCountIndex, _ := proof["uses_count_index"].(string)
	if hasCountIndex != "true" {
		return false, "", fmt.Errorf("you must use count.index in your resource configuration")
	}

	return true, "flag{c0unt_m3t4_4rgum3nt_m4st3r}", nil
}

func validateForEachChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for for_each usage
	foreachType, hasForeach := proof["foreach_type"].(string)
	if !hasForeach {
		return false, "", fmt.Errorf("missing 'foreach_type' - specify 'map' or 'set'")
	}

	if foreachType != "map" && foreachType != "set" {
		return false, "", fmt.Errorf("foreach_type must be 'map' or 'set', got: %s", foreachType)
	}

	// Check for all required difficulty levels
	difficulties, hasDiff := proof["difficulties"].(string)
	if !hasDiff {
		return false, "", fmt.Errorf("missing 'difficulties' - provide comma-separated list")
	}

	diffList := strings.Split(difficulties, ",")
	requiredDiffs := map[string]bool{
		"beginner":     false,
		"intermediate": false,
		"advanced":     false,
	}

	for _, diff := range diffList {
		diff = strings.TrimSpace(diff)
		if _, exists := requiredDiffs[diff]; exists {
			requiredDiffs[diff] = true
		}
	}

	for diff, found := range requiredDiffs {
		if !found {
			return false, "", fmt.Errorf("missing difficulty level: %s", diff)
		}
	}

	// Check for each.key or each.value usage
	usesEach, _ := proof["uses_each"].(string)
	if usesEach != "true" {
		return false, "", fmt.Errorf("you must use each.key or each.value in your configuration")
	}

	return true, "flag{f0r_34ch_l00p_m4g1c}", nil
}

func validateDependsOnChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for dependency chain
	chainLength, hasChain := proof["dependency_chain_length"].(string)
	if !hasChain {
		return false, "", fmt.Errorf("specify how many resources are in your dependency chain")
	}

	length, err := strconv.Atoi(chainLength)
	if err != nil || length < 3 {
		return false, "", fmt.Errorf("dependency chain must have at least 3 resources, got: %s", chainLength)
	}

	// Check for explicit depends_on usage
	usesDependsOn, _ := proof["uses_depends_on"].(string)
	if usesDependsOn != "true" {
		return false, "", fmt.Errorf("you must use explicit depends_on meta-argument")
	}

	// Check for resource names in chain
	resourceChain, hasResources := proof["resource_chain"].(string)
	if !hasResources {
		return false, "", fmt.Errorf("missing 'resource_chain' - provide comma-separated resource names")
	}

	resources := strings.Split(resourceChain, ",")
	if len(resources) < 3 {
		return false, "", fmt.Errorf("resource chain must include at least 3 resources")
	}

	// Verify dependency order is documented
	if _, hasOrder := proof["dependency_order"].(string); !hasOrder {
		return false, "", fmt.Errorf("missing 'dependency_order' - document your dependency sequence")
	}

	return true, "flag{d3p3nd3ncy_ch41n_m4st3r}", nil
}

func validateLifecycleChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for create_before_destroy
	usesCreateBefore, _ := proof["uses_create_before_destroy"].(string)
	if usesCreateBefore != "true" {
		return false, "", fmt.Errorf("you must use lifecycle.create_before_destroy")
	}

	// Check for ignore_changes
	ignoreChanges, hasIgnore := proof["ignore_changes"].(string)
	if !hasIgnore || ignoreChanges == "" {
		return false, "", fmt.Errorf("you must specify lifecycle.ignore_changes with at least one attribute")
	}

	// Check that both lifecycle rules are in the same or different resources
	lifecycleCount, hasCount := proof["lifecycle_rules_count"].(string)
	if !hasCount {
		return false, "", fmt.Errorf("missing 'lifecycle_rules_count' - how many lifecycle rules did you use?")
	}

	count, err := strconv.Atoi(lifecycleCount)
	if err != nil || count < 2 {
		return false, "", fmt.Errorf("you must use at least 2 lifecycle rules")
	}

	// Check for documentation of why lifecycle rules were needed
	justification, hasJust := proof["lifecycle_justification"].(string)
	if !hasJust || len(justification) < 10 {
		return false, "", fmt.Errorf("provide 'lifecycle_justification' explaining why you used these lifecycle rules")
	}

	return true, "flag{l1f3cycl3_rul3s_3xp3rt}", nil
}

func validateMetaGrandmasterChallenge(proof map[string]interface{}) (bool, string, error) {
	// This challenge requires combining multiple meta-arguments
	requiredMetaArgs := map[string]bool{
		"count":      false,
		"for_each":   false,
		"depends_on": false,
		"lifecycle":  false,
	}

	metaArgsUsed, hasMetaArgs := proof["meta_arguments_used"].(string)
	if !hasMetaArgs {
		return false, "", fmt.Errorf("missing 'meta_arguments_used' - provide comma-separated list")
	}

	args := strings.Split(metaArgsUsed, ",")
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if _, exists := requiredMetaArgs[arg]; exists {
			requiredMetaArgs[arg] = true
		}
	}

	var missing []string
	for arg, used := range requiredMetaArgs {
		if !used {
			missing = append(missing, arg)
		}
	}

	if len(missing) > 0 {
		return false, "", fmt.Errorf("missing meta-arguments: %s", strings.Join(missing, ", "))
	}

	// Check for minimum resource count
	totalResources, hasTotal := proof["total_resources"].(string)
	if !hasTotal {
		return false, "", fmt.Errorf("missing 'total_resources' count")
	}

	count, err := strconv.Atoi(totalResources)
	if err != nil || count < 5 {
		return false, "", fmt.Errorf("you must create at least 5 resources, got: %s", totalResources)
	}

	// Check for configuration complexity
	configLines, hasLines := proof["config_lines"].(string)
	if !hasLines {
		return false, "", fmt.Errorf("missing 'config_lines' - how many lines is your configuration?")
	}

	lines, err := strconv.Atoi(configLines)
	if err != nil || lines < 50 {
		return false, "", fmt.Errorf("configuration must be at least 50 lines to demonstrate complexity")
	}

	// Verify architectural documentation
	architecture, hasArch := proof["architecture_description"].(string)
	if !hasArch || len(architecture) < 50 {
		return false, "", fmt.Errorf("provide detailed 'architecture_description' (min 50 chars) of your infrastructure")
	}

	return true, "flag{m3t4_4rgum3nt_gr4ndm4st3r_ultimate}", nil
}

func validateDynamicBlocksChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for dynamic block usage
	usesDynamic, _ := proof["uses_dynamic_blocks"].(string)
	if usesDynamic != "true" {
		return false, "", fmt.Errorf("you must use dynamic blocks in your configuration")
	}

	// Check for iteration
	dynamicCount, hasDynamic := proof["dynamic_iterations"].(string)
	if !hasDynamic {
		return false, "", fmt.Errorf("missing 'dynamic_iterations' - how many iterations in your dynamic block?")
	}

	iterations, err := strconv.Atoi(dynamicCount)
	if err != nil || iterations < 2 {
		return false, "", fmt.Errorf("dynamic block must iterate at least 2 times, got: %s", dynamicCount)
	}

	return true, "flag{dyn4m1c_bl0ck_4rch1t3ct}", nil
}

func validateLocalsCountChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for locals usage
	usesLocals, _ := proof["uses_locals"].(string)
	if usesLocals != "true" {
		return false, "", fmt.Errorf("you must define and use locals")
	}

	// Check for count usage
	countValue, hasCount := proof["count_value"].(string)
	if !hasCount {
		return false, "", fmt.Errorf("missing 'count_value'")
	}

	count, err := strconv.Atoi(countValue)
	if err != nil || count < 2 {
		return false, "", fmt.Errorf("count must be at least 2, got: %s", countValue)
	}

	// Check for computed resource names
	resourceNames, hasNames := proof["resource_names"].(string)
	if !hasNames {
		return false, "", fmt.Errorf("missing 'resource_names' - provide comma-separated list of generated names")
	}

	names := strings.Split(resourceNames, ",")
	if len(names) != count {
		return false, "", fmt.Errorf("expected %d resource names, got %d", count, len(names))
	}

	// Check that names follow a pattern (computed from locals + count.index)
	usesCountIndex, _ := proof["uses_count_index_in_locals"].(string)
	if usesCountIndex != "true" {
		return false, "", fmt.Errorf("you must use count.index with locals to compute names")
	}

	return true, "flag{l0c4ls_c0unt_c0mb0_m4st3r}", nil
}

func validateConditionalChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for conditional count usage
	usesConditional, _ := proof["uses_conditional_count"].(string)
	if usesConditional != "true" {
		return false, "", fmt.Errorf("you must use conditional count (count = condition ? 1 : 0)")
	}

	// Check for variable condition
	hasVariable, _ := proof["uses_variable_condition"].(string)
	if hasVariable != "true" {
		return false, "", fmt.Errorf("condition must be based on a variable")
	}

	// Check both true and false cases
	_, hasTrue := proof["condition_true_result"].(string)
	_, hasFalse := proof["condition_false_result"].(string)

	if !hasTrue || !hasFalse {
		return false, "", fmt.Errorf("you must demonstrate both true and false conditions")
	}

	// Verify the pattern
	pattern, hasPattern := proof["conditional_pattern"].(string)
	if !hasPattern {
		return false, "", fmt.Errorf("missing 'conditional_pattern' - document your ternary pattern")
	}

	if !strings.Contains(pattern, "?") || !strings.Contains(pattern, ":") {
		return false, "", fmt.Errorf("conditional_pattern must show ternary operator (? :)")
	}

	return true, "flag{c0nd1t10n4l_cr34t10n_m4st3r}", nil
}

// validateMetaArgumentStructure validates meta-argument challenges using structured proof
func validateMetaArgumentStructure(c *Challenge, proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	switch c.ID {
	case "count_master":
		return validateCountStructure(proof)
	case "foreach_wizard":
		return validateForEachStructure(proof)
	case "dependency_chain":
		return validateDependsOnStructure(proof)
	case "lifecycle_expert":
		return validateLifecycleStructure(proof)
	case "meta_grandmaster":
		return validateMetaGrandmasterStructure(proof)
	case "dynamic_block_architect":
		return validateDynamicBlocksStructure(proof)
	case "locals_count_combo":
		return validateLocalsCountStructure(proof)
	case "conditional_resources":
		return validateConditionalStructure(proof)
	default:
		// Fall back to legacy validator
		success, flag, err := c.Validator(proof.Manual)
		result.Success = success
		result.Flag = flag
		if err != nil {
			result.Message = err.Error()
		}
	}

	return result
}

// Add stub implementations for structure validators
func validateCountStructure(proof *ProofData) ValidationResult {
	// For now, fall back to manual validation
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateForEachStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateDependsOnStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateLifecycleStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateMetaGrandmasterStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateDynamicBlocksStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateLocalsCountStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

func validateConditionalStructure(proof *ProofData) ValidationResult {
	return ValidationResult{Message: "Use proof_of_work for this challenge"}
}

// validateModuleStructure validates module challenges
func validateModuleStructure(c *Challenge, proof *ProofData) ValidationResult {
	// Implemented in validation_conditions.go
	return validateModuleContractStructure(proof)
}
