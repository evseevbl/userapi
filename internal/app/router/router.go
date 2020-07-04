// a simple router using go-chi
package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/evseevbl/userapi/internal/app/userapi"
)

type userAPI interface {
	Register(ctx context.Context, req *userapi.RegisterRequest) (*userapi.RegisterResponse, error)
	Login(ctx context.Context, req *userapi.LoginRequest) (*userapi.LoginResponse, error)
}

type router struct {
	api userAPI
	http.Handler
}

// New constructor
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
	rt.Use(middleware.Recoverer) // keep going in case of panic
	rt.Use(middleware.Logger)    // log requests
	rt.Route("/v1/user", func(r chi.Router) {
		r.Post("/register", srv.fnRegister) // POST /v1/user/register
		r.Post("/login", srv.fnLogin)       // POST /v1/user/login
	})
	return rt
}

// fnRegister handler
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

// fnLogin handler
func (srv *router) fnLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return
	}

	req := new(userapi.LoginRequest)
	if err := json.Unmarshal(body, req); err != nil {
		writeError(w, err)
		return
	}

	response, err := srv.api.Login(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeResponse(w, response)
}

// writeResponse add data to response body
func writeResponse(w http.ResponseWriter, resp interface{}) {
	b, err := json.Marshal(resp)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Write(b)
}

// writeError add error message to response body
func writeError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
