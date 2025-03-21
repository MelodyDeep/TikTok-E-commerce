package service

import (
	"context"
	"testing"

	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/dal/redis"
	"github.com/MelodyDeep/TikTok-E-commerce/app/product/biz/model"
	product "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
)

var prdList = []model.Product{
	{StoreId: 1, Name: "apple", Description: "apple of test"},
	{StoreId: 1, Name: "banana", Description: "banana of test"},
	{StoreId: 1, Name: "juice", Description: "juice of test"},
}

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductService(ctx)
	// todo: edit your unit test
	dao := model.NewCacheDao(s.ctx, mysql.DB, redis.RedisClient)
	for _, prd := range prdList {
		id := dao.CreateProduct(&prd)
		if id == 0 {
			t.Errorf("fail to create test cases")
			continue
		}
		resp, err := s.Run(&product.DeleteProductReq{
			Id: id,
			StoreId: 1,
		})
		if err != nil || !resp.Success {
			t.Errorf("DeleteProduct(): %v", err)
			continue
		}
		_, err = dao.GetProductById(id)
		if err == nil  {
			t.Errorf("fail to delete product: %v", err)
		}
	}
	// test if evil operation can be stopped
	for _, prd := range prdList {
		id := dao.CreateProduct(&prd)
		if id == 0 {
			t.Errorf("fail to create test cases")
			continue
		}
		_, err := s.Run(&product.DeleteProductReq{
			Id: id,
			StoreId: 0,
		})
		if err == nil {
			t.Errorf("Fail to stop evil operation(store id not matched): %v", err)
			continue
		}
	}
}
