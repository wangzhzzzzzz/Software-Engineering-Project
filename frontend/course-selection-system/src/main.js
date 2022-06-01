import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import axios from 'axios';
import VueAxios from 'vue-axios';

import 'default-passive-events'
// 引入 element-ui
import ElementUI from 'element-ui'
// 引入 element-ui 的 css 文件
import 'element-ui/lib/theme-chalk/index.css';
import VueCookies from 'vue-cookies';

// 声明使用 element-ui
Vue.use(ElementUI);
Vue.use(VueAxios, axios);
Vue.use(VueCookies);
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
