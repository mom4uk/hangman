package internal

import (
	"fmt"
)

func Run() {
	var userInput string
	fmt.Println("Добро пожаловать! \nДля начала игры введите 'start' или введите любой символ для выхода и нажмите 'Enter'.")

	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	switch userInput {
	case "start":
		StartGame()
		return
	default:
		fmt.Println("Введите 'start'")
	}
}
