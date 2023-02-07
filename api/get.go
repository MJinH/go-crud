package api

import (
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type getRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) get(ctx *gin.Context, db *DB) {
	var req getRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	query := `SELECT * FROM users WHERE id = $1;`

	row := db.QueryRow(query, req.Id)
	err = row.Scan(&req.Id, &req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, row)
} 