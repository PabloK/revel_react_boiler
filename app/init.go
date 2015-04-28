package app

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"rrb/packages/configmanager"
)

// Globally available variables
var configManager configmanager.ConfigHash
var session mgo.Session

func init() {

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	// register startup functions with OnAppStart
	revel.OnAppStart(initConfig)
	revel.OnAppStart(initDB)

	// Copy to own controller "revel/modules/static/app/controllers/static.go"
	// Edit to set max-age
}

func initConfig() {
	var c = configmanager.New("conf/environment.json")
	configManager = c.GetConfig()
}

// Initialize the database connection pool
func initDB() {

	// TODO: Encapsulate session
	session, err := mgo.Dial(configManager["DB_CONNECTION_STRING"])
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//c := session.DB("test").C("people")
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
