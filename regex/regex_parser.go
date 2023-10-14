package regex

import ( 
	"errors"
	"github.com/dgoncas/lexic_analizer_generator/utils"
)


type RegexParser struct{
	alphabet_chars utils.Set
	single_operand_operators utils.Set
}

func NewRegexParser() RegexParser {
	result := RegexParser{
		utils.NewSet( " ", "\n", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_", "-"),
		utils.NewSet( Star, Optional, Cross ),
	}
	return result
}


func (parser RegexParser) parseUnionSubexpresion(input string) (RegExpresion, string, error){
	var prev_expresion *RegExpresion = nil
	var last_expresion RegExpresion
	var current_char string

	for len(input) > 0 {
		var current_expresion RegExpresion
		current_char, input = utils.Pop(input)

		if parser.alphabet_chars.Contains(current_char) {
			character_expresion := RegExpresion{ Characters, current_char, nil, nil}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &character_expresion}
		} else if current_char == Any {
			character_expresion := RegExpresion{ Any, "", nil, nil}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &character_expresion}
		} else if parser.single_operand_operators.Contains(current_char) {
			if prev_expresion == nil || prev_expresion.right_operand == nil {
				return RegExpresion{}, "", errors.New("Attempting to obtain right operand of nil")
			}
			prev_expresion.right_operand =  &RegExpresion{ current_char, "" , nil , prev_expresion.right_operand  }
			current_expresion = *prev_expresion
		} else if current_char == Parenthesis {
			var sub_expresion RegExpresion
			var err error
			sub_expresion, input, err  = parser.Parse(input)
			if err != nil {
				return RegExpresion{}, "", err
			}
			parenthesis_expresion := RegExpresion{ Parenthesis, "", &sub_expresion,  &sub_expresion}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &parenthesis_expresion}
		} else if current_char == ClosingParenthesis {
			return *prev_expresion, input, nil
		}else if current_char == Union  {
			return *prev_expresion, "|" + input, nil
		}

		prev_expresion = &current_expresion
		last_expresion = current_expresion
	}

	return last_expresion, input, nil
}


func (parser RegexParser) Parse(input string) (RegExpresion, string, error) {
	var prev_expresion *RegExpresion = nil
	var last_expresion RegExpresion
	var current_char string

	for len(input) > 0 {
		var current_expresion RegExpresion
		current_char, input = utils.Pop(input)

		if parser.alphabet_chars.Contains(current_char) {
			character_expresion := RegExpresion{ Characters, current_char, nil, nil}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &character_expresion}
		} else if current_char == Any {
			character_expresion := RegExpresion{ Any, "", nil, nil}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &character_expresion}
		} else if parser.single_operand_operators.Contains(current_char) {
			if prev_expresion == nil || prev_expresion.right_operand == nil {
				return RegExpresion{}, "", errors.New("Attempting to obtain right operand of nil")
			}
			prev_expresion.right_operand =  &RegExpresion{ current_char, "" , nil , prev_expresion.right_operand  }
			current_expresion = *prev_expresion
		} else if current_char == Parenthesis {
			var sub_expresion RegExpresion
			var err error
			sub_expresion, input, err  = parser.Parse(input)
			if err != nil {
				return RegExpresion{}, "", err
			}
			parenthesis_expresion := RegExpresion{ Parenthesis, "", &sub_expresion,  &sub_expresion}
			current_expresion = RegExpresion{ Concatenation, "", prev_expresion,  &parenthesis_expresion}
		} else if current_char == ClosingParenthesis {
			return *prev_expresion, input, nil
		} else if current_char == Union {
			var sub_expresion RegExpresion
			var err error
			sub_expresion, input, err  = parser.parseUnionSubexpresion(input)
			if err != nil {
				return RegExpresion{}, "", err
			}
			current_expresion = RegExpresion{ Union, "", prev_expresion,  &sub_expresion}
		}

		prev_expresion = &current_expresion
		last_expresion = current_expresion
	}

	return last_expresion, input, nil
}