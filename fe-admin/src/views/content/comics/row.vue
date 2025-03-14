<template>
  <TableRow>
    <TableCell>
      <img :src="data.cover" alt="Comic cover" width="100" preview />
    </TableCell>
    <TableCell>
      <div class="flex flex-col gap-2 max-w-[200px]">
        <Label>{{ data.name }}</Label>
        <p class="text-xs">{{ data.description }}</p>
      </div>
    </TableCell>
    <TableCell>
      <div class=" flex flex-col gap-1">
        <div v-for="author in data.authors">
          <Badge variant="outline" class="text-xs">{{ author.name }}</Badge>
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
      <div class="flex flex-col gap-1 min-w-[320px]">
        <div>
          <Label>Code: </Label>
          <Badge variant="outline" class="text-xs">
            {{ data.code }}
          </Badge>
        </div>
        <div>
          <Label>Original Language: </Label>
          <Badge variant="outline" class="text-xs">
            {{ langConverter(data.original_language) }}
          </Badge>
        </div>
        <div>
          <Label>Language: </Label>
          <Badge variant="outline" class="text-xs">
            {{ langConverter(data.lang) }}
          </Badge>
        </div>
        <div>
          <Label>Audience: </Label>
          <Badge variant="outline" class="text-xs">{{ data.audience }}</Badge>
        </div>
        <div>
          <Label>Status: </Label>
          <Badge variant="outline" class="text-xs">{{ data.status }}</Badge>
        </div>
      </div>
    </TableCell>
    <TableCell>
      <Badge variant="secondary" class="text-sm">{{
        data.created_by_user.username
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
        data.updated_by_user.username
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
              code: data.code,
              cover: data.cover,
              description: data.description,
              active: data.active,
              lang: data.lang,
              original_language: data.original_language,
              audience: data.audience,
              authors: data.authors,
              genres: data.genres,
              status: data.status,
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
import { Comic } from "@/stores/comicStore";
import { formatDate, getTimeAgo, langConverter } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil, FilePenLine } from "lucide-vue-next";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
defineProps<{
  data: Comic;
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
