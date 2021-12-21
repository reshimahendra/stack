package cart

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cart Handler
type cartHandler struct {
    service CartService
}

func NewCartHandler(service CartService) *cartHandler {
    return &cartHandler{service}
}

func (ch *cartHandler) CreateCartHandler(c *gin.Context) {
    var cart Cart

    err := c.ShouldBindJSON(&cart).Error
    log.Println(err)
    if err != nil {
        // c.JSON(http.StatusBadRequest, gin.H{
        //     "error": "Something wrong while retreiving data.",
        // })
        panic(err)
        // return
    }

    cr, e := ch.service.Create(cart)
    if e != nil {
        c.JSON(http.StatusOK, gin.H{
            "error": "Something wrong while processing data.",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": cr,
    })
}
