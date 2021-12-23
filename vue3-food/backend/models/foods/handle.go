package foods

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
   PRODUCT TYPE HANDLES
*/
type productTypeHandler struct {
    productTypeService ProductTypeService
}

func NewProductTypeHandler(productTypeService ProductTypeService) *productTypeHandler {
    return &productTypeHandler{productTypeService}
}

// create handle 
func (h *productTypeHandler) CreateProductTypeHandler(c *gin.Context) {
    var pts ProdTypeBase

    err := c.ShouldBindJSON(&pts)
    if err != nil {
        fmt.Println(err)
        msg := []string{}
        for _, e := range err.(validator.ValidationErrors) {
            aMsg := fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag())
            msg = append(msg, aMsg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })

        return
    }

    pt, err := h.productTypeService.Create(pts)
    fmt.Println(err)
    if err != nil {
        eMsg := "Error occur while creating 'product type' record."
        fmt.Println(eMsg)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    // show data to user 
    c.JSON(http.StatusOK, gin.H{
        "data": pt,
    })
}

// update handle
func (h *productTypeHandler) UpdateProductTypeHandler(c *gin.Context) {
   var ptr ProdTypeBase

   // Check if required record available
   ptID, err := strconv.Atoi(c.Param("id"))
   if err != nil || ptID < 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Error occur while checking the required params!",
       })
       return
   }

   // Validate the input request 
   err = c.ShouldBindJSON(&ptr)
   if err != nil {
       msg := []string{}
       for _, e := range err.(validator.ValidationErrors) {
           msg = append(msg, fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag()))
       }
       c.JSON(http.StatusBadRequest, gin.H{
           "error": msg,
       })
   }

   // Update the data based on the 'ProductTypeID/ ptID' found on 'ProductTypeRequest/ ptr' struct 
   pt, err := h.productTypeService.Update(ptID, ptr)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Update operation aborted. Something wrong with the data.",
       })
       return
   }

   
   // Show updated data 
   pt.ID = ptID
   c.JSON(http.StatusOK, gin.H{
       "data": pt,
   })
}

func (h *productTypeHandler) DeleteProductTypeHandler(c *gin.Context) {
   // Check if the apropriate record available
   ptID, err := strconv.Atoi(c.Param("id"))
   if err != nil || ptID < 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Error occur while checking the required params",
       })
       return
   }

   // Check the data that want to be deleted. 
   _, err = h.productTypeService.FindByID(ptID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying to find the data!",
       })
       return
   }

   // Delete the data 
   err = h.productTypeService.Delete(ptID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying deleting the data!",
       })
       return
   }

   // If all good, response to the user request 
   c.JSON(http.StatusOK, gin.H{
       "info": fmt.Sprintf("Data with id: '%d' has been deleted.", ptID),
   })
}

// First record handler 
func (h *productTypeHandler) FirstProductTypeHandler(c *gin.Context) {
    // Looking for the first data on the table
    pt, err := h.productTypeService.First()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // show the data to the client
    c.JSON(http.StatusOK, gin.H{
        "data": pt, 
    })
}

// Last record handler 
func (h *productTypeHandler) LastProductTypeHandler(c *gin.Context) {
    // looking for the last data on the table 
    pt, err := h.productTypeService.Last()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to the client 
    c.JSON(http.StatusOK, gin.H{
        "data" : pt,
    })
}

// Find Product type by ID
func (h *productTypeHandler) FindByIDProductTypeHandler(c *gin.Context) {
    // check the parameter 
    ptID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Println("Error converting string to integer for the 'url param'.")
        return
    }
    pt, err := h.productTypeService.FindByID(ptID)

    // Looking for the data based on the ID
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to client 
    c.JSON(http.StatusOK, gin.H{
        "data": pt,
    })
}

// Show all data 
func (h *productTypeHandler) ShowProductTypeHandler(c *gin.Context) {
    // Get all product type data 
    allPt, err := h.productTypeService.All()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while trying to get the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": allPt,
    })
}


/*
   PRODUCT CATEGORY HANDLES
*/
type productCategoryHandler struct {
    productCategoryService ProductCategoryService
}

func NewProductCategoryHandler(productCategoryService ProductCategoryService) *productCategoryHandler{
    return &productCategoryHandler{productCategoryService}
}

// create handle 
func (h *productCategoryHandler) CreateProductCategoryHandler(c *gin.Context) {
    var pcr ProdCatBase 

    err := c.ShouldBindJSON(&pcr)
    if err != nil {
        fmt.Println(err)
        msg := []string{}
        for _, e := range err.(validator.ValidationErrors) {
            aMsg := fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag())
            msg = append(msg, aMsg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })

        return
    }

    pc, err := h.productCategoryService.Create(pcr)
    if err != nil {
        eMsg := "Error occur while creating 'product category' record."
        fmt.Println(eMsg)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": pc,
    })
}

// update handle
func (h *productCategoryHandler) UpdateProductCategoryHandler(c *gin.Context) {
   var pcr ProdCatBase 

   // Check if required record available
   pcID, err := strconv.Atoi(c.Param("id"))
   if err != nil || pcID < 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Error occur while checking the required params!",
       })
       return
   }

   // Validate the input request 
   err = c.ShouldBindJSON(&pcr)
   if err != nil {
       msg := []string{}
       for _, e := range err.(validator.ValidationErrors) {
           msg = append(msg, fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag()))
       }
       c.JSON(http.StatusBadRequest, gin.H{
           "error": msg,
       })
   }

   // Update the data based on the 'ProductCategoryID/ pcID' found on 'ProductCategoryRequest/ pcr' struct 
   pc, err := h.productCategoryService.Update(pcID, pcr)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Update operation aborted. Something wrong with the data.",
       })
       return
   }

   // Show updated data 
   pc.ID = pcID
   c.JSON(http.StatusOK, gin.H{
       "data": pc,
   })
}

func (h *productCategoryHandler) DeleteProductCategoryHandler(c *gin.Context) {
   // Check if the apropriate record available
   pcID, err := strconv.Atoi(c.Param("id"))
   if err != nil || pcID < 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Error occur while checking the required params",
       })
       return
   }

   // Check the data that want to be deleted. 
   _, err = h.productCategoryService.FindByID(pcID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying to find the data!",
       })
       return
   }

   // Delete the data 
   err = h.productCategoryService.Delete(pcID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying deleting the data!",
       })
       return
   }

   // If all good, response to the user request 
   c.JSON(http.StatusOK, gin.H{
       "info": fmt.Sprintf("Data with id: '%d' has been deleted.", pcID),
   })
}

// First record handler 
func (h *productCategoryHandler) FirstProductCategoryHandler(c *gin.Context) {
    // Looking for the first data on the table
    pc, err := h.productCategoryService.First()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": pc, 
    })
}

// Last record handler 
func (h *productCategoryHandler) LastProductCategoryHandler(c *gin.Context) {
    // looking for the last data on the table 
    pc, err := h.productCategoryService.Last()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to the client 
    c.JSON(http.StatusOK, gin.H{
        "data" : pc,
    })
}

// Find By ID 
func (h *productCategoryHandler) FindByIDProductCategoryHandler(c *gin.Context) {
    // check the parameter 
    pcID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Println("Error converting string to integer for the 'url param'.")
        return
    }
    pc, err := h.productCategoryService.FindByID(pcID)

    // Looking for the data based on the ID
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to client 
    c.JSON(http.StatusOK, gin.H{
        "data": pc,
    })
}

// Show all data 
func (h *productCategoryHandler) ShowProductCategoryHandler(c *gin.Context) {
    // Get all product category data 
    allPc, err := h.productCategoryService.All()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while trying to get the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": allPc,
    })
}


/*
   PRODUCT HANDLES
*/
type productHandler struct {
    productService ProductService
}

func NewProductHandler(productService ProductService) *productHandler{
    return &productHandler{productService}
}

// create handle 
func (h *productHandler) CreateProductHandler(c *gin.Context) {
    var ProdReq ProductBase

    err := c.ShouldBindJSON(&ProdReq)
    fmt.Println("Product handler create 1:\n", ProdReq)
    if err != nil {
        log.Println(err)
        msg := []string{}
        for _, e := range err.(validator.ValidationErrors) {
            aMsg := fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag())
            msg = append(msg, aMsg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })

        return
    }

    Prod, err := h.productService.Create(ProdReq)
    fmt.Println("Product handler create 2:\n", Prod)
    if err != nil {
        eMsg := "Error occur while creating 'product' record."
        fmt.Println(eMsg)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": Prod,
    })
}

// update handle
func (h *productHandler) UpdateProductHandler(c *gin.Context) {
   var prodReq ProductBase

   prodID, _ := strconv.Atoi(c.Param("id"))

   // Validate the input request 
   err := c.ShouldBindJSON(&prodReq)
   fmt.Println("handler update 1:\n",prodReq)
   if err != nil {
       msg := []string{}
       for _, e := range err.(validator.ValidationErrors) {
           msg = append(msg, fmt.Sprintf("Field error: '%s', condition: '%s'.", e.Field(), e.ActualTag()))
       }
       c.JSON(http.StatusBadRequest, gin.H{
           "error": msg,
       })
   }

   // Update the data based on the 'ProductID/ ProdID' found on 'ProductRequest/ ProdReq' struct 
   prod, err := h.productService.Update(prodID, prodReq)
   fmt.Println("handler update 2:\n",prod)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Update operation aborted. Something wrong with the data.",
       })
       return
   }

   // Update the record data 
   prod.ID = uint(prodID)

   // Show updated data 
   c.JSON(http.StatusOK, gin.H{
       "data": prod,
   })
}

func (h *productHandler) DeleteProductHandler(c *gin.Context) {
   // Check if the apropriate record available
   ProdID, err := strconv.Atoi(c.Param("id"))
   if err != nil || ProdID < 0 {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Error occur while checking the required params",
       })
       return
   }

   // Check the data that want to be deleted. 
   _, err = h.productService.FindByID(ProdID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying to find the data!",
       })
       return
   }

   // Delete the data 
   err = h.productService.Delete(ProdID)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
           "error": "Something wrong while trying deleting the data!",
       })
       return
   }

   // If all good, response to the user request 
   c.JSON(http.StatusOK, gin.H{
       "info": fmt.Sprintf("Data with id: '%d' has been deleted.", ProdID),
   })
}

// First record handler 
func (h *productHandler) FirstProductHandler(c *gin.Context) {
    // Looking for the first data on the table
    Prod, err := h.productService.First()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // show the data to the client
    c.JSON(http.StatusOK, gin.H{
        "data": Prod, 
    })
}

// Last record handler 
func (h *productHandler) LastProductHandler(c *gin.Context) {
    // looking for the last data on the table 
    Prod, err := h.productService.Last()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to the client 
    c.JSON(http.StatusOK, gin.H{
        "data" : Prod,
    })
}

// Find By ID 
func (h *productHandler) FindByIDProductHandler(c *gin.Context) {
    // check the parameter 
    ProdID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Println("Error converting string to integer for the 'url param'.")
        return
    }
    Prod, err := h.productService.FindByID(ProdID)

    // Looking for the data based on the ID
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to client 
    c.JSON(http.StatusOK, gin.H{
        "data": Prod,
    })
}

// Show all data 
func (h *productHandler) ShowProductHandler(c *gin.Context) {
    // Get all product category data 
    allProd, err := h.productService.All()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while trying to get the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": allProd,
    })
}

// Show limited list data 
func (h *productHandler) LimitProductHandler(c *gin.Context) {
    // Get all product data 
    maxRec, _ := strconv.Atoi(c.Param("limit"))
    allProd, err := h.productService.Limit(maxRec)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while trying to get the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": allProd,
    })
}

// Show offseted data 
func (h *productHandler) OffsetProductHandler(c *gin.Context) {
    // Get all product Data 
    offset, _ := strconv.Atoi(c.Param("offset"))
    allProd, err := h.productService.Offset(offset)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while trying to get the data!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": allProd,
    })
}

// Find Name 
func (h *productHandler) FindNameProductHandler(c *gin.Context) {
    // Looking for data with 'LIKE' clausal on 'name' field 
    filterProd, err := h.productService.FindName(c.Param("name"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Something wrong while looking up the data!",
        })
        return
    }

    // Show data to client 
    c.JSON(http.StatusOK, gin.H{
        "data": filterProd,
    })
}
