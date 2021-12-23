package foods

import (
	"time"
)

/*
   Models Struct Objects
*/

// Product
type ProductBase struct {
    Sku         string              `json:"sku"`
    Name        string              `json:"name"`
    CatID       int                 `json:"category_id"`
    Category    *ProductCategory    `json:"category" gorm:"foreignKey:CatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    TypeID      int                 `json:"type_id"`
    Type        *ProductType        `json:"type" gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
    Price       int                 `json:"price"`
    IsReady     bool                `json:"is_ready"`
    IsAvailable bool                `json:"is_available"`
    Picture     string              `json:"picture"`
}

type Product struct {
    ID          uint            `json:"id"`
    Prod        ProductBase     `json:"prod" gorm:"embedded;embeddedPrefix:"`
    UpdatedAt   time.Time       `json:"updated_at" gorm:"<-:update"`
    CreatedAt   time.Time       `json:"created_at" gorm:"<-:create"`
    DeletedAt   time.Time       `json:"deleted_at" gorm:"default:NULL"`
}

type ProductResponse struct {
    ID          uint            `json:"id"`
    Prod        ProductBase     `json:"prod" gorm:"embedded"`
}

// Product category 
type ProductCategory struct {
    ID          int             `json:"id" gorm:"primaryKey;<-:create;autoIncrement"`
    ProdCat     ProdCatBase     `json:"cat" gorm:"embedded"`
}

type ProdCatBase struct {
    Name        string          `json:"name" binding:"required" gorm:"unique;type char(75)"`
    Description string          `json:"description"`
}

// Product type 
type ProductType struct {
    ID          int             `json:"id" gorm:"primaryKey;<-:create;autoIncrement;"`
    ProdType    ProdTypeBase    `json:"type" gorm:"embedded"` 
}

type ProdTypeBase struct {
    Name        string          `json:"name" binding:"required" gorm:"unique;type char(75)"`
    Description string          `json:"description"`
}


