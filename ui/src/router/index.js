import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { loadLayoutMiddleware } from './middleware/loadLayoutMiddleware'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      meta:{
        layout:'Admin'
      },
      component: HomeView
    },
    {
      path: '/initial',
      name: 'initial',
      meta:{
        layout:'Blank'
      },
      component: () => import('../views/InitialView.vue')
    },
    {
      path: '/about',
      name: 'about',
      meta:{
        layout:'Blank'
      },
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

router.beforeEach(loadLayoutMiddleware)

export default router
