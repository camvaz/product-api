package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/camvaz/product-api/data"
	"github.com/camvaz/product-api/models"
)


type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		p.readProducts(rw,r)
	case http.MethodPost:
		p.createProducts(rw,r)
	case http.MethodPut:
		path := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(path, -1)
		
		if len(g) != 1 {
			p.l.Println("Invalid URI more than 1 id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than 1 capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
 
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil{
			p.l.Println("Invalid URI invalid number conversion")
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) readProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw) 

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) createProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST Products")

	prod := &models.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json",http.StatusBadRequest)
	}

	p.l.Printf("Prod %#v", prod)

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle PUT Products")
	p.l.Println("got id: ",id)

	prod := &models.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json",http.StatusBadRequest)
	}

	p.l.Printf("Prod %#v", prod)

	err := data.UpdateProduct(id,prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}