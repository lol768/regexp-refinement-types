// Code generated from PocLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PocLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePocLangListener is a complete listener for a parse tree produced by PocLang.
type BasePocLangListener struct{}

var _ PocLangListener = &BasePocLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePocLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePocLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePocLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePocLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBody is called when production body is entered.
func (s *BasePocLangListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasePocLangListener) ExitBody(ctx *BodyContext) {}

// EnterFunction_body is called when production function_body is entered.
func (s *BasePocLangListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BasePocLangListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterBody_line is called when production body_line is entered.
func (s *BasePocLangListener) EnterBody_line(ctx *Body_lineContext) {}

// ExitBody_line is called when production body_line is exited.
func (s *BasePocLangListener) ExitBody_line(ctx *Body_lineContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BasePocLangListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BasePocLangListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BasePocLangListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BasePocLangListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterOptional_else is called when production optional_else is entered.
func (s *BasePocLangListener) EnterOptional_else(ctx *Optional_elseContext) {}

// ExitOptional_else is called when production optional_else is exited.
func (s *BasePocLangListener) ExitOptional_else(ctx *Optional_elseContext) {}

// EnterVar_decl is called when production var_decl is entered.
func (s *BasePocLangListener) EnterVar_decl(ctx *Var_declContext) {}

// ExitVar_decl is called when production var_decl is exited.
func (s *BasePocLangListener) ExitVar_decl(ctx *Var_declContext) {}

// EnterUnsignedIntType is called when production UnsignedIntType is entered.
func (s *BasePocLangListener) EnterUnsignedIntType(ctx *UnsignedIntTypeContext) {}

// ExitUnsignedIntType is called when production UnsignedIntType is exited.
func (s *BasePocLangListener) ExitUnsignedIntType(ctx *UnsignedIntTypeContext) {}

// EnterStringType is called when production StringType is entered.
func (s *BasePocLangListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production StringType is exited.
func (s *BasePocLangListener) ExitStringType(ctx *StringTypeContext) {}

// EnterVoidType is called when production VoidType is entered.
func (s *BasePocLangListener) EnterVoidType(ctx *VoidTypeContext) {}

// ExitVoidType is called when production VoidType is exited.
func (s *BasePocLangListener) ExitVoidType(ctx *VoidTypeContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BasePocLangListener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BasePocLangListener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BasePocLangListener) EnterType_specifier(ctx *Type_specifierContext) {}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BasePocLangListener) ExitType_specifier(ctx *Type_specifierContext) {}

// EnterLessThanConstraint is called when production LessThanConstraint is entered.
func (s *BasePocLangListener) EnterLessThanConstraint(ctx *LessThanConstraintContext) {}

// ExitLessThanConstraint is called when production LessThanConstraint is exited.
func (s *BasePocLangListener) ExitLessThanConstraint(ctx *LessThanConstraintContext) {}

// EnterGreaterThanConstraint is called when production GreaterThanConstraint is entered.
func (s *BasePocLangListener) EnterGreaterThanConstraint(ctx *GreaterThanConstraintContext) {}

// ExitGreaterThanConstraint is called when production GreaterThanConstraint is exited.
func (s *BasePocLangListener) ExitGreaterThanConstraint(ctx *GreaterThanConstraintContext) {}

// EnterLessThanEqualsConstraint is called when production LessThanEqualsConstraint is entered.
func (s *BasePocLangListener) EnterLessThanEqualsConstraint(ctx *LessThanEqualsConstraintContext) {}

// ExitLessThanEqualsConstraint is called when production LessThanEqualsConstraint is exited.
func (s *BasePocLangListener) ExitLessThanEqualsConstraint(ctx *LessThanEqualsConstraintContext) {}

// EnterGreaterThanEqualsConstraint is called when production GreaterThanEqualsConstraint is entered.
func (s *BasePocLangListener) EnterGreaterThanEqualsConstraint(ctx *GreaterThanEqualsConstraintContext) {
}

// ExitGreaterThanEqualsConstraint is called when production GreaterThanEqualsConstraint is exited.
func (s *BasePocLangListener) ExitGreaterThanEqualsConstraint(ctx *GreaterThanEqualsConstraintContext) {
}

// EnterString_constraint is called when production string_constraint is entered.
func (s *BasePocLangListener) EnterString_constraint(ctx *String_constraintContext) {}

// ExitString_constraint is called when production string_constraint is exited.
func (s *BasePocLangListener) ExitString_constraint(ctx *String_constraintContext) {}

// EnterFunction_sig is called when production function_sig is entered.
func (s *BasePocLangListener) EnterFunction_sig(ctx *Function_sigContext) {}

// ExitFunction_sig is called when production function_sig is exited.
func (s *BasePocLangListener) ExitFunction_sig(ctx *Function_sigContext) {}

// EnterArgument_decl is called when production argument_decl is entered.
func (s *BasePocLangListener) EnterArgument_decl(ctx *Argument_declContext) {}

// ExitArgument_decl is called when production argument_decl is exited.
func (s *BasePocLangListener) ExitArgument_decl(ctx *Argument_declContext) {}

// EnterFunction is called when production function is entered.
func (s *BasePocLangListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasePocLangListener) ExitFunction(ctx *FunctionContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePocLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePocLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePocLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePocLangListener) ExitExpr(ctx *ExprContext) {}

// EnterVar_assignment is called when production var_assignment is entered.
func (s *BasePocLangListener) EnterVar_assignment(ctx *Var_assignmentContext) {}

// ExitVar_assignment is called when production var_assignment is exited.
func (s *BasePocLangListener) ExitVar_assignment(ctx *Var_assignmentContext) {}

// EnterValue_ref is called when production value_ref is entered.
func (s *BasePocLangListener) EnterValue_ref(ctx *Value_refContext) {}

// ExitValue_ref is called when production value_ref is exited.
func (s *BasePocLangListener) ExitValue_ref(ctx *Value_refContext) {}

// EnterIdentifier_ref is called when production identifier_ref is entered.
func (s *BasePocLangListener) EnterIdentifier_ref(ctx *Identifier_refContext) {}

// ExitIdentifier_ref is called when production identifier_ref is exited.
func (s *BasePocLangListener) ExitIdentifier_ref(ctx *Identifier_refContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BasePocLangListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BasePocLangListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterJava_call is called when production java_call is entered.
func (s *BasePocLangListener) EnterJava_call(ctx *Java_callContext) {}

// ExitJava_call is called when production java_call is exited.
func (s *BasePocLangListener) ExitJava_call(ctx *Java_callContext) {}

// EnterRe is called when production re is entered.
func (s *BasePocLangListener) EnterRe(ctx *ReContext) {}

// ExitRe is called when production re is exited.
func (s *BasePocLangListener) ExitRe(ctx *ReContext) {}

// EnterUnion_prime is called when production union_prime is entered.
func (s *BasePocLangListener) EnterUnion_prime(ctx *Union_primeContext) {}

// ExitUnion_prime is called when production union_prime is exited.
func (s *BasePocLangListener) ExitUnion_prime(ctx *Union_primeContext) {}

// EnterSimple_re is called when production simple_re is entered.
func (s *BasePocLangListener) EnterSimple_re(ctx *Simple_reContext) {}

// ExitSimple_re is called when production simple_re is exited.
func (s *BasePocLangListener) ExitSimple_re(ctx *Simple_reContext) {}

// EnterConcat_prime is called when production concat_prime is entered.
func (s *BasePocLangListener) EnterConcat_prime(ctx *Concat_primeContext) {}

// ExitConcat_prime is called when production concat_prime is exited.
func (s *BasePocLangListener) ExitConcat_prime(ctx *Concat_primeContext) {}

// EnterBasic_re is called when production basic_re is entered.
func (s *BasePocLangListener) EnterBasic_re(ctx *Basic_reContext) {}

// ExitBasic_re is called when production basic_re is exited.
func (s *BasePocLangListener) ExitBasic_re(ctx *Basic_reContext) {}

// EnterKleene_star is called when production kleene_star is entered.
func (s *BasePocLangListener) EnterKleene_star(ctx *Kleene_starContext) {}

// ExitKleene_star is called when production kleene_star is exited.
func (s *BasePocLangListener) ExitKleene_star(ctx *Kleene_starContext) {}

// EnterPlus is called when production plus is entered.
func (s *BasePocLangListener) EnterPlus(ctx *PlusContext) {}

// ExitPlus is called when production plus is exited.
func (s *BasePocLangListener) ExitPlus(ctx *PlusContext) {}

// EnterElementary_re is called when production elementary_re is entered.
func (s *BasePocLangListener) EnterElementary_re(ctx *Elementary_reContext) {}

// ExitElementary_re is called when production elementary_re is exited.
func (s *BasePocLangListener) ExitElementary_re(ctx *Elementary_reContext) {}

// EnterGroup is called when production group is entered.
func (s *BasePocLangListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BasePocLangListener) ExitGroup(ctx *GroupContext) {}

// EnterRange_re is called when production range_re is entered.
func (s *BasePocLangListener) EnterRange_re(ctx *Range_reContext) {}

// ExitRange_re is called when production range_re is exited.
func (s *BasePocLangListener) ExitRange_re(ctx *Range_reContext) {}

// EnterPositive_range is called when production positive_range is entered.
func (s *BasePocLangListener) EnterPositive_range(ctx *Positive_rangeContext) {}

// ExitPositive_range is called when production positive_range is exited.
func (s *BasePocLangListener) ExitPositive_range(ctx *Positive_rangeContext) {}

// EnterNegative_range is called when production negative_range is entered.
func (s *BasePocLangListener) EnterNegative_range(ctx *Negative_rangeContext) {}

// ExitNegative_range is called when production negative_range is exited.
func (s *BasePocLangListener) ExitNegative_range(ctx *Negative_rangeContext) {}

// EnterLax_character is called when production lax_character is entered.
func (s *BasePocLangListener) EnterLax_character(ctx *Lax_characterContext) {}

// ExitLax_character is called when production lax_character is exited.
func (s *BasePocLangListener) ExitLax_character(ctx *Lax_characterContext) {}

// EnterCharacter is called when production character is entered.
func (s *BasePocLangListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BasePocLangListener) ExitCharacter(ctx *CharacterContext) {}

// EnterRange_items is called when production range_items is entered.
func (s *BasePocLangListener) EnterRange_items(ctx *Range_itemsContext) {}

// ExitRange_items is called when production range_items is exited.
func (s *BasePocLangListener) ExitRange_items(ctx *Range_itemsContext) {}

// EnterRange_item is called when production range_item is entered.
func (s *BasePocLangListener) EnterRange_item(ctx *Range_itemContext) {}

// ExitRange_item is called when production range_item is exited.
func (s *BasePocLangListener) ExitRange_item(ctx *Range_itemContext) {}
