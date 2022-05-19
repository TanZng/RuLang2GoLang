// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ru

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRuVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitSentencia(ctx *SentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitSentenciaIf(ctx *SentenciaIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitBloqueCondicional(ctx *BloqueCondicionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitBloqueDeSentencia(ctx *BloqueDeSentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitSentenciaWhile(ctx *SentenciaWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitImprimir(ctx *ImprimirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitMenosUnarioExpr(ctx *MenosUnarioExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitMultiplicacionExpr(ctx *MultiplicacionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitIgualdadExpr(ctx *IgualdadExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitRelacionalExpr(ctx *RelacionalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitAditivaExpr(ctx *AditivaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
