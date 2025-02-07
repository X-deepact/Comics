import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "../stores/authStore";

// Layout
const DashboardLayout = () => import("../layouts/DashboardLayout.vue");

// Views
const Login = () => import("../views/auth/Login.vue");
const Dashboard = () => import("../views/Dashboard.vue");
const Comics = () => import("../views/content/Comics.vue");
const Subjects = () => import("../views/content/Subjects.vue");
const Ads = () => import("../views/ads/Ads.vue");
const Recommendation = () => import("../views/ads/Recommendation.vue");
const Profile = () => import("../views/profile/Profile.vue")
const Users = () => import("../views/system/Users.vue");

const routes = [
  { 
    path: "/", 
    redirect: "/dashboard" 
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
    ]
  },
  {
    path: "/:pathMatch(.*)*",
    name: "not-found",
    component: () => import("../views/NotFound.vue")
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
