package cart

import (
	"lbw-resto/models/foods"

	"gorm.io/gorm"
)

type Cart struct {
    gorm.Model
    // ID          int         `json:"id"`
    TableNo     uint         `json:"table_no" binding:"required"`
    OrderList   []OrderList  `json:"order_list" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OrderList struct {
    gorm.Model
    Quantity    int             `json:"quantity,default=1" binding:"required"`
    Note        string          `json:"note"`
    Product     foods.Product   `json:"products,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelet:SET NULL;"` 
    CartID      uint            `json:"id"`
}
