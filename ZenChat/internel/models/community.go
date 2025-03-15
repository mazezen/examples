package models

type CCommunity struct {
	Model   `xorm:"extends"`
	Name    string `xorm:"varchar(50) notnull unique 'name' comment('群聊名称')" json:"name"`
	OwnerId uint64 `xorm:"bigint notnull 'owner_id' comment('群主ID')" json:"owner_id"`
	Type    uint32 `xorm:"tinyint notnull 'type' comment('群类型')" json:"type"`
	Avatar  string `xorm:"varchar(255) 'avatar' comment('群头像')" json:"avatar"`
	Desc    string `xorm:"varchar(255) 'desc' comment('群描述')" json:"desc"`
}

func (*CCommunity) TableName() string {
	return "c_community"
}
