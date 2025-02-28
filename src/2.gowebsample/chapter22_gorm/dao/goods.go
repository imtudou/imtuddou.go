package dao

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Good struct {
	Id         int       //表字段名为：id
	Name       string    //表字段名为：name
	Price      float64   //表字段名为：price
	TypeId     int       //表字段名为：type_id
	CreateTime time.Time `gorm:"column:createtime"` //表字段名为：createtime
}

func (*Good) TableName() string {
	return "gorm_goods"
}

// 钩子函数 钩子方法的函数签名应该是 func(*gorm.DB) error
func (g *Good) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate......")
	return nil
}

func (g *Good) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate......")
	return nil
}
func (g *Good) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave......")
	return nil
}

func (g *Good) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave......")
	return nil
}

func Save(good *Good) error {
	err := DB.Create(good).Error
	if err != nil {
		return err
	}
	return nil
}

// 自动事务
func Transcation1() error {
	// 执行事务的相关操作使用tx 而不是 db'
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&Good{
			Name:       "小米15",
			Price:      5999,
			TypeId:     1,
			CreateTime: time.Now(),
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&Good{
			Name:       "苹果25",
			Price:      5999,
			TypeId:     1,
			CreateTime: time.Now(),
		}).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// 手动事务
func Transcation2() error {

	// 开始事务
	tx := DB.Begin()
	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	err := tx.Create(&Good{
		Name:       "小米52",
		Price:      5999,
		TypeId:     1,
		CreateTime: time.Now(),
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.SavePoint("sp1")
	err = tx.Create(&Good{
		Name:       "小米253456987643",
		Price:      5999,
		TypeId:     1,
		CreateTime: time.Now(),
	}).Error

	if err != nil {
		//tx.Rollback()
		tx.RollbackTo("sp1") // 保存点 回滚到指定位置  这个代码预期效果就是只会保存上面的一条记录
		return err
	}

	tx.Commit()
	return nil
}

// 嵌套事务
func Transcation3() error {
	//tx.Create(&user1)
	//
	//tx.Transaction(func(tx2 *gorm.DB) error {
	//	tx2.Create(&user2)
	//	return errors.New("rollback user2") // Rollback user2
	//})
	//
	//tx.Transaction(func(tx3 *gorm.DB) error {
	//	tx3.Create(&user3)
	//	return nil
	//})
	//
	return nil
}
