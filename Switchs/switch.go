package switchs

func DayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Sábado"
	case 3:
		return "Segunda-Feira"
	case 4:
		return "Terça-Feira"
	case 5:
		return "Quarta-Feira"
	case 6:
		return "Quinta-Feira"
	case 7:
		return "Sexta-Feira"
	default:
		return "Número Inválido"
	}
}

func DayOfTheWeek2(number int) (dayOfTheWeek string) {

	switch {
	case number == 1:
		dayOfTheWeek = "Domingo"
		fallthrough
	case number == 2:
		dayOfTheWeek = "Sábado"
	case number == 3:
		dayOfTheWeek = "Segunda-Feira"
	case number == 4:
		dayOfTheWeek = "Terça-Feira"
	case number == 5:
		dayOfTheWeek = "Quarta-Feira"
	case number == 6:
		dayOfTheWeek = "Quinta-Feira"
	case number == 7:
		dayOfTheWeek = "Sexta-Feira"
	default:
		return "Número Inválido"
	}

	return
}
