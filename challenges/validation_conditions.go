package challenges

import (
	"fmt"
	"strings"
)

func init() {
	registerValidationConditionChallenges()
}

func registerValidationConditionChallenges() {
	// Precondition Basics
	Challenges["precondition_guardian"] = &Challenge{
		ID:          "precondition_guardian",
		Name:        "Precondition Guardian",
		Description: "Use preconditions to validate inputs before resource creation",
		Points:      150,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validatePreconditionChallenge,
	}

	// Postcondition Basics
	Challenges["postcondition_validator"] = &Challenge{
		ID:          "postcondition_validator",
		Name:        "Postcondition Validator",
		Description: "Use postconditions with 'self' to validate resource attributes after creation",
		Points:      175,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validatePostconditionChallenge,
	}

	// Combined Pre/Post Conditions
	Challenges["condition_master"] = &Challenge{
		ID:          "condition_master",
		Name:        "Condition Master",
		Description: "Combine preconditions and postconditions in a single resource",
		Points:      200,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateCombinedConditionsChallenge,
	}

	// Data Source Validation
	Challenges["data_validator"] = &Challenge{
		ID:          "data_validator",
		Name:        "Data Source Validator",
		Description: "Use postconditions to validate data source outputs",
		Points:      160,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateDataSourceConditionChallenge,
	}

	// Output Validation
	Challenges["output_contract"] = &Challenge{
		ID:          "output_contract",
		Name:        "Output Contract Enforcer",
		Description: "Use preconditions in output blocks to enforce module contracts",
		Points:      180,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateOutputConditionChallenge,
	}

	// Complex Validation Chain
	Challenges["validation_chain"] = &Challenge{
		ID:          "validation_chain",
		Name:        "Validation Chain Architect",
		Description: "Create a chain of resources with interconnected pre/postconditions",
		Points:      250,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateValidationChainChallenge,
	}

	// Module Contract Challenge
	Challenges["module_contract"] = &Challenge{
		ID:          "module_contract",
		Name:        "Module Contract Designer",
		Description: "Design a module with comprehensive pre/postconditions for input validation and output guarantees",
		Points:      300,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateModuleContractChallenge,
	}

	// Self-Reference Master
	Challenges["self_reference_master"] = &Challenge{
		ID:          "self_reference_master",
		Name:        "Self-Reference Master",
		Description: "Master the use of 'self' in postconditions to validate multiple attributes",
		Points:      190,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateSelfReferenceChallenge,
	}

	// Conditional Validation
	Challenges["conditional_validation"] = &Challenge{
		ID:          "conditional_validation",
		Name:        "Conditional Validation Expert",
		Description: "Use complex boolean logic in condition blocks with multiple checks",
		Points:      220,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateConditionalValidationChallenge,
	}

	// Error Message Designer
	Challenges["error_message_designer"] = &Challenge{
		ID:          "error_message_designer",
		Name:        "Error Message Designer",
		Description: "Create helpful, informative error messages for all validation failures",
		Points:      140,
		Difficulty:  "beginner",
		Category:    "validation",
		Validator:   validateErrorMessageChallenge,
	}
}

// validateValidationChallengeStructure validates validation challenges using structured proof
func validateValidationChallengeStructure(c *Challenge, proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	switch c.ID {
	case "precondition_guardian":
		return validatePreconditionStructure(proof)
	case "postcondition_validator":
		return validatePostconditionStructure(proof)
	case "condition_master":
		return validateCombinedConditionsStructure(proof)
	case "data_validator":
		return validateDataSourceConditionStructure(proof)
	case "output_contract":
		return validateOutputConditionStructure(proof)
	case "validation_chain":
		return validateValidationChainStructure(proof)
	case "module_contract":
		return validateModuleContractStructure(proof)
	case "self_reference_master":
		return validateSelfReferenceStructure(proof)
	case "conditional_validation":
		return validateConditionalValidationStructure(proof)
	case "error_message_designer":
		return validateErrorMessageStructure(proof)
	default:
		result.Message = "Unknown validation challenge"
	}

	return result
}

func validatePreconditionStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.Resources) == 0 {
		result.Message = "No resources provided. You must create a resource with a precondition in its lifecycle block."
		result.Details = append(result.Details, "Expected: At least one resource with lifecycle.precondition")
		return result
	}

	// Check each resource for preconditions
	var validResource *ResourceProof
	for i := range proof.Resources {
		r := &proof.Resources[i]
		if r.Lifecycle != nil && len(r.Lifecycle.Preconditions) > 0 {
			validResource = r
			break
		}
	}

	if validResource == nil {
		result.Message = "No preconditions found in any resource lifecycle block"
		result.Details = append(result.Details, "Add a lifecycle block with precondition to your resource")
		result.Details = append(result.Details, "Example: lifecycle { precondition { condition = ..., error_message = ... } }")
		return result
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ Found resource '%s' with lifecycle block", validResource.ResourceName))

	// Validate each precondition
	for i, precond := range validResource.Lifecycle.Preconditions {
		result.Details = append(result.Details, fmt.Sprintf("Checking precondition %d...", i+1))

		// Check condition expression exists and is not empty
		if precond.Condition == "" {
			result.Message = fmt.Sprintf("Precondition %d has empty condition expression", i+1)
			result.Details = append(result.Details, "✗ Condition expression is required")
			return result
		}
		result.Details = append(result.Details, fmt.Sprintf("  ✓ Has condition: %s", truncate(precond.Condition, 60)))

		// Check that precondition does NOT use 'self' (resource doesn't exist yet)
		if conditionUsesSelf(precond.Condition) {
			result.Message = fmt.Sprintf("Precondition %d incorrectly uses 'self'", i+1)
			result.Details = append(result.Details, "  ✗ Preconditions should NOT use 'self' - the resource doesn't exist yet")
			result.Details = append(result.Details, "  Hint: Use var.* or local.* to validate inputs before creation")
			return result
		}
		result.Details = append(result.Details, "  ✓ Does not use 'self' (correct for precondition)")

		// Check error message
		if !validateErrorMessage(precond.ErrorMessage, 10) {
			result.Message = fmt.Sprintf("Precondition %d has inadequate error message", i+1)
			result.Details = append(result.Details, "  ✗ Error message must be at least 10 characters and descriptive")
			return result
		}
		result.Details = append(result.Details, fmt.Sprintf("  ✓ Has descriptive error message (%d chars)", len(precond.ErrorMessage)))
	}

	result.Success = true
	result.Flag = "flag{pr3c0nd1t10n_gu4rd14n_m4st3r}"
	result.Message = "✓ Precondition challenge completed! Your resource properly validates inputs before creation."
	return result
}

func validatePostconditionStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.Resources) == 0 && len(proof.DataSources) == 0 {
		result.Message = "No resources or data sources provided"
		result.Details = append(result.Details, "You must create a resource or data source with a postcondition")
		return result
	}

	// Check resources for postconditions
	var hasPostcondition bool
	var postconditions []ConditionBlock
	var resourceName string

	for i := range proof.Resources {
		r := &proof.Resources[i]
		if r.Lifecycle != nil && len(r.Lifecycle.Postconditions) > 0 {
			hasPostcondition = true
			postconditions = r.Lifecycle.Postconditions
			resourceName = r.ResourceName
			break
		}
	}

	// Also check data sources
	if !hasPostcondition {
		for i := range proof.DataSources {
			ds := &proof.DataSources[i]
			if ds.Lifecycle != nil && len(ds.Lifecycle.Postconditions) > 0 {
				hasPostcondition = true
				postconditions = ds.Lifecycle.Postconditions
				resourceName = ds.DataSourceName
				break
			}
		}
	}

	if !hasPostcondition {
		result.Message = "No postconditions found in lifecycle blocks"
		result.Details = append(result.Details, "Add a lifecycle block with postcondition to your resource/data source")
		result.Details = append(result.Details, "Postconditions validate resource state AFTER creation")
		return result
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ Found '%s' with postcondition", resourceName))

	// Validate each postcondition
	for i, postcond := range postconditions {
		result.Details = append(result.Details, fmt.Sprintf("Checking postcondition %d...", i+1))

		if postcond.Condition == "" {
			result.Message = fmt.Sprintf("Postcondition %d has empty condition", i+1)
			return result
		}
		result.Details = append(result.Details, fmt.Sprintf("  ✓ Has condition: %s", truncate(postcond.Condition, 60)))

		// Postconditions MUST use 'self' to reference the resource
		if !conditionUsesSelf(postcond.Condition) {
			result.Message = fmt.Sprintf("Postcondition %d does not use 'self'", i+1)
			result.Details = append(result.Details, "  ✗ Postconditions must use 'self' to reference resource attributes")
			result.Details = append(result.Details, "  Example: self.solved == true or self.status == \"active\"")
			return result
		}
		result.Details = append(result.Details, "  ✓ Uses 'self' to reference resource (correct for postcondition)")

		// Extract what attribute is being validated
		selfRefs := extractSelfReferences(postcond.Condition)
		if len(selfRefs) > 0 {
			result.Details = append(result.Details, fmt.Sprintf("  ✓ Validates attributes: %s", strings.Join(selfRefs, ", ")))
		}

		// Check error message
		if !validateErrorMessage(postcond.ErrorMessage, 10) {
			result.Message = "Postcondition error message too short"
			result.Details = append(result.Details, "  ✗ Error message must be descriptive (min 10 chars)")
			return result
		}
		result.Details = append(result.Details, "  ✓ Has descriptive error message")
	}

	result.Success = true
	result.Flag = "flag{p0stc0nd1t10n_v4l1d4t0r_3xp3rt}"
	result.Message = "✓ Postcondition challenge completed! Your configuration properly validates resource state after creation."
	return result
}

func validateCombinedConditionsStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.Resources) == 0 {
		result.Message = "No resources provided"
		result.Details = append(result.Details, "Create a resource with BOTH precondition and postcondition in lifecycle")
		return result
	}

	// Find a resource with both pre and post conditions
	var validResource *ResourceProof
	for i := range proof.Resources {
		r := &proof.Resources[i]
		if r.Lifecycle != nil &&
			len(r.Lifecycle.Preconditions) > 0 &&
			len(r.Lifecycle.Postconditions) > 0 {
			validResource = r
			break
		}
	}

	if validResource == nil {
		result.Message = "No resource found with both precondition and postcondition"
		result.Details = append(result.Details, "You must use BOTH in the same resource's lifecycle block")

		// Give specific feedback
		hasPre := false
		hasPost := false
		for i := range proof.Resources {
			r := &proof.Resources[i]
			if r.Lifecycle != nil {
				if len(r.Lifecycle.Preconditions) > 0 {
					hasPre = true
					result.Details = append(result.Details, fmt.Sprintf("  Resource '%s' has preconditions", r.ResourceName))
				}
				if len(r.Lifecycle.Postconditions) > 0 {
					hasPost = true
					result.Details = append(result.Details, fmt.Sprintf("  Resource '%s' has postconditions", r.ResourceName))
				}
			}
		}

		if hasPre && !hasPost {
			result.Details = append(result.Details, "Missing: postconditions (add postcondition block to validate after creation)")
		} else if hasPost && !hasPre {
			result.Details = append(result.Details, "Missing: preconditions (add precondition block to validate before creation)")
		}

		return result
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ Found resource '%s' with both condition types", validResource.ResourceName))

	// Validate preconditions don't use self
	for i, precond := range validResource.Lifecycle.Preconditions {
		if conditionUsesSelf(precond.Condition) {
			result.Message = fmt.Sprintf("Precondition %d incorrectly uses 'self'", i+1)
			result.Details = append(result.Details, "  ✗ Preconditions should NOT use 'self'")
			return result
		}
	}
	result.Details = append(result.Details, fmt.Sprintf("  ✓ %d precondition(s) correctly validate inputs", len(validResource.Lifecycle.Preconditions)))

	// Validate postconditions DO use self
	selfCount := 0
	for _, postcond := range validResource.Lifecycle.Postconditions {
		if conditionUsesSelf(postcond.Condition) {
			selfCount++
		}
	}

	if selfCount == 0 {
		result.Message = "Postconditions must use 'self' to reference resource attributes"
		result.Details = append(result.Details, "  ✗ None of your postconditions use 'self'")
		return result
	}
	result.Details = append(result.Details, fmt.Sprintf("  ✓ %d postcondition(s) correctly use 'self'", selfCount))

	// Verify error messages are descriptive
	for _, pre := range validResource.Lifecycle.Preconditions {
		if len(pre.ErrorMessage) < 10 {
			result.Message = "Precondition error messages must be descriptive (min 10 chars)"
			return result
		}
	}

	for _, post := range validResource.Lifecycle.Postconditions {
		if len(post.ErrorMessage) < 10 {
			result.Message = "Postcondition error messages must be descriptive (min 10 chars)"
			return result
		}
	}

	result.Details = append(result.Details, "  ✓ All error messages are descriptive")

	result.Success = true
	result.Flag = "flag{c0mb1n3d_c0nd1t10ns_m4st3r}"
	result.Message = "✓ Combined conditions challenge completed! You've mastered both pre and post validation."
	return result
}

func validateDataSourceConditionStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.DataSources) == 0 {
		result.Message = "No data sources provided"
		result.Details = append(result.Details, "Create a data source with postcondition to validate fetched data")
		return result
	}

	var validDataSource *DataSourceProof
	for i := range proof.DataSources {
		ds := &proof.DataSources[i]
		if ds.Lifecycle != nil && len(ds.Lifecycle.Postconditions) > 0 {
			validDataSource = ds
			break
		}
	}

	if validDataSource == nil {
		result.Message = "No data source with postconditions found"
		result.Details = append(result.Details, "Data sources use postconditions to validate fetched data")
		return result
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ Found data source '%s' with postconditions", validDataSource.DataSourceName))

	// Validate postconditions use self
	for i, postcond := range validDataSource.Lifecycle.Postconditions {
		if !conditionUsesSelf(postcond.Condition) {
			result.Message = fmt.Sprintf("Data source postcondition %d must use 'self'", i+1)
			result.Details = append(result.Details, "  ✗ Use 'self' to reference data source attributes")
			return result
		}

		selfRefs := extractSelfReferences(postcond.Condition)
		result.Details = append(result.Details, fmt.Sprintf("  ✓ Validates data attributes: %s", strings.Join(selfRefs, ", ")))

		if !validateErrorMessage(postcond.ErrorMessage, 15) {
			result.Message = "Data source error messages should be detailed (min 15 chars)"
			return result
		}
	}

	result.Success = true
	result.Flag = "flag{d4t4_s0urc3_v4l1d4t0r_pr0}"
	result.Message = "✓ Data source validation completed! You're ensuring data quality at query time."
	return result
}

func validateOutputConditionStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Message: "Output condition validation requires manual proof_of_work",
		Details: []string{"Use proof_of_work with uses_output_block, uses_precondition, etc."},
	}
	return result
}

func validateValidationChainStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.Resources) < 3 {
		result.Message = fmt.Sprintf("Validation chain requires at least 3 resources (found %d)", len(proof.Resources))
		result.Details = append(result.Details, "Create a chain of resources with interconnected conditions")
		return result
	}

	// Count total conditions
	totalConditions := 0
	for _, r := range proof.Resources {
		if r.Lifecycle != nil {
			totalConditions += len(r.Lifecycle.Preconditions)
			totalConditions += len(r.Lifecycle.Postconditions)
		}
	}

	if totalConditions < 4 {
		result.Message = fmt.Sprintf("Validation chain needs at least 4 conditions (found %d)", totalConditions)
		result.Details = append(result.Details, "Add more pre/postconditions across your resources")
		return result
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ %d resources in chain", len(proof.Resources)))
	result.Details = append(result.Details, fmt.Sprintf("✓ %d total conditions", totalConditions))

	// Check for depends_on usage
	hasDependencies := false
	for _, r := range proof.Resources {
		if deps, ok := r.MetaArguments["depends_on"]; ok && deps != "" {
			hasDependencies = true
			break
		}
	}

	if !hasDependencies {
		result.Message = "Validation chain should use depends_on for proper ordering"
		result.Details = append(result.Details, "  ✗ Add depends_on meta-argument to establish dependencies")
		return result
	}

	result.Details = append(result.Details, "✓ Uses depends_on for proper ordering")
	result.Success = true
	result.Flag = "flag{v4l1d4t10n_ch41n_4rch1t3ct}"
	result.Message = "✓ Validation chain completed! You've built an interconnected validation system."
	return result
}

func validateModuleContractStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if proof.Module == nil {
		result.Message = "No module proof provided"
		result.Details = append(result.Details, "Create a module with input/output validations")
		return result
	}

	m := proof.Module
	result.Details = append(result.Details, fmt.Sprintf("✓ Module '%s' provided", m.ModuleName))

	if len(m.InputValidations) < 2 {
		result.Message = fmt.Sprintf("Module needs at least 2 input validations (found %d)", len(m.InputValidations))
		result.Details = append(result.Details, "  ✗ Add variable validation blocks or preconditions")
		return result
	}
	result.Details = append(result.Details, fmt.Sprintf("  ✓ %d input validations", len(m.InputValidations)))

	if len(m.OutputValidations) < 2 {
		result.Message = fmt.Sprintf("Module needs at least 2 output validations (found %d)", len(m.OutputValidations))
		result.Details = append(result.Details, "  ✗ Add postconditions to outputs")
		return result
	}
	result.Details = append(result.Details, fmt.Sprintf("  ✓ %d output validations", len(m.OutputValidations)))

	// Verify error messages
	shortMessages := 0
	for _, val := range append(m.InputValidations, m.OutputValidations...) {
		if len(val.ErrorMessage) < 20 {
			shortMessages++
		}
	}

	if shortMessages > 0 {
		result.Message = fmt.Sprintf("%d validation(s) have error messages that are too short", shortMessages)
		result.Details = append(result.Details, "  ✗ Error messages should be consumer-friendly (min 20 chars)")
		return result
	}

	result.Details = append(result.Details, "  ✓ All error messages are consumer-friendly")
	result.Success = true
	result.Flag = "flag{m0dul3_c0ntr4ct_d3s1gn3r_m4st3r}"
	result.Message = "✓ Module contract completed! Your module has a clear, validated interface."
	return result
}

func validateSelfReferenceStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	if len(proof.Resources) == 0 {
		result.Message = "No resources provided"
		return result
	}

	// Find resource with multiple self references
	allSelfRefs := make(map[string]bool)
	postcondCount := 0

	for _, r := range proof.Resources {
		if r.Lifecycle != nil {
			for _, postcond := range r.Lifecycle.Postconditions {
				postcondCount++
				refs := extractSelfReferences(postcond.Condition)
				for _, ref := range refs {
					allSelfRefs[ref] = true
				}

				// Check for complex logic
				if strings.Contains(postcond.Condition, "&&") || strings.Contains(postcond.Condition, "||") {
					result.Details = append(result.Details, "  ✓ Uses complex boolean logic")
				}
			}
		}
	}

	if len(allSelfRefs) < 3 {
		result.Message = fmt.Sprintf("Must reference at least 3 different attributes with 'self' (found %d)", len(allSelfRefs))
		result.Details = append(result.Details, "Add more postconditions that validate different attributes")
		return result
	}

	if postcondCount < 2 {
		result.Message = "Must have at least 2 postconditions"
		return result
	}

	refList := make([]string, 0, len(allSelfRefs))
	for ref := range allSelfRefs {
		refList = append(refList, ref)
	}

	result.Details = append(result.Details, fmt.Sprintf("✓ %d unique self references: %s", len(refList), strings.Join(refList, ", ")))
	result.Details = append(result.Details, fmt.Sprintf("✓ %d postconditions", postcondCount))

	result.Success = true
	result.Flag = "flag{s3lf_r3f3r3nc3_m4st3r_pr0}"
	result.Message = "✓ Self-reference mastery achieved! You're validating multiple attributes effectively."
	return result
}

func validateConditionalValidationStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	hasAnd := false
	hasOr := false
	hasFunctions := false
	maxComplexity := 0

	for _, r := range proof.Resources {
		if r.Lifecycle != nil {
			for _, cond := range append(r.Lifecycle.Preconditions, r.Lifecycle.Postconditions...) {
				condition := cond.Condition

				if strings.Contains(condition, "&&") {
					hasAnd = true
				}
				if strings.Contains(condition, "||") {
					hasOr = true
				}

				// Check for function usage
				functions := []string{"length(", "can(", "try(", "contains(", "alltrue(", "anytrue("}
				for _, fn := range functions {
					if strings.Contains(condition, fn) {
						hasFunctions = true
						break
					}
				}

				// Rough complexity score
				complexity := strings.Count(condition, "&&") + strings.Count(condition, "||") + strings.Count(condition, "!")
				if complexity > maxComplexity {
					maxComplexity = complexity
				}
			}
		}
	}

	if !hasAnd {
		result.Message = "Must use && operator in conditions"
		return result
	}
	result.Details = append(result.Details, "✓ Uses && operator")

	if !hasOr {
		result.Message = "Must use || operator in conditions"
		return result
	}
	result.Details = append(result.Details, "✓ Uses || operator")

	if !hasFunctions {
		result.Message = "Must use Terraform functions in conditions"
		result.Details = append(result.Details, "  Try: length(), can(), try(), contains(), etc.")
		return result
	}
	result.Details = append(result.Details, "✓ Uses Terraform functions")

	if maxComplexity < 2 {
		result.Message = "Conditions not complex enough (need multiple operators)"
		return result
	}
	result.Details = append(result.Details, fmt.Sprintf("✓ Complexity score: %d/10", min(maxComplexity+3, 10)))

	result.Success = true
	result.Flag = "flag{c0nd1t10n4l_v4l1d4t10n_3xp3rt}"
	result.Message = "✓ Conditional validation mastery! Your logic is sophisticated and robust."
	return result
}

func validateErrorMessageStructure(proof *ProofData) ValidationResult {
	result := ValidationResult{
		Success: false,
		Details: []string{},
	}

	allMessages := []string{}
	hasInterpolation := false
	hasContext := false

	// Collect all error messages
	for _, r := range proof.Resources {
		if r.Lifecycle != nil {
			for _, cond := range append(r.Lifecycle.Preconditions, r.Lifecycle.Postconditions...) {
				allMessages = append(allMessages, cond.ErrorMessage)
			}
		}
	}

	for _, ds := range proof.DataSources {
		if ds.Lifecycle != nil {
			for _, cond := range append(ds.Lifecycle.Preconditions, ds.Lifecycle.Postconditions...) {
				allMessages = append(allMessages, cond.ErrorMessage)
			}
		}
	}

	if len(allMessages) < 3 {
		result.Message = fmt.Sprintf("Need at least 3 error messages (found %d)", len(allMessages))
		return result
	}

	// Check message quality
	shortCount := 0
	for _, msg := range allMessages {
		if len(msg) < 20 {
			shortCount++
		}

		// Check for interpolation (contains ${...})
		if strings.Contains(msg, "${") && strings.Contains(msg, "}") {
			hasInterpolation = true
		}

		// Check for context words
		contextWords := []string{"must", "should", "expected", "required", "invalid", "failed"}
		for _, word := range contextWords {
			if strings.Contains(strings.ToLower(msg), word) {
				hasContext = true
				break
			}
		}
	}

	if shortCount > 0 {
		result.Message = fmt.Sprintf("%d error message(s) too short (min 20 chars)", shortCount)
		return result
	}
	result.Details = append(result.Details, fmt.Sprintf("✓ %d error messages, all descriptive", len(allMessages)))

	if !hasInterpolation {
		result.Message = "Error messages should use interpolation to show actual values"
		result.Details = append(result.Details, "  Example: \"Value must be positive, got ${var.count}\"")
		return result
	}
	result.Details = append(result.Details, "✓ Uses interpolation to show actual values")

	if !hasContext {
		result.Message = "Error messages should provide context about what failed"
		return result
	}
	result.Details = append(result.Details, "✓ Messages provide helpful context")

	result.Success = true
	result.Flag = "flag{3rr0r_m3ss4g3_d3s1gn3r_pr0}"
	result.Message = "✓ Error message design mastered! Your messages guide users effectively."
	return result
}

// Helper functions

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func extractSelfReferences(condition string) []string {
	refs := []string{}
	parts := strings.Split(condition, "self.")
	for i := 1; i < len(parts); i++ {
		// Extract the attribute name (up to space, operator, or parenthesis)
		attr := ""
		for _, ch := range parts[i] {
			if ch == ' ' || ch == '=' || ch == '!' || ch == '>' || ch == '<' || ch == ')' || ch == ',' {
				break
			}
			attr += string(ch)
		}
		if attr != "" {
			refs = append(refs, "self."+attr)
		}
	}
	return refs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Legacy validators (for backward compatibility with manual proof_of_work)
func validatePreconditionChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validatePostconditionChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateCombinedConditionsChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateDataSourceConditionChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateOutputConditionChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateValidationChainChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateModuleContractChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use module_proof instead of manual proof_of_work")
}

func validateSelfReferenceChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateConditionalValidationChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}

func validateErrorMessageChallenge(proof map[string]interface{}) (bool, string, error) {
	return false, "", fmt.Errorf("use resource_proof with lifecycle configuration instead of manual proof_of_work")
}
