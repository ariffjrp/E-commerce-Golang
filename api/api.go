package api

import (
	repo "a21hc3NpZ25tZW50/repository"
	"fmt"
	"net/http"
	"path"
	"html/template"
	"encoding/json"
	"a21hc3NpZ25tZW50/model"
)

type API struct {
	usersRepo    repo.UserRepository
	sessionsRepo repo.SessionsRepository
	products     repo.ProductRepository
	cartsRepo    repo.CartRepository
	mux          *http.ServeMux
}

type Page struct {
	File string
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", p.File)
	data, Alert := template.ParseFiles(filepath)

	if (Alert != nil){
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Alert := model.ErrorResponse{
			Error: "Server Internal Error",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}else{
		
	}

	Alert = data.Execute(w, nil)
	// TODO: answer here
}

func NewAPI(usersRepo repo.UserRepository, sessionsRepo repo.SessionsRepository, products repo.ProductRepository, cartsRepo repo.CartRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		sessionsRepo,
		products,
		cartsRepo,
		mux,
	}

	index := Page{File: "index.html"}
	mux.Handle("/", api.Get(index))

	// Please create routing for:
	// - Register page with endpoint `/page/register`, GET method and render `register.html` file on views folder
	// - Login page with endpoint `/page/login`, GET method and render `login.html` file on views folder
	// TODO: answer here
	mux.Handle("/page/register", api.Get(Page{File: "register.html"}))
	mux.Handle("/page/login", api.Get(Page{File: "login.html"}))

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/user/img/profile", api.Get(api.Auth(http.HandlerFunc(api.ImgProfileView))))
	mux.Handle("/user/img/update-profile", api.Post(api.Auth(http.HandlerFunc(api.ImgProfileUpdate))))

	// Please create routing for endpoint `/cart/add` with POST method, Authentication and handle api.AddCart
	// TODO: answer here
	mux.Handle("/cart/add", api.Post(api.Auth(http.HandlerFunc(api.AddCart))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
