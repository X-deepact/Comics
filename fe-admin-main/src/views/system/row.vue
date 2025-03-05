<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>{{ data.username }}</TableCell>
    <TableCell>{{ data.display_name }}</TableCell>
    <TableCell>{{ data.full_name }}</TableCell>
    <TableCell>{{ data.phone }}</TableCell>
    <TableCell>{{ data.email }}</TableCell>
    <TableCell>{{ data.role_name }}</TableCell>
    <TableCell>{{ data.birthday ? formatDateSafe(data.birthday, false) : data.birthday }}</TableCell>
    <!-- <TableCell>{{ data.active }}</TableCell> -->
    <TableCell>
      <Switch v-model="data.active" :checked="data.active" @click="
        () => {
          $emit('clickActive', data);
        }
      " />
    </TableCell>

    <!-- <TableCell>
      <img :src="data.avatar" alt="Avatar" class="w-10 h-10 rounded-full" />
    </TableCell> -->

    <!-- <TableCell>{{ data.birthday }}</TableCell> -->

    <!-- <TableCell>
      <p class="font-medium">{{ data.created_by_name }}</p>
    </TableCell> -->

    <!-- <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.created_at, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell> -->

    <!-- <TableCell>
      <p class="font-medium">{{ data.updated_by_name }}</p>
    </TableCell> -->

    <!-- <TableCell>
      <div>
        <Label>{{ formatDateSafe(data.updated_at, true) }}</Label>
        <p class="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.updated_at) }}
        </p>
      </div>
    </TableCell> -->

    <!-- <TableCell>{{ data.description }}</TableCell> -->

    <!-- <TableCell>{{ data.first_name }}</TableCell> -->

    <!-- <TableCell>{{ data.last_name }}</TableCell> -->

    <!-- <TableCell>{{ data.tier_code }}</TableCell> -->

    <!-- <TableCell>{{ data.tier_id }}</TableCell> -->

    <!-- Actions -->
    <TableCell>
      <div class="flex gap-3">
        <Button variant="outline" size="icon" @click="
          () => {
            $emit('clickUpdate', {
              id: data.id,
              username: data.username,
              password: data.password,
              phone: data.phone,
              email: data.email,
              FirstName: data.first_name,
              LastName: data.last_name,
              DisplayName: data.display_name,
              TierId: data.tier_id,
              Birthday: data.birthday,
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
import { User } from "@/stores/userStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil } from "lucide-vue-next";
import { Switch } from "@/components/ui/switch";

defineProps<{
  data: User;
}>();
defineEmits(["clickDelete", "clickUpdate", "clickActive"]);

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
