package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func createTemplateCache(map[string]*template.Template, error){
	myCache := map[string]*template.Template{}
	

}

// var templateCache = make(map[string]*template.Template)
// func RenderTemplateTest(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have the template in cache
// 	_, inMap := templateCache[t]
// 	if !inMap{
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err !=nil{
// 			log.Println(err)
// 		}
// 	}else{
// 		log.Println("using cached template")
// 	}

// 	tmpl = templateCache[t]

// 	err = tmpl.Execute(w, nil)
// 	if err !=nil{
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error{
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err !=nil{
// 		return err
// 	}
// 	templateCache[t] = tmpl

// 	return nil
// }