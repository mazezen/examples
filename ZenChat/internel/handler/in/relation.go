package in

type RelationListPayload struct {
	Page     uint   `json:"page"`
	PageSize uint   `json:"pageSize"`
	UserId   uint64 `json:"user_id"`
}

type FriendAddByNamePayload struct {
	UserId     uint64 `json:"user_id"`
	TargetName string `json:"target_name"`
}

type RemoveRelationByName = FriendAddByNamePayload
