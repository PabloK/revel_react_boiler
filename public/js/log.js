define('log', [] ,function () {
  if (console){
    return console;
  }
    return {
        info: function(msg){},
        warning: function(msg){},
        debug: function(msg){},
        info: function(msg){},
        log: function(msg){},
    };
});
