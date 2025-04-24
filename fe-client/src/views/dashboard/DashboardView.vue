<template>
  <Dashboard />
</template>
<script setup lang="ts">
import Dashboard from "@/components/module/dashboard/Dashboard.vue";
import {useTitle} from "@vueuse/core";
import {i18n} from "@/main.ts";

useTitle(`${i18n.global.t("dashboard.pageTitle")}`)
</script>
