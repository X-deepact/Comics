<template>
  <RecentlyUpdate />
</template>
<script setup lang="ts">
import RecentlyUpdate from "@/components/module/recently-update/RecentlyUpdate.vue";
import {useTitle} from "@vueuse/core";
import {i18n} from "@/main.ts";

useTitle(`${i18n.global.t("recentlyUpdate.pageTitle")}`)
</script>
