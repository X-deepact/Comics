<template>
  <TableRow>
    <TableCell>{{ data.number }}</TableCell>
    <TableCell>{{ data.name }}</TableCell>
    <TableCell>{{ data.active }}</TableCell>
    <TableCell>{{ data.cover }}</TableCell>
    <TableCell>
      <p class="font-medium">{{ data.created_by_name }}</p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at, true) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <p class="font-medium">{{ data.updated_by_name }}</p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.updated_at, true) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.updated_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div class="flex gap-3">
        <Button variant="outline" size="icon" @click="
          () => {
            $emit('clickAction', data.id);
          }
        ">
          <FilePenLine />
        </Button>
        <Button variant="outline" size="icon" @click="
          () => {
            $emit('clickUpdate', {
              id: data.id,
              name: data.name,
              number: data.number,
              active: data.active,
              comic_id: data.comic_id,
              cover: data.cover,
              birth_day: data.birth_day,
            });
          }
        ">
          <Pencil />
        </Button>
        <Button variant="destructive" size="icon" @click="
          () => {
            $emit('clickDelete', data);
          }
        ">
          <Trash2 />
        </Button>
      </div>
    </TableCell>
  </TableRow>
</template>

<script setup lang="ts">
import { Chapter } from "@/stores/chapterStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil, FilePenLine } from "lucide-vue-next";

defineProps<{
  data: Chapter;
}>();
defineEmits(["clickDelete", "clickUpdate", "clickAction"]);
const formatDateSafe = (
  date: Date | string | number,
  includeTime: boolean
): string => {
  try {
    return formatDate(date, includeTime);
  } catch (error) {
    console.error("Date formatting error:", error);
    return "Invalid date";
  }
};

const getTimeAgoSafe = (date: Date | string | number): string => {
  try {
    return getTimeAgo(date);
  } catch (error) {
    console.error("Time ago formatting error:", error);
    return "Invalid date";
  }
};
</script>
