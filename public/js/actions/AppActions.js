define(["react", "dispatcher/AppDispatcher", "constants/AppConstants"],
function(React, AppDispatcher, AppConstants) {

  var AppActions = {
  
    setMessage: function (data) {
      AppDispatcher.handleAction({
        actionType: AppConstants.SET_MESSAGE,
        data: data
      });
    }
  };

  return AppActions;
});
