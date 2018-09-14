package lexer

import (
	"github.com/adwd/monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置 (現在の文字の位置)
	readPosition int  // これから読み込む位置 (現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tt token.TokenType

	switch l.ch {
	case '=':
		tt = token.ASSIGN
	case ';':
		tt = token.SEMICOLON
	case '(':
		tt = token.LPAREN
	case ')':
		tt = token.RPAREN
	case ',':
		tt = token.COMMA
	case '+':
		tt = token.PLUS
	case '{':
		tt = token.LBRACE
	case '}':
		tt = token.RBRACE
	}

	var tok token.Token
	if l.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
	} else {
		tok = newToken(tt, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCIIコードのNULに対応している
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
