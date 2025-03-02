<template>
  <TableRow>
    <TableCell>{{ data.title }}</TableCell>
    <TableCell>
      <img 
        v-if="data.cover" 
        :src="data.cover" 
        alt="cover" 
        class="w-20 h-20 object-cover" 
        @error="handleImageError"
      />
      <span v-else>No image</span>
    </TableCell>
    <TableCell>{{ data.position }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_from) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_to) }}</TableCell>
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
        <Button
          variant="outline"
          size="icon"
          @click="
            () => {
              $emit('clickUpdate', {
                id: data.id,
                title: data.title,
                cover: data.cover,
                position: data.position,
                active_from: data.active_from,
                active_to: data.active_to,
              });
            }
          "
        >
          <Pencil />
        </Button>
        <Button
          variant="destructive"
          size="icon"
          @click="
            () => {
              $emit('clickDelete', data);
            }
          "
        >
          <Trash2 />
        </Button>
      </div>
    </TableCell>
  </TableRow>
</template>

<script setup lang="ts">
import { Recommend } from "@/stores/recommendStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil } from "lucide-vue-next";
import { ref } from 'vue';

defineProps<{
  data: Recommend;
}>();
defineEmits(["clickDelete", "clickUpdate"]);

const formatDateSafe = (
  date: Date | string | number | null | undefined,
  includeTime = false
): string => {
  if (!date) return "N/A";
  try {
    return formatDate(date, includeTime);
  } catch (error) {
    console.error("Date formatting error:", error);
    return "N/A";
  }
};

const getTimeAgoSafe = (date: Date | string | number | null | undefined): string => {
  if (!date) return "N/A";
  try {
    return getTimeAgo(date);
  } catch (error) {
    console.error("Time ago formatting error:", error);
    return "N/A";
  }
};

// Add this function to handle image loading errors
const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement;
  target.src = ''; // You could set a default image here
  target.alt = 'Image not available';
};
</script> 