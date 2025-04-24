<template>
  <div class="fixed bottom-0 w-full z-10 bg-white dark:bg-gray-900 shadow-lg border-t border-gray-200 dark:border-gray-900">
    <div class="flex">
      <Button class="flex-1 p-3 sm:p-5 hover:text-orange-500"
              v-for="item in items"
              :key="item.label"
              :icon="item.icon"
              icon-pos="left"
              :class="getButtonClass(item.action)"
              @click="item.command()"
      >
        {{ item.label }}
      </Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { useRoute } from "vue-router";
import {i18n} from "@/main.ts";

const router = useRouter();
const route = useRoute();

const items = computed(() => [
  {
    label: `${i18n.global.t("navigator.homepage")}`,
    icon: 'pi pi-home',
    action: '/',
    command: () => router.push('/'),
    class: route.path === '/' ? 'active-nav-item' : ''
  },
  {
    label: `${i18n.global.t("navigator.updates")}`,
    icon: 'pi pi-refresh',
    action: "/recently-updated",
    command: () => router.push('/recently-updated'),
    class: route.path === '/recently-updated' ? 'active-nav-item' : ''
  },
  {
    label: `${i18n.global.t("navigator.categories")}`,
    icon: 'pi pi-list',
    action: "/category",
    command: () => router.push('/category'),
    class: route.path === '/category' ? 'active-nav-item' : ''
  },
  {
    label: `${i18n.global.t("navigator.short")}`,
    icon: 'pi pi-home',
    action: '/short',
    command: () => router.push('/short'),
    class: route.path === '/short' ? 'active-nav-item' : ''
  },
]);

const getButtonClass = (action: string) => {
  return route.path === action
      ? "text-orange-500 border-b-2 border-orange-500"
      : "text-black hover:text-orange-500 dark:text-white";
};

</script>
