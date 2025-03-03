<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>{{ data.title }}</TableCell>
    <TableCell>
      <img v-if="data.image" :src="data.image" alt="Ad Image" class="w-20 h-20 object-cover" />
      <span v-else>No image</span>
    </TableCell>
    <TableCell>{{ formatDateSafe(data.active_from) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_to) }}</TableCell>
    <TableCell>{{ data.type }}</TableCell>
    <TableCell>{{ data.direct_url }}</TableCell>
    <TableCell>
      <Badge variant="secondary">N/A</Badge>
    </TableCell>
    <TableCell>
      <p class="font-medium">{{ data.created_by_name }}</p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at) }}</Label>
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
        <Label>{{ formatDateSafe(data.updated_at) }}</Label>
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
                image: data.image,
                active_from: data.active_from,
                active_to: data.active_to,
                type: data.type,
                direct_url: data.direct_url,
                comic_id: data.comic_id,
                status: data.status,
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
import { Ad, useAdStore } from "@/stores/adStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2, Pencil } from "lucide-vue-next";

defineProps<{
  data: Ad;
}>();
defineEmits(["clickDelete", "clickUpdate"]);

const formatDateSafe = (
  date: Date | string | number | null | undefined,
  includeTime: boolean = true
): string => {
  if (!date) return "N/A";
  
  try {
    return formatDate(date, includeTime);
  } catch (error) {
    console.error("Date formatting error:", error);
    return "Invalid date";
  }
};

const getTimeAgoSafe = (date: Date | string | number | null | undefined): string => {
  if (!date) return "N/A";
  
  try {
    return getTimeAgo(date);
  } catch (error) {
    console.error("Time ago formatting error:", error);
    return "Invalid date";
  }
};

const adStore = useAdStore();
console.log('Current ad data:', adStore.adData);
</script> 