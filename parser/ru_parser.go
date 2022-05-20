// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ru

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type RuParser struct {
	*antlr.BaseParser
}

var ruParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func ruParserInit() {
  staticData := &ruParserStaticData
  staticData.literalNames = []string{
    "", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'<='", "'>='", "'+'", 
    "'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "'='", "'('", "')'", 
    "'{'", "'}'", "'true'", "'false'", "'nil'", "'if'", "'else'", "'while'", 
    "'log'", "'imprime'",
  }
  staticData.symbolicNames = []string{
    "", "OR", "AND", "IGUAL", "DIFQ", "MAYORQ", "MENORQ", "MENIG", "MAYIG", 
    "MAS", "MENOS", "MULT", "DIV", "MOD", "POW", "NOT", "PTOCOMA", "ASIGNA", 
    "APAR", "CPAR", "ALLAVE", "CLLAVE", "TRUE", "FALSE", "NIL", "IF", "ELSE", 
    "WHILE", "LOG", "IMPRIMIR", "ID", "INT", "FLOAT", "STRING", "COMENTARIO", 
    "ESPACIO", "OTRO",
  }
  staticData.ruleNames = []string{
    "programa", "bloque", "sentencia", "asignacion", "sentenciaIf", "bloqueCondicional", 
    "bloqueDeSentencia", "sentenciaWhile", "log", "imprimir", "expr", "atomo",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 36, 128, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 29, 8, 1, 10, 1, 12, 1, 
	32, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 41, 8, 2, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 53, 8, 4, 
	10, 4, 12, 4, 56, 9, 4, 1, 4, 1, 4, 3, 4, 60, 8, 4, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 
	8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 3, 10, 89, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 
	1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 5, 10, 112, 8, 10, 10, 10, 12, 10, 115, 9, 10, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 126, 
	8, 11, 1, 11, 0, 1, 20, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	0, 6, 1, 0, 11, 13, 1, 0, 9, 10, 1, 0, 5, 8, 1, 0, 3, 4, 1, 0, 31, 32, 
	1, 0, 22, 23, 137, 0, 24, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 40, 1, 0, 0, 
	0, 6, 42, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 66, 1, 
	0, 0, 0, 14, 70, 1, 0, 0, 0, 16, 74, 1, 0, 0, 0, 18, 78, 1, 0, 0, 0, 20, 
	88, 1, 0, 0, 0, 22, 125, 1, 0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 0, 
	0, 1, 26, 1, 1, 0, 0, 0, 27, 29, 3, 4, 2, 0, 28, 27, 1, 0, 0, 0, 29, 32, 
	1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 
	32, 30, 1, 0, 0, 0, 33, 41, 3, 6, 3, 0, 34, 41, 3, 8, 4, 0, 35, 41, 3, 
	14, 7, 0, 36, 41, 3, 16, 8, 0, 37, 41, 3, 18, 9, 0, 38, 39, 5, 36, 0, 0, 
	39, 41, 6, 2, -1, 0, 40, 33, 1, 0, 0, 0, 40, 34, 1, 0, 0, 0, 40, 35, 1, 
	0, 0, 0, 40, 36, 1, 0, 0, 0, 40, 37, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 
	5, 1, 0, 0, 0, 42, 43, 5, 30, 0, 0, 43, 44, 5, 17, 0, 0, 44, 45, 3, 20, 
	10, 0, 45, 46, 5, 16, 0, 0, 46, 7, 1, 0, 0, 0, 47, 48, 5, 25, 0, 0, 48, 
	54, 3, 10, 5, 0, 49, 50, 5, 26, 0, 0, 50, 51, 5, 25, 0, 0, 51, 53, 3, 10, 
	5, 0, 52, 49, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 
	1, 0, 0, 0, 55, 59, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 58, 5, 26, 0, 0, 
	58, 60, 3, 12, 6, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 9, 1, 
	0, 0, 0, 61, 62, 5, 18, 0, 0, 62, 63, 3, 20, 10, 0, 63, 64, 5, 19, 0, 0, 
	64, 65, 3, 12, 6, 0, 65, 11, 1, 0, 0, 0, 66, 67, 5, 20, 0, 0, 67, 68, 3, 
	2, 1, 0, 68, 69, 5, 21, 0, 0, 69, 13, 1, 0, 0, 0, 70, 71, 5, 27, 0, 0, 
	71, 72, 3, 20, 10, 0, 72, 73, 3, 12, 6, 0, 73, 15, 1, 0, 0, 0, 74, 75, 
	5, 28, 0, 0, 75, 76, 3, 20, 10, 0, 76, 77, 5, 16, 0, 0, 77, 17, 1, 0, 0, 
	0, 78, 79, 5, 29, 0, 0, 79, 80, 3, 20, 10, 0, 80, 81, 5, 16, 0, 0, 81, 
	19, 1, 0, 0, 0, 82, 83, 6, 10, -1, 0, 83, 84, 5, 10, 0, 0, 84, 89, 3, 20, 
	10, 9, 85, 86, 5, 15, 0, 0, 86, 89, 3, 20, 10, 8, 87, 89, 3, 22, 11, 0, 
	88, 82, 1, 0, 0, 0, 88, 85, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 113, 1, 
	0, 0, 0, 90, 91, 10, 10, 0, 0, 91, 92, 5, 14, 0, 0, 92, 112, 3, 20, 10, 
	10, 93, 94, 10, 7, 0, 0, 94, 95, 7, 0, 0, 0, 95, 112, 3, 20, 10, 8, 96, 
	97, 10, 6, 0, 0, 97, 98, 7, 1, 0, 0, 98, 112, 3, 20, 10, 7, 99, 100, 10, 
	5, 0, 0, 100, 101, 7, 2, 0, 0, 101, 112, 3, 20, 10, 6, 102, 103, 10, 4, 
	0, 0, 103, 104, 7, 3, 0, 0, 104, 112, 3, 20, 10, 5, 105, 106, 10, 3, 0, 
	0, 106, 107, 5, 2, 0, 0, 107, 112, 3, 20, 10, 4, 108, 109, 10, 2, 0, 0, 
	109, 110, 5, 1, 0, 0, 110, 112, 3, 20, 10, 3, 111, 90, 1, 0, 0, 0, 111, 
	93, 1, 0, 0, 0, 111, 96, 1, 0, 0, 0, 111, 99, 1, 0, 0, 0, 111, 102, 1, 
	0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 108, 1, 0, 0, 0, 112, 115, 1, 0, 0, 
	0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 21, 1, 0, 0, 0, 115, 
	113, 1, 0, 0, 0, 116, 117, 5, 18, 0, 0, 117, 118, 3, 20, 10, 0, 118, 119, 
	5, 19, 0, 0, 119, 126, 1, 0, 0, 0, 120, 126, 7, 4, 0, 0, 121, 126, 7, 5, 
	0, 0, 122, 126, 5, 30, 0, 0, 123, 126, 5, 33, 0, 0, 124, 126, 5, 24, 0, 
	0, 125, 116, 1, 0, 0, 0, 125, 120, 1, 0, 0, 0, 125, 121, 1, 0, 0, 0, 125, 
	122, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126, 23, 1, 
	0, 0, 0, 8, 30, 40, 54, 59, 88, 111, 113, 125,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// RuParserInit initializes any static state used to implement RuParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRuParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuParserInit() {
  staticData := &ruParserStaticData
  staticData.once.Do(ruParserInit)
}

// NewRuParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRuParser(input antlr.TokenStream) *RuParser {
	RuParserInit()
	this := new(RuParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &ruParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Ru.g4"

	return this
}


// RuParser tokens.
const (
	RuParserEOF = antlr.TokenEOF
	RuParserOR = 1
	RuParserAND = 2
	RuParserIGUAL = 3
	RuParserDIFQ = 4
	RuParserMAYORQ = 5
	RuParserMENORQ = 6
	RuParserMENIG = 7
	RuParserMAYIG = 8
	RuParserMAS = 9
	RuParserMENOS = 10
	RuParserMULT = 11
	RuParserDIV = 12
	RuParserMOD = 13
	RuParserPOW = 14
	RuParserNOT = 15
	RuParserPTOCOMA = 16
	RuParserASIGNA = 17
	RuParserAPAR = 18
	RuParserCPAR = 19
	RuParserALLAVE = 20
	RuParserCLLAVE = 21
	RuParserTRUE = 22
	RuParserFALSE = 23
	RuParserNIL = 24
	RuParserIF = 25
	RuParserELSE = 26
	RuParserWHILE = 27
	RuParserLOG = 28
	RuParserIMPRIMIR = 29
	RuParserID = 30
	RuParserINT = 31
	RuParserFLOAT = 32
	RuParserSTRING = 33
	RuParserCOMENTARIO = 34
	RuParserESPACIO = 35
	RuParserOTRO = 36
)

// RuParser rules.
const (
	RuParserRULE_programa = 0
	RuParserRULE_bloque = 1
	RuParserRULE_sentencia = 2
	RuParserRULE_asignacion = 3
	RuParserRULE_sentenciaIf = 4
	RuParserRULE_bloqueCondicional = 5
	RuParserRULE_bloqueDeSentencia = 6
	RuParserRULE_sentenciaWhile = 7
	RuParserRULE_log = 8
	RuParserRULE_imprimir = 9
	RuParserRULE_expr = 10
	RuParserRULE_atomo = 11
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_programa
	return p
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) Bloque() IBloqueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuParserEOF, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Programa() (localctx IProgramaContext) {
	this := p
	_ = this

	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuParserRULE_programa)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Bloque()
	}
	{
		p.SetState(25)
		p.Match(RuParserEOF)
	}



	return localctx
}


// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuParserRULE_bloque)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 25)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 25))) & ((1 << (RuParserIF - 25)) | (1 << (RuParserWHILE - 25)) | (1 << (RuParserLOG - 25)) | (1 << (RuParserIMPRIMIR - 25)) | (1 << (RuParserID - 25)) | (1 << (RuParserOTRO - 25)))) != 0) {
		{
			p.SetState(27)
			p.Sentencia()
		}


		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_OTRO returns the _OTRO token.
	Get_OTRO() antlr.Token 


	// Set_OTRO sets the _OTRO token.
	Set_OTRO(antlr.Token) 


	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_OTRO antlr.Token
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_sentencia
	return p
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) Get_OTRO() antlr.Token { return s._OTRO }


func (s *SentenciaContext) Set_OTRO(v antlr.Token) { s._OTRO = v }


func (s *SentenciaContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *SentenciaContext) SentenciaIf() ISentenciaIfContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaIfContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaIfContext)
}

func (s *SentenciaContext) SentenciaWhile() ISentenciaWhileContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaWhileContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaWhileContext)
}

func (s *SentenciaContext) Log() ILogContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogContext)
}

func (s *SentenciaContext) Imprimir() IImprimirContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImprimirContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImprimirContext)
}

func (s *SentenciaContext) OTRO() antlr.TerminalNode {
	return s.GetToken(RuParserOTRO, 0)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterSentencia(s)
	}
}

func (s *SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitSentencia(s)
	}
}

func (s *SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitSentencia(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Sentencia() (localctx ISentenciaContext) {
	this := p
	_ = this

	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuParserRULE_sentencia)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Asignacion()
		}


	case RuParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.SentenciaIf()
		}


	case RuParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.SentenciaWhile()
		}


	case RuParserLOG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Log()
		}


	case RuParserIMPRIMIR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(37)
			p.Imprimir()
		}


	case RuParserOTRO:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(38)

			var _m = p.Match(RuParserOTRO)

			localctx.(*SentenciaContext)._OTRO = _m
		}
		fmt.Println("caracter desconocido: ", (func() string { if localctx.(*SentenciaContext).Get_OTRO() == nil { return "" } else { return localctx.(*SentenciaContext).Get_OTRO().GetText() }}()));



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(RuParserID, 0)
}

func (s *AsignacionContext) ASIGNA() antlr.TerminalNode {
	return s.GetToken(RuParserASIGNA, 0)
}

func (s *AsignacionContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionContext) PTOCOMA() antlr.TerminalNode {
	return s.GetToken(RuParserPTOCOMA, 0)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuParserRULE_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(RuParserID)
	}
	{
		p.SetState(43)
		p.Match(RuParserASIGNA)
	}
	{
		p.SetState(44)
		p.expr(0)
	}
	{
		p.SetState(45)
		p.Match(RuParserPTOCOMA)
	}



	return localctx
}


// ISentenciaIfContext is an interface to support dynamic dispatch.
type ISentenciaIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenciaIfContext differentiates from other interfaces.
	IsSentenciaIfContext()
}

type SentenciaIfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaIfContext() *SentenciaIfContext {
	var p = new(SentenciaIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_sentenciaIf
	return p
}

func (*SentenciaIfContext) IsSentenciaIfContext() {}

func NewSentenciaIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaIfContext {
	var p = new(SentenciaIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_sentenciaIf

	return p
}

func (s *SentenciaIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaIfContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(RuParserIF)
}

func (s *SentenciaIfContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(RuParserIF, i)
}

func (s *SentenciaIfContext) AllBloqueCondicional() []IBloqueCondicionalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueCondicionalContext); ok {
			len++
		}
	}

	tst := make([]IBloqueCondicionalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueCondicionalContext); ok {
			tst[i] = t.(IBloqueCondicionalContext)
			i++
		}
	}

	return tst
}

func (s *SentenciaIfContext) BloqueCondicional(i int) IBloqueCondicionalContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueCondicionalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueCondicionalContext)
}

func (s *SentenciaIfContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(RuParserELSE)
}

func (s *SentenciaIfContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(RuParserELSE, i)
}

func (s *SentenciaIfContext) BloqueDeSentencia() IBloqueDeSentenciaContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueDeSentenciaContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueDeSentenciaContext)
}

func (s *SentenciaIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SentenciaIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterSentenciaIf(s)
	}
}

func (s *SentenciaIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitSentenciaIf(s)
	}
}

func (s *SentenciaIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitSentenciaIf(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) SentenciaIf() (localctx ISentenciaIfContext) {
	this := p
	_ = this

	localctx = NewSentenciaIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RuParserRULE_sentenciaIf)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(RuParserIF)
	}
	{
		p.SetState(48)
		p.BloqueCondicional()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				p.Match(RuParserELSE)
			}
			{
				p.SetState(50)
				p.Match(RuParserIF)
			}
			{
				p.SetState(51)
				p.BloqueCondicional()
			}


		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == RuParserELSE {
		{
			p.SetState(57)
			p.Match(RuParserELSE)
		}
		{
			p.SetState(58)
			p.BloqueDeSentencia()
		}

	}



	return localctx
}


// IBloqueCondicionalContext is an interface to support dynamic dispatch.
type IBloqueCondicionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueCondicionalContext differentiates from other interfaces.
	IsBloqueCondicionalContext()
}

type BloqueCondicionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueCondicionalContext() *BloqueCondicionalContext {
	var p = new(BloqueCondicionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_bloqueCondicional
	return p
}

func (*BloqueCondicionalContext) IsBloqueCondicionalContext() {}

func NewBloqueCondicionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueCondicionalContext {
	var p = new(BloqueCondicionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_bloqueCondicional

	return p
}

func (s *BloqueCondicionalContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueCondicionalContext) APAR() antlr.TerminalNode {
	return s.GetToken(RuParserAPAR, 0)
}

func (s *BloqueCondicionalContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BloqueCondicionalContext) CPAR() antlr.TerminalNode {
	return s.GetToken(RuParserCPAR, 0)
}

func (s *BloqueCondicionalContext) BloqueDeSentencia() IBloqueDeSentenciaContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueDeSentenciaContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueDeSentenciaContext)
}

func (s *BloqueCondicionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueCondicionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BloqueCondicionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterBloqueCondicional(s)
	}
}

func (s *BloqueCondicionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitBloqueCondicional(s)
	}
}

func (s *BloqueCondicionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitBloqueCondicional(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) BloqueCondicional() (localctx IBloqueCondicionalContext) {
	this := p
	_ = this

	localctx = NewBloqueCondicionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuParserRULE_bloqueCondicional)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(RuParserAPAR)
	}
	{
		p.SetState(62)
		p.expr(0)
	}
	{
		p.SetState(63)
		p.Match(RuParserCPAR)
	}
	{
		p.SetState(64)
		p.BloqueDeSentencia()
	}



	return localctx
}


// IBloqueDeSentenciaContext is an interface to support dynamic dispatch.
type IBloqueDeSentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueDeSentenciaContext differentiates from other interfaces.
	IsBloqueDeSentenciaContext()
}

type BloqueDeSentenciaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueDeSentenciaContext() *BloqueDeSentenciaContext {
	var p = new(BloqueDeSentenciaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_bloqueDeSentencia
	return p
}

func (*BloqueDeSentenciaContext) IsBloqueDeSentenciaContext() {}

func NewBloqueDeSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueDeSentenciaContext {
	var p = new(BloqueDeSentenciaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_bloqueDeSentencia

	return p
}

func (s *BloqueDeSentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueDeSentenciaContext) ALLAVE() antlr.TerminalNode {
	return s.GetToken(RuParserALLAVE, 0)
}

func (s *BloqueDeSentenciaContext) Bloque() IBloqueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *BloqueDeSentenciaContext) CLLAVE() antlr.TerminalNode {
	return s.GetToken(RuParserCLLAVE, 0)
}

func (s *BloqueDeSentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueDeSentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BloqueDeSentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterBloqueDeSentencia(s)
	}
}

func (s *BloqueDeSentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitBloqueDeSentencia(s)
	}
}

func (s *BloqueDeSentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitBloqueDeSentencia(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) BloqueDeSentencia() (localctx IBloqueDeSentenciaContext) {
	this := p
	_ = this

	localctx = NewBloqueDeSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuParserRULE_bloqueDeSentencia)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(RuParserALLAVE)
	}
	{
		p.SetState(67)
		p.Bloque()
	}
	{
		p.SetState(68)
		p.Match(RuParserCLLAVE)
	}



	return localctx
}


// ISentenciaWhileContext is an interface to support dynamic dispatch.
type ISentenciaWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenciaWhileContext differentiates from other interfaces.
	IsSentenciaWhileContext()
}

type SentenciaWhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaWhileContext() *SentenciaWhileContext {
	var p = new(SentenciaWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_sentenciaWhile
	return p
}

func (*SentenciaWhileContext) IsSentenciaWhileContext() {}

func NewSentenciaWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaWhileContext {
	var p = new(SentenciaWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_sentenciaWhile

	return p
}

func (s *SentenciaWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RuParserWHILE, 0)
}

func (s *SentenciaWhileContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SentenciaWhileContext) BloqueDeSentencia() IBloqueDeSentenciaContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueDeSentenciaContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueDeSentenciaContext)
}

func (s *SentenciaWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SentenciaWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterSentenciaWhile(s)
	}
}

func (s *SentenciaWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitSentenciaWhile(s)
	}
}

func (s *SentenciaWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitSentenciaWhile(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) SentenciaWhile() (localctx ISentenciaWhileContext) {
	this := p
	_ = this

	localctx = NewSentenciaWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuParserRULE_sentenciaWhile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(RuParserWHILE)
	}
	{
		p.SetState(71)
		p.expr(0)
	}
	{
		p.SetState(72)
		p.BloqueDeSentencia()
	}



	return localctx
}


// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) LOG() antlr.TerminalNode {
	return s.GetToken(RuParserLOG, 0)
}

func (s *LogContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogContext) PTOCOMA() antlr.TerminalNode {
	return s.GetToken(RuParserPTOCOMA, 0)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitLog(s)
	}
}

func (s *LogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitLog(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Log() (localctx ILogContext) {
	this := p
	_ = this

	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuParserRULE_log)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(RuParserLOG)
	}
	{
		p.SetState(75)
		p.expr(0)
	}
	{
		p.SetState(76)
		p.Match(RuParserPTOCOMA)
	}



	return localctx
}


// IImprimirContext is an interface to support dynamic dispatch.
type IImprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImprimirContext differentiates from other interfaces.
	IsImprimirContext()
}

type ImprimirContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImprimirContext() *ImprimirContext {
	var p = new(ImprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_imprimir
	return p
}

func (*ImprimirContext) IsImprimirContext() {}

func NewImprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimirContext {
	var p = new(ImprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_imprimir

	return p
}

func (s *ImprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimirContext) IMPRIMIR() antlr.TerminalNode {
	return s.GetToken(RuParserIMPRIMIR, 0)
}

func (s *ImprimirContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ImprimirContext) PTOCOMA() antlr.TerminalNode {
	return s.GetToken(RuParserPTOCOMA, 0)
}

func (s *ImprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ImprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterImprimir(s)
	}
}

func (s *ImprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitImprimir(s)
	}
}

func (s *ImprimirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitImprimir(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuParser) Imprimir() (localctx IImprimirContext) {
	this := p
	_ = this

	localctx = NewImprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuParserRULE_imprimir)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(RuParserIMPRIMIR)
	}
	{
		p.SetState(79)
		p.expr(0)
	}
	{
		p.SetState(80)
		p.Match(RuParserPTOCOMA)
	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type MenosUnarioExprContext struct {
	*ExprContext
	right IExprContext 
}

func NewMenosUnarioExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MenosUnarioExprContext {
	var p = new(MenosUnarioExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *MenosUnarioExprContext) GetRight() IExprContext { return s.right }


func (s *MenosUnarioExprContext) SetRight(v IExprContext) { s.right = v }

func (s *MenosUnarioExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MenosUnarioExprContext) MENOS() antlr.TerminalNode {
	return s.GetToken(RuParserMENOS, 0)
}

func (s *MenosUnarioExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *MenosUnarioExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterMenosUnarioExpr(s)
	}
}

func (s *MenosUnarioExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitMenosUnarioExpr(s)
	}
}

func (s *MenosUnarioExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitMenosUnarioExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type NotExprContext struct {
	*ExprContext
	right IExprContext 
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *NotExprContext) GetRight() IExprContext { return s.right }


func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(RuParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MultiplicacionExprContext struct {
	*ExprContext
	left IExprContext 
	op antlr.Token
	right IExprContext 
}

func NewMultiplicacionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicacionExprContext {
	var p = new(MultiplicacionExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *MultiplicacionExprContext) GetOp() antlr.Token { return s.op }


func (s *MultiplicacionExprContext) SetOp(v antlr.Token) { s.op = v }


func (s *MultiplicacionExprContext) GetLeft() IExprContext { return s.left }

func (s *MultiplicacionExprContext) GetRight() IExprContext { return s.right }


func (s *MultiplicacionExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultiplicacionExprContext) SetRight(v IExprContext) { s.right = v }

func (s *MultiplicacionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicacionExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicacionExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplicacionExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(RuParserMULT, 0)
}

func (s *MultiplicacionExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuParserDIV, 0)
}

func (s *MultiplicacionExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(RuParserMOD, 0)
}


func (s *MultiplicacionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterMultiplicacionExpr(s)
	}
}

func (s *MultiplicacionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitMultiplicacionExpr(s)
	}
}

func (s *MultiplicacionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitMultiplicacionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type PowExprContext struct {
	*ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PowExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(RuParserPOW, 0)
}


func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AtomExprContext struct {
	*ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atomo() IAtomoContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomoContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomoContext)
}


func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type OrExprContext struct {
	*ExprContext
	left IExprContext 
	right IExprContext 
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *OrExprContext) GetLeft() IExprContext { return s.left }

func (s *OrExprContext) GetRight() IExprContext { return s.right }


func (s *OrExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *OrExprContext) SetRight(v IExprContext) { s.right = v }

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(RuParserOR, 0)
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type IgualdadExprContext struct {
	*ExprContext
	left IExprContext 
	op antlr.Token
	right IExprContext 
}

func NewIgualdadExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IgualdadExprContext {
	var p = new(IgualdadExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *IgualdadExprContext) GetOp() antlr.Token { return s.op }


func (s *IgualdadExprContext) SetOp(v antlr.Token) { s.op = v }


func (s *IgualdadExprContext) GetLeft() IExprContext { return s.left }

func (s *IgualdadExprContext) GetRight() IExprContext { return s.right }


func (s *IgualdadExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *IgualdadExprContext) SetRight(v IExprContext) { s.right = v }

func (s *IgualdadExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgualdadExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IgualdadExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IgualdadExprContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(RuParserIGUAL, 0)
}

func (s *IgualdadExprContext) DIFQ() antlr.TerminalNode {
	return s.GetToken(RuParserDIFQ, 0)
}


func (s *IgualdadExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterIgualdadExpr(s)
	}
}

func (s *IgualdadExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitIgualdadExpr(s)
	}
}

func (s *IgualdadExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitIgualdadExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type RelacionalExprContext struct {
	*ExprContext
	left IExprContext 
	op antlr.Token
	right IExprContext 
}

func NewRelacionalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelacionalExprContext {
	var p = new(RelacionalExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *RelacionalExprContext) GetOp() antlr.Token { return s.op }


func (s *RelacionalExprContext) SetOp(v antlr.Token) { s.op = v }


func (s *RelacionalExprContext) GetLeft() IExprContext { return s.left }

func (s *RelacionalExprContext) GetRight() IExprContext { return s.right }


func (s *RelacionalExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *RelacionalExprContext) SetRight(v IExprContext) { s.right = v }

func (s *RelacionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelacionalExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelacionalExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelacionalExprContext) MAYIG() antlr.TerminalNode {
	return s.GetToken(RuParserMAYIG, 0)
}

func (s *RelacionalExprContext) MENIG() antlr.TerminalNode {
	return s.GetToken(RuParserMENIG, 0)
}

func (s *RelacionalExprContext) MENORQ() antlr.TerminalNode {
	return s.GetToken(RuParserMENORQ, 0)
}

func (s *RelacionalExprContext) MAYORQ() antlr.TerminalNode {
	return s.GetToken(RuParserMAYORQ, 0)
}


func (s *RelacionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterRelacionalExpr(s)
	}
}

func (s *RelacionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitRelacionalExpr(s)
	}
}

func (s *RelacionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitRelacionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AditivaExprContext struct {
	*ExprContext
	left IExprContext 
	op antlr.Token
	right IExprContext 
}

func NewAditivaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AditivaExprContext {
	var p = new(AditivaExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *AditivaExprContext) GetOp() antlr.Token { return s.op }


func (s *AditivaExprContext) SetOp(v antlr.Token) { s.op = v }


func (s *AditivaExprContext) GetLeft() IExprContext { return s.left }

func (s *AditivaExprContext) GetRight() IExprContext { return s.right }


func (s *AditivaExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *AditivaExprContext) SetRight(v IExprContext) { s.right = v }

func (s *AditivaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AditivaExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AditivaExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AditivaExprContext) MAS() antlr.TerminalNode {
	return s.GetToken(RuParserMAS, 0)
}

func (s *AditivaExprContext) MENOS() antlr.TerminalNode {
	return s.GetToken(RuParserMENOS, 0)
}


func (s *AditivaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterAditivaExpr(s)
	}
}

func (s *AditivaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitAditivaExpr(s)
	}
}

func (s *AditivaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitAditivaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AndExprContext struct {
	*ExprContext
	left IExprContext 
	right IExprContext 
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *AndExprContext) GetLeft() IExprContext { return s.left }

func (s *AndExprContext) GetRight() IExprContext { return s.right }


func (s *AndExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *AndExprContext) SetRight(v IExprContext) { s.right = v }

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(RuParserAND, 0)
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *RuParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RuParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, RuParserRULE_expr, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuParserMENOS:
		localctx = NewMenosUnarioExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(83)
			p.Match(RuParserMENOS)
		}
		{
			p.SetState(84)

			var _x = p.expr(9)

			localctx.(*MenosUnarioExprContext).right = _x
		}


	case RuParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(85)
			p.Match(RuParserNOT)
		}
		{
			p.SetState(86)

			var _x = p.expr(8)

			localctx.(*NotExprContext).right = _x
		}


	case RuParserAPAR, RuParserTRUE, RuParserFALSE, RuParserNIL, RuParserID, RuParserINT, RuParserFLOAT, RuParserSTRING:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)
			p.Atomo()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(91)
					p.Match(RuParserPOW)
				}
				{
					p.SetState(92)
					p.expr(10)
				}


			case 2:
				localctx = NewMultiplicacionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplicacionExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(94)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicacionExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RuParserMULT) | (1 << RuParserDIV) | (1 << RuParserMOD))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicacionExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)

					var _x = p.expr(8)

					localctx.(*MultiplicacionExprContext).right = _x
				}


			case 3:
				localctx = NewAditivaExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AditivaExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AditivaExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuParserMAS || _la == RuParserMENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AditivaExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)

					var _x = p.expr(7)

					localctx.(*AditivaExprContext).right = _x
				}


			case 4:
				localctx = NewRelacionalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelacionalExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(100)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelacionalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RuParserMAYORQ) | (1 << RuParserMENORQ) | (1 << RuParserMENIG) | (1 << RuParserMAYIG))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelacionalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(101)

					var _x = p.expr(6)

					localctx.(*RelacionalExprContext).right = _x
				}


			case 5:
				localctx = NewIgualdadExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IgualdadExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(103)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*IgualdadExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuParserIGUAL || _la == RuParserDIFQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*IgualdadExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)

					var _x = p.expr(5)

					localctx.(*IgualdadExprContext).right = _x
				}


			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(106)
					p.Match(RuParserAND)
				}
				{
					p.SetState(107)

					var _x = p.expr(4)

					localctx.(*AndExprContext).right = _x
				}


			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrExprContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(109)
					p.Match(RuParserOR)
				}
				{
					p.SetState(110)

					var _x = p.expr(3)

					localctx.(*OrExprContext).right = _x
				}

			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}



	return localctx
}


// IAtomoContext is an interface to support dynamic dispatch.
type IAtomoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomoContext differentiates from other interfaces.
	IsAtomoContext()
}

type AtomoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomoContext() *AtomoContext {
	var p = new(AtomoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuParserRULE_atomo
	return p
}

func (*AtomoContext) IsAtomoContext() {}

func NewAtomoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomoContext {
	var p = new(AtomoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuParserRULE_atomo

	return p
}

func (s *AtomoContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomoContext) CopyFrom(ctx *AtomoContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ParExprContext struct {
	*AtomoContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) APAR() antlr.TerminalNode {
	return s.GetToken(RuParserAPAR, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(RuParserCPAR, 0)
}


func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type BooleanAtomContext struct {
	*AtomoContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuParserFALSE, 0)
}


func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type IdAtomContext struct {
	*AtomoContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(RuParserID, 0)
}


func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

func (s *IdAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitIdAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type StringAtomContext struct {
	*AtomoContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(RuParserSTRING, 0)
}


func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type NilAtomContext struct {
	*AtomoContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(RuParserNIL, 0)
}


func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitNilAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type NumberAtomContext struct {
	*AtomoContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	p.AtomoContext = NewEmptyAtomoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomoContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(RuParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RuParserFLOAT, 0)
}


func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (s *NumberAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuVisitor:
		return t.VisitNumberAtom(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *RuParser) Atomo() (localctx IAtomoContext) {
	this := p
	_ = this

	localctx = NewAtomoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RuParserRULE_atomo)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuParserAPAR:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(RuParserAPAR)
		}
		{
			p.SetState(117)
			p.expr(0)
		}
		{
			p.SetState(118)
			p.Match(RuParserCPAR)
		}


	case RuParserINT, RuParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RuParserINT || _la == RuParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case RuParserTRUE, RuParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RuParserTRUE || _la == RuParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case RuParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Match(RuParserID)
		}


	case RuParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Match(RuParserSTRING)
		}


	case RuParserNIL:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Match(RuParserNIL)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


func (p *RuParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

