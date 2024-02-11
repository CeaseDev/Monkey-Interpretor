package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // tracks current position in input points to current charecter
	readPosition int  //current reading position in input (after current char)
	ch           byte // current character which we are looking at
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // intitializing a new Lexer with input field as the input provided -> l contains the address to the Lexer
	l.readChar()
	return l // return the pointer containing address // lets say xyz
}

// this is a reciever fucniton which enbles us to attach to the specific object -> here Lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition

	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.ASSIGN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = " "
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// focuses only on ascii character not on the entire unicode for simplicity
