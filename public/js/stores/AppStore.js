define(["react","eventemitter","dispatcher/AppDispatcher","constants/AppConstants"], 
function(React, EventEmitter, AppDispatcher, AppConstants) {

  var _message = {}

  function setMessage(data) {
    _message = data[0]
  }

  var AppStore = $.extend({},EventEmitter.prototype, {

    getMessage: function() {
      return _message;
    },

    emitChange: function(){
      this.emit('change')
    },

    addChangeListner: function(callback) {
      this.on('change', callback);
    },

    removeChangeListner: function(callback) {
      this.removeListner('change', callback) 
    }

  });

  AppDispatcher.register(function(payload){
    var action = payload.action;
    var text;

    switch(action.actionType) {
      case AppConstants.SET_MESSAGE:
        alert(setMessage(action.data));
      break;
      default:
        return true;
    }

    AppStore.emitChange();

    return true;

  });

  return AppStore;
});
