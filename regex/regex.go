package regex

const (
	Parenthesis			= "("
	ClosingParenthesis	= ")"
	Concatenation		= "||"
	Union				= "|"
	Optional			= "?"
	Star				= "*"
	Cross				= "+"
	Characters			= "a"
	Any					= "."
)

type RegExpresion struct {
	regex_type string
	value string
	left_operand *RegExpresion
	right_operand *RegExpresion
}


func ptrToStr(expr *RegExpresion) string {
	if expr == nil {
		return ""
	}
	return expr.String()
}

func (expr RegExpresion) String() string {

	switch expr.regex_type {
	case Characters:
		return expr.value
	case Any:
		return "."
	case Concatenation:
		return ptrToStr(expr.left_operand) + ptrToStr(expr.right_operand)
	case Star:
		return ptrToStr(expr.right_operand) + "*"
	case Cross:
		return ptrToStr(expr.right_operand) + "+"
	case Optional:
		return ptrToStr(expr.right_operand) + "?"
	case Parenthesis:
		return "(" + ptrToStr(expr.right_operand) + ")"
	case Union:
		return ptrToStr(expr.left_operand) + "|" + ptrToStr(expr.right_operand)
	}

	return " "
}