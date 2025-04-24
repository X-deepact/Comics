<template>
  <About />
</template>
<script setup lang="ts">
import About from "@/components/module/about/About.vue";
import {useTitle} from "@vueuse/core";
import {i18n} from "@/main.ts";

useTitle(`${i18n.global.t("footer.aboutUs")}`)
</script>
