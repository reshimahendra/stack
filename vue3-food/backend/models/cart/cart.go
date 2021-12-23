package cart

import "lbw-resto/models/foods"

type Cart struct {
    ID          uint            `json:"id" gorm:"primaryKey;autoInc"`
    TableNo     uint            `json:"table_no" binding:"required"`
    CartID      uint            `json:"cart_id"`
    CartItems   []CartItem      `json:"cart_items" gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CartItem struct {
    CartId      uint            `json:"cart_id" gorm:"not null"`
    Quantity    uint            `json:"quantity,default=1" binding:"required, number"`
    Note        string          `json:"note"`
    ProductID   uint            `json:"product_id"`
    Product     foods.Product   `json:"products" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelet:SET NULL;"` 
}
