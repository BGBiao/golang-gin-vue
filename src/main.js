// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

// 引入element-ui框架
// import ElementUI from 'element-ui'

Vue.config.productionTip = false

// Vue.vue(ElementUI)

/* eslint-disable no-new */
// 注释要以空格开头;末尾行不能有空行
// 实例化vue
new Vue({
  // '#app'选择器用来选择html中DOM 元素中的 id,所有的变动仅生效在web的dav中
  el: '#web',
  router,
  components: { App },
  template: '<App/>'
})
