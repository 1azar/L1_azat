package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// УСЛОВИЕ:
// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ввод выражения от пользователя
		fmt.Print("Введите выражение (для выхода введите 'exit'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		// Разбиваем введенное выражение на операнды и оператор
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Некорректный ввод. Пожалуйста, введите выражение в формате '<число> <оператор> <число>'.")
			continue
		}

		operand1 := new(big.Float)
		operand2 := new(big.Float)

		// Преобразуем операнды в big.Float
		if _, success := operand1.SetString(parts[0]); !success {
			fmt.Println("Ошибка при преобразовании первого операнда.")
			continue
		}

		operator := parts[1]

		if _, success := operand2.SetString(parts[2]); !success {
			fmt.Println("Ошибка при преобразовании второго операнда.")
			continue
		}

		// Выполняем операцию в зависимости от введенного оператора
		result := new(big.Float)
		switch operator {
		case "+":
			result.Add(operand1, operand2)
		case "-":
			result.Sub(operand1, operand2)
		case "*":
			result.Mul(operand1, operand2)
		case "/":
			if operand2.Cmp(big.NewFloat(0)) == 0 {
				fmt.Println("Ошибка: деление на ноль.")
				continue
			}
			result.Quo(operand1, operand2)
		default:
			fmt.Println("Неподдерживаемая операция. Поддерживаемые операции: +, -, *, /.")
			continue
		}

		// Вывод результата
		fmt.Printf("Результат: %s %s %s = %s\n", operand1.String(), operator, operand2.String(), result.String())
	}

}
