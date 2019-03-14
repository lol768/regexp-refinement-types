// Code generated from PocLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PocLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePocLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePocLangVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitFunction_body(ctx *Function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitBody_line(ctx *Body_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitOptional_else(ctx *Optional_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitVar_decl(ctx *Var_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitUnsignedIntType(ctx *UnsignedIntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitVoidType(ctx *VoidTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitType_specifier(ctx *Type_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitLessThanConstraint(ctx *LessThanConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitGreaterThanConstraint(ctx *GreaterThanConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitLessThanEqualsConstraint(ctx *LessThanEqualsConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitGreaterThanEqualsConstraint(ctx *GreaterThanEqualsConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitString_constraint(ctx *String_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitFunction_sig(ctx *Function_sigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitArgument_decl(ctx *Argument_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitVar_assignment(ctx *Var_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitValue_ref(ctx *Value_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitIdentifier_ref(ctx *Identifier_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitJava_call(ctx *Java_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitRe(ctx *ReContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitUnion_prime(ctx *Union_primeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitSimple_re(ctx *Simple_reContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitConcat_prime(ctx *Concat_primeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitBasic_re(ctx *Basic_reContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitKleene_star(ctx *Kleene_starContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitPlus(ctx *PlusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitElementary_re(ctx *Elementary_reContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitRange_re(ctx *Range_reContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitPositive_range(ctx *Positive_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitNegative_range(ctx *Negative_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitLax_character(ctx *Lax_characterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitCharacter(ctx *CharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitRange_items(ctx *Range_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePocLangVisitor) VisitRange_item(ctx *Range_itemContext) interface{} {
	return v.VisitChildren(ctx)
}
