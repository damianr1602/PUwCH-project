package main

import (
	"log"
	"os"

	"github.com/damianr1602/chmuryrest/gen/models"

	. "github.com/damianr1602/chmuryrest/config"
	. "github.com/damianr1602/chmuryrest/dao"
	logger "github.com/damianr1602/chmuryrest/logging"

	"github.com/damianr1602/chmuryrest/gen/restapi"
	"github.com/damianr1602/chmuryrest/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
)

var conf = Config{}
var dao = MoviesDAO{}

func getAllMoviesHandler(params operations.AllMoviesParams) middleware.Responder {
	logger.Log.Println("getAllMoviesHandler ", params)

	mergedParams := operations.NewAllMoviesParams()
	mergedParams.Skip = swag.Int32(0)
	if params.Skip != nil {
		mergedParams.Skip = params.Skip
	}
	if params.Limit != nil {
		mergedParams.Limit = params.Limit
	} else {
		limit := int32(2)
		mergedParams.Limit = &limit
	}
	logger.Log.Println("getAllMoviesHandler ", mergedParams)
	// logger.Log.Println("getAllMoviesHandler ", dao.)

	movies, err := dao.FindAll(0, 2)
	logger.Log.Println("getAllMoviesHandler movies: ", movies, "len movies: ", len(movies))

	// if len(movies) >= int(mergedParams.Limit) {
	// 	return
	// }
	if err != nil {
		code := int64(404)
		errPayload := &models.Result{
			Code:    &code,
			Message: swag.String("Not found"),
		}
		return operations.NewAllMoviesNotFound().WithPayload(errPayload)
	}
	return operations.NewAllMoviesOK().WithPayload(movies)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	conf.Read()
	dao.Server = conf.Server
	dao.Database = conf.Database
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMoviesAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.AllMoviesHandler = operations.AllMoviesHandlerFunc(getAllMoviesHandler)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "A db movie api"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
	// Configure the server port
	server.Port = 3000

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
