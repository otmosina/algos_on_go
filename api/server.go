package api

import (
	"learn/algos_on_go/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()
	router.GET("/ping", server.Anagrams)
	server.router = router
	return server
}

type AnagramsRequestParams struct {
	Target string `form:"target" binding:"required,min=3"`
	Words  string `form:"words" binding:"required,min=3"`
}

func (s *Server) Anagrams(ctx *gin.Context) {
	var req AnagramsRequestParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	wordsFormatted := strings.Split(req.Words, ",")
	anagramsArr := lib.AnagramFilter(req.Target, wordsFormatted)

	ctx.JSON(http.StatusOK, anagramsArr)
}

func (s *Server) Start() error {
	return s.router.Run()
}

func ErrorResponse(e error) map[string]interface{} {
	return gin.H{
		"error": e.Error(),
	}
}
