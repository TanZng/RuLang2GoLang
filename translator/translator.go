package translator

import (
	"RuLang2GoLang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"os"
)

func Init(prog string) string {
	rescueStdout := os.Stdout
	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	re, wr, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = wr

	// this gets captured
	input := antlr.NewInputStream(prog)
	lexer := parser.NewRuLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRuParser(tokens)
	p.BuildParseTrees = true
	tree := p.Programa()
	eval := Visitor{memory: make(map[string]Value)}
	eval.Visit(tree)

	w.Close()
	wr.Close()
	out, _ := ioutil.ReadAll(r)
	outErr, _ := ioutil.ReadAll(re)
	os.Stdout = rescueStdout
	os.Stderr = rescueStderr

	log := string(out)
	if string(outErr) != "" {
		log += "\nERRORS:\n" + string(outErr)
	}
	return log
}
