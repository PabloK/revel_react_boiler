define(["jquery", "react", "cookies", "log", "flux"], 
function($, React, cookies, log, flux) {

  log.info("Initializing application");
  require(["components/maincom"], function(m){
    React.render(
      React.createElement(m),document.getElementById('main'));
  });

});
