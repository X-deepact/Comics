<template>
  <Table>
    <TableHeader>
      <TableRow>
        <TableHead v-for="column in columns">{{ column.label }}</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <component v-for="subject in subjects" :is="ComicRow" :data="subject" />
    </TableBody>
  </Table>
</template>

<script setup>
import { defineProps } from "vue";
import ComicRow from "../views/content/comics/row.vue";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
const props = defineProps({
  subjects: Array,
  columns: Array,
});
</script>
