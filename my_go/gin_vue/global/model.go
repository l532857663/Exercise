package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type OrderPayment struct {
	Flag         bool  // 订单的循环控制
	Index        int64 // 循环队列的查询定位
	Temp         int64 // 放置数据的位置
	TempInterval int64 // 间隔时间对应的步数
	Head         int64 // 队列开始位置
	Rear         int64 // 队列末尾位置
}
