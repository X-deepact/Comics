<template>
  <div
    ref="scrollTopButton"
    class="sticky w-full flex justify-end bottom-10 pb-3 pr-5 transition"
  >
    <Button
      @click="scrollToTop"
      icon="pi pi-arrow-up"
      class="w-10 h-10 rounded-full border-1 text-white bg-[#ff620e] font-bold transition"
    />
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref} from "vue";

const scrollTopButton = ref<HTMLButtonElement | null>(null);
const handleScroll = () => {
  if (window.scrollY > 0) {
    scrollTopButton.value?.classList.remove("invisible");
  } else {
    scrollTopButton.value?.classList.add("invisible");
  }
};

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: "smooth" });
}

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
</script>
