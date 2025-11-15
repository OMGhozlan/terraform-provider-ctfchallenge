package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registerValidationConditionChallenges()
}

func registerValidationConditionChallenges() {
	// Precondition Basics
	Challenges["precondition_guardian"] = Challenge{
		ID:          "precondition_guardian",
		Name:        "Precondition Guardian",
		Description: "Use preconditions to validate inputs before resource creation",
		Points:      150,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validatePreconditionChallenge,
	}

	// Postcondition Basics
	Challenges["postcondition_validator"] = Challenge{
		ID:          "postcondition_validator",
		Name:        "Postcondition Validator",
		Description: "Use postconditions with 'self' to validate resource attributes after creation",
		Points:      175,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validatePostconditionChallenge,
	}

	// Combined Pre/Post Conditions
	Challenges["condition_master"] = Challenge{
		ID:          "condition_master",
		Name:        "Condition Master",
		Description: "Combine preconditions and postconditions in a single resource",
		Points:      200,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateCombinedConditionsChallenge,
	}

	// Data Source Validation
	Challenges["data_validator"] = Challenge{
		ID:          "data_validator",
		Name:        "Data Source Validator",
		Description: "Use postconditions to validate data source outputs",
		Points:      160,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateDataSourceConditionChallenge,
	}

	// Output Validation
	Challenges["output_contract"] = Challenge{
		ID:          "output_contract",
		Name:        "Output Contract Enforcer",
		Description: "Use preconditions in output blocks to enforce module contracts",
		Points:      180,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateOutputConditionChallenge,
	}

	// Complex Validation Chain
	Challenges["validation_chain"] = Challenge{
		ID:          "validation_chain",
		Name:        "Validation Chain Architect",
		Description: "Create a chain of resources with interconnected pre/postconditions",
		Points:      250,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateValidationChainChallenge,
	}

	// Module Contract Challenge
	Challenges["module_contract"] = Challenge{
		ID:          "module_contract",
		Name:        "Module Contract Designer",
		Description: "Design a module with comprehensive pre/postconditions for input validation and output guarantees",
		Points:      300,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateModuleContractChallenge,
	}

	// Self-Reference Master
	Challenges["self_reference_master"] = Challenge{
		ID:          "self_reference_master",
		Name:        "Self-Reference Master",
		Description: "Master the use of 'self' in postconditions to validate multiple attributes",
		Points:      190,
		Difficulty:  "intermediate",
		Category:    "validation",
		Validator:   validateSelfReferenceChallenge,
	}

	// Conditional Validation
	Challenges["conditional_validation"] = Challenge{
		ID:          "conditional_validation",
		Name:        "Conditional Validation Expert",
		Description: "Use complex boolean logic in condition blocks with multiple checks",
		Points:      220,
		Difficulty:  "advanced",
		Category:    "validation",
		Validator:   validateConditionalValidationChallenge,
	}

	// Error Message Designer
	Challenges["error_message_designer"] = Challenge{
		ID:          "error_message_designer",
		Name:        "Error Message Designer",
		Description: "Create helpful, informative error messages for all validation failures",
		Points:      140,
		Difficulty:  "beginner",
		Category:    "validation",
		Validator:   validateErrorMessageChallenge,
	}
}

func validatePreconditionChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for precondition usage
	usesPrecondition, _ := proof["uses_precondition"].(string)
	if usesPrecondition != "true" {
		return false, "", fmt.Errorf("you must use a precondition block in lifecycle")
	}

	// Check for condition expression
	conditionExpr, hasCondition := proof["condition_expression"].(string)
	if !hasCondition || conditionExpr == "" {
		return false, "", fmt.Errorf("missing 'condition_expression' - provide your condition logic")
	}

	// Verify it checks input before resource creation
	checksInput, _ := proof["checks_input"].(string)
	if checksInput != "true" {
		return false, "", fmt.Errorf("precondition must validate input values before resource creation")
	}

	// Check for error message
	errorMessage, hasError := proof["error_message"].(string)
	if !hasError || len(errorMessage) < 10 {
		return false, "", fmt.Errorf("provide a descriptive error_message (min 10 characters)")
	}

	// Verify placement in lifecycle block
	inLifecycle, _ := proof["in_lifecycle_block"].(string)
	if inLifecycle != "true" {
		return false, "", fmt.Errorf("precondition must be placed in a lifecycle block")
	}

	// Check what is being validated
	validatesWhat, hasValidates := proof["validates"].(string)
	if !hasValidates {
		return false, "", fmt.Errorf("specify what you're validating (e.g., 'variable', 'input', 'parameter')")
	}

	return true, "flag{pr3c0nd1t10n_gu4rd14n_m4st3r}", nil
}

func validatePostconditionChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for postcondition usage
	usesPostcondition, _ := proof["uses_postcondition"].(string)
	if usesPostcondition != "true" {
		return false, "", fmt.Errorf("you must use a postcondition block in lifecycle")
	}

	// Check for self reference
	usesSelf, _ := proof["uses_self"].(string)
	if usesSelf != "true" {
		return false, "", fmt.Errorf("postcondition must use 'self' to reference the resource")
	}

	// Check what attribute is validated
	validatedAttr, hasAttr := proof["validated_attribute"].(string)
	if !hasAttr || validatedAttr == "" {
		return false, "", fmt.Errorf("specify which attribute you validated with 'self' (e.g., 'self.solved')")
	}

	// Verify it validates after creation
	validatesAfter, _ := proof["validates_after_creation"].(string)
	if validatesAfter != "true" {
		return false, "", fmt.Errorf("postcondition validates the resource state AFTER creation")
	}

	// Check for error message
	errorMessage, hasError := proof["error_message"].(string)
	if !hasError || len(errorMessage) < 10 {
		return false, "", fmt.Errorf("provide a descriptive error_message (min 10 characters)")
	}

	// Verify placement in lifecycle block
	inLifecycle, _ := proof["in_lifecycle_block"].(string)
	if inLifecycle != "true" {
		return false, "", fmt.Errorf("postcondition must be placed in a lifecycle block")
	}

	return true, "flag{p0stc0nd1t10n_v4l1d4t0r_3xp3rt}", nil
}

func validateCombinedConditionsChallenge(proof map[string]interface{}) (bool, string, error) {
	// Must have both precondition and postcondition
	usesPre, _ := proof["uses_precondition"].(string)
	usesPost, _ := proof["uses_postcondition"].(string)

	if usesPre != "true" || usesPost != "true" {
		return false, "", fmt.Errorf("you must use BOTH precondition and postcondition in the same resource")
	}

	// Check for self in postcondition only
	usesSelf, _ := proof["postcondition_uses_self"].(string)
	if usesSelf != "true" {
		return false, "", fmt.Errorf("postcondition must use 'self' to reference resource attributes")
	}

	// Verify precondition doesn't use self
	preUsesSelf, _ := proof["precondition_uses_self"].(string)
	if preUsesSelf == "true" {
		return false, "", fmt.Errorf("precondition should NOT use 'self' (resource doesn't exist yet)")
	}

	// Check that they validate different things
	preValidates, hasPre := proof["precondition_validates"].(string)
	postValidates, hasPost := proof["postcondition_validates"].(string)

	if !hasPre || !hasPost {
		return false, "", fmt.Errorf("specify what each condition validates")
	}

	if preValidates == postValidates {
		return false, "", fmt.Errorf("precondition and postcondition should validate different aspects")
	}

	// Check for distinct error messages
	preError, hasPreError := proof["precondition_error"].(string)
	postError, hasPostError := proof["postcondition_error"].(string)

	if !hasPreError || !hasPostError {
		return false, "", fmt.Errorf("provide error messages for both conditions")
	}

	if len(preError) < 10 || len(postError) < 10 {
		return false, "", fmt.Errorf("both error messages must be descriptive (min 10 characters)")
	}

	return true, "flag{c0mb1n3d_c0nd1t10ns_m4st3r}", nil
}

func validateDataSourceConditionChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for data source usage
	usesDataSource, _ := proof["uses_data_source"].(string)
	if usesDataSource != "true" {
		return false, "", fmt.Errorf("you must use a data source block")
	}

	// Data sources typically use postconditions
	usesPostcondition, _ := proof["uses_postcondition"].(string)
	if usesPostcondition != "true" {
		return false, "", fmt.Errorf("data sources should use postconditions to validate fetched data")
	}

	// Check for self reference to data attributes
	usesSelf, _ := proof["uses_self"].(string)
	if usesSelf != "true" {
		return false, "", fmt.Errorf("use 'self' to reference data source attributes")
	}

	// Verify what data attribute is validated
	validatedAttr, hasAttr := proof["validated_data_attribute"].(string)
	if !hasAttr || validatedAttr == "" {
		return false, "", fmt.Errorf("specify which data attribute you validated")
	}

	// Check validation purpose
	validationPurpose, hasPurpose := proof["validation_purpose"].(string)
	if !hasPurpose || len(validationPurpose) < 15 {
		return false, "", fmt.Errorf("explain why this data validation is important (min 15 chars)")
	}

	return true, "flag{d4t4_s0urc3_v4l1d4t0r_pr0}", nil
}

func validateOutputConditionChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for output block usage
	usesOutput, _ := proof["uses_output_block"].(string)
	if usesOutput != "true" {
		return false, "", fmt.Errorf("you must use an output block with conditions")
	}

	// Outputs typically use preconditions to validate before outputting
	usesPrecondition, _ := proof["uses_precondition"].(string)
	if usesPrecondition != "true" {
		return false, "", fmt.Errorf("output blocks typically use preconditions to validate before output")
	}

	// Check what is being validated
	validatesWhat, hasValidates := proof["validates_before_output"].(string)
	if !hasValidates || validatesWhat == "" {
		return false, "", fmt.Errorf("specify what you're validating before output")
	}

	// Check for module contract enforcement
	enforcesContract, _ := proof["enforces_module_contract"].(string)
	if enforcesContract != "true" {
		return false, "", fmt.Errorf("output conditions should enforce module contracts")
	}

	// Verify error message helps module consumers
	errorMessage, hasError := proof["consumer_friendly_error"].(string)
	if !hasError || len(errorMessage) < 20 {
		return false, "", fmt.Errorf("provide a consumer-friendly error message (min 20 chars)")
	}

	return true, "flag{0utput_c0ntr4ct_3nf0rc3r}", nil
}

func validateValidationChainChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for multiple resources
	resourceCount, hasCount := proof["resource_count"].(string)
	if !hasCount {
		return false, "", fmt.Errorf("specify how many resources are in your validation chain")
	}

	count, err := strconv.Atoi(resourceCount)
	if err != nil || count < 3 {
		return false, "", fmt.Errorf("validation chain must have at least 3 resources")
	}

	// Check for condition count
	totalConditions, hasConditions := proof["total_conditions"].(string)
	if !hasConditions {
		return false, "", fmt.Errorf("specify total number of conditions (pre + post)")
	}

	conditions, err := strconv.Atoi(totalConditions)
	if err != nil || conditions < 4 {
		return false, "", fmt.Errorf("must have at least 4 total conditions in the chain")
	}

	// Verify conditions are interconnected
	interconnected, _ := proof["conditions_interconnected"].(string)
	if interconnected != "true" {
		return false, "", fmt.Errorf("conditions must validate outputs from previous resources")
	}

	// Check for validation flow documentation
	validationFlow, hasFlow := proof["validation_flow"].(string)
	if !hasFlow || len(validationFlow) < 30 {
		return false, "", fmt.Errorf("document your validation flow (min 30 chars)")
	}

	// Verify use of depends_on for proper ordering
	usesDependsOn, _ := proof["uses_depends_on"].(string)
	if usesDependsOn != "true" {
		return false, "", fmt.Errorf("use depends_on to ensure proper validation order")
	}

	return true, "flag{v4l1d4t10n_ch41n_4rch1t3ct}", nil
}

func validateModuleContractChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for module creation
	createdModule, _ := proof["created_module"].(string)
	if createdModule != "true" {
		return false, "", fmt.Errorf("you must create a reusable module")
	}

	// Check for input validation (preconditions on variables)
	inputValidations, hasInput := proof["input_validations"].(string)
	if !hasInput {
		return false, "", fmt.Errorf("specify number of input validations (variable preconditions)")
	}

	inputs, err := strconv.Atoi(inputValidations)
	if err != nil || inputs < 2 {
		return false, "", fmt.Errorf("module must validate at least 2 inputs")
	}

	// Check for output guarantees (postconditions)
	outputGuarantees, hasOutput := proof["output_guarantees"].(string)
	if !hasOutput {
		return false, "", fmt.Errorf("specify number of output guarantees (postconditions)")
	}

	outputs, err := strconv.Atoi(outputGuarantees)
	if err != nil || outputs < 2 {
		return false, "", fmt.Errorf("module must guarantee at least 2 outputs with postconditions")
	}

	// Check for module documentation
	moduleDoc, hasDoc := proof["module_documentation"].(string)
	if !hasDoc || len(moduleDoc) < 50 {
		return false, "", fmt.Errorf("provide comprehensive module documentation (min 50 chars)")
	}

	// Verify contract clarity
	contractClear, _ := proof["contract_clearly_defined"].(string)
	if contractClear != "true" {
		return false, "", fmt.Errorf("module contract must be clearly defined with conditions")
	}

	// Check for helpful error messages
	errorMessagesCount, hasErrors := proof["helpful_error_messages"].(string)
	if !hasErrors {
		return false, "", fmt.Errorf("specify number of helpful error messages")
	}

	errors, err := strconv.Atoi(errorMessagesCount)
	if err != nil || errors < 4 {
		return false, "", fmt.Errorf("provide at least 4 helpful error messages for consumers")
	}

	return true, "flag{m0dul3_c0ntr4ct_d3s1gn3r_m4st3r}", nil
}

func validateSelfReferenceChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for multiple self references
	selfReferences, hasRefs := proof["self_references"].(string)
	if !hasRefs {
		return false, "", fmt.Errorf("specify how many 'self' references you used")
	}

	refs, err := strconv.Atoi(selfReferences)
	if err != nil || refs < 3 {
		return false, "", fmt.Errorf("must use 'self' to reference at least 3 different attributes")
	}

	// Check which attributes were validated
	attributes, hasAttrs := proof["validated_attributes"].(string)
	if !hasAttrs {
		return false, "", fmt.Errorf("list the attributes you validated with 'self'")
	}

	attrList := strings.Split(attributes, ",")
	if len(attrList) < 3 {
		return false, "", fmt.Errorf("must validate at least 3 attributes")
	}

	// Verify complex validation logic
	usesComplexLogic, _ := proof["uses_complex_logic"].(string)
	if usesComplexLogic != "true" {
		return false, "", fmt.Errorf("use complex boolean logic (&&, ||, !) in your conditions")
	}

	// Check for multiple postconditions
	postconditionCount, hasPost := proof["postcondition_count"].(string)
	if !hasPost {
		return false, "", fmt.Errorf("specify number of postconditions")
	}

	postCount, err := strconv.Atoi(postconditionCount)
	if err != nil || postCount < 2 {
		return false, "", fmt.Errorf("use at least 2 postconditions with different 'self' validations")
	}

	return true, "flag{s3lf_r3f3r3nc3_m4st3r_pr0}", nil
}

func validateConditionalValidationChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for complex boolean expressions
	usesComplexBool, _ := proof["uses_complex_boolean"].(string)
	if usesComplexBool != "true" {
		return false, "", fmt.Errorf("use complex boolean expressions with multiple operators")
	}

	// Check which operators were used
	operators, hasOps := proof["boolean_operators"].(string)
	if !hasOps {
		return false, "", fmt.Errorf("list boolean operators used (&&, ||, !)")
	}

	requiredOps := map[string]bool{
		"&&": false,
		"||": false,
	}

	opList := strings.Split(operators, ",")
	for _, op := range opList {
		op = strings.TrimSpace(op)
		if _, exists := requiredOps[op]; exists {
			requiredOps[op] = true
		}
	}

	for op, used := range requiredOps {
		if !used {
			return false, "", fmt.Errorf("must use operator: %s", op)
		}
	}

	// Check for multiple conditions in single expression
	multipleChecks, hasMulti := proof["multiple_checks_in_condition"].(string)
	if !hasMulti {
		return false, "", fmt.Errorf("specify number of checks in your condition expression")
	}

	checks, err := strconv.Atoi(multipleChecks)
	if err != nil || checks < 3 {
		return false, "", fmt.Errorf("condition must check at least 3 things")
	}

	// Verify use of functions in conditions
	usesFunctions, _ := proof["uses_functions"].(string)
	if usesFunctions != "true" {
		return false, "", fmt.Errorf("use Terraform functions in your conditions (length, can, try, etc.)")
	}

	// Check for validation complexity
	complexityScore, hasScore := proof["complexity_score"].(string)
	if !hasScore {
		return false, "", fmt.Errorf("rate complexity of your validation (1-10)")
	}

	score, err := strconv.Atoi(complexityScore)
	if err != nil || score < 7 {
		return false, "", fmt.Errorf("validation must be sufficiently complex (score >= 7)")
	}

	return true, "flag{c0nd1t10n4l_v4l1d4t10n_3xp3rt}", nil
}

func validateErrorMessageChallenge(proof map[string]interface{}) (bool, string, error) {
	// Check for multiple error messages
	errorCount, hasCount := proof["error_message_count"].(string)
	if !hasCount {
		return false, "", fmt.Errorf("specify number of error messages created")
	}

	count, err := strconv.Atoi(errorCount)
	if err != nil || count < 3 {
		return false, "", fmt.Errorf("must create at least 3 error messages")
	}

	// Check error message quality
	messages, hasMessages := proof["error_messages"].(string)
	if !hasMessages {
		return false, "", fmt.Errorf("provide your error messages (separated by |)")
	}

	msgList := strings.Split(messages, "|")
	if len(msgList) < 3 {
		return false, "", fmt.Errorf("must provide at least 3 error messages")
	}

	// Verify messages are descriptive
	for i, msg := range msgList {
		if len(strings.TrimSpace(msg)) < 20 {
			return false, "", fmt.Errorf("error message %d is too short (min 20 chars)", i+1)
		}
	}

	// Check for helpful elements in messages
	includesContext, _ := proof["includes_context"].(string)
	if includesContext != "true" {
		return false, "", fmt.Errorf("error messages must include context about what failed")
	}

	includesSolution, _ := proof["includes_solution_hint"].(string)
	if includesSolution != "true" {
		return false, "", fmt.Errorf("error messages should hint at how to fix the issue")
	}

	// Verify messages use interpolation
	usesInterpolation, _ := proof["uses_interpolation"].(string)
	if usesInterpolation != "true" {
		return false, "", fmt.Errorf("use string interpolation in error messages to show actual values")
	}

	return true, "flag{3rr0r_m3ss4g3_d3s1gn3r_pr0}", nil
}
