<script setup lang="ts">
import { 
  BookOpenIcon, 
  UsersIcon, 
  FilmIcon, 
  RectangleGroupIcon 
} from '@heroicons/vue/24/outline';

interface Stat {
  name: string;
  value: number;
  icon: string;
}

const props = defineProps<{
  stat: Stat;
}>();

const getIcon = (iconName: string) => {
  const icons = {
    'Total Comics': BookOpenIcon,
    'Total Users': UsersIcon,
    'Total Dramas': FilmIcon,
    'Total Ads': RectangleGroupIcon
  };
  return icons[iconName as keyof typeof icons];
};
</script>

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-600">{{ stat.name }}</p>
        <p class="text-2xl font-semibold mt-1">{{ stat.value }}</p>
      </div>
      <div class="bg-primary/10 p-3 rounded-full">
        <component 
          :is="getIcon(stat.name)"
          class="w-6 h-6 text-primary"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.text-primary {
  color: rgb(75, 192, 192);
}

.bg-primary\/10 {
  background-color: rgba(75, 192, 192, 0.1);
}
</style> 