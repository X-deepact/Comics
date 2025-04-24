<template>
  <header class="fixed top-0 left-0 w-full z-50 bg-white dark:bg-gray-900 shadow-md flex justify-between items-center p-4">
    <div class="flex gap-10">
      <div v-if="!showMobileSearch" class="flex gap-3 items-center relative">
        <i v-if="isSmallAndMiddleScreen" class="pi pi-search text-[#ff620e] ml-2" @click="toggleMobileSearch" />
        <div class="logo">
          <img
              src="@/assets/images/logo.png"
              alt="漫血屋 Logo"
              class="h-10 cursor-pointer"
              @click="router.push('/')"
          />
        </div>
      </div>
      <div v-if="showMobileSearch" class="flex gap-3 items-center relative">
        <i class="pi pi-angle-left text-[#ff620e]" @click="toggleMobileSearch"/>
      </div>
      <nav v-if="isLargeScreen" class="flex gap-4">
        <Button
          size="large"
          v-for="button in buttons"
          :key="button.label"
          :label="$t(button.label)"
          :class="getButtonClass(button.action)"
          @click="handleButtonClick(button.action)"
        />
      </nav>
    </div>

    <SearchBox v-if="isLargeScreen || showMobileSearch" />

    <div class="flex gap-3 items-center relative">
      <AppLocaleSwitch v-if="isLargeScreen" />
      <div v-if="isLargeScreen" class="flex gap-3 relative">
        <Avatar
          icon="pi pi-user"
          class="mr-2 cursor-pointer"
          size="large"
          shape="circle"
          @mouseover="handleMouseOver('profile')"
        />
      </div>
    </div>
    <AppLocaleSwitch v-if="isSmallAndMiddleScreen && (!showMobileSearch)" />
  </header>

  <div class="border-b-2 border-gray-200"></div>

</template>

<script lang="ts" setup>
import {ref, onMounted, onUnmounted, computed} from "vue";
import { useRouter, useRoute } from "vue-router";
import AppLocaleSwitch from "@/components/AppLocaleSwitch.vue";
import SearchBox from "@/components/module/search/SearchBox.vue";
import { useResponsive } from "@/composables/useResponsive";

interface ButtonItem {
  label: string;
  action: string;
}
const { isLargeScreen, isSmallAndMiddleScreen } = useResponsive();
const showMobileSearch = ref(false);

const router = useRouter();
const route = useRoute();
const searchContainer = ref<HTMLElement | null>(null);
const showHotSearch = ref(false);

const showProfilePopup = ref(false);

const buttons = ref<ButtonItem[]>([
  { label: "navigator.homepage", action: "/" },
  { label: "navigator.updates", action: "/recently-updated" },
  { label: "navigator.categories", action: "/category" },
  { label: "navigator.short", action: "/short" },
]);

const getButtonClass = (action: string) => {
  return route.path === action
    ? "text-orange-500 border-b-2 border-orange-500"
    : "text-black hover:text-orange-500 dark:text-white";
};

const handleButtonClick = (action: string) => {
  router.push(action);
};

const handleMouseOver = (action: string) => {
  showProfilePopup.value = action === "profile";
};

const handleClickOutside = (event: MouseEvent) => {
  if (
    searchContainer.value &&
    !searchContainer.value.contains(event.target as Node)
  ) {
    showHotSearch.value = false;
  }
};

const toggleMobileSearch = () => {
  showMobileSearch.value = !showMobileSearch.value;
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>
