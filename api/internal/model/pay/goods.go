package pay

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"xiaozhu/internal/config/mysql"
)

type Goods struct {
	Id           int    `json:"id" gorm:"id"`
	GameId       int    `json:"game_id" gorm:"game_id"`
	GoodsType    int8   `json:"goods_type" gorm:"goods_type"`       // 商品类型：1（内购）、2（点数卡）
	Developer    string `json:"developer" gorm:"developer"`         // 苹果开发者
	GoodsName    string `json:"goods_name" gorm:"goods_name"`       // 内购商品名
	GoodsDesc    string `json:"goods_desc" gorm:"goods_desc"`       // 内购商品描述
	GoodsId      string `json:"goods_id" gorm:"goods_id"`           // 内购ID
	Amount       int    `json:"amount" gorm:"amount"`               // 商品价格 单位分
	ActualAmount int    `json:"actual_amount" gorm:"actual_amount"` // 实付金额单位分
	Currency     string `json:"currency" gorm:"currency"`           // 货币单位:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...
	Status       int8   `json:"status" gorm:"status"`               // 商品状态: 关闭(1)、正常(0)
	CreatedAt    int64  `json:"created_at" gorm:"created_at"`       // 添加时间
}

// TableName 表名称
func (*Goods) TableName() string {
	return "pay_goods"
}

func NewGoods() *Goods {
	return &Goods{}
}

func (g *Goods) Find(ctx context.Context) error {
	if g.GoodsId == "" {
		return fmt.Errorf("商品id不能为空")
	}

	if g.GameId == 0 {
		return fmt.Errorf("游戏id不能为空")
	}

	err := mysql.PlatformDB.Model(&g).WithContext(ctx).First(&g).Error
	if err != nil {
		// 判断是否是记录不存在的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("该商品不存在")
		}

		// 其他数据库错误
		return fmt.Errorf("查询商品失败: %v", err)
	}

	return nil
}
