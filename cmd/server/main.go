package main

import (
	"fmt"
	"net/http"

	"github.com/unbeatable-abayomi/ProductionGoRestApi/internal/comment"
	transportHTTP "github.com/unbeatable-abayomi/ProductionGoRestApi/internal/transport/http"
	"github.com/unbeatable-abayomi/ProductionGoRestApi/internal/transport/http/database"
)

//App struct which contains application information
type App struct{}

func (app *App) Run() error{
    fmt.Println("Setting Up Our App");

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
        fmt.Println("Failed to set up server")
		return err; 
	}
	return nil;                               
}

func main(){
	fmt.Println("Go Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
       fmt.Println("Error starting Our RestApi");
	   fmt.Println(err);
	}
}