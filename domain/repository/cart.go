package repository

import (
	"errors"
	"micro_cart/domain/model"

	"gorm.io/gorm"
)

type ICartRepository interface {
	//初始化数据表
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDB: db}
}

func (u *CartRepository) InitTable() error {
	if u.mysqlDB.Migrator().HasTable(&model.Cart{}) {
		return nil
	}
	return u.mysqlDB.Migrator().CreateTable(&model.Cart{})
}

func (u *CartRepository) FindCartByID(cartId int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDB.First(cart, cartId).Error
}

func (u *CartRepository) CreateCart(cart *model.Cart) (pid int64, err error) {
	db := u.mysqlDB.FirstOrCreate(cart, model.Cart{ProductId: cart.ProductId, SizeId: cart.SizeId, UserId: cart.UserId})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车创建失败")
	}
	return cart.Id, nil
}

func (u *CartRepository) DeleteCartByID(cartId int64) error {
	return u.mysqlDB.Where("id = ?", cartId).Delete(&model.Cart{}).Error
}

func (u *CartRepository) UpdateCart(cart *model.Cart) error {
	return u.mysqlDB.Model(cart).Save(&cart).Error
}

func (u *CartRepository) FindAll(userId int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDB.Where("uid = ?", userId).Find(&cartAll).Error
}

func (u *CartRepository) CleanCart(userId int64) error {
	return u.mysqlDB.Where("user_id = ? ", userId).Delete(&model.Cart{}).Error
}
func (u *CartRepository) IncrNum(cartId int64, num int64) error {
	cart := &model.Cart{Id: cartId}
	return u.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}
func (u *CartRepository) DecrNum(cartId int64, num int64) error {
	cart := &model.Cart{Id: cartId}
	db := u.mysqlDB.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("购物车减少失败")
	}
	return nil
}
