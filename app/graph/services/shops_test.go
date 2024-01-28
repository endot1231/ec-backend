package services

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/graph/model"
	"github.com/endot1231/ec-backend/test_utils"
	"github.com/stretchr/testify/suite"
)

type TestShopStruct struct {
	suite.Suite
	db      *ent.Client
	Clocker clock.Clocker

	shops        []*ent.Shops
	shop_name    string
	shop_address string

	shop1_name    string
	shop1_address string
}

func TestShopStructSuite(t *testing.T) {
	suite.Run(t, new(TestShopStruct))
}

func (s *TestShopStruct) SetupTest() {
	s.db = test_utils.NewTestDbClient(s.T())
	s.Clocker = clock.FixedClocker{}

	s.shop_name = "admin"
	s.shop_address = "example"

	s.shop1_name = "admin"
	s.shop1_address = "example1"

	ctx := context.Background()
	shop, _ := s.db.Shops.Create().SetName(s.shop_name).SetAddress(s.shop_address).Save(ctx)
	shop1, _ := s.db.Shops.Create().SetName(s.shop1_name).SetAddress(s.shop1_address).Save(ctx)
	s.shops = append(s.shops, shop, shop1)
}

func (s *TestShopStruct) Test_convertShop() {
	want := &model.Shop{ID: strconv.Itoa(s.shops[0].ID), Name: *s.shops[0].Name, Address: *s.shops[0].Address}

	type args struct {
		shop *ent.Shops
	}
	tests := []struct {
		name string
		args args
		want *model.Shop
	}{
		{
			name: "return shop model",
			args: args{shop: s.shops[0]},
			want: want,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := convertShop(tt.args.shop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertShop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestShopStruct) Test_convertShops() {

	var want []*model.Shop
	shop := &model.Shop{ID: strconv.Itoa(s.shops[0].ID), Name: *s.shops[0].Name, Address: *s.shops[0].Address}
	shop1 := &model.Shop{ID: strconv.Itoa(s.shops[1].ID), Name: *s.shops[1].Name, Address: *s.shops[1].Address}
	want = append(want, shop, shop1)

	type args struct {
		shops []*ent.Shops
	}
	tests := []struct {
		name string
		args args
		want []*model.Shop
	}{
		{
			name: "return shop model array",
			args: args{shops: s.shops},
			want: want,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := convertShops(tt.args.shops); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertShops() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestShopStruct) Test_shopService_GetShops() {
	type fields struct {
		exec    ent.Client
		Clocker clock.Clocker
	}
	type args struct {
		ctx     context.Context
		name    string
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Shop
		wantErr bool
	}{}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s := &shopService{
				exec:    tt.fields.exec,
				Clocker: tt.fields.Clocker,
			}
			got, err := s.GetShops(tt.args.ctx, tt.args.name, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("shopService.GetShops() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shopService.GetShops() = %v, want %v", got, tt.want)
			}
		})
	}
}
