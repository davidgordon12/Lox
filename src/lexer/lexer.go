package lexer

import "lox/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tkn token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			tkn = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tkn = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		tkn = newToken(token.PLUS, lexer.ch)
	case '-':
		tkn = newToken(token.MINUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			literal := string(ch) + string(lexer.ch)
			tkn = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tkn = newToken(token.BANG, lexer.ch)
		}
	case '/':
		tkn = newToken(token.SLASH, lexer.ch)
	case '*':
		tkn = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tkn = newToken(token.LT, lexer.ch)
	case '>':
		tkn = newToken(token.GT, lexer.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, lexer.ch)
	case ',':
		tkn = newToken(token.COMMA, lexer.ch)
	case '{':
		tkn = newToken(token.LBRACE, lexer.ch)
	case '}':
		tkn = newToken(token.RBRACE, lexer.ch)
	case '(':
		tkn = newToken(token.LPAREN, lexer.ch)
	case ')':
		tkn = newToken(token.RPAREN, lexer.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tkn.Literal = lexer.readIdentifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		} else if isDigit(lexer.ch) {
			tkn.Type = token.INT
			tkn.Literal = lexer.readNumber()
			return tkn
		} else {
			tkn = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tkn
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
