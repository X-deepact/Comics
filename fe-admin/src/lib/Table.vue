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
        :is="ComicRow"
        :data="subject"
        @clickDelete="
          (data) => {
            $emit('clickDelete', data);
          }
        "
        @clickUpdate="
          (data) => {
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
import ComicRow from "../views/content/comics/row.vue";
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
  columns: Array,
  isLoading: Boolean,
});
defineEmits(["clickDelete", "clickUpdate"]);
</script>
