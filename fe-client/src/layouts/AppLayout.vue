<template>
  <div>
    <AppHeader v-once />
    <div class="h-[72px]"></div>
    <div class="h-fit min-h-screen">
      <router-view />
    </div>
    <AppFooter v-once />
    <Toast />
  </div>
  <MenuMobileBar v-if="!isLargeScreen" />
</template>

<script setup lang="ts">
import Toast from "primevue/toast";
import AppFooter from "@/components/AppFooter.vue"
import AppHeader from "@/components/AppHeader.vue";
import MenuMobileBar from "@/components/common/MenuMobileBar.vue";
import {useResponsive} from "@/composables/useResponsive.ts";

const { isLargeScreen } = useResponsive();
</script>
