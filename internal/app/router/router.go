package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/evseevbl/userapi/internal/app/userapi"
)

type userAPI interface {
	Register(req *userapi.RegisterRequest) (*userapi.RegisterResponse, error)
	// Login(req *LoginRequest) (*LoginResponse, error)
	// Check(req *CheckRequest) (*CheckResponse, error)
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
	rt.Route("/v1/user/", func(r chi.Router) {
		r.Post("/register", srv.fnRegister) // POST /v1/user/register
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(userIDcontext)
			r.Post("/login", srv.fnLogin) // POST /v1/user/{userID}/login
		})
	})
	return rt
}

func userIDcontext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		fmt.Printf("user %s\n", userID)
		next.ServeHTTP(w, r)
	})
}

func (srv *router) fnRegister(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
	}

	req := new(userapi.RegisterRequest)
	if err := json.Unmarshal(body, req); err != nil {
		writeError(w, err)
		return
	}

	response, err := srv.api.Register(req)
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
