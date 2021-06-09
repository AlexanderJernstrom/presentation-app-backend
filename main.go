package main

import (
	"net/http"
	"server/db"
	"server/middlewares"
	elementRoutes "server/routes/element"
	"server/routes/login"
	presentationRoutes "server/routes/presentation"
	"server/routes/register"
	slideRoutes "server/routes/slide"
)




func main() {

	db.Connect()



	
	http.HandleFunc("/register", (http.HandlerFunc(register.Register)))
	http.HandleFunc("/login",(http.HandlerFunc(login.Login)))

	http.Handle("/presentations",middlewares.SetCors( middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.GetAllPresentations))))
	http.Handle("/createPresentation", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.CreatePresentation))))
	http.Handle("/presentation", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.GetPresentation))))
	http.Handle("/updatePresentation", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.UpdatePresentation))))
	http.Handle("/deletePresentation", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.DeletePresentation))))
	http.Handle("/changeName", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(presentationRoutes.ChangeName))))

	http.Handle("/createSlide", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(slideRoutes.CreateSlide))))

	http.Handle("/createElement", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(elementRoutes.CreateElement))))
	http.Handle("/updateElement", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(elementRoutes.UpdateElement))))
	http.Handle("/deleteElement", middlewares.SetCors(middlewares.CheckAuth(http.HandlerFunc(elementRoutes.DeleteElement))))

	http.ListenAndServe(":5000", nil)


	
}