package main

import (
	"bufio"
	"fmt"
	"os"
)

type Parser struct {
	pos           int
	input         string
	stickynessMap map[byte]int
}

type Expression interface {
	expression()
	String() string
}

type IntExpression struct {
	Value int
}

type InfixExpression struct {
	left  Expression
	right Expression
	infix byte
}

func main() {
	file, err := os.Open("18/input.txt")
	// file, err := os.Open("18/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum1 := 0
	p1 := Parser{stickynessMap: map[byte]int{
		0:   0,
		'*': 1,
		'+': 1,
	}}
	sum2 := 0
	p2 := Parser{stickynessMap: map[byte]int{
		0:   0,
		'*': 1,
		'+': 2,
	}}

	for scanner.Scan() {
		p1.input, p1.pos = scanner.Text(), 0
		sum1 += eval(p1.parseExpression(0))

		p2.input, p2.pos = scanner.Text(), 0
		sum2 += eval(p2.parseExpression(0))
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func (i IntExpression) expression()   {}
func (o InfixExpression) expression() {}

func (i IntExpression) String() string {
	return fmt.Sprint(i.Value)
}
func (o InfixExpression) String() string {
	return fmt.Sprintf("(%s %c %s)", o.left.String(), o.infix, o.right.String())
}

func (p *Parser) parseExpression(leftStickyness int) Expression {
	var left Expression
	if p.currentToken() == '(' {
		left = p.parseBracedExpression()
	} else {
		left = IntExpression{Value: int(p.currentToken() - '0')}
	}

	if p.peekToken() == 0 || p.peekToken() == ')' {
		return left
	}
	rightStickyness := p.stickynessMap[p.peekToken()]

	for p.peekToken() != 0 && p.peekToken() != ')' && rightStickyness > leftStickyness {
		p.advance()
		left = p.parseInfixExpression(left)
		rightStickyness = p.stickynessMap[p.peekToken()]
	}
	return left
}

func (p *Parser) parseBracedExpression() Expression {
	p.advance()
	e := p.parseExpression(-1)
	p.advance()
	return e
}

func (p *Parser) parseInfixExpression(left Expression) InfixExpression {
	infix := p.currentToken()
	leftStickyness := p.stickynessMap[infix]
	p.advance()
	return InfixExpression{
		left:  left,
		infix: infix,
		right: p.parseExpression(leftStickyness),
	}
}

func (p *Parser) currentToken() byte {
	return p.input[p.pos]
}

func (p *Parser) peekToken() byte {
	if p.pos >= len(p.input)-1 {
		return 0
	}
	pos := p.pos + 1
	for ; p.input[pos] == ' '; pos++ {
	}
	return p.input[pos]
}

func (p *Parser) advance() bool {
	if p.pos >= len(p.input)-1 {
		return false
	}
	p.pos++
	for ; p.input[p.pos] == ' '; p.pos++ {
	}
	return true
}

func eval(e Expression) int {
	switch e := e.(type) {
	case IntExpression:
		return e.Value
	case InfixExpression:
		switch e.infix {
		case '+':
			return eval(e.left) + eval(e.right)
		case '*':
			return eval(e.left) * eval(e.right)
		}
	}
	return -1
}
