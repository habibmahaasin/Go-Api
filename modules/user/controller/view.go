package controller

import (
	"gop-api/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pagesView struct {
	service service.UserService
}

func NewPagesView(service service.UserService) *pagesView {
	return &pagesView{service}
}

func (h *pagesView) Indexview(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"page": "Index Page",
	})
}

func (h *pagesView) Others(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"page": "Others",
	})
}
