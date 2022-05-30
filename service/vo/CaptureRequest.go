package vo

type CaptureRequest struct {
	Request struct {

		//UID транзакции авторизации
		ParentUid string `json:"parent_uid"`

		//сумма списания в минимальных денежных единицах, например 1000 для $10.00
		Amount int64 `json:"amount"`

		//(необязательный) true или false. Параметр управляет процессом проверки входящего запроса на уникальность.
		//Если в течение 30 секунд придет запрос на списание средств с одинаковыми amount и parent_uid, то запрос будет отклонен.
		//По умолчанию, этот параметр имеет значение true
		DuplicateCheck *bool `json:"duplicate_check,omitempty"`
	} `json:"request"`
}

// NewCaptureRequest creates CaptureRequest with mandatory fields

func NewCaptureRequest(parentUid string, amount int64) *CaptureRequest {
	r := &CaptureRequest{}

	r.Request.ParentUid = parentUid
	r.Request.Amount = amount

	return r
}

func (cr *CaptureRequest) WithDuplicateCheck(duplicateCheck bool) *CaptureRequest {
	cr.Request.DuplicateCheck = &duplicateCheck
	return cr
}
