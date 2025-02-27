<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>
      <img :src="data.image" alt="Ad image" width="100" preview />
    </TableCell>
    <TableCell>{{ data.title }}</TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.active_from) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.active_from) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.active_to) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.active_to) }}
        </p>
      </div>
    </TableCell>
    <TableCell>{{ data.type }}</TableCell>
    <TableCell>{{ data.direct_url }}</TableCell>
    <TableCell>
      <div>
        <Badge 
          :variant="data.status?.toLowerCase() === 'active' ? 'outline' : 'destructive'"
        >
          {{ data.status === 'active' ? 'Active' : 'Inactive' }}
        </Badge>
      </div>
    </TableCell>
    <TableCell>
      <div class="flex gap-3">
        <Button
          variant="outline"
          size="icon"
          @click="$emit('clickUpdate', data)"
        >
          <Pencil />
        </Button> 
        <Button
          variant="destructive"
          size="icon"
          @click="$emit('clickDelete', data)"
        >
          <Trash2 />
        </Button>
      </div>
    </TableCell>
  </TableRow>
</template>

<script setup lang="ts">
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2, Pencil } from "lucide-vue-next";
import { formatDate, getTimeAgo } from "@/lib/utils";
import type { Ad } from "@/stores/adsStore";

defineProps<{
  data: Ad;
}>();
defineEmits(["clickDelete", "clickUpdate"]);

const formatDateSafe = (date: Date | string | number | null): string => {
  if (!date) return 'N/A';
  try {
    return formatDate(date, true);
  } catch (error) {
    console.error("Date formatting error:", error);
    return "Invalid date";
  }
};

const getTimeAgoSafe = (date: Date | string | number | null): string => {
  if (!date) return 'N/A';
  try {
    return getTimeAgo(date);
  } catch (error) {
    console.error("Time ago formatting error:", error);
    return "Invalid date";
  }
};
</script> 