// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ru

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by RuParser.
type RuVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by RuParser#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by RuParser#sentencia.
	VisitSentencia(ctx *SentenciaContext) interface{}

	// Visit a parse tree produced by RuParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by RuParser#sentenciaIf.
	VisitSentenciaIf(ctx *SentenciaIfContext) interface{}

	// Visit a parse tree produced by RuParser#bloqueCondicional.
	VisitBloqueCondicional(ctx *BloqueCondicionalContext) interface{}

	// Visit a parse tree produced by RuParser#bloqueDeSentencia.
	VisitBloqueDeSentencia(ctx *BloqueDeSentenciaContext) interface{}

	// Visit a parse tree produced by RuParser#sentenciaWhile.
	VisitSentenciaWhile(ctx *SentenciaWhileContext) interface{}

	// Visit a parse tree produced by RuParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by RuParser#imprimir.
	VisitImprimir(ctx *ImprimirContext) interface{}

	// Visit a parse tree produced by RuParser#MenosUnarioExpr.
	VisitMenosUnarioExpr(ctx *MenosUnarioExprContext) interface{}

	// Visit a parse tree produced by RuParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by RuParser#multiplicacionExpr.
	VisitMultiplicacionExpr(ctx *MultiplicacionExprContext) interface{}

	// Visit a parse tree produced by RuParser#PowExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by RuParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by RuParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by RuParser#igualdadExpr.
	VisitIgualdadExpr(ctx *IgualdadExprContext) interface{}

	// Visit a parse tree produced by RuParser#relacionalExpr.
	VisitRelacionalExpr(ctx *RelacionalExprContext) interface{}

	// Visit a parse tree produced by RuParser#aditivaExpr.
	VisitAditivaExpr(ctx *AditivaExprContext) interface{}

	// Visit a parse tree produced by RuParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by RuParser#parExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by RuParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by RuParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by RuParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}

	// Visit a parse tree produced by RuParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by RuParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}

}