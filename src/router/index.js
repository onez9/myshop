import { createRouter, createWebHistory } from 'vue-router'
import ProductsView from '../views/Products.vue'
import AboutView from '../views/About.vue'
import HomeView from '../views/Home.vue'
import RegView from '../views/Reg.vue'
import AuthView from '../views/Auth.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about', 
      name: 'about',
      component: AboutView
    },
    {
      path: '/products',
      name: 'products',
      component: ProductsView
    },
    {
      path: '/auth',
      name: 'auth',
      component: AuthView
    },
    {
      path: '/reg',
      name: 'reg',
      component: RegView
    }
  ]
})

export default router