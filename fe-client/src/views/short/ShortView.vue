<template>
  <Short />
</template>
<script setup lang="ts">
import Short from "@/components/module/short/Short.vue";
import {useTitle} from "@vueuse/core";
import {i18n} from "@/main.ts";

useTitle(`${i18n.global.t("short.pageTitle")}`)
</script>
