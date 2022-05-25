package vo

const (
	URLRequest = "https://gateway.bepaid.by/transactions/captures"	
)

type CaptureRequest struct {
	Amount int
	Parent_uid string
	Status string
	Message string
	Uid	string
	Gateway_id int
}

func NewCaptureRequest(amount int, parent_uid string) *CaptureRequest {
	return &CaptureRequest{Amount: amount, Parent_uid: parent_uid}
}
