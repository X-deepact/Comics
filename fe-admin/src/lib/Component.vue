<template>
  <div class="flex justify-between items-center mb-6">
    <h1 class="text-2xl font-semibold">{{ title }} Management</h1>
  </div>
  <div class="flex justify-between items-center mb-4 gap-4">
    <div class="flex items-center gap-4">
      <Input
        placeholder="Search..."
        class="w-[280px]"
        @input="$emit('update:search', $event.target.value)"
      />
      <Button
        @click="$emit('clickAdd')"
        class="bg-blue-500 hover:bg-blue-600 text-white"
        >Add {{ title }}</Button
      >
      <Button
        @click="$emit('clickRefresh')"
        class="bg-blue-500 hover:bg-blue-600 text-white"
        size="icon"
        ><RotateCw />
      </Button>
    </div>
    <div class="flex items-center gap-4">
      <div class="flex items-center gap-4">
        <Select
          :value="pageSize"
          @update:modelValue="
            (value: string) => {
              $emit('update:pageSize', Number(value));
            }
          "
        >
          <SelectTrigger class="w-[70px]">
            <SelectValue :placeholder="`${pageSize}`" />
          </SelectTrigger>
          <SelectContent class="bg-white">
            <SelectItem value="10">10</SelectItem>
            <SelectItem value="20">20</SelectItem>
            <SelectItem value="30">30</SelectItem>
            <SelectItem value="50">50</SelectItem>
          </SelectContent>
        </Select>
        items / page
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { RotateCw } from "lucide-vue-next";
defineProps({
  title: String,
  pageSize: Number,
  onAddClick: Function,
});

const emit = defineEmits([
  "update:search",
  "clickAdd",
  "update:pageSize",
  "clickRefresh",
]);
</script>
