<template>
  <div class="flex items-center justify-between mt-4">
    <div class="flex items-center gap-4">
      <span class="text-sm">total {{ totalItems }} items</span>
    </div>
    <div class="flex items-center gap-2">
      <Button
        variant="outline"
        size="icon"
        :disabled="currentPage === 1"
        @click="$emit('update:currentPage', currentPage - 1)"
      >
        <ChevronLeft />
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
        size="icon"
        :disabled="currentPage === totalPages"
        @click="$emit('update:currentPage', currentPage + 1)"
      >
        <ChevronRight />
      </Button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits } from "vue";
import { Button } from "@/components/ui";
import { Input } from "@/components/ui/input";
import { ChevronLeft, ChevronRight } from "lucide-vue-next";
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
