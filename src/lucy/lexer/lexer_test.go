package lexer

import (
	"lucy/token"

	"testing"
)

// func TestNextToken(t *testing.T) {
// 	input := `=+(){},;`

// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.ASSIGN, "="},
// 		{token.PLUS, "+"},
// 		{token.LPAREN, "("},
// 		{token.RPAREN, ")"},
// 		{token.LBRACE, "{"},
// 		{token.RBRACE, "}"},
// 		{token.COMMA, ","},
// 		{token.SEMICOLON, ";"},
// 		{token.EOF, ""},
// 	}

// 	l := New(input)

// 	for i, tt := range tests {
// 		tok := l.NextToken()

// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("tests[%d] - tokentype wrong.  expected %q, got %q",
// 				i, tt.expectedType, tok.Type)
// 		}

// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("tests[%d] - literal wrong.  expected %q, got %q",
// 				i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}
// }

func TestNextToken(t *testing.T) {
	input := `let five = 5;
    let ten = 10;

    let add = fn(x, y) {
        x+y;
    };

    let result = add(five, ten);
    !-/*5;
    5 < 10 > 5;

    if (5 < 10){
        return true;
    } else {
        return false;
    }

    10 == 10;
    10 != 9;

    let dec = 1.234
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.GT, ">"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LT, "<"},
		{token.NUMBER, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.NUMBER, "10"},
		{token.EQ, "=="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.NUMBER, "10"},
		{token.NOT_EQ, "!="},
		{token.NUMBER, "9"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "dec"},
		{token.ASSIGN, "="},
		{token.NUMBER, "1.234"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d - literal wrong. expected=%q, got %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
