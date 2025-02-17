<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>
      <Image
        :src="`${API_URL}${data.cover}`"
        alt="Comic cover"
        width="50"
        preview
      />
    </TableCell>
    <TableCell>{{ data.title }}</TableCell>
    <TableCell>{{ data.description }}</TableCell>
    <TableCell>{{ data.active }}</TableCell>
    <TableCell>{{ data.code }}</TableCell>
    <TableCell>{{ data.created_by }}</TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at, true) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>{{ data.updated_by }}</TableCell>
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
        <Button variant="outline" class="w-20">Edit</Button>
        <Button variant="destructive" class="w-20">Delete</Button>
      </div>
    </TableCell>
  </TableRow>
</template>

<script setup lang="ts">
import { Comic } from "@/stores/comicStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
const API_URL = import.meta.env.VITE_API_URL;
defineProps<{
  data: Comic;
}>();

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
