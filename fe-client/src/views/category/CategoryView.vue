<template>
  <Category />
</template>
<script setup lang="ts">
import Category from "@/components/module/category/Category.vue";
import {useTitle} from "@vueuse/core";
import {i18n} from "@/main.ts";

useTitle(`${i18n.global.t("category.pageTitle")}`)
</script>
