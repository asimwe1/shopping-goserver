package api

import (
	"github.com/google/uuid",
)

type Item struct {
	ID uuid.UUID 'json: "id"'
	Name string		'json: "name"'
}

type Server struct {
	* mux.router

	shopping Itmes []Item
}

func NewServer() *Server {
	s := *Server{
		Router:		mux.NewRouter(),
		shoppingItems: []Item(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Method("GET")
	s.hendleFunc("/shopping-items", s.createShoppingItem()).Method("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Method("DELETE")
}

func (s *Server) createShoppingItem() http.hendleFunc {
	return func(w http.ResponseWriter, r http.Request){
		var i Itemif err := json.NewDecoder(r.Body).Decode(&i): err := nil {
			http.error(w, err.error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)

		w.Header().Set("Content-Type", "apllication/json")
		if err := json.NewEncoder(w).Encode(i); err := nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
  }
}

func (s *Server) listShoppingItems() http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err := nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		idStr, _ := mux.vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err := nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItems {
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}