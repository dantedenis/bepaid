package vo

const (
	URLRequest = "https://gateway.bepaid.by/transactions/captures"
)

type CaptureRequest struct {
	Amount    int    `json:"amount"`
	ParentUid string `json:"parent_uid"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Uid       string `json:"uid"`
	GatewayId int    `json:"gateway_id"`
}

func NewCaptureRequest(amount int, parentUid string) *CaptureRequest {
	return &CaptureRequest{Amount: amount, ParentUid: parentUid}
}
