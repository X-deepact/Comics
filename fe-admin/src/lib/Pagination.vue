<template>
  <div class="flex items-center justify-between mt-4">
    <div class="flex items-center gap-4">
      <span class="text-sm">total {{ totalItems }} items</span>
    </div>
    <div class="flex items-center gap-2">
      <Button
        variant="outline"
        :disabled="currentPage === 1"
        class="h-[40px] w-[70px]"
        @click="$emit('update:currentPage', currentPage - 1)"
      >
        Previous
      </Button>
      <Input
        type="number"
        v-model="props.currentPage"
        :min="1"
        :max="totalPages"
        class="text-center h-[40px] w-[70px]"
        @input="$emit('update:currentPage', $event.target.value)"
      />
      /{{ totalPages }} pages
      <Button
        variant="outline"
        :disabled="currentPage === totalPages"
        class="h-[40px] w-[70px]"
        @click="$emit('update:currentPage', currentPage + 1)"
      >
        Next
      </Button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits } from "vue";
import { Button } from "@/components/ui";
import { Input } from "@/components/ui/input";
const props = defineProps({
  currentPage: {
    type: Number,
    required: true,
  },
  totalPages: {
    type: Number,
    required: true,
  },
  totalItems: { type: Number, required: true },
});
defineEmits(["update:currentPage"]);
</script>
