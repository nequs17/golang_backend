package main

import (
	"backend/api"
	appAuth "backend/app/auth"
	appLogs "backend/app/logs"
	_ "backend/docs"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	Log.Info("Server has been started")

	routerRun()
}

func routerRun() {

	router := mux.NewRouter()

	//router.Handle("/", http.FileServer(http.Dir("./client/public"))) // Путь до Frontend части. |СОБРАННОЙ!|

	// Маршрут для документации Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Маршрут для отображения документации на странице /openapi
	router.HandleFunc("/openapi", api.OpenAPI).Methods("GET")

	// admin page
	router.HandleFunc("/api/admin/users", api.AllUsers).Methods("POST")       // <- in Release
	router.HandleFunc("/api/admin/changerole", api.ChangeRole).Methods("GET") // <- in Release

	//engineer page
	router.HandleFunc("/api/engineer/addbasestation", api.AddBaseStation).Methods("POST")       // <- in Release
	router.HandleFunc("/api/engineer/verifybasestation", api.VerifyBaseStation).Methods("POST") // <- in Release
	//router.HandleFunc("/api/engineer/deletebasestation", api.VerifyBaseStation).Methods("POST") // <- in Release

	// user
	router.HandleFunc("/api/user/register", api.UserRegister).Methods("POST")       // <- in Release
	router.HandleFunc("/api/user/auth", api.UserAuth).Methods("POST")               // <- in Release
	router.HandleFunc("/api/user/verify", api.UserVerify).Methods("GET")            // <- in Release
	router.HandleFunc("/api/user/logout", api.UserLogout).Methods("GET")            // <- in Release
	router.HandleFunc("/api/user/postapptrafic", api.PostAppTrafic).Methods("POST") // <- in Test
	router.HandleFunc("/api/user/getapptrafic", api.GetAppTraffic).Methods("GET")   // <- in Test

	// for testing jwt
	router.HandleFunc("/api/jwt/test", api.JwtTest).Methods("POST")     // <- in Release
	router.HandleFunc("/api/jwt/verify", api.JwtVerify).Methods("POST") // <- in Release

	// Thermalmap
	router.HandleFunc("/api/sockets/thermalmap", api.SocketThermal).Methods("GET")                // <- in Release
	router.HandleFunc("/api/sockets/thermalmapdataall", api.SocketThermalOut).Methods("GET")      // <- in Release
	router.HandleFunc("/api/sockets/thermalmapdata", api.SocketThermalOutByParams).Methods("GET") // <- in Release
	router.HandleFunc("/api/picture/thermalmappic", api.ThermalMapPicture).Methods("GET")         // <- in Release

	// Home page

	router.Use(appLogs.Handler)
	router.Use(appAuth.Handler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

/*

	Добавить роль для ответа при авторизации

*/
