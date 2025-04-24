<template>
  <div class="flex flex-col items-center justify-center bg-white text-center">
    <div>
      <img class="w-full" src="@/assets/images/not-found.png" alt="404" />
      <p class="message">{{ $t("message.pageNotFound") }}</p>
      <p>{{ $t("message.pageNotFoundMessage") }}</p>
      <Button class="pt-10" :label="$t('message.backToHome')" @click="goHome" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
const router = useRouter();
const goHome = () => {
  router.push("/");
};
</script>

<style scoped>
.message {
  font-size: 1.5rem;
  margin: 0.5rem 0;
}
</style>
