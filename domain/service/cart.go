package service

import (
	"micro_cart/domain/model"
	"micro_cart/domain/repository"
)

type ICartService interface {
	AddCart(*model.Cart) (int64, error)
	UpdateCart(cart *model.Cart) error
	DeleteCartById(int64) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{CartRepository: cartRepository}
}

func (u *CartService) AddCart(cart *model.Cart) (cid int64, err error) {
	return u.CartRepository.CreateCart(cart)
}

func (u *CartService) DeleteCartById(cartId int64) error {
	return u.CartRepository.DeleteCartByID(cartId)
}

func (u *CartService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

func (u *CartService) FindCartByID(cartId int64) (cart *model.Cart, err error) {
	return u.CartRepository.FindCartByID(cartId)
}

func (u *CartService) FindAllCart(userId int64) (cartList []model.Cart, err error) {
	return u.CartRepository.FindAll(userId)
}

func (u *CartService) CleanCart(userId int64) error {
	return u.CartRepository.CleanCart(userId)
}
func (u *CartService) DecrNum(carId int64, num int64) error {
	return u.CartRepository.DecrNum(carId, num)
}

func (u *CartService) IncrNum(carId int64, num int64) error {
	return u.CartRepository.IncrNum(carId, num)
}
