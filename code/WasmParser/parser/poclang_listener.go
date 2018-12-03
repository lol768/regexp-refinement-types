// Code generated from PocLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PocLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PocLangListener is a complete listener for a parse tree produced by PocLang.
type PocLangListener interface {
	antlr.ParseTreeListener

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterFunction_body is called when entering the function_body production.
	EnterFunction_body(c *Function_bodyContext)

	// EnterBody_line is called when entering the body_line production.
	EnterBody_line(c *Body_lineContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterOptional_else is called when entering the optional_else production.
	EnterOptional_else(c *Optional_elseContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterUnsignedIntType is called when entering the UnsignedIntType production.
	EnterUnsignedIntType(c *UnsignedIntTypeContext)

	// EnterStringType is called when entering the StringType production.
	EnterStringType(c *StringTypeContext)

	// EnterVoidType is called when entering the VoidType production.
	EnterVoidType(c *VoidTypeContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterType_specifier is called when entering the type_specifier production.
	EnterType_specifier(c *Type_specifierContext)

	// EnterLessThanConstraint is called when entering the LessThanConstraint production.
	EnterLessThanConstraint(c *LessThanConstraintContext)

	// EnterGreaterThanConstraint is called when entering the GreaterThanConstraint production.
	EnterGreaterThanConstraint(c *GreaterThanConstraintContext)

	// EnterLessThanEqualsConstraint is called when entering the LessThanEqualsConstraint production.
	EnterLessThanEqualsConstraint(c *LessThanEqualsConstraintContext)

	// EnterGreaterThanEqualsConstraint is called when entering the GreaterThanEqualsConstraint production.
	EnterGreaterThanEqualsConstraint(c *GreaterThanEqualsConstraintContext)

	// EnterString_constraint is called when entering the string_constraint production.
	EnterString_constraint(c *String_constraintContext)

	// EnterFunction_sig is called when entering the function_sig production.
	EnterFunction_sig(c *Function_sigContext)

	// EnterArgument_decl is called when entering the argument_decl production.
	EnterArgument_decl(c *Argument_declContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterVar_assignment is called when entering the var_assignment production.
	EnterVar_assignment(c *Var_assignmentContext)

	// EnterValue_ref is called when entering the value_ref production.
	EnterValue_ref(c *Value_refContext)

	// EnterIdentifier_ref is called when entering the identifier_ref production.
	EnterIdentifier_ref(c *Identifier_refContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterRe is called when entering the re production.
	EnterRe(c *ReContext)

	// EnterUnion_prime is called when entering the union_prime production.
	EnterUnion_prime(c *Union_primeContext)

	// EnterSimple_re is called when entering the simple_re production.
	EnterSimple_re(c *Simple_reContext)

	// EnterConcat_prime is called when entering the concat_prime production.
	EnterConcat_prime(c *Concat_primeContext)

	// EnterBasic_re is called when entering the basic_re production.
	EnterBasic_re(c *Basic_reContext)

	// EnterKleene_star is called when entering the kleene_star production.
	EnterKleene_star(c *Kleene_starContext)

	// EnterPlus is called when entering the plus production.
	EnterPlus(c *PlusContext)

	// EnterElementary_re is called when entering the elementary_re production.
	EnterElementary_re(c *Elementary_reContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterRange_re is called when entering the range_re production.
	EnterRange_re(c *Range_reContext)

	// EnterPositive_range is called when entering the positive_range production.
	EnterPositive_range(c *Positive_rangeContext)

	// EnterNegative_range is called when entering the negative_range production.
	EnterNegative_range(c *Negative_rangeContext)

	// EnterLax_character is called when entering the lax_character production.
	EnterLax_character(c *Lax_characterContext)

	// EnterCharacter is called when entering the character production.
	EnterCharacter(c *CharacterContext)

	// EnterRange_items is called when entering the range_items production.
	EnterRange_items(c *Range_itemsContext)

	// EnterRange_item is called when entering the range_item production.
	EnterRange_item(c *Range_itemContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitFunction_body is called when exiting the function_body production.
	ExitFunction_body(c *Function_bodyContext)

	// ExitBody_line is called when exiting the body_line production.
	ExitBody_line(c *Body_lineContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitOptional_else is called when exiting the optional_else production.
	ExitOptional_else(c *Optional_elseContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitUnsignedIntType is called when exiting the UnsignedIntType production.
	ExitUnsignedIntType(c *UnsignedIntTypeContext)

	// ExitStringType is called when exiting the StringType production.
	ExitStringType(c *StringTypeContext)

	// ExitVoidType is called when exiting the VoidType production.
	ExitVoidType(c *VoidTypeContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitType_specifier is called when exiting the type_specifier production.
	ExitType_specifier(c *Type_specifierContext)

	// ExitLessThanConstraint is called when exiting the LessThanConstraint production.
	ExitLessThanConstraint(c *LessThanConstraintContext)

	// ExitGreaterThanConstraint is called when exiting the GreaterThanConstraint production.
	ExitGreaterThanConstraint(c *GreaterThanConstraintContext)

	// ExitLessThanEqualsConstraint is called when exiting the LessThanEqualsConstraint production.
	ExitLessThanEqualsConstraint(c *LessThanEqualsConstraintContext)

	// ExitGreaterThanEqualsConstraint is called when exiting the GreaterThanEqualsConstraint production.
	ExitGreaterThanEqualsConstraint(c *GreaterThanEqualsConstraintContext)

	// ExitString_constraint is called when exiting the string_constraint production.
	ExitString_constraint(c *String_constraintContext)

	// ExitFunction_sig is called when exiting the function_sig production.
	ExitFunction_sig(c *Function_sigContext)

	// ExitArgument_decl is called when exiting the argument_decl production.
	ExitArgument_decl(c *Argument_declContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitVar_assignment is called when exiting the var_assignment production.
	ExitVar_assignment(c *Var_assignmentContext)

	// ExitValue_ref is called when exiting the value_ref production.
	ExitValue_ref(c *Value_refContext)

	// ExitIdentifier_ref is called when exiting the identifier_ref production.
	ExitIdentifier_ref(c *Identifier_refContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitRe is called when exiting the re production.
	ExitRe(c *ReContext)

	// ExitUnion_prime is called when exiting the union_prime production.
	ExitUnion_prime(c *Union_primeContext)

	// ExitSimple_re is called when exiting the simple_re production.
	ExitSimple_re(c *Simple_reContext)

	// ExitConcat_prime is called when exiting the concat_prime production.
	ExitConcat_prime(c *Concat_primeContext)

	// ExitBasic_re is called when exiting the basic_re production.
	ExitBasic_re(c *Basic_reContext)

	// ExitKleene_star is called when exiting the kleene_star production.
	ExitKleene_star(c *Kleene_starContext)

	// ExitPlus is called when exiting the plus production.
	ExitPlus(c *PlusContext)

	// ExitElementary_re is called when exiting the elementary_re production.
	ExitElementary_re(c *Elementary_reContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitRange_re is called when exiting the range_re production.
	ExitRange_re(c *Range_reContext)

	// ExitPositive_range is called when exiting the positive_range production.
	ExitPositive_range(c *Positive_rangeContext)

	// ExitNegative_range is called when exiting the negative_range production.
	ExitNegative_range(c *Negative_rangeContext)

	// ExitLax_character is called when exiting the lax_character production.
	ExitLax_character(c *Lax_characterContext)

	// ExitCharacter is called when exiting the character production.
	ExitCharacter(c *CharacterContext)

	// ExitRange_items is called when exiting the range_items production.
	ExitRange_items(c *Range_itemsContext)

	// ExitRange_item is called when exiting the range_item production.
	ExitRange_item(c *Range_itemContext)
}
