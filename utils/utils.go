package utils

import "math/rand"

func GenerateRandomInt(maxInt int) int {
	return rand.Intn(maxInt)
}

func GetNames() []string {
	return []string{"Boris", "Kirill", "Andrey", "Svetlana", "Dmitry", "Olga", "Ivan", "Elena", "Alexey", "Natalia", "Vladimir", "Anastasia", "Sergey", "Yulia", "Artem", "Maria", "Maxim", "Ekaterina", "Nikita", "Anna", "Pavel", "Victoria", "Anton", "Eva", "Denis", "Tatiana", "Roman", "Julia", "Evgeny", "Larisa"}
}
