package calc
import (
	"fmt"
	"strconv"
	"strings"
	"tracker/cmd/task/errorhandler"
)

const (
	operAddition       = "+"
	operSubtraction    = "-"
	operMultiplication = "*"
	operDivision       = "/"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) CalculateTwoArg(number1 float64, number2 float64, operator string)(float64, error) {
	switch operator {
		case operAddition:
			return c.add(number1, number2)
	 	case operSubtraction:
			return c.subtract(number1, number2)
	 	case operMultiplication:
			return c.multiply(number1, number2)
	 	case operDivision:
			return c.divide(number1, number2)
	 	default:
		 	return 0, fmt.Errorf("Ошибка: неподдерживаемый оператор")
	}
}

func (c *calculator) add(number1, number2 float64) (float64, error) {
	return number1 + number2, nil
}
func (c *calculator) subtract(number1, number2 float64) (float64, error) {
	return number1 - number2, nil
}
func (c *calculator) multiply(number1, number2 float64) (float64, error) {
	return number1 * number2, nil
}
func (c *calculator) divide(number1, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errorhandler.PrintError(fmt.Errorf("некорректное выражение - деление на ноль"))  //fmt.Errorf("некорректное выражение - деление на ноль")
	}else{
		return number1 / number2, nil
	}
}

func (c *calculator) CalculateExpression(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
   
	// 1: Проверяем корректность скобок в выражении
	if !areParenthesesBalanced(expression) {
	 	return 0, fmt.Errorf("некорректное выражение - непарные скобки")
	}
   
	// 2: Разделяем выражение на отдельные элементы
	elements := tokenize(expression)
   
	// 3: Вычисляем результат
	result, err := c.evaluate(elements)
	if err != nil {
	 	return 0, err
	}

	return result, nil
}

// 1: Проверяем корректность скобок в выражении
func areParenthesesBalanced(expression string) bool {
	var stack []rune
   
	for _, char := range expression {
	 	if char == '(' {
	  		stack = append(stack, char)
	 	} else if char == ')' {
	  		if len(stack) == 0 || stack[len(stack)-1] != '(' {
	   			return false
	  		}
	  		stack = stack[:len(stack)-1] // Удаляем последнюю открывающую скобку из стека
	 	}
	}
   
	return len(stack) == 0
}

// 2: Разделяем выражение на отдельные элементы
func tokenize(expression string) []string {
	var tokens []string
   
	var currentToken string
	for _, char := range expression {
	 	if isOperator(string(char)) || char == '(' || char == ')' {
	  		if len(currentToken) > 0 {
	   			tokens = append(tokens, currentToken)
	  		}
	  		tokens = append(tokens, string(char))
	  		currentToken = ""
	 	} else {
	  		currentToken += string(char)
	 	}
	}
   
	if len(currentToken) > 0 {
	 	tokens = append(tokens, currentToken)
	}
   
	return tokens
}

// 3: Вычисляем результат
func (c *calculator)evaluate(tokens []string) (float64, error) {
	var numbers []float64
	var operators []string
   
	for _, token := range tokens {
	 	if isNumber(token) {
	  		number, err := strconv.ParseFloat(token, 64)
	  		if err != nil {
	   			return 0, fmt.Errorf("некорректное выражение - ошибка в числе")
	  		}
	  		numbers = append(numbers, number)
	 		} else if isOperator(token) {
	  			for len(operators) > 0 && hasHigherPrecedence(operators[len(operators)-1], token) {
	   				err := c.applyOperation(&numbers, &operators)
	   				if err != nil {
						return 0, err
	   				}
	  			}
	  			operators = append(operators, token)
	 		} else if token == "(" {
	  			operators = append(operators, token)
	 		} else if token == ")" {
	  			for len(operators) > 0 && operators[len(operators)-1] != "(" {
	   				err := c.applyOperation(&numbers, &operators)
	   				if err != nil {
						return 0, err
	   				}
	  			}
	  			if len(operators) == 0 || operators[len(operators)-1] != "(" {
	   				return 0, fmt.Errorf("некорректное выражение - непарные скобки")
	  			}
	  			operators = operators[:len(operators)-1] // Удаляем последнюю открывающую скобку из стека
	 		}
	}
   
	for len(operators) > 0 {
	 	err := c.applyOperation(&numbers, &operators)
	 	if err != nil {
	  		return 0, err
	 	}
	}
   
	if len(numbers) == 0 {
	 	return 0, fmt.Errorf("некорректное выражение - отсутствуют числа")
	}
  
	return numbers[0], nil
}

// Проверяем число на валидность
func isNumber(token string) bool {
    _, err := strconv.ParseFloat(token, 64)
    return err == nil
}

// Получаем число в нужном формате
func GetNumber(token string) (float64, error) {
    val, err := strconv.ParseFloat(token, 64)
    return val, err
} 

// Проверяем оператор на валидность
func isOperator(token string) bool {
	operators := map[string]bool{operAddition: true, operSubtraction: true, operMultiplication: true, operDivision: true}
	return operators[token]
}

// Определяем приоритеты для операций
func hasHigherPrecedence(operator1, operator2 string) bool {
	precedence := map[string]int{operAddition: 1, operSubtraction: 1, operMultiplication: 2, operDivision: 2}
	return precedence[operator1] >= precedence[operator2]
}

// Делаем необходимые вычисления  
func (c *calculator)applyOperation(numbers *[]float64, operators *[]string) error {
	if len(*numbers) < 2 {
		return fmt.Errorf("некорректное выражение - недостаточно чисел для операции")
	}
	  
	b := (*numbers)[len(*numbers)-1]
	(*numbers) = (*numbers)[:len(*numbers)-1]
	  
	a := (*numbers)[len(*numbers)-1]
	(*numbers) = (*numbers)[:len(*numbers)-1]
	  
	operator := (*operators)[len(*operators)-1]
	(*operators) = (*operators)[:len(*operators)-1]

	var result float64
	var err error
	result, err = c.CalculateTwoArg(a,  b, operator)

	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))
	}else{
			(*numbers) = append((*numbers), result)
	}
	  
	return nil
}
	  
  








 