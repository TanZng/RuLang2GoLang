package translator

import "math"

func realizarOpMath(l, r Value, op string, line int) Value {
	switch l.Value.(type) {
	case int64:
		switch r.Value.(type) {
		case int64:
			//
			return realizarOpInt(l.Value.(int64), r.Value.(int64), op, line)
		case float64:
			return SyntaxError("Imposible to operate INT"+op+"FLOAT", line)
		default:
			return SyntaxError("Imposible to operate INT"+op+"INT", line)
		}
	case float64:
		switch r.Value.(type) {
		case float64:
			return realizarOpFloat(l.Value.(float64), r.Value.(float64), op, line)
		case int64:
			return SyntaxError("Imposible to operate FLOAT"+op+"INT", line)
		default:
			return SyntaxError("Imposible to operate FLOAT"+op+"FLOAT", line)
		}
	default:
		return SyntaxError("Operator"+op+"only allows FLOAT or INT", line)
	}
	return Value{Value: false}
}

func realizarOpInt(l, r int64, op string, line int) Value {
	switch op {
	case "<":
		return Value{Value: l < r}
	case ">":
		return Value{Value: l > r}
	case ">=":
		return Value{Value: l >= r}
	case "<=":
		return Value{Value: l <= r}
	case "+":
		return Value{Value: l + r}
	case "-":
		return Value{Value: l - r}
	case "*":
		return Value{Value: l * r}
	case "/":
		return Value{Value: l / r}
	case "%":
		return Value{Value: l % r}
	case "^":
		return Value{Value: math.Pow(float64(l), float64(r))}
	}
	return SyntaxError("Invalid operation"+op, line)
}

func realizarOpFloat(l, r float64, op string, line int) Value {
	switch op {
	case "<":
		return Value{Value: l < r}
	case ">":
		return Value{Value: l > r}
	case ">=":
		return Value{Value: l >= r}
	case "<=":
		return Value{Value: l <= r}
	case "+":
		return Value{Value: l + r}
	case "-":
		return Value{Value: l - r}
	case "*":
		return Value{Value: l * r}
	case "/":
		return Value{Value: l / r}
	case "%":
		return SyntaxError("FLOAT % FLOAT invalid operation ", line)
	case "^":
		return Value{Value: math.Pow(l, r)}
	}
	return SyntaxError("Invalid operation"+op, line)
}

func realizarOpBool(l, r Value, op string, line int) Value {
	switch l.Value.(type) {
	case bool:
		switch r.Value.(type) {
		case bool:
			rv := r.Value.(bool)
			lv := l.Value.(bool)
			if op == "&&" {
				return Value{Value: rv && lv}
			} else if op == "||" {
				return Value{Value: rv || lv}
			} else {
				return SyntaxError("BOOL"+op+"BOOL invalid operation", line)
			}
		default:
			return SyntaxError("Operator "+op+" operates only between BOOLs", line)
		}
	default:
		return SyntaxError("Operator "+op+" operates only between BOOLs", line)
	}
	return Value{Value: false}
}

func realizarIgualdad(l, r Value, op string, line int) Value {
	switch l.Value.(type) {
	case int64:
		switch r.Value.(type) {
		case int64:
			//
			return realizarIgualdadInt(l.Value.(int64), r.Value.(int64))
		default:
			return SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case float64:
		switch r.Value.(type) {
		case float64:
			return realizarIgualdadFloat(l.Value.(float64), r.Value.(float64))
		default:
			return SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case bool:
		switch r.Value.(type) {
		case bool:
			return realizarIgualdadBool(l.Value.(bool), r.Value.(bool))
		default:
			return SyntaxError("Operation "+op+" only works with same type,", line)
		}
	case string:
		switch r.Value.(type) {
		case string:
			return realizarIgualdadString(l.Value.(string), r.Value.(string))
		default:
			return SyntaxError("Operation "+op+" only works with same type,", line)
		}
	default:
		return SyntaxError("Operator"+op+"only allows FLOAT or INT", line)
	}
}

func realizarIgualdadString(s, s2 string) Value {
	return Value{Value: s == s2}
}

func realizarIgualdadBool(b bool, b2 bool) Value {
	return Value{Value: b == b2}
}

func realizarIgualdadFloat(f float64, f2 float64) Value {
	return Value{Value: f == f2}
}

func realizarIgualdadInt(i int64, i2 int64) Value {
	return Value{Value: i == i2}
}

func realizarNot(l Value, line int) Value {
	switch l.Value.(type) {
	case bool:
		return Value{Value: !l.Value.(bool)}
	}
	return SyntaxError(" ! (not) operation is only for BOOL", line)
}

func realizarNegativo(l Value, line int) Value {
	switch l.Value.(type) {
	case int64:
		return Value{Value: -l.Value.(int64)}
	case float64:
		return Value{Value: -l.Value.(float64)}
	}
	return SyntaxError(" - (negative) operation is only for INT or FLOAT", line)
}
