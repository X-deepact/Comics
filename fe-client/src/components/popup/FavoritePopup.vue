<template>
  <div class="history-popup absolute top-14 right-15 mt-2">
    <div
      class="popup-box bg-white shadow-md rounded-lg border-t-4 border-orange-600"
    >
      <div class="flex justify-between items-center p-4">
        <span class="text-gray-700">{{ $t("navigator.myFavorites") }}</span>
        <span class="text-orange-500 text-sm cursor-pointer hover:underline">
          {{ $t("navigator.viewAll") }}<i class="pi pi-chevron-right"></i>
        </span>
      </div>

      <div
        class="flex flex-col items-center text-center text-gray-400 p-4 pt-0"
      >
        <i class="pi pi-globe text-gray-400 py-4" style="font-size: 3rem"></i>
        <p>
          {{ $t("navigator.logInHistory") }}，<span
            class="text-orange-500 cursor-pointer hover:underline"
            >{{ $t("navigator.goLogin") }}
          </span>
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style scoped></style>
