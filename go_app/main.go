package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

func main() {
	controllers.AuthController{}.Init()
	controllers.UserController{}.Init()
	controllers.ProjectController{}.Init()
	controllers.TaskController{}.Init()
	controllers.TaskProjectController{}.Init()
	controllers.ProjectMemberController{}.Init()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	corsUseCase := usecases.CorsUseCase{}
	corsUseCase.DisableCors(&w, r)

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
