<template>
  <div class="history-popup absolute top-14 right-5 mt-2">
    <div
      class="popup-box bg-white shadow-md rounded-lg border-t-4 border-orange-600"
    >
      <div class="flex justify-between items-center p-4">
        <span class="text-gray-700">{{ $t("navigator.guest") }}</span>
        <span class="text-orange-500 text-sm cursor-pointer hover:underline">
          {{ $t("navigator.loginGetMore") }}
        </span>
      </div>

      <div
        class="flex flex-col items-center text-center text-gray-400 p-4 pt-0"
      >
        {{ $t("navigator.VIPExpired") }} : {{ $t("navigator.trialExpired") }}
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style scoped></style>
