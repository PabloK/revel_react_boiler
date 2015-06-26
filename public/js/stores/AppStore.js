define(["react","eventemitter","dispatcher/AppDispatcher","constants/AppConstants","log"], 
function(React, EventEmitter, AppDispatcher, AppConstants, log) {

  var _message = {}

  function setMessage(data) {
    _message = data[0]
  }

  var AppStore = $.extend({},EventEmitter.prototype, {

    getMessage: function() {
      return _message;
    },

    setMessage: function(msg) {
      _message = msg;
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
    log.info(action);

    switch(action.actionType) {
      case AppConstants.SET_MESSAGE:
        AppStore.setMessage(action.data);
      break;
      default:
        return true;
    }

    AppStore.emitChange();

    return true;

  });

  return AppStore;
});
