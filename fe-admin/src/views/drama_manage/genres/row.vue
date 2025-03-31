<template>
  <TableRow>
    <TableCell>{{ data.id || 'N/A' }}</TableCell>
    <TableCell>{{ data.name || 'N/A' }}</TableCell>
    <TableCell>{{ data.position || 'N/A' }}</TableCell>
    <TableCell>
      <div v-if="data.translations && data.translations.length > 0" class="flex flex-col gap-1">
        <div v-for="translation in data.translations" :key="translation.id" class="text-sm">
          <Badge>{{ translation.language }}</Badge>
          <span class="ml-2">{{ translation.translated_name }}</span>
        </div>
      </div>
      <span v-else class="text-gray-500">No translations</span>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.updated_at) }}</Label>
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
import { Genre, useDramaGenreStore } from "@/stores/drama_genreStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2, Pencil } from "lucide-vue-next";

const props = defineProps<{
  data: Genre;
}>();

defineEmits(["clickDelete", "clickUpdate"]);

const formatDateSafe = (date: string | null | undefined): string => {
  if (!date) return 'N/A';
  try {
    return formatDate(date);
  } catch {
    return 'N/A';
  }
};

const getTimeAgoSafe = (date: string | null | undefined): string => {
  if (!date) return 'N/A';
  try {
    return getTimeAgo(date);
  } catch {
    return 'N/A';
  }
};
</script> 