<template>
  <div class="grid grid-cols-6 gap-x-4 gap-y-4">
    <div v-for="i in 12" :key="i">
      <Skeleton width="237px" height="300px" />
      <div class="mt-2 flex flex-col gap-2">
        <Skeleton width="50%" v-for="i in 3" :key="i" />
      </div>
    </div>
  </div>
</template>
<script></script>
