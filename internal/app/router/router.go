package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/evseevbl/userapi/internal/app/userapi"
)

type userAPI interface {
	Register(ctx context.Context, req *userapi.RegisterRequest) (*userapi.RegisterResponse, error)
	// Login(req *userapi.LoginRequest) (*userapi.LoginResponse, error)
	// Check(req *userapi.CheckRequest) (*userapi.CheckResponse, error)
}

type router struct {
	api userAPI
	http.Handler
}

func New(
	api userAPI,
) http.Handler {
	srv := &router{
		api: api,
	}

	srv.Handler = srv.setupChiRouter()
	return srv
}

func (srv *router) setupChiRouter() chi.Router {
	rt := chi.NewRouter()
	rt.Route("/v1/user", func(r chi.Router) {
		r.Post("/register", srv.fnRegister) // POST /v1/user/register
		r.Post("/login", srv.fnLogin)       // POST /v1/user/login
	})
	return rt
}

func (srv *router) fnRegister(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return
	}

	req := new(userapi.RegisterRequest)
	if err := json.Unmarshal(body, req); err != nil {
		writeError(w, err)
		return
	}

	response, err := srv.api.Register(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeResponse(w, response)
}

func (srv *router) fnLogin(w http.ResponseWriter, r *http.Request) {
	return
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	b, err := json.Marshal(resp)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Write(b)
}

func writeError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
