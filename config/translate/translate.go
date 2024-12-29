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

var currentLang map[string]string

var t = Translate{
	UA: map[string]string{
		"garageNumber": "№ гаражу",
		"fullName":     "Імʼя та прізвище",
		"phoneNumber":  "Телефон",
		"address":      "Адреса",

		"enterGarageNumber": "введіть № гаражу",
		"enterFullName":     "введіть Імʼя та прізвище",
		"enterPhoneNumber":  "введіть Телефон",
		"enterAddress":      "введіть Адреса",
		// "lastPayedYear":   "Останній сплачений рік",
		"edit":            "Редагувати",
		"paymentButton":   "Оплата",
		"paymentFormName": "Оплата",
		"create":          "Створити",
		"cancel":          "Відміна",
		"amount":          "Сума",
		"selectYearPromt": "Оберіть тариф",
		"showPayments":    "Оплати",
		"addAccount":      "Додати",
		"done":            "Застосувати",
		"createdAt":       "Дата створення запису",

		"garageNumberBlankError": "Номер гаражу повинен бути заповнений",
		"fullNameBlankError":     "Імʼя та прізвище повинне бути заповнене",
		"phoneNumberBlankError":  "Номер нелефону повинен бути заповнений",
		"addressBlankError":      "Адреса повинна бути заповнений",
		"searchSign":             "🔍",
	},
	EN: map[string]string{},
}

func SetLang(l string) error {
	switch l {
	case UA:
		currentLang = t.UA
	case EN:
		currentLang = t.EN
	default:
		return fmt.Errorf("unknown language %s", l)
	}

	return nil
}

func T(key string) string {
	if cu, ok := currentLang[key]; ok {
		return cu
	} else {
		return fmt.Sprintf("missing transtation['%s']", key)
	}
}
