package in

type ChatMsgPayload struct {
	UserId   uint64 `json:"user_id" validate:"required"`
	TargetId uint64 `json:"target_id" validate:"required"`
	Start    int64  `json:"start"`
	End      int64  `json:"end"`
	IsRev    bool   `json:"is_rev" validate:"required"`
}
