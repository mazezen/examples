package models

type CMessage struct {
	Model    `xorm:"extends"`
	FromId   uint64 `xorm:"bigint notnull 'from_id' comment('消息发送者')" json:"from_id"`
	TargetId uint64 `xorm:"bigint notnull 'target_id' comment('消息接受者')" json:"target_id"`
	Type     uint32 `xorm:"tinyint notnull 'type' comment('聊天类型：1:私聊 2:群聊 3:广播')"`
	Media    uint32 `xorm:"tinyint notnull 'media' comment('1:文字 2:图片 3:音频')" json:"media"`
	Content  string `xorm:"text notnull 'content' comment('消息内容')" json:"content"`
	Pic      string `xorm:"text 'pic' comment('图片相关地址')" json:"pic"`
	Url      string `xorm:"VARCHAR(255) comment('文件相关地址')" json:"url"`
	Desc     string `xorm:"VARCHAR(255) comment('文件描述')" json:"desc"`
	Amount   uint64 `xorm:"bigint comment('文件大小')" json:"amount"`
}

func (m *CMessage) TableName() string {
	return "c_message"
}
