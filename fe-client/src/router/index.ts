import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/dashboard/DashboardView.vue'),
    },
    {
      path: '/recently-updated',
      name: 'recently-updated',
      component: () => import('@/views/recently-updated/RecentlyUpdatedView.vue'),
    },
    {
      path: '/category',
      name: 'category',
      component: () => import('@/views/category/CategoryView.vue'),
    },
    {
      path: '/comic/:comicId',
      name: 'comic-detail',
      component: () => import('@/views/comic-details/ComicDetailView.vue'),
    },
    {
      path: '/comic/:comicId/chapter/:chapterId',
      name: 'chapter',
      component: () => import('@/views/chapter/ChapterView.vue'),
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/NotFoundView.vue'),
    },
    {
      path: '/search',
      name: 'search',
      component: () => import('@/views/search/SearchView.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/about/AboutView.vue'),
    },
    {
      path: '/short',
      name: 'short',
      component: () => import('@/views/short/ShortView.vue'),
    },
    {
      path: '/short/:shortId',
      name: 'short-detail',
      component: () => import('@/views/short-details/ShortDetailView.vue'),
    },
    {
      path: '/watch/:shortId/:episodeId',
      name: 'watch',
      component: () => import('@/views/watching/WatchingView.vue'),
    },
  ],
})

export default router;
