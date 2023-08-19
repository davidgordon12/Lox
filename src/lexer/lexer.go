package lexer

import "lox/token"

type Lexer struct {
	input        string // Lox source code
	position     int    // Current position in our input
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // The current character under examination

	/*
	* The reason for these two “pointers” (position, readPosition)
	* pointing into our input string is the fact that we will need to be able to “peek” further into
	* the input and look after the current character to see what comes up next. readPosition always
	* points to the “next” character in the input. position points to the character in the input that
	* corresponds to the ch byte.
	 */
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.ReadChar()
	return lexer
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) ReadChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tkn token.Token

	switch lexer.ch {
	case '=':
		tkn = NewToken(token.ASSIGN, lexer.ch)
	case ';':
		tkn = NewToken(token.SEMICOLON, lexer.ch)
	case '(':
		tkn = NewToken(token.LPAREN, lexer.ch)
	case ')':
		tkn = NewToken(token.RPAREN, lexer.ch)
	case ',':
		tkn = NewToken(token.COMMA, lexer.ch)
	case '+':
		tkn = NewToken(token.PLUS, lexer.ch)
	case '{':
		tkn = NewToken(token.LBRACE, lexer.ch)
	case '}':
		tkn = NewToken(token.RBRACE, lexer.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	}

	lexer.ReadChar()

	return tkn
}
