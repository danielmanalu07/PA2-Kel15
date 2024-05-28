package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id             uint           `json:"id"`
	Code           string         `json:"code" gorm:"type:varchar(50);idxcd" validate:"required"`
	CustomerID     uint           `json:"customer_id" gorm:"idxcs"`
	Customer       Customer       `json:"-" gorm:"foreignKey:customer_id"`
	Products       []Product      `json:"products" gorm:"many2many:order_products;"`
	Total          string         `json:"total" gorm:"type:varchar(50)"`
	Note           string         `json:"note" gorm:"type:text"`
	PaymentMethod  string         `json:"payment_method" gorm:"type:varchar(50)"`
	TableId        *uint          `json:"table_id" gorm:"idxtb"`
	Table          Table          `json:"-" gorm:"foreignKey:TableId"`
	PickUpType     string         `json:"pick_up_type" gorm:"type:varchar(50)" validate:"required"`
	ProofOfPayment string         `json:"proof_of_payment" gorm:"type:varchar(50)" validate:"required"`
	Status         int            `json:"status" gorm:"idxst"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type OrderProduct struct {
	OrderID   uint    `json:"order_id" gorm:"primaryKey"`
	Order     Order   `json:"-" gorm:"foreignKey:order_id"`
	ProductID uint    `json:"product_id" gorm:"primaryKey"`
	Product   Product `json:"-" gorm:"foreignKey:product_id"`
}
