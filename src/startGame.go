package src

import (
	"fmt"
	"hangman/utilities"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func StartGame() {

	words := []string{
		"тест",
		"солнце",
		"лес",
	}
	actualAnswer := words[rand.Intn(len(words))]
	usersMoves := 0
	numberOfAvaliableMoves := 6
	userErrors := 0
	userAnswer := []rune(
		strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return '_'
			}
			return r
		}, actualAnswer))

	fmt.Print("\nЯ загадал слово. Ваша задача угадать букву, которая в него входит.\n")

	for usersMoves <= numberOfAvaliableMoves {
		printCurrentGameState(&usersMoves, &userErrors)

		if isGameFinished(&usersMoves, &userAnswer, actualAnswer, numberOfAvaliableMoves) {
			printResult(&usersMoves, &userAnswer, actualAnswer, numberOfAvaliableMoves)
			promptRestart()
		}
		makeMove(&usersMoves, &userAnswer, actualAnswer, &userErrors)
	}

}

func makeMove(movesCounter *int, userAnswer *[]rune, rightAnswer string, errorsCounter *int) {
	var input string

	fmt.Printf("Слово: %v \n", string(*userAnswer))
	fmt.Print("Введите букву: \n")

	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	isInputValid, errorText := validateInput(&input, *userAnswer)

	if !isInputValid {
		fmt.Print(errorText)
		return
	}

	runeInput := []rune(input)[0]
	indexes := utilities.FindAllIndexes(rightAnswer, runeInput)
	isCharRight := len(indexes) != 0

	if isCharRight {
		utilities.ReplaceUnderscoreByChar(*userAnswer, runeInput, indexes)
	} else {
		fmt.Print("Этой буквы нет\n")
		*errorsCounter += 1
	}

	*movesCounter += 1
}

func promptRestart() {
	var userInput string
	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	switch userInput {
	case "yes":
		StartGame()
	default:
		os.Exit(0)
	}
}

func validateInput(input *string, userAnswer []rune) (bool, string) {
	doesAnswerHasInputChar := strings.ContainsRune(string(userAnswer), []rune(*input)[0])

	if doesAnswerHasInputChar {
		return false, "Данный символ уже отгадан\n"
	}

	hasInputOneChar := len([]rune(*input)) != 1

	if hasInputOneChar {
		fmt.Print("Ввод может содержать только одну букву\n")
		return false, "Ввод может содержать только одну букву\n"
	}

	isCyrillicChar := utilities.IsCyrillicChar(*input)

	if !isCyrillicChar {
		return false, "Вы можете вводить только кириллицу\n"
	}

	isCharInLowerRegister := unicode.IsLower([]rune(*input)[0])

	if !isCharInLowerRegister {
		return false, "Буква может быть только в нижнем регистре\n"
	}

	return true, ""
}

func printCurrentGameState(movesCounter *int, errorsCounter *int) {
	var pictures = []string{
		"assets/firstMove.txt",
		"assets/secondMove.txt",
		"assets/thirdMove.txt",
		"assets/fourthMove.txt",
		"assets/fifthMove.txt",
		"assets/sixthMove.txt",
		"assets/seventhMove.txt",
	}
	path := pictures[*movesCounter]
	text, _ := os.ReadFile(path)

	fmt.Print(string(text))
	fmt.Printf("\nЧисло ошибок: %v\n", *errorsCounter)
}

func printResult(movesCounter *int, userAnswer *[]rune, rightAnswer string, numOfAvaliableMoves int) {
	if doesUserWin(userAnswer, rightAnswer) {
		fmt.Printf(
			"Вы победили, поздарвляю! Это было слово %s\nЧтобы начать новую игру введите 'yes' или введите любой символ для выхода и нажмите 'Enter'.\n",
			rightAnswer,
		)
	}

	if doesUserLose(movesCounter, numOfAvaliableMoves) {
		fmt.Print("Вы проиграли.\nЧтобы закончить или начать новую игру введите n / y\nЧтобы начать новую игру введите 'yes' или введите любой символ для выхода и нажмите 'Enter'.\n")
	}
}

func isGameFinished(movesCounter *int, userAnswer *[]rune, rightAnswer string, numOfAvaliableMoves int) bool {
	return doesUserWin(userAnswer, rightAnswer) || doesUserLose(movesCounter, numOfAvaliableMoves)
}

func doesUserWin(userAnswer *[]rune, rightAnswer string) bool {
	return string(*userAnswer) == rightAnswer
}

func doesUserLose(movesCounter *int, numOfAvaliableMoves int) bool {
	return *movesCounter == numOfAvaliableMoves
}
