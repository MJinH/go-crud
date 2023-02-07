package api

import (
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) delete(ctx *gin.Context, db *DB) {
	var req deleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	query := `DELETE * FROM users WHERE id = $1;`

	result, err := db.Exec(query, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, result)
} 