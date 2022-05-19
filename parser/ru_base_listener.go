// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ru

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRuListener is a complete listener for a parse tree produced by RuParser.
type BaseRuListener struct{}

var _ RuListener = &BaseRuListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseRuListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseRuListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseRuListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseRuListener) ExitBloque(ctx *BloqueContext) {}

// EnterSentencia is called when production sentencia is entered.
func (s *BaseRuListener) EnterSentencia(ctx *SentenciaContext) {}

// ExitSentencia is called when production sentencia is exited.
func (s *BaseRuListener) ExitSentencia(ctx *SentenciaContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseRuListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseRuListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterSentenciaIf is called when production sentenciaIf is entered.
func (s *BaseRuListener) EnterSentenciaIf(ctx *SentenciaIfContext) {}

// ExitSentenciaIf is called when production sentenciaIf is exited.
func (s *BaseRuListener) ExitSentenciaIf(ctx *SentenciaIfContext) {}

// EnterBloqueCondicional is called when production bloqueCondicional is entered.
func (s *BaseRuListener) EnterBloqueCondicional(ctx *BloqueCondicionalContext) {}

// ExitBloqueCondicional is called when production bloqueCondicional is exited.
func (s *BaseRuListener) ExitBloqueCondicional(ctx *BloqueCondicionalContext) {}

// EnterBloqueDeSentencia is called when production bloqueDeSentencia is entered.
func (s *BaseRuListener) EnterBloqueDeSentencia(ctx *BloqueDeSentenciaContext) {}

// ExitBloqueDeSentencia is called when production bloqueDeSentencia is exited.
func (s *BaseRuListener) ExitBloqueDeSentencia(ctx *BloqueDeSentenciaContext) {}

// EnterSentenciaWhile is called when production sentenciaWhile is entered.
func (s *BaseRuListener) EnterSentenciaWhile(ctx *SentenciaWhileContext) {}

// ExitSentenciaWhile is called when production sentenciaWhile is exited.
func (s *BaseRuListener) ExitSentenciaWhile(ctx *SentenciaWhileContext) {}

// EnterLog is called when production log is entered.
func (s *BaseRuListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseRuListener) ExitLog(ctx *LogContext) {}

// EnterImprimir is called when production imprimir is entered.
func (s *BaseRuListener) EnterImprimir(ctx *ImprimirContext) {}

// ExitImprimir is called when production imprimir is exited.
func (s *BaseRuListener) ExitImprimir(ctx *ImprimirContext) {}

// EnterMenosUnarioExpr is called when production MenosUnarioExpr is entered.
func (s *BaseRuListener) EnterMenosUnarioExpr(ctx *MenosUnarioExprContext) {}

// ExitMenosUnarioExpr is called when production MenosUnarioExpr is exited.
func (s *BaseRuListener) ExitMenosUnarioExpr(ctx *MenosUnarioExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseRuListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseRuListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterMultiplicacionExpr is called when production multiplicacionExpr is entered.
func (s *BaseRuListener) EnterMultiplicacionExpr(ctx *MultiplicacionExprContext) {}

// ExitMultiplicacionExpr is called when production multiplicacionExpr is exited.
func (s *BaseRuListener) ExitMultiplicacionExpr(ctx *MultiplicacionExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseRuListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseRuListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseRuListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseRuListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterIgualdadExpr is called when production igualdadExpr is entered.
func (s *BaseRuListener) EnterIgualdadExpr(ctx *IgualdadExprContext) {}

// ExitIgualdadExpr is called when production igualdadExpr is exited.
func (s *BaseRuListener) ExitIgualdadExpr(ctx *IgualdadExprContext) {}

// EnterRelacionalExpr is called when production relacionalExpr is entered.
func (s *BaseRuListener) EnterRelacionalExpr(ctx *RelacionalExprContext) {}

// ExitRelacionalExpr is called when production relacionalExpr is exited.
func (s *BaseRuListener) ExitRelacionalExpr(ctx *RelacionalExprContext) {}

// EnterAditivaExpr is called when production aditivaExpr is entered.
func (s *BaseRuListener) EnterAditivaExpr(ctx *AditivaExprContext) {}

// ExitAditivaExpr is called when production aditivaExpr is exited.
func (s *BaseRuListener) ExitAditivaExpr(ctx *AditivaExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseRuListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseRuListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseRuListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseRuListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseRuListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseRuListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseRuListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseRuListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseRuListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseRuListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseRuListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseRuListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseRuListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseRuListener) ExitNilAtom(ctx *NilAtomContext) {}
