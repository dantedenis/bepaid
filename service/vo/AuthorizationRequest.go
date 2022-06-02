package vo

type AuthorizationRequest struct {
	Request struct {

		//стоимость в минимальных денежных единицах.
		//Например, $32.45 должна быть отправлена как 3245
		Amount int64 `json:"amount"`

		//валюта в ISO-4217 формате, например USD
		Currency string `json:"currency"`

		//описание заказа. Максимальная длина: 255 символов
		Description string `json:"description"`

		//id транзакции или заказа в вашей системе.
		//Максимальная длина: 255 символов.
		//Пожалуйста, используйте уникальное значение для того, чтобы при запросе статуса транзакции получить актуальную информацию.
		//В противном случае вы получите первую найденную по tracking_id транзакцию
		TrackingId string `json:"tracking_id"`

		//(необязательный) true или false.
		//Параметр управляет процессом проверки входящего запроса на уникальность.
		//Если в течение 30 секунд придет запрос на авторизацию с одинаковыми amount и number или token, то запрос будет отклонен.
		//По умолчанию, этот параметр имеет значение true
		DuplicateCheck *bool `json:"duplicate_check,omitempty"`

		//параметр обязателен, если 3-D Secure включен.
		//Обратитесь к менеджеру за информацией. return_url - это URL на стороне торговца,
		//на который bePaid будет перенаправлять клиента после возврата с 3-D Secure проверки
		ReturnUrl string `json:"return_url,omitempty"`

		//true или false. Транзакция будет тестовой, если значение true.
		Test bool `json:"test"`

		CreditCard CreditCard `json:"credit_card"`

		//секция, содержащая дополнительную информацию о платеже
		AdditionalData map[string]interface{} `json:"additional_data,omitempty"`

		Customer *Customer `json:"customer,omitempty"`
	} `json:"request"`
}

// NewAuthorizationRequest creates AuthorizationRequest with mandatory fields
func NewAuthorizationRequest(amount int64, currency, description, trackingId string, test bool, cc CreditCard) *AuthorizationRequest {

	r := &AuthorizationRequest{}
	r.Request.Amount = amount
	r.Request.Currency = currency
	r.Request.Description = description
	r.Request.TrackingId = trackingId
	r.Request.Test = test
	r.Request.CreditCard = cc

	return r
}

func (cr *AuthorizationRequest) WithDuplicateCheck(duplicateCheck bool) *AuthorizationRequest {
	cr.Request.DuplicateCheck = &duplicateCheck
	return cr
}

func (a *AuthorizationRequest) WithReturnUrl(returnUrl string) *AuthorizationRequest {
	a.Request.ReturnUrl = returnUrl
	return a
}

// WithAdditionalData saves argument to AuthorizationRequest.Request.AdditionalData field.
//
// Don't change content of additionalData after function call.
func (a *AuthorizationRequest) WithAdditionalData(additionalData map[string]interface{}) *AuthorizationRequest {
	a.Request.AdditionalData = additionalData
	return a
}

func (a *AuthorizationRequest) WithCustomer(customer Customer) *AuthorizationRequest {
	a.Request.Customer = &customer
	return a
}

func (a *AuthorizationRequest) SetTest(test bool) {
	a.Request.Test = test
}
