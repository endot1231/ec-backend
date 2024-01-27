package services

import (
	"context"
	"strconv"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/shops"
	"github.com/endot1231/ec-backend/graph/model"
)

type shopService struct {
	exec    ent.Client
	Clocker clock.Clocker
}

func convertShop(shop *ent.Shops) *model.Shop {
	return &model.Shop{
		ID:      strconv.Itoa(shop.ID),
		Name:    *shop.Name,
		Address: *shop.Address,
	}
}

func convertShops(shops []*ent.Shops) []*model.Shop {
	var shoprArray []*model.Shop
	for _, shop := range shops {
		shoprArray = append(shoprArray, convertShop(shop))
	}
	return shoprArray
}

func (s *shopService) GetShops(ctx context.Context, name string, address string) ([]*model.Shop, error) {

	query := s.exec.Shops.Query()

	if name != "" {
		query = query.Where(shops.NameHasPrefix(name))
	}

	if address != "" {
		query = query.Where(shops.AddressHasPrefix(address))
	}

	shops, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return convertShops(shops), nil
}
