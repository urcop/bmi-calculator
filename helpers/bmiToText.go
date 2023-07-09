package helpers

func BmiToText(bmi float64) string {
	if bmi <= 16 {
		return "Выраженный дефицит массы тела"
	}
	if bmi <= 18.5 {
		return "Недостаточная (дефицит) масса тела"
	}
	if bmi <= 25 {
		return "Норма"
	}
	if bmi <= 30 {
		return "Избыточная масса тела (предожирение)"
	}
	if bmi <= 35 {
		return "Ожирение 1 степени"
	}
	if bmi <= 40 {
		return "Ожирение 2 степени"
	}

	return "Ожирение 3 степени"

}
