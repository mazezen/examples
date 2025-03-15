package models

import "time"

type Model struct {
	ID        uint64    `xorm:"bigint not null pk autoincr 'id' comment('主键ID')" json:"id"`
	CreatedAt time.Time `xorm:"datetime 'created' comment('创建时间')" json:"created"`
	UpdatedAt time.Time `xorm:"datetime 'updated' comment('更新时间')" json:"updated"`
	DeletedAt time.Time `xorm:"datetime 'deleted' comment('删除时间')" json:"deleted"`
}
