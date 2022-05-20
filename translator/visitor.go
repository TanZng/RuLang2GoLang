package translator

import (
	"RuLang2GoLang/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
	"time"
)

type Value struct {
	Value interface{}
}

type Visitor struct {
	parser.RuVisitor
	memory map[string]Value
}

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.ProgramaContext:
		return v.VisitPrograma(val)

	case *parser.BloqueContext:
		return v.VisitBloque(val)

	case *parser.SentenciaContext:
		return v.VisitSentencia(val)
	case *parser.AsignacionContext:
		return v.VisitAsignacion(val)
	case *parser.ImprimirContext:
		return v.VisitImprimir(val)
	case *parser.BloqueDeSentenciaContext:
		return v.VisitBloqueDeSentencia(val)
	case *parser.SentenciaIfContext:
		return v.VisitSentenciaIf(val)
	case *parser.BloqueCondicionalContext:
		return v.VisitBloqueCondicional(val)
	case *parser.SentenciaWhileContext:
		return v.VisitSentenciaWhile(val)
	case *parser.LogContext:
		return v.VisitLog(val)

	case *parser.PowExprContext:
		return v.VisitPowExpr(val)
	case *parser.AditivaExprContext:
		return v.VisitAddExpr(val)
	case *parser.MultiplicacionExprContext:
		return v.VisitMultExpr(val)
	case *parser.IgualdadExprContext:
		return v.VisitIgualdadExpr(val)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val)
	case *parser.MenosUnarioExprContext:
		return v.VisitMenosUnarioExpr(val)
	case *parser.OrExprContext:
		return v.VisitOrExpr(val)
	case *parser.AndExprContext:
		return v.VisitAndExpr(val)
	case *parser.RelacionalExprContext:
		return v.VisitRelExpr(val)
	case *parser.AtomExprContext:
		return v.VisitAtomExpr(val)

	case *parser.ParExprContext:
		return v.VisitParExpr(val)
	case *parser.NumberAtomContext:
		return v.VisitNumAtom(val)
	case *parser.IdAtomContext:
		return v.VisitIdAtom(val)
	case *parser.StringAtomContext:
		return v.VisitStrAtom(val)
	case *parser.BooleanAtomContext:
		return v.VisitBoolAtom(val)
	case *parser.NilAtomContext:
		return v.VisitNilAtom(val)
	default:
		panic("Unknown context " + val.GetText())
	}
}

//   = = = = = = = = = = = = =
//   = = = = Programa  = = = =
//   = = = = = = = = = = = = =

func (v *Visitor) VisitPrograma(ctx *parser.ProgramaContext) Value {
	return v.Visit(ctx.Bloque())
}

//   = = = = = = = = = = = = =
//   = = = = Bloque  = = = =
//   = = = = = = = = = = = = =

func (v *Visitor) VisitBloque(ctx *parser.BloqueContext) Value {
	for i := 0; ctx.Sentencia(i) != nil; i++ {
		v.Visit(ctx.Sentencia(i))
	}
	return Value{Value: true}
}

//   = = = = = = = = = = = = = =
//   = = = = Sentencias  = = = =
//   = = = = = = = = = = = = = =

func (v *Visitor) VisitSentencia(ctx *parser.SentenciaContext) Value {
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())
	}
	if ctx.SentenciaIf() != nil {
		return v.Visit(ctx.SentenciaIf())
	}
	if ctx.Imprimir() != nil {
		return v.Visit(ctx.Imprimir())
	}
	if ctx.SentenciaWhile() != nil {
		return v.Visit(ctx.SentenciaWhile())
	}
	if ctx.SentenciaWhile() != nil {
		return v.Visit(ctx.SentenciaWhile())
	}
	if ctx.Log() != nil {
		return v.Visit(ctx.Log())
	}
	return Value{Value: true}
}

func (v *Visitor) VisitAsignacion(ctx *parser.AsignacionContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.memory[id] = value
	return Value{Value: true}
}

func (v *Visitor) VisitSentenciaIf(ctx *parser.SentenciaIfContext) Value {
	conditionChosen := Value{Value: true}
	conditionChosen = v.Visit(ctx.BloqueCondicional(0))
	if !conditionChosen.Value.(bool) {
		numIfElse := len(ctx.AllIF())
		i := 1
		for i <= numIfElse && !conditionChosen.Value.(bool) {
			conditionChosen = v.Visit(ctx.BloqueCondicional(i))
			i++
		}
		if !conditionChosen.Value.(bool) {
			conditionChosen = v.Visit(ctx.BloqueDeSentencia())
		}
	}

	return conditionChosen
}

func (v *Visitor) VisitSentenciaWhile(ctx *parser.SentenciaWhileContext) Value {
	value, ok := v.Visit(ctx.Expr()).Value.(bool)
	for ok && value {
		v.Visit(ctx.BloqueDeSentencia())
		value, ok = v.Visit(ctx.Expr()).Value.(bool)
	}
	return Value{Value: true}
}

func (v *Visitor) VisitLog(ctx *parser.LogContext) Value {
	val := v.Visit(ctx.Expr())
	timeNow := time.Now().Format("2006/01/02")
	line := ctx.GetStart().GetLine()
	fmt.Printf("LOG: %s, %T: %#v in Line %d\n", timeNow, val.Value, val.Value, line)
	return Value{Value: true}
}

func (v *Visitor) VisitImprimir(ctx *parser.ImprimirContext) Value {
	val := v.Visit(ctx.Expr()).Value
	switch val.(type) {
	case int64:
		fmt.Printf("%d\n", val)
	case float64:
		fmt.Printf("%.2f\n", val)
	case bool:
		fmt.Printf("%b\n", val)
	case string:
		fmt.Printf("%s\n", val)
	}

	return Value{Value: true}
}

func (v *Visitor) VisitBloqueDeSentencia(ctx *parser.BloqueDeSentenciaContext) Value {
	return v.Visit(ctx.Bloque())
}

func (v *Visitor) VisitBloqueCondicional(ctx *parser.BloqueCondicionalContext) Value {
	value, ok := v.Visit(ctx.Expr()).Value.(bool)
	if ok && value {
		return v.Visit(ctx.BloqueDeSentencia())
	}
	return Value{Value: false}
}

//   = = = = = = = = = = = = =
// = = = = Expresiones = = = =
//   = = = = = = = = = = = = =

func (v *Visitor) VisitPowExpr(ctx *parser.PowExprContext) Value {
	l := v.Visit(ctx.Expr(0))
	r := v.Visit(ctx.Expr(1))
	op := "^"

	line := ctx.GetStart().GetLine()
	return realizarOpMath(l, r, op, line)
}

func (v *Visitor) VisitAddExpr(ctx *parser.AditivaExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	line := ctx.GetStart().GetLine()
	return realizarOpMath(l, r, op, line)
}

func (v *Visitor) VisitMultExpr(ctx *parser.MultiplicacionExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	line := ctx.GetStart().GetLine()

	return realizarOpMath(l, r, op, line)
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) Value {
	r := v.Visit(ctx.GetRight())
	line := ctx.GetStart().GetLine()

	return realizarNot(r, line)
}

func (v *Visitor) VisitMenosUnarioExpr(ctx *parser.MenosUnarioExprContext) Value {
	r := v.Visit(ctx.GetRight())
	line := ctx.GetStart().GetLine()

	return realizarNegativo(r, line)
}

func (v *Visitor) VisitIgualdadExpr(ctx *parser.IgualdadExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	line := ctx.GetStart().GetLine()

	return realizarIgualdad(l, r, op, line)
}

func (v *Visitor) VisitRelExpr(ctx *parser.RelacionalExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	line := ctx.GetStart().GetLine()

	return realizarOpMath(l, r, op, line)
}

func (v *Visitor) VisitAtomExpr(ctx *parser.AtomExprContext) Value {
	if ctx.Atomo() != nil {
		return v.Visit(ctx.Atomo())
	}
	return Value{Value: true}
}

func (v *Visitor) VisitOrExpr(ctx *parser.OrExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := "&&"
	line := ctx.GetStart().GetLine()

	return realizarOpBool(l, r, op, line)
}

func (v *Visitor) VisitAndExpr(ctx *parser.AndExprContext) Value {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	line := ctx.GetStart().GetLine()
	op := "||"
	return realizarOpMath(l, r, op, line)
}

// = = = = = = = = = = =
// = = = Atomos = = = =
// = = = = = = = = = = =

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) Value {
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	return Value{Value: true}
}

func (v *Visitor) VisitNumAtom(ctx *parser.NumberAtomContext) Value {
	if ctx.INT() != nil {
		i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
		return Value{Value: i}
	}
	if ctx.FLOAT() != nil {
		i, _ := strconv.ParseFloat(ctx.GetText(), 64)
		return Value{Value: i}
	}
	return Value{Value: false}
}

func (v *Visitor) VisitBoolAtom(ctx *parser.BooleanAtomContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{Value: value}
}

func (v *Visitor) VisitIdAtom(ctx *parser.IdAtomContext) Value {
	id := ctx.GetText()
	value, ok := v.memory[id]
	if ok {
		return value
	} else {
		line := ctx.GetStart().GetLine()
		return TypeError("No such variable '"+id+"'", line)
	}
}

func (v *Visitor) VisitStrAtom(ctx *parser.StringAtomContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{Value: value}
}

func (v *Visitor) VisitNilAtom(ctx *parser.NilAtomContext) Value {
	return Value{Value: ctx.GetText()}
}
