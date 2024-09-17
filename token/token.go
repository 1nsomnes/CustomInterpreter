package token

type TokenType string

type Token struct {
  Type  TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"
  
  // operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  // comparators
  LT = "<"
  GT = ">"
  EQ = "=="
  NOT_EQ = "!="

  COMMA = ","
  SEMICOLON = ";"

  RPAREN = ")"
  LPAREN = "("
  RBRACE = "}" 
  LBRACE = "{"
  
  // keyword
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
)

var keywords = map[string]TokenType {
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
