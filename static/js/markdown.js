var temp = `
<div class='markdown' v-html="html"></div>
`
Vue.component('markdown', {
  template: temp,
  data: function() {
    return {
      html: ""
    }
  },
  created: function() {
    var url = getParameters()["url"];
    console.log("URL:" + url);

    var self = this;
    axios.get("/markdown/" + url).then(function(response) {
      self.html = response.data; //
    }).catch(function(error) { console.log(error); });
  }
});

function getParameters() {
  var vars = {};
  var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(m, key, value) {
    vars[key] = value;
  });
  return vars;
}
