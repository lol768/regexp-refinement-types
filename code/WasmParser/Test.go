package main

import (
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"html"
	"strconv"
	"syscall/js"
)

type pocLangListener struct {
	*parser.BasePocLangListener
	names []string
	tokenTypes []string
	Html  string
}

type pocLangErrorListener struct {
	*antlr.DefaultErrorListener

	Bad     bool
	LastMsg string
}

func (el *pocLangErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.LastMsg = "L" + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg
	el.Bad = true
}

func newPocLangListener(names []string, tokenTypes []string) *pocLangListener {
	return &pocLangListener{names: names, tokenTypes: tokenTypes}
}

func (s *pocLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if (len(ctx.GetText()) == 0) {
		return
	}

	lineCls := ""
	if ctx.GetStart().GetLine() == ctx.GetStop().GetLine() {
		lineCls = "l" + strconv.Itoa(ctx.GetStart().GetLine())
	}

	s.Html = s.Html + "<span class='"+s.names[ctx.GetRuleIndex()]+" " + lineCls + "'>"
}

func (s *pocLangListener) VisitTerminal(node antlr.TerminalNode) {

	if node.GetSymbol().GetTokenType() < 0 || len(s.tokenTypes) < (node.GetSymbol().GetTokenType()+1) {
		return
	}

	if node.GetText() != "<EOF>" {
		s.Html = s.Html + "<span class='"+s.tokenTypes[node.GetSymbol().GetTokenType()]+"'>" + html.EscapeString(node.GetText()) + "</span>"
	}
}

func (s *pocLangListener) VisitErrorNode(node antlr.ErrorNode) {
	s.Html = s.Html + "<span class='ERROR'>" + html.EscapeString(node.GetText()) + "</span>"
}

func (s *pocLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if len(ctx.GetText()) == 0 {
		return
	}
	s.Html = s.Html + "</span>"
}

func main() {
	done := make(chan struct{}, 0)
	doc := js.Global().Get("document")
	doc.Call("querySelector", "#loading").Set("innerHTML", "")
	changeEvent := js.NewCallback(func(args []js.Value) {
		highlight(doc.Call("getElementById", "input").Get("value").String(), doc)
	})
	doc.Call("getElementById", "input").Call("addEventListener", "keyup", changeEvent)
	highlight(doc.Call("getElementById", "input").Get("value").String(), doc)
	defer changeEvent.Release()
	<- done
}

func highlight(s string, doc js.Value) {
	is := antlr.NewInputStream(s)
	pocLexer := parser.NewPocLex(is)
	errorListener := &pocLangErrorListener{Bad: false}
	pocLexer.RemoveErrorListeners()
	pocLexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(pocLexer, antlr.TokenDefaultChannel)

	pocParse := parser.NewPocLang(stream)
	tree := pocParse.Program()

	listener := newPocLangListener(pocParse.RuleNames, pocLexer.GetSymbolicNames())
	if errorListener.Bad {
		println("Bad")
		tag := doc.Call("querySelector", "#syntax-error")
		tag.Get("classList").Call("remove", "hide")
		tag.Call("setAttribute", "data-content", "<code>"+html.EscapeString(errorListener.LastMsg)+"</code>")
	} else {
		tag := doc.Call("querySelector", "#syntax-error")
		tag.Get("classList").Call("add", "hide")
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	body := doc.Call("querySelector", "#output")
	body.Set("innerHTML", listener.Html)
}

