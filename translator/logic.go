package translator

import "RuLang2GoLang/utils"

func realizarOpMath(l, r Value, op string, line int) Value {
	switch l.value.(type) {
	case int64:
		switch r.value.(type) {
		case int64:
			//
			return realizarOpInt(l.value.(int64), r.value.(int64), op)
		case float64:
			utils.SyntaxError("Imposible to operate INT"+op+"FLOAT", line)
		default:
			utils.SyntaxError("Imposible to operate INT"+op+"INT", line)
		}
	case float64:
		switch r.value.(type) {
		case float64:
			return realizarOpFloat(l.value.(float64), r.value.(float64), op, line)
		case int64:
			utils.SyntaxError("Imposible to operate FLOAT"+op+"INT", line)
		default:
			utils.SyntaxError("Imposible to operate FLOAT"+op+"FLOAT", line)
		}
	default:
		utils.SyntaxError("Operator"+op+"only allows FLOAT or INT", line)
	}
	return Value{value: false}
}

func realizarOpInt(l, r int64, op string) Value {
	switch op {
	case "<":
		return Value{value: l < r}
	case ">":
		return Value{value: l > r}
	case ">=":
		return Value{value: l >= r}
	case "<=":
		return Value{value: l <= r}
	case "+":
		return Value{value: l + r}
	case "-":
		return Value{value: l - r}
	case "*":
		return Value{value: l * r}
	case "/":
		return Value{value: l / r}
	case "%":
		return Value{value: l % r}
	}
	return Value{value: false}
}

func realizarOpFloat(l, r float64, op string, line int) Value {
	switch op {
	case "<":
		return Value{value: l < r}
	case ">":
		return Value{value: l > r}
	case ">=":
		return Value{value: l >= r}
	case "<=":
		return Value{value: l <= r}
	case "+":
		return Value{value: l + r}
	case "-":
		return Value{value: l - r}
	case "*":
		return Value{value: l * r}
	case "/":
		return Value{value: l / r}
	case "%":
		utils.SyntaxError("FLOAT % FLOAT invalid operation ", line)
	}
	return Value{value: false}
}

func realizarOpBool(l, r Value, op string, line int) Value {
	switch l.value.(type) {
	case bool:
		switch r.value.(type) {
		case bool:
			rv := r.value.(bool)
			lv := l.value.(bool)
			if op == "&&" {
				return Value{value: rv && lv}
			} else if op == "||" {
				return Value{value: rv || lv}
			} else {
				utils.SyntaxError("BOOL"+op+"BOOL invalid operation", line)
			}
		default:
			utils.SyntaxError("Operator "+op+" operates only between BOOLs", line)
		}
	default:
		utils.SyntaxError("Operator "+op+" operates only between BOOLs", line)
	}
	return Value{value: false}
}

func realizarIgualdad(l, r Value, op string, line int) Value {
	switch l.value.(type) {
	case int64:
		switch r.value.(type) {
		case int64:
			//
			return realizarIgualdadInt(l.value.(int64), r.value.(int64))
		default:
			utils.SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case float64:
		switch r.value.(type) {
		case float64:
			return realizarIgualdadFloat(l.value.(float64), r.value.(float64))
		default:
			utils.SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case bool:
		switch r.value.(type) {
		case bool:
			return realizarIgualdadBool(l.value.(bool), r.value.(bool))
		default:
			utils.SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case string:
		switch r.value.(type) {
		case string:
			return realizarIgualdadString(l.value.(string), r.value.(string))
		default:
			utils.SyntaxError("Operation "+op+" only works with same type,", line)
		}
	default:
		utils.SyntaxError("Operator"+op+"only allows FLOAT or INT", line)
	}
	return Value{value: false}
}

func realizarIgualdadString(s, s2 string) Value {
	return Value{value: s == s2}
}

func realizarIgualdadBool(b bool, b2 bool) Value {
	return Value{value: b == b2}
}

func realizarIgualdadFloat(f float64, f2 float64) Value {
	return Value{value: f == f2}
}

func realizarIgualdadInt(i int64, i2 int64) Value {
	return Value{value: i == i2}
}

func realizarNot(l Value, line int) Value {
	switch l.value.(type) {
	case bool:
		return Value{value: !l.value.(bool)}
	}
	utils.SyntaxError(" ! (not) operation is only for BOOL", line)

	return Value{value: false}
}

func realizarNegativo(l Value, line int) Value {
	switch l.value.(type) {
	case int64:
		return Value{value: -l.value.(int64)}
	case float64:
		return Value{value: -l.value.(float64)}
	}
	utils.SyntaxError(" - (negative) operation is only for INT or FLOAT", line)

	return Value{value: false}
}
