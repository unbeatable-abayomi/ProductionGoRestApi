package main

import "fmt"

//App struct wjich contains things like pointers
//to database connections           
type App struct{}

func (app *App) Run() error{
    fmt.Println("Setting Up Our App");
	return nil;                               
}

func main(){
	fmt.Println("Go Rest API Coures")
	app := App{}
	if err := app.Run(); err != nil {
       fmt.Println("Error starting Our RestApi");
	   fmt.Println(err);
	}
}