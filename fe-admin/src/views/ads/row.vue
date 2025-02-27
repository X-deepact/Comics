<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>{{ data.title }}</TableCell>
    <TableCell>
      <img 
        :src="data.cover" 
        alt="Recommend cover" 
        width="100" 
        class="object-cover h-[60px]"
      />
    </TableCell>
    <TableCell>{{ data.position }}</TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.active_from, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.active_from) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.active_to, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.active_to) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.updated_at, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.updated_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div class="flex gap-3">
        <Button
          variant="outline"
          size="icon"
          @click="$emit('clickUpdate', data)"
        >
          <Pencil class="h-4 w-4" />
        </Button>
        <Button
          variant="destructive"
          size="icon"
          @click="$emit('clickDelete', data)"
        >
          <Trash2 class="h-4 w-4" />
        </Button>
      </div>
    </TableCell>
  </TableRow>
</template>

<script setup lang="ts">
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil } from "lucide-vue-next";
import { Recommend } from "@/stores/recommendStore";

defineProps<{
  data: Recommend;
}>();

defineEmits(["clickDelete", "clickUpdate"]);

const formatDateSafe = (date: Date | string | number | undefined, includeTime: boolean): string => {
  if (!date) return "N/A";
  try {
    return formatDate(date, includeTime);
  } catch (error) {
    console.error("Date formatting error:", error);
    return "N/A";
  }
};

const getTimeAgoSafe = (date: Date | string | number | undefined): string => {
  if (!date) return "N/A";
  try {
    return getTimeAgo(date);
  } catch (error) {
    console.error("Time ago formatting error:", error);
    return "N/A";
  }
};
</script> 