<template>
  <div class="flex justify-between items-center mb-6">
    <h1 class="text-2xl font-semibold">{{ title }} Management</h1>
  </div>
  <div class="flex justify-between items-center mb-4 gap-4">
    <Input
      placeholder="Search..."
      class="w-[280px]"
      @input="$emit('update:search', $event.target.value)"
    />
    <div class="flex items-center gap-4">
      <div>
        <span class="text-sm">total {{ totalItems }} items</span>
      </div>

      <Button
        @click="onAddClick"
        class="bg-blue-500 hover:bg-blue-600 text-white"
        >Add {{ title }}</Button
      >
    </div>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, watch } from "vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

const props = defineProps({
  title: String,
  totalItems: Number,
  onAddClick: Function,
});

const emit = defineEmits(["update:search"]);
</script>
