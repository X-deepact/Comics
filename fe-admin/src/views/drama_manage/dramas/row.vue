<template>
  <TableRow>
    <TableCell>
      <img :src="data.thumb" alt="Comic cover" width="100" preview />
    </TableCell>
    <TableCell>
      <div class="flex flex-col gap-1">
        <div v-for="translation in data.translations"
          class="flex flex-col gap-2 max-w-[400px] border rounded-lg p-2 gap-1">
          <div class="flex flex-row justify-between items-center gap-2">
            <Label>{{ translation.title }}</Label>
            <Badge variant="outline">{{ langConverter(translation.language) }}</Badge>
          </div>
          <p class="text-xs">{{ translation.description }}</p>
        </div>
      </div>

    </TableCell>
    <TableCell>
      <div class="flex flex-col gap-1">
        <div v-for="genre in data.genres">
          <Badge variant="outline" class="text-xs">{{ genre.name }}</Badge>
        </div>
      </div>
    </TableCell>
    <TableCell>
      <div class="flex flex-col gap-1">
        {{ formatDateSafe(data.release_date, false) }}
      </div>
    </TableCell>
    <TableCell>
      <Badge variant="secondary" class="text-sm">{{
        data.created_by_name
        }}</Badge>
      <div>
        <Label class="text-xs">{{
          formatDateSafe(data.created_at, true)
        }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <Badge variant="secondary" class="text-sm">{{
        data.updated_by_name
        }}</Badge>
      <div>
        <Label class="text-xs">{{
          formatDateSafe(data.updated_at, true)
        }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.updated_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <Switch v-model="data.active" :checked="data.active" @click="
        () => {
          $emit('clickActive', data);
        }
      " />
    </TableCell>
    <TableCell>
      <div class="flex gap-3">
        <Button variant="outline" size="icon" @click="
          () => {
            $emit('clickAction', data);
          }
        ">
          <FilePenLine />
        </Button>
        <Button variant="outline" size="icon" @click="
          () => {
            $emit('clickUpdate', data);
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
import { Drama } from "@/stores/dramaStore";
import { formatDate, getTimeAgo, langConverter } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil, FilePenLine } from "lucide-vue-next";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
defineProps<{
  data: Drama;
}>();
defineEmits(["clickDelete", "clickUpdate", "clickAction", "clickActive"]);
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
