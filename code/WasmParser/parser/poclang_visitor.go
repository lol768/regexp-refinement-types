// Code generated from PocLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PocLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PocLang.
type PocLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PocLang#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by PocLang#function_body.
	VisitFunction_body(ctx *Function_bodyContext) interface{}

	// Visit a parse tree produced by PocLang#body_line.
	VisitBody_line(ctx *Body_lineContext) interface{}

	// Visit a parse tree produced by PocLang#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by PocLang#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by PocLang#optional_else.
	VisitOptional_else(ctx *Optional_elseContext) interface{}

	// Visit a parse tree produced by PocLang#var_decl.
	VisitVar_decl(ctx *Var_declContext) interface{}

	// Visit a parse tree produced by PocLang#UnsignedIntType.
	VisitUnsignedIntType(ctx *UnsignedIntTypeContext) interface{}

	// Visit a parse tree produced by PocLang#StringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by PocLang#VoidType.
	VisitVoidType(ctx *VoidTypeContext) interface{}

	// Visit a parse tree produced by PocLang#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by PocLang#type_specifier.
	VisitType_specifier(ctx *Type_specifierContext) interface{}

	// Visit a parse tree produced by PocLang#LessThanConstraint.
	VisitLessThanConstraint(ctx *LessThanConstraintContext) interface{}

	// Visit a parse tree produced by PocLang#GreaterThanConstraint.
	VisitGreaterThanConstraint(ctx *GreaterThanConstraintContext) interface{}

	// Visit a parse tree produced by PocLang#LessThanEqualsConstraint.
	VisitLessThanEqualsConstraint(ctx *LessThanEqualsConstraintContext) interface{}

	// Visit a parse tree produced by PocLang#GreaterThanEqualsConstraint.
	VisitGreaterThanEqualsConstraint(ctx *GreaterThanEqualsConstraintContext) interface{}

	// Visit a parse tree produced by PocLang#string_constraint.
	VisitString_constraint(ctx *String_constraintContext) interface{}

	// Visit a parse tree produced by PocLang#function_sig.
	VisitFunction_sig(ctx *Function_sigContext) interface{}

	// Visit a parse tree produced by PocLang#argument_decl.
	VisitArgument_decl(ctx *Argument_declContext) interface{}

	// Visit a parse tree produced by PocLang#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by PocLang#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by PocLang#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PocLang#var_assignment.
	VisitVar_assignment(ctx *Var_assignmentContext) interface{}

	// Visit a parse tree produced by PocLang#value_ref.
	VisitValue_ref(ctx *Value_refContext) interface{}

	// Visit a parse tree produced by PocLang#identifier_ref.
	VisitIdentifier_ref(ctx *Identifier_refContext) interface{}

	// Visit a parse tree produced by PocLang#function_call.
	VisitFunction_call(ctx *Function_callContext) interface{}

	// Visit a parse tree produced by PocLang#java_call.
	VisitJava_call(ctx *Java_callContext) interface{}

	// Visit a parse tree produced by PocLang#re.
	VisitRe(ctx *ReContext) interface{}

	// Visit a parse tree produced by PocLang#union_prime.
	VisitUnion_prime(ctx *Union_primeContext) interface{}

	// Visit a parse tree produced by PocLang#simple_re.
	VisitSimple_re(ctx *Simple_reContext) interface{}

	// Visit a parse tree produced by PocLang#concat_prime.
	VisitConcat_prime(ctx *Concat_primeContext) interface{}

	// Visit a parse tree produced by PocLang#basic_re.
	VisitBasic_re(ctx *Basic_reContext) interface{}

	// Visit a parse tree produced by PocLang#kleene_star.
	VisitKleene_star(ctx *Kleene_starContext) interface{}

	// Visit a parse tree produced by PocLang#plus.
	VisitPlus(ctx *PlusContext) interface{}

	// Visit a parse tree produced by PocLang#elementary_re.
	VisitElementary_re(ctx *Elementary_reContext) interface{}

	// Visit a parse tree produced by PocLang#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by PocLang#range_re.
	VisitRange_re(ctx *Range_reContext) interface{}

	// Visit a parse tree produced by PocLang#positive_range.
	VisitPositive_range(ctx *Positive_rangeContext) interface{}

	// Visit a parse tree produced by PocLang#negative_range.
	VisitNegative_range(ctx *Negative_rangeContext) interface{}

	// Visit a parse tree produced by PocLang#lax_character.
	VisitLax_character(ctx *Lax_characterContext) interface{}

	// Visit a parse tree produced by PocLang#character.
	VisitCharacter(ctx *CharacterContext) interface{}

	// Visit a parse tree produced by PocLang#range_items.
	VisitRange_items(ctx *Range_itemsContext) interface{}

	// Visit a parse tree produced by PocLang#range_item.
	VisitRange_item(ctx *Range_itemContext) interface{}
}
