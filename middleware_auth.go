package main

import (
	"fmt"
	"net/http"

	"github.com/scaranaraa/rss_aggregator/internal/auth"
	"github.com/scaranaraa/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		resopondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		resopondWithError(w, 400, fmt.Sprintf("Couldnt get user: %v", err))
		return
	}

	handler(w,r, user)
	}
}