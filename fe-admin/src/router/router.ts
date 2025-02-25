import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "../stores/authStore";

// Layout
const DashboardLayout = () => import("../layouts/DashboardLayout.vue");

// Views
const Login = () => import("../views/auth/Login.vue");
const Dashboard = () => import("../views/Dashboard.vue");
const Comics = () => import("../views/content/comics/Comics.vue");
const Subjects = () => import("../views/content/subjects/Subjects.vue");
const Authors = () => import("../views/content/authors/Authors.vue");
const Ads = () => import("../views/ads/Ads.vue");
const Recommendation = () => import("../views/ads/Recommendation.vue");
const Profile = () => import("../views/profile/Profile.vue");
const Users = () => import("../views/system/Users.vue");

const routes = [
  {
    path: "/",
    redirect: "/dashboard",
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/dashboard",
    component: DashboardLayout,
    children: [
      {
        path: "",
        name: "dashboard",
        component: Dashboard,
      },
      // Content Management
      {
        path: "comics",
        name: "comics",
        component: Comics,
      },
      {
        path: "subjects",
        name: "subjects",
        component: Subjects,
      },
      {
        path: "authors",
        name: "authors",
        component: Authors,
      },
      // Ads and Recommendation
      {
        path: "ads",
        name: "ads",
        component: Ads,
      },
      {
        path: "recommendation",
        name: "recommendation",
        component: Recommendation,
      },
      // System Management
      {
        path: "users",
        name: "users",
        component: Users,
      },
      // Profile
      {
        path: "profile",
        name: "profile",
        component: Profile,
      },
    ],
  },
  {
    path: "/:pathMatch(.*)*",
    name: "not-found",
    component: () => import("../views/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
router.beforeEach(async (to) => {
  const authStore = useAuthStore();
  const publicPages = ["/login"];
  const authRequired = !publicPages.includes(to.path);
  if (authRequired && !authStore.user) {
    return "/login";
  }
});
export default router;
