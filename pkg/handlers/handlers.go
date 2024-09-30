package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stephieandrade/go-course/pkg/render"
)
	
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
	//fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	// sum := addValues(2,2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, and 2+2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))

}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		error := errors.New("cannot divide by zero")
		return 0, error
	}
	result := x / y
	return result, nil
}


func addValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}