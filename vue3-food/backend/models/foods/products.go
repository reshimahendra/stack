package foods

import (
	"time"

	"gorm.io/gorm"
)

/*
   Models Struct Objects
*/

// Product
type Product struct {
    gorm.Model
    ID          int             `json:"id"`
    Sku         string          `json:"sku" binding:"required"`
    Name        string          `json:"name" binding:"required"`
    CategoryID  int             `json:"category_id"`
    Category    ProductCategory `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    TypeID      int             `json:"type_id"`
    Type        ProductType     `json:"type" gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
    Price       int             `json:"price" binding:"required,number"`
    IsReady     bool            `json:"is_ready,default=true"`
    IsAvailable bool            `json:"is_available,default=true"`
    Picture     string          `json:"picture"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
}

type ProductRequest struct {
    Sku         string          `json:"sku" binding:"required"`
    Name        string          `json:"name" binding:"required"`
    CategoryID  int             `json:"category_id"`
    TypeID      int             `json:"type_id"`
    Price       int             `json:"price" binding:"required,number"`
    IsReady     bool            `json:"is_ready,default=true"`
    IsAvailable bool            `json:"is_available,default=true"`
    Picture     string          `json:"picture"`
}

type ProductReqUpdate struct {
    Sku         string          `json:"sku"`
    Name        string          `json:"name"`
    CategoryID  int             `json:"category_id"`
    TypeID      int             `json:"type_id"`
    Price       int             `json:"price" binding:"number"`
    IsReady     bool            `json:"is_ready"`
    IsAvailable bool            `json:"is_available"`
    Picture     string          `json:"picture"`
}

type ProductResponse struct {
    ID          int             `json:"id"`
    Sku         string          `json:"sku"`
    Name        string          `json:"name"`
    CategoryID  int             `json:"category_id"`
    TypeID      int             `json:"type_id"`
    Price       int             `json:"price"`
    IsReady     bool            `json:"is_ready"`
    IsAvailable bool            `json:"is_available"`
    Picture     string          `json:"picture"`
}


// Product category 
type ProductCategory struct {
    gorm.Model
    ID          int         `json:"id"`
    Code        string      `json:"code" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Description string      `json:"description"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

type ProductCategoryRequest struct {
    ID          int         `json:"id"`
    Code        string      `json:"code" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Description string      `json:"description"`
}


// Product type 
type ProductType struct {
    gorm.Model
    ID          int         `json:"id"`
    Code        string      `json:"code" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Description string      `json:"description"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

type ProductTypeRequest struct {
    Code        string      `json:"code" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Description string      `json:"description"`
}

type ProductCategoryResponse struct {
    ID          int         `json:"id"`
    Code        string      `json:"code" binding:"required"`
    Name        string      `json:"name" binding:"required"`
    Description string      `json:"description"`
}

type ProductTypeResponse struct {
    ID          int         `json:"id"`
    Code        string      `json:"code"`
    Name        string      `json:"name"`
    Description string      `json:"description"`
}


/*
    Converter to RESPONSE object 
*/
func productToResponse(p Product) (ProductResponse) {
    return ProductResponse{
        ID: p.ID,
        Sku: p.Sku,
        Name: p.Name,
        CategoryID: p.CategoryID,
        TypeID: p.TypeID,
        Price: p.Price,
        IsReady: p.IsReady,
        IsAvailable: p.IsAvailable,
        Picture: p.Picture,
    }
}

func prodCategoryToResponse(pc ProductCategory) (ProductCategoryResponse) {
    return ProductCategoryResponse{
        ID:          pc.ID,
        Code:        pc.Code,
        Name:        pc.Name,
        Description: pc.Description,
    }
}

func prodTypeToResponse(pt ProductType) (ProductTypeResponse) {
    return ProductTypeResponse{
        ID:          pt.ID,
        Code:        pt.Code,
        Name:        pt.Name,
        Description: pt.Description,
    }
}
