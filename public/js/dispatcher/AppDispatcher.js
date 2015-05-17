define(["react", "flux"],
function(React, flux) {

  var AppDispatcher = new flux.Dispatcher();

  return AppDispatcher.handleAction = function(action) {
    this.dispatch({
      source: 'SET_MESSAGE',
      action: action
    })
  };
});
