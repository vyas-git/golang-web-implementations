package examples

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v2"
)

type UrlShortner struct {
}
type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func (u UrlShortner) Run() {
	mux := mux.NewRouter()

	//V1 using golang map
	pathsToUrls := map[string]string{
		"/wrsv":  "https://pkg.go.dev/bufio@go1.17",
		"/wrsqe": "https://golang.org/doc/tutorial/getting-started",
	}
	mapHandler := u.MapHandler(pathsToUrls, mux)

	//V2 uses yaml to parse urls
	yaml := `
- path: /qrsv
  url: https://golang.org/doc/tutorial/web-service-gin
- path: /pqrs
  url: https://golang.org/doc/tutorial/web-service-gin
`
	yamlHandler, err := u.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")

	mux.HandleFunc("/v1/{url}", mapHandler)
	mux.HandleFunc("/v2/{url}", yamlHandler)

	http.ListenAndServe(":8080", mux)
}

func (UrlShortner) MapHandler(pathstourls map[string]string, fallback *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		path := "/" + params["url"]

		if dest, ok := pathstourls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func (u UrlShortner) YAMLHandler(yamlBytes []byte, fallback *mux.Router) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return u.MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []PathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func parseYaml(data []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
