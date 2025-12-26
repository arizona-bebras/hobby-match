package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"

	"shumi/internal/database"
)

type PagesHandler struct {
	DB *gorm.DB
}

// GetPage
// @Summary Получить анкету(страницу) одного пользователя
// @Produce json
// @Param page_id path string true "id анкеты"
// @Success 200 {object} handlers.PageData
// @Failure 500 {object} database.Error "Внутренняя ошибка сервера"
// @Router /api/pages/{page_id} [get]
func (h *PagesHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	selfId := r.Context().Value(database.AuthContextKey).(string)
	otherId := r.PathValue("page_id")

	var user PageData

	err := h.DB.Preload("Widgets").Preload("TgUser").First(&user.User, "id = ?", otherId).Error
	if err != nil {
		log.Printf("pages handler: failed to get user, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("pages handler: failed to get user, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}
	user.Username = user.TgUser.TgUsername

	err = h.DB.Raw(`
		WITH user_vectors AS (
		  SELECT u.id, info_embedding, personality_test, AVG(i.embedding) as avg_interest_embedding FROM users AS u
		  LEFT JOIN user_interests ui ON u.id = ui.user_id
		  LEFT JOIN interests i ON i.id = ui.interest_id
		  GROUP BY u.id
		), self_user AS (
		  SELECT * FROM user_vectors WHERE id = @self
		), other_user AS (
		  SELECT * FROM user_vectors WHERE id = @other
		), scores AS (
		  SELECT 
			coalesce(1 - ((SELECT info_embedding FROM self_user) <=> (SELECT info_embedding FROM other_user)), 0) AS info_score,
			coalesce(1 - ((SELECT avg_interest_embedding FROM self_user) <=> (SELECT avg_interest_embedding FROM other_user)), 0) AS interests_score,
			coalesce(1 - (((SELECT personality_test FROM self_user) <-> (SELECT personality_test FROM other_user)) / SQRT(500)), 0) AS personality_score
		)
		SELECT info_score, interests_score, personality_score,
			(0.1 * info_score + 0.5 * interests_score + 0.4 * personality_score) AS similarity 
		FROM scores;
	`, sql.Named("self", selfId), sql.Named("other", otherId)).Scan(&user.SimilarityData).Error
	if err != nil {
		log.Printf("pages handler: failed to calculate similarity, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("pages handler: failed to calculate similarity, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	err = h.DB.Raw(`
		WITH self_user_interests AS (
		  SELECT i.id, i.embedding FROM user_interests AS ui
		  LEFT JOIN interests i ON i.id = ui.interest_id
		  WHERE ui.user_id = @self
		), 
		self_user AS (
		  SELECT u.id, AVG(i.embedding) as avg_interest_embedding FROM users AS u
		  LEFT JOIN self_user_interests i ON TRUE
		  WHERE u.id = @self
		  GROUP BY u.id
		),
		q_interests AS (
		  SELECT i.id, i.tag, MAX(1 - (i.embedding <=> sui.embedding)) AS similarity
		  FROM user_interests AS ui 
		  LEFT JOIN interests i ON i.id = ui.interest_id
		  LEFT JOIN self_user_interests sui ON TRUE
		  WHERE ui.user_id = @other
		  GROUP BY i.id
		  ORDER BY similarity DESC
		)
		SELECT * FROM q_interests;
	`, sql.Named("self", selfId), sql.Named("other", otherId)).Scan(&user.Interests).Error
	if err != nil {
		log.Printf("pages handler: failed to calculate interest similarity, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("pages handler: failed to calculate interest similarity, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("pages handler: failed to marshal user data, %v", err)
		http.Error(
			w,
			database.JSONErr(
				http.StatusInternalServerError,
				fmt.Sprintf("pages handler: failed to marshal user data, %v", err),
			),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(userJSON)
}

func (h PagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetPage(w, r)
	}
}
