package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, Alert := r.Cookie("session_token")
		
		if (Alert != nil){
			w.WriteHeader(401)
			w.WriteHeader(http.StatusUnauthorized)
			Alert := model.ErrorResponse{
				Error: "http: named cookie not present",
			}
			dataAlert, _ := json.Marshal(Alert)
			w.Write(dataAlert)
			return
		}else{
		}// TODO: replace this

		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionToken.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if (r.Method != http.MethodGet){
			w.WriteHeader(405)
			w.WriteHeader(http.StatusMethodNotAllowed)
			Alert := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			dataAlert, _ := json.Marshal(Alert)
			w.Write(dataAlert)
			return
		}else{

		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method != http.MethodPost){
			w.WriteHeader(405)
			w.WriteHeader(http.StatusMethodNotAllowed)
			Alert := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			dataAlert, _ := json.Marshal(Alert)
			w.Write(dataAlert)
			return
		}else{
			
		}
		next.ServeHTTP(w, r)
	})
}
