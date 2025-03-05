<script setup lang="ts">
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "./stores/authStore";
import Toaster from "@/components/ui/toast/Toaster.vue";
const router = useRouter();
const authStore = useAuthStore();

onMounted(async () => {
  // Check if there's a token in localStorage or wherever you store it
  if (authStore.token) {
    try {
      await authStore.fetchProfile();
    } catch (error) {
      // If token is invalid or expired, clear auth and redirect to login
      authStore.logout();
      router.push("/login");
    }
  }
});
</script>

<template>
  <div class="w-full">
    <router-view></router-view>
  </div>
  <Toaster />
</template>

<style>
html,
body {
  width: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

#app {
  width: 100%;
}
</style>
