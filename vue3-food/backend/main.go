package main

import (
	"fmt"
	"lbw-resto/middleware"
	"lbw-resto/models/foods"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main()  {
    db, err := gorm.Open(sqlite.Open("resto.db"), &gorm.Config{})
    if err != nil {
        panic("Fail to connect to Database.")
    }
    // automigrate models 
    db.AutoMigrate(
        &foods.ProductType{},
        &foods.ProductCategory{},
        &foods.Product{},
    )

    // product Type
    prodTypeRepo    := foods.NewProductTypeRepo(db)
    prodTypeService := foods.NewProductTypeService(prodTypeRepo)
    prodTypeHandler := foods.NewProductTypeHandler(prodTypeService)

    // product category 
    prodCatRepo    := foods.NewProductCategoryRepo(db)
    prodCatService := foods.NewProductCategoryService(prodCatRepo)
    prodCatHandler := foods.NewProductCategoryHandler(prodCatService)

    // product 
    prodRepo    := foods.NewProductRepo(db)
    prodService := foods.NewProductService(prodRepo)
    prodHandler := foods.NewProductHandler(prodService)

    // gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.CORSMiddleware())
    r.SetTrustedProxies([]string{
        "192.168.43.26",
    })

    // GROUP for the 'API Version.1' routes
    v1 := r.Group("/v1")
    
    // GROUP for PRODUCT TYPE
    prodType := v1.Group("/product-type")
    prodType.GET("/", prodTypeHandler.ShowProductTypeHandler)
    prodType.POST("/create", prodTypeHandler.CreateProductTypeHandler)
    prodType.PUT("/:id/update", prodTypeHandler.UpdateProductTypeHandler)
    prodType.GET("/first", prodTypeHandler.FirstProductTypeHandler)
    prodType.GET("/last", prodTypeHandler.LastProductTypeHandler)
    prodType.GET("/:id", prodTypeHandler.FindByIDProductTypeHandler)

    // GROUP for PRODUCT CATEGORY 
    prodCategory := v1.Group("/product-category")
    prodCategory.GET("/", prodCatHandler.ShowProductCategoryHandler)
    prodCategory.POST("/create", prodCatHandler.CreateProductCategoryHandler)
    prodCategory.PUT("/:id/update", prodCatHandler.UpdateProductCategoryHandler)
    prodCategory.GET("/first", prodCatHandler.FirstProductCategoryHandler)
    prodCategory.GET("/last", prodCatHandler.LastProductCategoryHandler)
    prodCategory.GET("/:id", prodCatHandler.FindByIDProductCategoryHandler)

    // GROUP for PRODUCT CATEGORY 
    prod := v1.Group("/product")
    prod.GET("/", prodHandler.ShowProductHandler)
    prod.POST("/create", prodHandler.CreateProductHandler)
    prod.PUT("/:id/update", prodHandler.UpdateProductHandler)
    prod.GET("/first", prodHandler.FirstProductHandler)
    prod.GET("/last", prodHandler.LastProductHandler)
    prod.GET("/:id", prodHandler.FindByIDProductHandler)

    fmt.Println("Starting server at port '8000'") 
    r.Run(":8000")
}
