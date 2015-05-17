define(["react","actions/AppActions","stores/AppStore"], 
function(React, AppDispatcher, EventEmitter, AppConstants, EventEmitter) {
// var EventEmitter = require('events').EventEmitter;

  var _message = {}

  function setMessage(data) {
    _message = data[0]
  }

  EventEmitter = {};

  return AppStore = $.extend({},EventEmitter.prototype, {

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

    ProductStore.emitChange();

    return true;

  });
});
