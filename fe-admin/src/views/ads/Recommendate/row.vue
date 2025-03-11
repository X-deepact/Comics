<template>
  <TableRow>
    <TableCell>{{ data.id }}</TableCell>
    <TableCell>{{ data.title || 'N/A' }}</TableCell>
    <TableCell>
      <img 
        v-if="data.cover" 
        :src="getImageUrl(data.cover)" 
        alt="Recommendation Image"
        class="w-12 h-12 xs:w-14 xs:h-14 sm:w-16 sm:h-16 md:w-20 md:h-20 lg:w-24 lg:h-24 object-cover rounded-md transition-all duration-300" 
      />
      <span v-else class="text-xs xs:text-sm sm:text-base text-gray-500">No image</span>
    </TableCell>
 
    <TableCell>{{ getPositionLabel(data.position) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_from) }}</TableCell>
    <TableCell>{{ formatDateSafe(data.active_to) }}</TableCell>
    <TableCell>
      <p class="font-medium">{{ data.created_by_name || 'N/A' }}</p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateTimeSafe(data.created_at) }}</Label>
        <p className="text-xs text-muted-foreground">
          {{ getTimeAgoSafe(data.created_at) }}
        </p>
      </div>
    </TableCell>
    <TableCell>
      <p class="font-medium">{{ data.updated_by_name || 'N/A' }}</p>
    </TableCell>
    <TableCell>
      <div>
        <Label>{{ formatDateTimeSafe(data.updated_at) }}</Label>
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
          @click="$emit('clickUpdate', data)"
        >
          <Pencil class="h-4 w-4" />
        </Button>
        <Button
          variant="destructive"
          size="icon"
          @click="$emit('clickDelete', data)"
        >
          <Trash2 class="h-4 w-4" />
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
import { positionLabels } from "@/stores/recommendStore";
import { useToast } from "@/components/ui/toast/use-toast";

const { toast } = useToast();
const props = defineProps<{
  data: Recommend;
}>();

const formatDateSafe = (dateString: string | number | null | undefined): string => {
  if (!dateString) return "N/A";
  try {
    // For Unix timestamps (numbers), convert from seconds to milliseconds
    const timestamp = typeof dateString === 'number' 
      ? dateString * 1000 
      : Date.parse(dateString);

    const date = new Date(timestamp);
    if (isNaN(date.getTime())) return "N/A";

    return date.toLocaleDateString('en-GB', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
    });
  } catch (error) {
    toast({
      description: `Date formatting error for value: ${dateString}, ${error}`,
      variant: "destructive"
    });
    return "N/A";
  }
};

const formatDateTimeSafe = (dateString: string | number | null | undefined): string => {
  if (!dateString) return "N/A";
  try {
    // For Unix timestamps (numbers), convert from seconds to milliseconds
    const timestamp = typeof dateString === 'number' 
      ? dateString * 1000 
      : Date.parse(dateString);

    const date = new Date(timestamp);
    if (isNaN(date.getTime())) return "N/A";

    return date.toLocaleString('en-GB', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  } catch (error) {
    toast({
      description: `DateTime formatting error for value: ${dateString}, ${error}`,
      variant: "destructive"
    });
    return "N/A";
  }
};

const getTimeAgoSafe = (dateString: string | number | null | undefined): string => {
  if (!dateString) return "N/A";
  try {
    // Convert Unix timestamp (seconds) to milliseconds if it's a number
    const timestamp = typeof dateString === 'number' 
      ? dateString * 1000  // Convert seconds to milliseconds
      : new Date(dateString).getTime();
    
    const date = new Date(timestamp);
    if (isNaN(date.getTime())) return "N/A";
    return getTimeAgo(date);
  } catch (error) {
    toast({
      description: `Time ago formatting error: ${error}`,
      variant: "destructive"
    });
    return "N/A";
  }
};

const getPositionLabel = (position: number): string => {
  return positionLabels[position] || `Position ${position}`;
};

// Function to handle image URLs
const getImageUrl = (url: string): string => {
  // If the URL is already absolute (starts with http:// or https://)
  if (url.startsWith('http://') || url.startsWith('https://')) {
    return url;
  }
  
  // If it's a relative URL, prepend the API base URL
  const API_URL = import.meta.env.VITE_API_URL;
  // Remove any leading slashes from the URL
  const cleanUrl = url.replace(/^\/+/, '');
  return `${API_URL}/${cleanUrl}`;
};

// Add this function to handle image loading errors
const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement;
  target.src = ''; // You could set a default image here
  target.alt = 'Image not available';
};

defineEmits(["clickDelete", "clickUpdate"]);
</script>