package api

import (
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Id string `uri:"id" binding:"required"`
	Password string `json:"password" binding: "required, min=6"`
}

func (server *Server) update(ctx *gin.Context, db *DB) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	query := `UPDATE users SET password = $1;`

	result, err := db.Exec(query, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, result)
} 