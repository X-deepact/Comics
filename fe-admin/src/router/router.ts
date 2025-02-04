import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Login from "../views/Login.vue";
import { useAuthStore } from "../stores/authStore";

const routes = [
  { 
    path: "/", 
    redirect: "/dashboard" 
  },
  // { 
  //   path: "/login", 
  //   name: "login",
  //   component: Login,
  //   meta: { requiresAuth: false },
  //   beforeEnter: (to: any, from: any, next: any) => {
  //     const authStore = useAuthStore();
  //     if (authStore.token) {
  //       next('/dashboard');
  //     } else {
  //       next();
  //     }
  //   }
  // },
  { 
    path: "/dashboard", 
    name: "dashboard",
    component: Dashboard,
    // meta: { requiresAuth: true }
  },
  // Catch-all route for 404
  {
    path: "/:pathMatch(.*)*",
    redirect: "/dashboard"
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Global navigation guard
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();
  const requiresAuth = to.meta.requiresAuth;

  // If route requires auth and no token is present
  if (requiresAuth && !authStore.token) {
    next('/login');
    return;
  }

  // If route requires auth and token is present, verify token validity
  if (requiresAuth && authStore.token) {
    try {
      await authStore.fetchProfile();
      next();
    } catch (error) {
      authStore.logout();
      next('/login');
    }
    return;
  }

  // For all other cases, allow navigation
  next();
});

export default router;
