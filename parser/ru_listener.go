// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ru

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuListener is a complete listener for a parse tree produced by RuParser.
type RuListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterSentencia is called when entering the sentencia production.
	EnterSentencia(c *SentenciaContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterSentenciaIf is called when entering the sentenciaIf production.
	EnterSentenciaIf(c *SentenciaIfContext)

	// EnterBloqueCondicional is called when entering the bloqueCondicional production.
	EnterBloqueCondicional(c *BloqueCondicionalContext)

	// EnterBloqueDeSentencia is called when entering the bloqueDeSentencia production.
	EnterBloqueDeSentencia(c *BloqueDeSentenciaContext)

	// EnterSentenciaWhile is called when entering the sentenciaWhile production.
	EnterSentenciaWhile(c *SentenciaWhileContext)

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterImprimir is called when entering the imprimir production.
	EnterImprimir(c *ImprimirContext)

	// EnterMenosUnarioExpr is called when entering the MenosUnarioExpr production.
	EnterMenosUnarioExpr(c *MenosUnarioExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterMultiplicacionExpr is called when entering the multiplicacionExpr production.
	EnterMultiplicacionExpr(c *MultiplicacionExprContext)

	// EnterPowExpr is called when entering the PowExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterIgualdadExpr is called when entering the igualdadExpr production.
	EnterIgualdadExpr(c *IgualdadExprContext)

	// EnterRelacionalExpr is called when entering the relacionalExpr production.
	EnterRelacionalExpr(c *RelacionalExprContext)

	// EnterAditivaExpr is called when entering the aditivaExpr production.
	EnterAditivaExpr(c *AditivaExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterParExpr is called when entering the parExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// EnterStringAtom is called when entering the stringAtom production.
	EnterStringAtom(c *StringAtomContext)

	// EnterNilAtom is called when entering the nilAtom production.
	EnterNilAtom(c *NilAtomContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitSentencia is called when exiting the sentencia production.
	ExitSentencia(c *SentenciaContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitSentenciaIf is called when exiting the sentenciaIf production.
	ExitSentenciaIf(c *SentenciaIfContext)

	// ExitBloqueCondicional is called when exiting the bloqueCondicional production.
	ExitBloqueCondicional(c *BloqueCondicionalContext)

	// ExitBloqueDeSentencia is called when exiting the bloqueDeSentencia production.
	ExitBloqueDeSentencia(c *BloqueDeSentenciaContext)

	// ExitSentenciaWhile is called when exiting the sentenciaWhile production.
	ExitSentenciaWhile(c *SentenciaWhileContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitImprimir is called when exiting the imprimir production.
	ExitImprimir(c *ImprimirContext)

	// ExitMenosUnarioExpr is called when exiting the MenosUnarioExpr production.
	ExitMenosUnarioExpr(c *MenosUnarioExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitMultiplicacionExpr is called when exiting the multiplicacionExpr production.
	ExitMultiplicacionExpr(c *MultiplicacionExprContext)

	// ExitPowExpr is called when exiting the PowExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitIgualdadExpr is called when exiting the igualdadExpr production.
	ExitIgualdadExpr(c *IgualdadExprContext)

	// ExitRelacionalExpr is called when exiting the relacionalExpr production.
	ExitRelacionalExpr(c *RelacionalExprContext)

	// ExitAditivaExpr is called when exiting the aditivaExpr production.
	ExitAditivaExpr(c *AditivaExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitParExpr is called when exiting the parExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)

	// ExitStringAtom is called when exiting the stringAtom production.
	ExitStringAtom(c *StringAtomContext)

	// ExitNilAtom is called when exiting the nilAtom production.
	ExitNilAtom(c *NilAtomContext)
}
