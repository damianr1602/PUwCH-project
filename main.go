package main

import (
	"log"
	"os"

	. "github.com/damianr1602/chmuryrest/config"
	. "github.com/damianr1602/chmuryrest/dao"

	"github.com/damianr1602/chmuryrest/gen/restapi"
	"github.com/damianr1602/chmuryrest/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jessevdk/go-flags"
)

var conf = Config{}
var dao = MoviesDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	conf.Read()

	dao.Server = conf.Server
	dao.Database = conf.Database
	dao.Connect()
}

// main parses the arguments from the CLI as specified
// by our configuration described in `cliArgs` and then
// populates the `args` reference we defined in the `vars`
// section above.
func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMoviesAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.GetVersionHandler = operations.GetVersionHandlerFunc(
		func() middleware.Responder {
			return operations.NewGetVersionOK().WithPayload("1.0.0")
		})

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

	server.ConfigureAPI()

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
