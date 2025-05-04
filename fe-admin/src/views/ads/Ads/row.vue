<template>
  <TableRow>
    <TableCell>{{ data.id || 'N/A' }}</TableCell>
    <TableCell>{{ data.title || 'N/A' }}</TableCell>
    <TableCell>
      <img 
        v-if="data.image"
        :src="getImageUrl(data.image)" 
        alt="Ad Image" 
        width="100"
        preview
      />
      <span v-else class="text-xs xs:text-sm sm:text-base text-gray-500">No image</span>
    </TableCell>
    <TableCell>{{ formatDateSafe(data.active_from) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_to) }}</TableCell>
    <TableCell>{{ data.type || 'N/A' }}</TableCell>
    <TableCell>
      <div class="w-[200px] flex items-center gap-2" :title="data.direct_url">
        <span class="truncate">{{ truncateUrl(data.direct_url) }}</span>
        <Button 
          v-if="data.direct_url" 
          variant="ghost" 
          size="icon" 
          class="h-6 w-6 flex-shrink-0" 
          @click="copyToClipboard(data.direct_url)"
        >
          <Copy class="h-4 w-4" />
        </Button>
      </div>
    </TableCell>
    <TableCell> 
      <Badge :variant="data.status === 'active' ? 'default' : 'secondary'">
        {{ data.status || 'N/A' }}
      </Badge>
    </TableCell>
    <TableCell>
      <p class="font-medium">{{ data.created_by_name || 'N/A' }}</p>
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
import type { Ad } from "@/stores/adStore";
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Trash2, Pencil, Copy } from "lucide-vue-next";
import { toast } from "@/components/ui/toast/use-toast";

defineProps<{
  data: Ad;
}>();

defineEmits(["clickDelete", "clickUpdate"]);

const getImageUrl = (imageUrl: string): string => {
  if (imageUrl.startsWith('http://') || imageUrl.startsWith('https://')) {
    return imageUrl;
  }
  
  const API_URL = import.meta.env.VITE_API_URL;
  return `${API_URL}/${imageUrl}`;
};

const truncateUrl = (url: string | null | undefined): string => {
  if (!url) return 'N/A';
  return url.length > 40 ? url.substring(0, 40) + '.........' : url;
};

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    toast({
      description: "URL copied to clipboard",
      duration: 2000,
    });
  } catch (err) {
    toast({
      description: "Failed to copy URL",
      variant: "destructive",
    });
  }
};

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