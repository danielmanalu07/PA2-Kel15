package response

import (
	"api/the_deck/Models/entity"
	"time"

	"gorm.io/gorm"
)

type OrderResponse struct {
	Id             uint             `json:"id"`
	Code           string           `json:"code" gorm:"type:varchar(50);idxcd" validate:"required"`
	CustomerID     uint             `json:"customer_id" gorm:"idxcs"`
	Products       []entity.Product `json:"products" gorm:"many2many:order_products;"`
	Total          string           `json:"total" gorm:"type:varchar(50)"`
	Note           string           `json:"note" gorm:"type:text"`
	PaymentMethod  string           `json:"payment_method" gorm:"type:varchar(50)"`
	TableId        *uint            `json:"table_id" gorm:"idxtb"`
	PickUpType     string           `json:"pick_up_type" gorm:"type:varchar(50)" validate:"required"`
	ProofOfPayment string           `json:"proof_of_payment" gorm:"type:varchar(50)" validate:"required"`
	Status         int              `json:"status" gorm:"idxst"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt      gorm.DeletedAt   `gorm:"index"`
}
