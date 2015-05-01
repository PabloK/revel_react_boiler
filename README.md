### Start the web server:

    revel run rrb

   Run with <tt>--help</tt> for options.
   
### Run the tests:
To run integration tests
    revel test rrb
To run package unit tests
		go test -v ./packages/*
	 
### View the package documentation

If changes has been made in the packages folder
		mkdir <GOROOT>\packages
		copy packages <GOROOT>\packages
Run the doc server
    godoc -http:6060
Browse to localhost:6060 to see the documentation

### Description of Contents

The directory structure:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      packages		Custom go packages
      LICENSES		external licenses
      LICENSE		this projects license
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.
    * environment.json, custom configuration which can be overriden with environment variables use the
      packages/configurationmanager to get these in the application.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory and the package directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### Prerequisites

go installation
revel installation
mgo "mongoDB library for go" (go get gopkg.in/mgo.v2)
MongoDB database service

### Development Installation

#### Installing go
See these instructions on how to install go (https://golang.org/doc/install)
Make sure to use default directory.

You can also take a look at thsese instructions (https://revel.github.io/tutorial/gettingstarted.html)

#### Installing Revel
If you did not already install revel according to these instructions (https://revel.github.io/tutorial/gettingstarted.html)
Don't forget to install Git and Mercurial as mentioned in the guide above

You do not need to create a new revel application.

#### Install mgo
Run the following command
	go get gopkg.in/mgo.v2

####Installing MongoDB as a service (Windows)
For instructions see http://docs.mongodb.org/manual/tutorial/install-mongodb-on-windows/

####Clone the rrb git repo into ../gocode/src/ (Using a git terminal)
git clone git@github.com:PabloK/revel_react_boiler.git

####Verifying the installation
Start the webserver
    revel run rrb
Open a browser and surf to http://localhost:9000

### For more information see:
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](http://revel.github.io/docs/godoc/index.html).
* The [React docs](http://facebook.github.io/react/docs/getting-started.html).
* The [Flux docs](http://facebook.github.io/flux/docs/overview.html).
* The [jQuery docs](http://api.jquery.com/).
* The [Requirejs webpage](http://requirejs.org/).
