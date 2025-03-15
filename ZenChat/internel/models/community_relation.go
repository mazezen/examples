package models

type CCommunityRelation struct {
	Model    `xorm:"extends"`
	OwnerID  uint64 `xorm:"bigint 'owner_id' comment('群成员ID')" json:"owner_id"`
	TargetID uint64 `xorm:"bigint 'target_id' comment('群ID')" json:"target_id"`
	Desc     string `xorm:"varchar(255) 'desc' comment('描述')" json:"desc"`
}

func (CCommunityRelation) TableName() string {
	return "c_community_relation"
}
