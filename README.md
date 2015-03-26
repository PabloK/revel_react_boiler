### Start the web server:

    revel run rrb

   Run with <tt>--help</tt> for options.
   
### Run the tests:

    revel test rrb
	 
### Generate documentation

    godoc -html rrb/package/<package_to_generate_doc_for> > doc/<output.html>

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


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### For more information see:
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](http://revel.github.io/docs/godoc/index.html).
* The [react docs](http://facebook.github.io/react/docs/getting-started.html).
* The [flux docs](http://facebook.github.io/flux/docs/overview.html).
* The [jquery docs](http://api.jquery.com/).
* The [requiurejs webpage](http://requirejs.org/).
