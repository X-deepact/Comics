<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>
      <img :src="data.cover" alt="Comic cover" width="100" preview />
    </TableCell>
    <TableCell>{{ data.name }}</TableCell>
    <TableCell>{{ data.description }}</TableCell>
    <TableCell>{{ data.lang }}</TableCell>
    <TableCell>{{ data.audience }}</TableCell>
    <TableCell>
      <div v-if="data.active"><Badge variant="outline">On</Badge></div>
      <div v-if="!data.active"><Badge variant="destructive">Off</Badge></div>
    </TableCell>
    <TableCell>{{ data.code }}</TableCell>
    <TableCell>
      <p class="font-medium">{{ data.created_by_user.username }}</p>
      <p class="text-muted-foreground text-xs">
        {{ data.created_by_user?.email }}
      </p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at, true) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell
      ><p class="font-medium">{{ data.updated_by_user.username }}</p>
      <p class="text-muted-foreground text-xs">
        {{ data.updated_by_user?.email }}
      </p></TableCell
    >
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
                code: data.code,
                cover: data.cover,
                description: data.description,
                active: data.active,
                language: data.language,
                audience: data.audience,
              });
            }
          "
          ><Pencil
        /></Button>
        <Button
          variant="destructive"
          size="icon"
          @click="
            () => {
              $emit('clickDelete', data);
            }
          "
          ><Trash2
        /></Button>
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
import { Trash2, Pencil } from "lucide-vue-next";
import { Badge } from "@/components/ui/badge";

const API_URL = import.meta.env.VITE_API_URL;
defineProps<{
  data: Comic;
}>();
defineEmits(["clickDelete", "clickUpdate"]);
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
