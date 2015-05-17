define(["react", "flux"],
function(React, flux) {

  var AppDispatcher = new flux.Dispatcher();

  AppDispatcher.handleAction = function(action) {
    this.dispatch({
      source: 'SET_MESSAGE',
      action: action
    });
  };

  return AppDispatcher;
});
