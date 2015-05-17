define(["react","actions/AppActions","stores/AppStore"], 
function(React, AppActions) {

return React.createClass({

    setMessage: function() {
      AppActions.setMessage("LALA!")
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
        onClick: AppActions.setMessage, 
      },"");
    },

    _onChange: function() {
      alert("");
    }
  });
});
