package models

type CUserRelation struct {
	Model    `xorm:"extends"`
	OwnerID  uint64 `xorm:"bigint 'owner_id' comment('谁的关系信息')" json:"owner_id"`
	TargetID uint64 `xorm:"bigint 'target_id' comment('对应的谁')" json:"target_id"`
	Type     int    `xorm:"tinyint 'type' comment('关系类型： 1表示好友关系')" json:"type"`
	Desc     string `xorm:"varchar(255) 'desc' comment('描述')" json:"desc"`
}

func (*CUserRelation) TableName() string {
	return "c_user_relation"
}
