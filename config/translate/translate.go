package translate

import "fmt"

const (
	EN = "EN"
	UA = "UA"
)

type Translate struct {
	EN map[string]string
	UA map[string]string
}

var T map[string]string

var t = Translate{
	UA: map[string]string{
		"garageNumber":    "№ гаражу",
		"fullName":        "Імʼя та прізвище",
		"phoneNumber":     "Телефон",
		"address":         "Адреса",
		"lastPayedYear":   "Останній сплачений рік",
		"edit":            "Редагувати",
		"paymentButton":   "Оплата",
		"paymentFormName": "Оплата",
		"create":          "Створити",
		"cancel":          "Відміна",
		"amount":          "Сума",
		"selectYearPromt": "Оберіть тариф",

		"garageNumberBlankError": "Номер гаражу повинен бути заповнений",
		"fullNameBlankError":     "Імʼя та прізвище повинне бути заповнене",
		"phoneNumberBlankError":  "Номер нелефону повинен бути заповнений",
		"addressBlankError":      "Адреса повинна бути заповнений",
	},
	EN: map[string]string{},
}

func SetLang(l string) error {
	switch l {
	case UA:
		T = t.UA
	case EN:
		T = t.EN
	default:
		return fmt.Errorf("unknown language %s", l)
	}

	return nil
}
