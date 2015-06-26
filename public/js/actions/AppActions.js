define(["react", "dispatcher/AppDispatcher", "constants/AppConstants","log"],
function(React, AppDispatcher, AppConstants,log) {

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
