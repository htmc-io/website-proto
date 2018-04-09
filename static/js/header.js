var header = `
<div class='header-wrapper'>
  <div id='logo'>
    <img src="/img/logo.png">
    <span>汉唐明创</span>
  </div>
  <nav id="nav-list">
    <ul>
      <li class='nav-home-url'><a href="/page/home.html">首页</a></li>
      <li class='nav-solution-url'><a href="/page/solution.html">方案</a></li>
      <li class='nav-case-url'><a href="/page/case.html">案例</a></li>
      <li class='nav-about-url'><a href="/page/about.html">关于</a></li>
    </ul>
  </nav>
</div>
`
Vue.component('v-header', {
  template: header
});
