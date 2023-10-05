package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unbeatable-abayomi/ProductionGoRestApi/internal/comment"
	log	"github.com/sirupsen/logrus"
)

//Handler - stores pointer to our comments service
type Handler struct{
  Router *mux.Router
  Service *comment.Service
}

//Response --- an object to store responses from our APIs
type Reponse struct{
	Message string
	Error string
}

//NewHandler - returns a pointer to a Handler                                         
func NewHandler(service *comment.Service) *Handler{
	return &Handler{Service: service }
}

// SetUpRoutes - sets up all the routes for our application
func(h *Handler) SetUpRoutes() {
    fmt.Println("Setting Up Routes")
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment",h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment",h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}",h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}",h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}",h.DeleteComment).Methods("DELETE")
	
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
	  //fmt.Fprintf(w, "I am alive");	
	 // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	  //w.WriteHeader(http.StatusOK)
	  //if err := json.NewEncoder(w).Encode(Reponse{Message: "I am Alive"}); err != nil{
		//panic(err)
	  //}
	  if err := sendOKResponse(w, Reponse{Message: "I am Alive"}); err != nil{
		panic(err)
	  }
	})
} 

// GetComment -- retrieve a comment by ID
func (h *Handler)GetComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	i,err := strconv.ParseUint(id,10,64);
	if err != nil {
		//fmt.Fprintf(w,"Unable to parse UNIT from ID");
		sendErrorResponse(w,"Unable to parse UNIT from ID", err);
		return 
	}
  comment, err := h.Service.GetComment(uint(i));
  if err != nil {
	//fmt.Fprintf(w, "Error Retriving Comment by Id");
	sendErrorResponse(w, "Error Retriving Comment by Id",err);
	return
  }
 //if err:= json.NewEncoder(w).Encode(comment); err != nil {
	//panic(err)
 //}
 if err:= sendOKResponse(w, comment); err != nil {
	panic(err)
 }
  //fmt.Fprintf(w, "%+v", comment)
}       

//GetAllComments -- retrives all comments from the comment service
func (h *Handler)GetAllComments(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		//fmt.Fprintf(w, "Failed to Retreive all comments")
		sendErrorResponse(w, "Failed to Retreive all comments",err)
		return
	}
	//if err:= json.NewEncoder(w).Encode(comments); err != nil {
		//panic(err)
	 //}
	 if err:= sendOKResponse(w, comments); err != nil {
		panic(err)
	 }
	//fmt.Fprintf(w, "%+v", comments)
}

//PostComment --adds a new comment

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		//fmt.Fprintf(w, "Failed to Decode JSON Body");
		sendErrorResponse(w, "Failed to Decode JSON Body",err);
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		//fmt.Fprintf(w, "Failed to post new comment")
		sendErrorResponse(w, "Failed to post new comment",err)
	}
	//if err:= json.NewEncoder(w).Encode(comment); err != nil {
		//panic(err)
	 //}
	 if err:= sendOKResponse(w, comment); err != nil {
		panic(err)
	 }
	
  //fmt.Fprintf(w, "%+v", comment)
}   

//UpdateComment Updates a comment by ID'
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		//fmt.Fprintf(w, "Failed to Decode JSON Body");
		sendErrorResponse(w, "Failed to Decode JSON Body", err);
		return
	}

	vars := mux.Vars(r);
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		//fmt.Fprintf(w,"Failed to parse Unit from ID")
		sendErrorResponse(w,"Failed to parse Unit from ID",err)
		return
	}
	comment, err = h.Service.UpdateComment(uint(commentID), comment)

	   if err != nil {
		//fmt.Fprintf(w, "Failed to Update Comments")
		sendErrorResponse(w, "Failed to Update Comments",err)
		return
	   }
	   //fmt.Fprintf(w, "%+v", comment)
	  // if err:= json.NewEncoder(w).Encode(comment); err != nil {
		//panic(err)
	 //}

	 	   //fmt.Fprintf(w, "%+v", comment)
	   if err:= sendOKResponse(w, comment); err != nil {
		panic(err)
	 }
	 //sendOKResponse
}

//DeleteComment deletes comments by Id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r);
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		//fmt.Fprintf(w,"Failed to parse Unit from ID")
		sendErrorResponse(w,"Failed to parse Unit from ID",err)
		return
	}
   err = h.Service.DeleteComment(uint(commentID))
   if err != nil {
	//fmt.Fprintf(w, "Failed to delete comment by Comment ID")
	sendErrorResponse(w, "Failed to delete comment by Comment ID",err)
	return
   }

    if err = sendOKResponse(w,Reponse{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
   //if err := json.NewEncoder(w).Encode(Reponse{Message: "Comment successfully deleted"}); err != nil{
	// panic(err)
   //}
   
   //fmt.Fprintf(w, "Success fully deleted comment by Comment ID")
}

func sendOKResponse(w http.ResponseWriter, resp interface{})error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
   if err := json.NewEncoder(w).Encode(Reponse{Message: message,Error: err.Error()}); err != nil {
	panic(err);
   }
}

  