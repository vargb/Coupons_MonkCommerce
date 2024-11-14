package services

import (
	"monkCommerce/config"
	"monkCommerce/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Server *gin.Engine
}

func (repo *Repository) SetupRoutes() error {
	repo.Server.POST("/products", repo.CreateProduct)
	repo.Server.GET("/products", repo.GetAllProducts)
	repo.Server.POST("/coupons", repo.CreateCoupon)
	repo.Server.GET("/coupons", repo.GetAllCoupons)
	repo.Server.GET("/coupons/:id", repo.GetCoupon)
	repo.Server.PUT("/coupons/:id", repo.UpdateCoupon)
	repo.Server.POST("/applicable-coupons", repo.GetApplicableCoupons)
	repo.Server.POST("/apply-coupon/:id", repo.ApplyCoupon)
	return nil
}

func Initialize(conf *config.Config) (*gin.Engine, error) {
	loggy := config.GetLogger()
	db, err := storage.NewDBConn(conf)
	if err != nil {
		loggy.Error("error in db initializing")
		return nil, err
	}
	err = storage.MigrateCoupons(db)
	if err != nil {
		loggy.Error("error in db initializing")
		return nil, err
	}
	repo := &Repository{
		DB:     db,
		Server: gin.Default(),
	}
	repo.SetupRoutes()
	return repo.Server, nil
}
