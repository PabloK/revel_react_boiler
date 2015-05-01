define(["jquery", "react", "cookies", "log", "flux"], 
function($, React, cookies, log, flux) {

  require(["components/maincom"], function(m){
    React.render(
      React.createElement(m),document.getElementById('main'));
  });

});
