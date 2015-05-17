define(["react","actions/AppActions","stores/AppStore"], 
function(React, AppActions) {

return React.createClass({

    setMessage: function() {
      AppActions.setMessage("LALA!")
    },
    componentDidMount: function(){
      AppStore.addChangeListner(this._onchange);
    },
    compnentWillUnmount: function() {
      AppStore.removeChangeListner(this._onchange);
    },

    render:  function(){
      return React.createElement('input', {
        type: "button", 
        value:"Click Me",
        onClick: "setMessage", 
      },"");
    },

    _onChange: function() {
      alert("");
    }
  });
});
