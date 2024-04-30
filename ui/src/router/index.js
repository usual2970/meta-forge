import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { loadLayoutMiddleware } from './middleware/loadLayoutMiddleware'
import { initializeMiddleware } from './middleware/initializeMiddleware'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      meta:{
        layout:'Admin'
      },
      component: HomeView,
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
    },
    {
      path: '/entity/:name',
      name: 'entity',
      meta:{
        layout:'Admin'
      },
      component: () => import('../views/EntityView.vue')
    }
    ,
    {
      path: '/entity/:name/config/:type',
      name: 'entity-setting',
      meta:{
        layout:'Admin'
      },
      component: () => import('../views/EntitySettingView.vue')
    },
    {
      path: '/entity/:name/:id',
      name: 'entity-detail',
      meta:{
        layout:'Admin'
      },
      component: () => import('../views/EntityDetailView.vue')
    }
  ]
})

router.beforeEach(initializeMiddleware)
router.beforeEach(loadLayoutMiddleware)

export default router
