package main

import (
	"fmt"
	"net/http"

	"github.com/unbeatable-abayomi/ProductionGoRestApi/internal/comment"
	transportHTTP "github.com/unbeatable-abayomi/ProductionGoRestApi/internal/transport/http"
	"github.com/unbeatable-abayomi/ProductionGoRestApi/internal/database"
    log	"github.com/sirupsen/logrus"
)     

//App struct which contains application information
type App struct{
	Name string
	Version string
}

func (app *App) Run() error{
    //fmt.Println("Setting Up Our App");
log.SetFormatter(&log.JSONFormatter{})
log.WithFields(
	log.Fields{
		"AppName" : app.Name,
		"AppVersion" : app.Version,
	}).Info("Setting Up Application")
	var err error
	db, err := database.NewDataBase()
	if err != nil {
		   return err
	}
    err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db);
	handler := transportHTTP.NewHandler(commentService);
	handler.SetUpRoutes();

	if err := http.ListenAndServe(":8080", handler.Router); err != nil{
       // fmt.Println("Failed to set up server")
	   log.Error("Failed to set up server")
		return err; 
	}
	return nil;                               
}

func main(){
	fmt.Println("Go Rest API Course")
	app := App{Name: "Commenting Servcie", Version: "1.0.0"}
	if err := app.Run(); err != nil {
       //fmt.Println("Error starting Our RestApi");
	   log.Error("Error starting Our RestApi")
	   //fmt.Println(err);
	   log.Fatal(err)
	}
}