package main

import (
	"fmt"
)

func main() {
	var str = "1+3-1"

	var tokens []Token
	var length = len(str)
	var token Token
	var stage int

	for i:=0;i<length;i++{
		s := str[i]

		switch stage {
		case 0:
			stage = initStage(s)
			token = Token{Type: stage}
			token.Append(s)
			continue
		case Int:
			if ScanInt(s){
				token.Append(s)
				continue
			}

		case Id:
			if ScanId(s){
				token.Append(s)
				continue
			}

		case Operate:

		default:
			panic(fmt.Sprintf("unknown token:%v", s))
		}

		tokens = append(tokens, token)

		stage = initStage(s)
		token = Token{Type: stage}
		token.Append(s)
	}

	tokens = append(tokens, token)


	for _,to := range tokens{
		fmt.Println(to.String())
	}
}

type Token struct {
	Type int
	Text string
}

func (t *Token)Append(s uint8)  {
	t.Text += string(s)
}

func (t *Token)String()string  {
	return fmt.Sprintf("%s %s", m[t.Type], t.Text)
}

const (
	Int = iota+1
	Id
	Operate
)

var m = map[int]string{
	Int: "int",
	Id: "identify",
	Operate: "operate",
}

func initStage(s uint8) int {
	switch s {
	case Add,Sub,Mul,Div,QuoR,QuoL:
		return Operate
	default:
		if ScanInt(s){
			return Int
		}

		if ScanId(s){
			return Id
		}

		panic(fmt.Sprintf("unknown s: %s", s))
	}
}

func ScanInt(s uint8) bool {
	return s >= '0' && s <= '9'
}

func ScanId(s uint8) bool {
	return s >= 'a' && s <= 'Z'
}

const (
	Add = '+'
	Sub = '-'
	Mul = '*'
	Div = '/'
	QuoL = '('
	QuoR = ')'
)
