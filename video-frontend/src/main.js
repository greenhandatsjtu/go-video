import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueClipboard from 'vue-clipboard2'
import router from './router'
import store from "./store"

Vue.use(VueClipboard)

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store: store,
  render: h => h(App)
}).$mount('#app')
