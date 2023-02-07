package api

import (
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Id string `json:"id" binding:"required"`
	Password string `json:"password" binding: "required, min=6"`
}

type createResponse struct {
	Id string
	CreatedAt time.Time
}

func (server *Server) create(ctx *gin.Context, db *DB) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	query := `INSERT INTO users (
		id,
		password,
	) VALUES ($1, $2) RETURNING *;`

	row := db.QueryRow(query, req.Id, req.Password)
	err = row.Scan(&req.Id, &req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)
} 