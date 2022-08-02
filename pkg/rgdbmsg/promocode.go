package rgdbmsg

import "time"

type (
	CreatePromocodeRequest struct {
		InvokerId  int64
		Promocode  string
		Payload    []byte
		UsageCount *int
		EndingAt   *time.Time
	}

	DeletePromocodeRequest struct {
		InvokerId   int64
		PromocodeId int64
	}
)
