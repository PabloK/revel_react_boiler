define(["react","actions/AppActions","stores/AppStore","log"], 
function(React, AppActions, AppStore,log) {

return React.createClass({

    setMessage: function() {
      AppActions.setMessage("Test message");
    },
    componentDidMount: function(){
      AppStore.addChangeListner(this._onChange);
    },
    compnentWillUnmount: function() {
      AppStore.removeChangeListner(this._onChange);
    },

    render:  function(){
      return React.createElement('input', {
        type: "button", 
        value:"Click Me",
        onClick: this.setMessage, 
      });
    },

    _onChange: function() {
      alert(AppStore.getMessage());
    }
  });
});
