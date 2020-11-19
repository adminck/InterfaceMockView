import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '@/views/Login.vue'
import about from '@/views/home'
import Domain from '@/views/Domain'
import error from '@/views/error.vue'
import register from '@/views/register.vue'
import { store } from '@/store/index'

Vue.use(VueRouter)
const whiteList = ['Login','register']
const routes = [
  { path: '/', redirect: { name: 'Login' } },
  { path: '/login', name: 'Login', component: Login },
  { path: '*', name: '404', component: error },
  { path: '/register', name: 'register', component: register },
  { path: '/about',
    name: 'About',
    component: about,
    children:[
      { path: '', redirect: { name: 'person' }},
      { path: 'person', name: 'person', component: () => import('@/views/ApiInfo/person.vue'),meta:{title:"接口管理"} },
      { path: 'Domain', name: 'Domain', component: Domain ,meta:{title:"域名管理"}}
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})


router.beforeEach(async(to, from, next) => {
  const token = store.getters['user/token']
 /* if (token) {
      const expiresAt = store.getters['user/expiresAt']
      const nowUnix = new Date().getTime()
      const hasExpires = (expiresAt - nowUnix) < 0
      if (hasExpires) {
          store.dispatch['user/claerAll']
      }
  }*/
  // 在白名单中的判断情况
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      next({ path: '/about' })
    } else {
      next()
    }
  } else {
    // 不在白名单中并且已经登陆的时候
    if (token) {
        next()
    }else {
      next({name: "Login"})
    }
  }
})



export default router
