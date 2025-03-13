<template>
  <TableRow>
    <TableCell>{{ data.id || 'N/A' }}</TableCell>
    <TableCell>{{ data.title || 'N/A' }}</TableCell>
    <TableCell>
      <img 
        v-if="data.image" 
        :src="data.image" 
        alt="Ad Image"
        class="w-12 h-12 xs:w-14 xs:h-14 sm:w-16 sm:h-16 md:w-20 md:h-20 lg:w-24 lg:h-24 object-cover rounded-md transition-all duration-300" 
      />
      <span v-else class="text-xs xs:text-sm sm:text-base text-gray-500">No image</span>
    </TableCell>
    <TableCell>{{ formatDateSafe(data.active_from) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_to) }}</TableCell>
    <TableCell>{{ data.type || 'N/A' }}</TableCell>
    <TableCell>{{ data.direct_url || 'N/A' }}</TableCell>
    <TableCell> 
      <Badge :variant="data.status === 'active' ? 'success' : 'secondary'">
        {{ data.status || 'N/A' }}
      </Badge>
    </TableCell>
    <TableCell>
      <p v-if="data.created_by_name" class="font-medium">{{ data.created_by_name }}</p>
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
      <p v-if="data.updated_by_name" class="font-medium">{{ data.updated_by_name }}</p>
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
import { Ad, useAdStore } from "@/stores/adStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2, Pencil } from "lucide-vue-next";
import { toast } from "@/components/ui/toast/use-toast";

const props = defineProps<{
  data: Ad;
}>();

defineEmits(["clickDelete", "clickUpdate"]);

const adStore = useAdStore();

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

const updateStatus = async (id: number, status: 'active' | 'inactive') => {
  try {
    await adStore.updateAdStatus(id, status);
    adStore.getAdData();
  } catch (error: any) {
    toast({
      description: error.response?.data?.message || error.message,
      variant: "destructive",
    });
  }
};
</script> 