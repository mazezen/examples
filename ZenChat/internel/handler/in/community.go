package in

type CreateCommunityPayload struct {
	OwnerId uint64 `json:"owner_id" validate:"required,numeric"` // 群主ID
	Name    string `json:"name" validate:"required,max=20"`      // 群名称
	Type    uint32 `json:"type" validate:"required,numeric"`     // 群类型
	Avatar  string `json:"avatar"`                               // 群头像
	Desc    string `json:"desc"`                                 // 群描述
}

type UpdateCommunityPayload struct {
	CommunityId uint64 `json:"community_id" validate:"required,numeric"` // 群ID
	Name        string `json:"name"`                                     // 群名称
	Type        uint32 `json:"type"`                                     // 群类型
	Avatar      string `json:"avatar"`                                   // 群头像
	Desc        string `json:"desc"`                                     // 群描述
}

type JoinCommunityPayload struct {
	CommunityId uint64 `json:"community_id" validate:"required,numeric"` // 群ID
	OwnerId     uint64 `json:"owner_id" validate:"required,numeric"`     // 用户ID
}

type ListByOwnerPayload struct {
	Page     uint   `json:"page"`
	PageSize uint   `json:"page_size"`
	OwnerId  uint64 `json:"owner_id" validate:"required,numeric"` // 用户ID
}
