package simple_parser

import (
	"errors"
	"io"
	"lexer"
)

type SimpleParser struct {
	lexer lexer.Lexer
}

func NewSimpleParser(lexer lexer.Lexer) *SimpleParser {
	return &SimpleParser{
		lexer: lexer,
	}
}

func (s *SimpleParser) list() (*SyntaxNode, error) {
	//根据读取的第一个字符决定选取哪个生产式
	token, err := s.lexer.Scan()
	if err != nil {
		return nil, err
	}

	current_list_node := NewSyntaxNode() //newly add

	if token.Tag == lexer.LEFT_BRACKET {
		//选择 list -> ( list )
		child_list_node, err := s.list() //newly add

		if err != nil { //newly add
			return nil, err
		}
		if child_list_node != nil {
			current_list_node.AddChild(child_list_node) //newly add , add child noe
		}

		token, err = s.lexer.Scan()
		if token.Tag != lexer.RIGHT_BRACKET {
			err := errors.New("Missinf of right bracket")
			return nil, err
		}
	}

	if token.Tag == lexer.NUM {
		// list -> number
		child_list_node := NewSyntaxNode()
		child_number_node, err := s.number() // newly add
		if child_number_node != nil {
			child_list_node.AddChild(child_number_node)
			current_list_node.AddChild(child_list_node)
		}
		if err != nil {
			if err == io.EOF {
				return current_list_node, err
			}
			return nil, err
		}

	}

	token, err = s.lexer.Scan()
	if err != nil {
		if err == io.EOF {
			return current_list_node, err
		}

		return nil, err
	}

	if token.Tag == lexer.PLUS || token.Tag == lexer.MINUS {
		current_list_node.T = s.lexer.Lexeme
		child_list_node, err := s.list() // list -> list + list , list -> list - list
		if child_list_node != nil {
			current_list_node.AddChild(child_list_node)
		}

		if err != nil {
			if err == io.EOF {
				return current_list_node, err
			}

			return nil, err
		}

	} else {
		s.lexer.ReverseScan()
	}

	return current_list_node, nil
}

func (s *SimpleParser) number() (*SyntaxNode, error) {
	if len(s.lexer.Lexeme) > 1 {
		err := errors.New("Number only allow 0-9")
		return nil, err
	}

	current_node := NewSyntaxNode()
	current_node.T = s.lexer.Lexeme

	return current_node, nil
}

func (s *SimpleParser) Parse() (*SyntaxNode, error) {
	return s.list()
}
