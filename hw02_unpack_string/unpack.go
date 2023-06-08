package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var b strings.Builder
	prevSubstr := ""

	for _, currentRune := range str {
		currentChar := string(currentRune)
		currentDigit, err := strconv.Atoi(currentChar)
		currentRuneIsDigit := err == nil

		switch prevSubstr {
		case "":
			// Если предыдущая подстрока пустая, то текущий символ не может быть цифрой
			if currentRuneIsDigit {
				return "", ErrInvalidString
			}
			// Если не цифра, то записываем символ в качестве предыдущей подстроки и переходим к следующей итерации
			prevSubstr = currentChar
			continue
		case "\\":
			// Если в предыдущей строке символ экранирования, то текущий символ должен быть или слешем или цифрой
			if !currentRuneIsDigit && currentChar != "\\" {
				return "", ErrInvalidString
			}
			// Если всё ок, то добавляем текущий символ в предыдущую подстроку и переходим к следующей итерации
			prevSubstr += currentChar
			continue
		default:
			// В данном кейсе предыдущая подстрока содержит символ (может быть и цифровой, и буквенный, и спецсимвол)
			// или 2 символа: экранирующий слеш + символ (цифра или слеш). Убираем экранирующий слеш
			prevSubstr = strings.Replace(prevSubstr, "\\", "", 1)

			// Если текущий символ не цифра, то добавляем предыдущую подстроку в итоговую строку 1 раз
			// и сохраняем текущий символ в качестве предыдущей подстроки
			// В противном случае добавляем currentDigit раз и зачищаем предыдущую подстроку
			if currentRuneIsDigit {
				b.WriteString(strings.Repeat(prevSubstr, currentDigit))
				prevSubstr = ""
			} else {
				b.WriteString(prevSubstr)
				prevSubstr = currentChar
			}
		}
	}
	// После прогона цикла в подстроке могли остаться не записанные в итоговую строку символы
	// Если там только 1 слеш (экранирование) - то это ошибка
	if prevSubstr == "\\" {
		return "", ErrInvalidString
	}
	// Остальные варианты нужно добавить в итоговую строку, удалив возможный экранирующий слеш
	b.WriteString(strings.Replace(prevSubstr, "\\", "", 1))

	return b.String(), nil
}
