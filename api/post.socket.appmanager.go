package api

import (
	"net/http"
	"strings"
)

// PostAppTrafic godoc
// @Summary PostAppTrafic
// @Tags FORTEST!!!!
// @ID PostAppTrafic
// @Accept json
// @Produce json
// @Success 200 Role change
// @Failure 400 {object} string "You not admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/user/postapptrafic [post]

func isUniqueConstraintError(err error) bool {
	// Проверьте, является ли ошибка нарушением уникального ограничения
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func PostAppTrafic(w http.ResponseWriter, r *http.Request) {

}
