// controllers/index_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexController handles actions for the index page
type IndexController struct{}

// NewIndexController creates a new instance of IndexController
func NewIndexController() *IndexController {
	return &IndexController{}
}

// Index displays the index page
func (ic *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
