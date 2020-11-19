import Vue from 'vue'
import App from './App.vue'
import router from './router'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
ElementUI.Dialog.props.closeOnClickModal.default = false
Vue.use(ElementUI);


import { store } from '@/store/index'
Vue.config.productionTip = false

const Bus = (Vue) => {
  const Bus = new Vue({
    methods: {
      emit(event, ...args) {
        this.$emit(event, ...args)
      },
      on(event, cb) {
        this.$on(event, cb)
      },
      off(event, cb) {
        this.$off(event, cb)
      }
    },
  })
  Vue.prototype.$bus = Bus
}

Vue.use(Bus)


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
