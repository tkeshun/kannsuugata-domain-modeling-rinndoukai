package userhandlers

import (
	"encoding/json"
	"net/http"

	shareinfra "example-ch7_8/internal/share/infra"
	userpersistence "example-ch7_8/internal/user/persistence"
	userusecases "example-ch7_8/internal/user/usecases"
)

type RegisterUserHandler struct{}

func (h RegisterUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.req2user(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := userusecases.NewRegisterUserUsecase(
		userpersistence.NewIsEmailTaken(ctx),
	)

	events, err := usecase(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(events) == 0 {
		http.Error(w, "No events generated", http.StatusInternalServerError)
		return
	}
	for _, event := range events {
		shareinfra.PublishEvent(event)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h RegisterUserHandler) req2user(r *http.Request) (userusecases.RegisterUser, error) {
	var user userusecases.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return userusecases.RegisterUser{}, err
	}
	return user, nil
}
