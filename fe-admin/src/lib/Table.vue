<template>
  <Table>
    <TableHeader>
      <TableRow>
        <TableHead v-for="column in columns">{{ column.label }}</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <component
        v-if="!isLoading"
        v-for="subject in subjects"
        :is="row"
        :data="subject"
        @clickDelete="
          (data: any) => {
            $emit('clickDelete', data);
          }
        "
        @clickUpdate="
          (data: any) => {
            $emit('clickUpdate', data);
          }
        "
      />
    </TableBody>
  </Table>
  <div v-if="isLoading" class="flex items-center justify-center py-10 gap-3">
    <img :src="loadingImg" class="w-[40px] h-[40px]" />
    <Label class="justify-self-center text-3xl"> Loading...</Label>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import loadingImg from "@/assets/loading.svg";
import {
  Table,
  TableBody,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Label } from "@/components/ui/label";
defineProps({
  subjects: Array,
  columns: Object,
  isLoading: Boolean,
  row: Object,
});
defineEmits(["clickDelete", "clickUpdate"]);
</script>
