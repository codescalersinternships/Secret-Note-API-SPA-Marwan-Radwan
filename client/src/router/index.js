import { createRouter, createWebHistory } from 'vue-router'
import CreateNote from '../views/CreateNote.vue'
import ViewNote from '../views/ViewNote.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'CreateNote',
      component: CreateNote,
    },
    { path: '/:id', name: 'ViewNote', component: ViewNote, props: true },
  ],
})

export default router
