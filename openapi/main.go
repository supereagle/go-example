package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

const (
	apiVersion = "/api/v1"

	// Response
	errorCode      = 1
	successCode    = 0
	successMessage = "success"
)

type Book struct {
	Title  string `json:"title" description:"title of book"`
	Author string `json:"author" description:"author of book"`
}

type Consumer struct {
	Name    string `json:"name" description:"name of consumer"`
	Contact string `json:"contact" description:"contact of consumer"`
}

type ResponseEntity struct {
	ErrorCode    int         `json:"errorCode" description:"error code"`
	ErrorMessage string      `json:"errorMessage,omitempty" description:"error message"`
	Data         interface{} `json:"data,omitempty" description:"response data"`
}

func main() {
	ws := new(restful.WebService)
	bookTags := []string{"Book"}
	ws.Path(apiVersion).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	// POST /api/v1/books/{bookId}
	ws.Route(ws.POST("/books/{bookId}").To(createBook).
		Doc("Add a book").
		Param(ws.HeaderParameter("X-AUTH-TOKEN", "token").DataType("string").Required(true)).
		Param(ws.PathParameter("bookId", "id of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, bookTags).Reads(Book{}).
		Writes(ResponseEntity{}))

	// PUT /api/v1/books/{bookId}
	ws.Route(ws.PUT("/books/{bookId}").To(updateBook).
		Doc("Update a book").
		Param(ws.HeaderParameter("X-AUTH-TOKEN", "token").DataType("string").Required(true)).
		Param(ws.PathParameter("bookId", "id of the book").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, bookTags).Reads(Book{}).
		Writes(ResponseEntity{}))

	consumerTags := []string{"Consumer"}
	ws.Path(apiVersion).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	// POST /api/v1/consumers/{consumerId}
	ws.Route(ws.POST("/consumers/{consumerId}").To(createBook).
		Doc("Add a consumer").
		Param(ws.HeaderParameter("X-AUTH-TOKEN", "token").DataType("string").Required(true)).
		Param(ws.PathParameter("consumerId", "id of the consumer").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, consumerTags).Reads(Consumer{}).
		Writes(ResponseEntity{}))

	// PUT /api/v1/consumers/{consumerId}
	ws.Route(ws.PUT("/consumers/{consumerId}").To(updateBook).
		Doc("Update a consumer").
		Param(ws.HeaderParameter("X-AUTH-TOKEN", "token").DataType("string").Required(true)).
		Param(ws.PathParameter("consumerId", "id of the consumer").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, consumerTags).Reads(Consumer{}).
		Writes(ResponseEntity{}))

	// NOTE Must add this web service befor openapi service.
	restful.DefaultContainer.Add(ws)

	// Enable openapi.
	enableOpenApiService()
	addOpenApiStaticHandler("./swagger-ui/dist")

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err.Error())
	}
}

// enableOpenApiService enables the openapi service.
func enableOpenApiService() {
	config := restfulspec.Config{
		WebServices:    restful.RegisteredWebServices(),
		WebServicesURL: "http://localhost:8080",
		APIPath:        "/apidocs.json",
	}
	openapiService := restfulspec.NewOpenAPIService(config)
	restful.DefaultContainer.Add(openapiService)

}

func addOpenApiStaticHandler(swaggerUIDir string) {
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/",
		http.StripPrefix("/apidocs/",
			http.FileServer(http.Dir(swaggerUIDir))))
}

func createBook(request *restful.Request, response *restful.Response) {
	//bookId := request.PathParameter("bookId")
	book := &Book{}
	err := request.ReadEntity(book)
	if err != nil {
		re := ResponseEntity{
			ErrorCode:    successCode,
			ErrorMessage: err.Error(),
		}
		response.WriteEntity(re)
		return
	}
	re := ResponseEntity{
		ErrorCode:    1,
		ErrorMessage: "",
		Data:         book,
	}
	response.WriteEntity(re)
}

func updateBook(request *restful.Request, response *restful.Response) {
	//bookId := request.PathParameter("bookId")
	book := &Book{}
	err := request.ReadEntity(book)
	if err != nil {
		re := ResponseEntity{
			ErrorCode:    successCode,
			ErrorMessage: err.Error(),
		}
		response.WriteEntity(re)
		return
	}
	re := ResponseEntity{
		ErrorCode:    1,
		ErrorMessage: "",
		Data:         book,
	}
	response.WriteEntity(re)
}

func createConsumer(request *restful.Request, response *restful.Response) {
	//consumerId := request.PathParameter("consumerId")
	consumer := &Consumer{}
	err := request.ReadEntity(consumer)
	if err != nil {
		re := ResponseEntity{
			ErrorCode:    successCode,
			ErrorMessage: err.Error(),
		}
		response.WriteEntity(re)
		return
	}
	re := ResponseEntity{
		ErrorCode:    1,
		ErrorMessage: "",
		Data:         consumer,
	}
	response.WriteEntity(re)
}

func updateConsumer(request *restful.Request, response *restful.Response) {
	//consumerId := request.PathParameter("consumerId")
	consumer := &Consumer{}
	err := request.ReadEntity(consumer)
	if err != nil {
		re := ResponseEntity{
			ErrorCode:    successCode,
			ErrorMessage: err.Error(),
		}
		response.WriteEntity(re)
		return
	}
	re := ResponseEntity{
		ErrorCode:    1,
		ErrorMessage: "",
		Data:         consumer,
	}
	response.WriteEntity(re)
}
