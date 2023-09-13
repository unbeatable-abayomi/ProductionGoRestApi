package main

import (
	"fmt"
	"net/http"
	transportHTTP "github.com/unbeatable-abayomi/ProductionGoRestApi/internal/transport/http"
)

//App struct wjich contains things like pointers to database connections
type App struct{}

func (app *App) Run() error{
    fmt.Println("Setting Up Our App");
	handler := transportHTTP.NewHandler();
	handler.SetUpRoutes();

	if err := http.ListenAndServe(":8080", handler.Router); err != nil{
        fmt.Println("Falied to set up server")
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