package storage

import (
	"fmt"
	"monkCommerce/config"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	Name    string         `json:"name" gorm:"not null"`
	Code    string         `json:"code" gorm:"uniqueIndex;not null"`
	Price   float64        `json:"price" gorm:"not null"`
	Coupons pq.StringArray `json:"coupons" gorm:"type:text[]"`
}

type Coupon struct {
	Id         string `json:"id" gorm:"uniqueIndex;not null"`
	Code       string `json:"code"`
	Type       string `json:"type"`
	ValidFrom  string `json:"valid_from"`
	ValidUntil string `json:"valid_until"`
	UsageLimit int    `json:"usage_limit"`
	UsageCount int    `json:"usage_count"`
	IsActive   bool   `json:"is_active"`

	// Type-specific rules (only one will be used based on Type)
	CartWiseRule    *CartWiseRule    `json:"cart_wise_rule,omitempty" gorm:"type:jsonb;"`
	ProductWiseRule *ProductWiseRule `json:"product_wise_rule,omitempty" gorm:"type:jsonb;"`
	BxGyRule        *BxGyRule        `json:"bxgy_rule,omitempty" gorm:"type:jsonb;"`
}

type CartWiseRule struct {
	MinimumAmount float64 `json:"minimum_amount"`
	Discount      float64 `json:"discount"`
	IsPercentage  bool    `json:"is_percentage"`
}

type ProductWiseRule struct {
	Discount     float64   `json:"discount"`
	Products     []Product `json:"products"`
	IsPercentage bool      `json:"is_percentage"`
}

type BxGyProduct struct {
	Quantity int       `json:"quantity"`
	Products []Product `json:"products"`
}

type BxGyRule struct {
	BuyProducts     []BxGyProduct `json:"buy_products"`
	GetProducts     []BxGyProduct `json:"get_products"`
	RepetitionLimit int           `json:"repetition_limit"`
}

type CartItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}

func MigrateCoupons(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Coupon{})
	return err
}

func NewDBConn(conf *config.Config) (*gorm.DB, error) {
	loggy := config.GetLogger()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		conf.Postgres.Host, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Dbname, conf.Postgres.Port, conf.Postgres.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		loggy.Error("error in connecting to db")
		return nil, err
	}
	return db, nil
}
