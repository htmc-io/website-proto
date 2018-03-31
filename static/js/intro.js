var temp = `
<div class='intro'>
  <div class='content-wrapper'>
    <div class='content'>
      <h2 class='title'>{{intro.title}}</h2>
      <ul>
        <li v-for="paragraph in intro.abstract">
          {{ paragraph}}
        </li>
      </ul>
      <a :href="intro.moreUrl">了解更多...</a>
    </div>
  </div>
  <div class='cover-wrapper'>
    <img :src="intro.image">
  </div>
</div>
`
Vue.component('intro', {
  template: temp,
  props: ['url'],
  data: function() {
    return {
      intro: {
        title: "",
        abstract: "",
        image: "",
        moreUrl: ""
      }
    }
  },
  created: function() {
    var self = this;
    axios.get("/intro" + this.url).then(function(response) {
      self.intro.title = response.data.title;
      self.intro.abstract = response.data.abstract.split("；");
      self.intro.image = response.data.image;
      self.intro.moreUrl = "/page/article.html?url=" + self.url;
      console.log(response.data.title);
    }).catch(function(error) { console.log(error); });
  }
});
