package handler

import (
	"context"
	"log"

	"micro_cart/common"
	"micro_cart/domain/model"
	"micro_cart/domain/service"
	"micro_cart/proto/cart"
	proto "micro_cart/proto/cart"
)

type Cart struct {
	CartService service.ICartService
}

// Return a new handler
// func New() *Cart {
// 	return &Cart{}
// }

func (e *Cart) AddCart(ctx context.Context, req *proto.CartInfo, rsp *proto.CartResponse) error {
	cart := &model.Cart{}
	err := common.SwapTo(req, cart)
	if err != nil {
		return err
	}
	cartId, err := e.CartService.AddCart(cart)
	if err != nil {
		return err
	}
	rsp.Id = cartId
	rsp.Msg = "添加成功"
	return nil
}

func (e *Cart) CleanCart(ctx context.Context, req *proto.CleanInfo, rsp *proto.Response) error {
	if err := e.CartService.CleanCart(req.Uid); err != nil {
		return err
	}
	rsp.Msg = "购物车清空成功"
	return nil
}

func (e *Cart) IncrItem(ctx context.Context, req *proto.ItemInfo, rsp *proto.Response) error {
	if err := e.CartService.IncrNum(req.Id, req.Num); err != nil {
		return err
	}
	rsp.Msg = "购物车增加成功"
	return nil
}

func (e *Cart) DecrItem(ctx context.Context, req *proto.ItemInfo, rsp *proto.Response) error {
	if err := e.CartService.DecrNum(req.Id, req.Num); err != nil {
		return err
	}
	rsp.Msg = "购物车减少成功"
	return nil
}

func (e *Cart) DeleteItemById(ctx context.Context, req *proto.CartIdRequest, rsp *proto.Response) error {
	err := e.CartService.DeleteCartById(req.Id)
	if err != nil {
		return err
	}
	rsp.Msg = "购物车删除成功"
	return nil
}

func (e *Cart) FindCartById(ctx context.Context, req *proto.CartIdRequest, rsp *proto.CartInfo) error {
	cart, err := e.CartService.FindCartByID(req.Id)
	if err != nil {
		return err
	}
	return common.SwapTo(cart, rsp)
}

func (e *Cart) GetAllCart(ctx context.Context, req *proto.FindAllRequest, rsp *proto.FindAllResponse) error {
	cartSlice, err := e.CartService.FindAllCart(req.Uid)
	if err != nil {
		return err
	}
	cartToResponse(cartSlice, rsp)
	return nil
}

func cartToResponse(cartSlice []model.Cart, rsp *cart.FindAllResponse) {
	for _, item := range cartSlice {
		cart := &cart.CartInfo{}
		err := common.SwapTo(item, cart)
		if err != nil {
			log.Fatal(err)
			break
		}
		rsp.CartInfo = append(rsp.CartInfo, cart)
	}
}
