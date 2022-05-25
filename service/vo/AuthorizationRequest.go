package vo

import (
	"fmt"
)

type AuthorizationRequest struct {
	Request struct {

		//стоимость в минимальных денежных единицах.
		//Например, $32.45 должна быть отправлена как 3245
		Amount int `json:"amount"`

		// валюта в ISO-4217 формате, например USD
		Currency string `json:"currency"`

		// описание заказа. Максимальная длина: 255 символов
		Description string `json:"description"`

		//id транзакции или заказа в вашей системе.
		//Максимальная длина: 255 символов.
		//Пожалуйста, используйте уникальное значение для того, чтобы при запросе статуса транзакции получить актуальную информацию.
		//В противном случае вы получите первую найденную по tracking_id транзакцию
		TrackingId string `json:"tracking_id"`

		//параметр обязателен, если 3-D Secure включен.
		//Обратитесь к менеджеру за информацией. return_url - это URL на стороне торговца,
		//на который bePaid будет перенаправлять клиента после возврата с 3-D Secure проверки
		ReturnUrl string `json:"return_url"`

		//true или false. Транзакция будет тестовой, если значение true.
		Test bool `json:"test"`

		CreditCard CreditCard `json:"credit_card"`

		// секция, содержащая дополнительную информацию о платеже
		AdditionalData map[string]interface{} `json:"additional_data"`

		Customer Customer `json:"customer"`
	} `json:"request"`
}

func NewAuthorizationRequest() *AuthorizationRequest {
	return &AuthorizationRequest{}
}

type CreditCard struct {

	// номер карты, длина - от 12 до 19 цифр
	Number Number `json:"number"`

	//3-х или 4-х цифровой код безопасности (CVC2, CVV2 или CID, в зависимости от бренда карты).
	//Может быть отправлен вместе с параметром token и bePaid доставит банку-эквайеру данные карты с CVC2/CVV2/CID
	VerificationValue string `json:"verification_value"`

	//имя владельца карты. Максимальная длина: 32 символа
	Holder string `json:"holder"`

	//месяц окончания срока действия карты, представленный двумя цифрами (например, 01)
	ExpMonth ExpMonth `json:"exp_month"`

	//год срока окончания действия карты, представленный четырьмя цифрами (например, 2007)
	ExpYear ExpYear `json:"exp_year"`

	//опционально
	//вместо 5 параметров выше можно отправить токен карты, который был получен в ответе первой оплаты.
	//Если используется токен карты, то необходимо обязательно указывать параметр additional_data.contract
	Token string `json:"token"`

	//опционально
	//если значение параметра true, bePaid не выполняет 3-D Secure проверку.
	//Это полезно если вы, например, не хотите чтобы клиент проходил 3-D Secure проверку снова.
	//Уточните у службы поддержки, можете ли вы использовать этот параметр.
	SkipThreeDSecureVerification bool `json:"skip_three_d_secure_verification"`
}

type ExpMonth string

const (
	ExpMonthJan ExpMonth = "01"
	ExpMonthFeb ExpMonth = "02"
	ExpMonthMar ExpMonth = "03"
	ExpMonthApr ExpMonth = "04"
	ExpMonthMay ExpMonth = "05"
	ExpMonthJun ExpMonth = "06"
	ExpMonthJul ExpMonth = "07"
	ExpMonthAug ExpMonth = "08"
	ExpMonthSep ExpMonth = "09"
	ExpMonthOct ExpMonth = "10"
	ExpMonthNov ExpMonth = "11"
	ExpMonthDec ExpMonth = "12"
)

type Number string

func NewNumber(number string) (Number, error) {
	if len(number) < 12 || len(number) > 19 {
		return "", fmt.Errorf("invalid Number: string length should be between 12 and 19")
	}
	return Number(number), nil
}

type ExpYear string

func NewExpYear(year string) (ExpYear, error) {
	if len(year) != 4 {
		return "", fmt.Errorf("invalid ExpYear: string lenght should be exactly 4")
	}
	return ExpYear(year), nil
}

type Customer struct {

	//IP-адрес клиента, производящего оплату в вашем магазине
	Ip string `json:"ip"`

	//email клиента, производящего оплату в вашем магазине
	Email string `json:"email"`

	//id устройства клиента, производящего оплату в вашем магазине
	// необязательный (из примеров запросов)
	DeviceId string `json:"device_id"`

	//(необязательный) дата рождения клиента в формате ISO 8601 YYYY-MM-DD
	BirthDate string `json:"birth_date"`
}
