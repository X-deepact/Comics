<script setup lang="ts">
defineProps<{
  title: string;
  loading: boolean;
  items: any[];
  type: 'comics' | 'users' | 'dramas' | 'ads';
}>();
</script>

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <h3 class="text-lg font-semibold mb-4">{{ title }}</h3>
    
    <div v-if="loading" class="space-y-4">
      <div v-for="i in 5" :key="i" class="animate-pulse">
        <div class="h-4 bg-gray-200 rounded w-3/4"></div>
      </div>
    </div>
    
    <div v-else class="space-y-4">
      <div v-for="item in items" :key="item.id" class="flex items-center space-x-3">
        <div v-if="['comics', 'dramas', 'ads'].includes(type) && item.cover || item.thumb || item.image" 
             class="w-10 h-10 rounded bg-gray-200 overflow-hidden">
          <img 
            :src="item.cover || item.thumb || item.image" 
            class="w-full h-full object-cover"
            alt=""
          />
        </div>
        <div class="flex-1">
          <p class="text-sm font-medium">
            {{ item.name || item.title || item.username || 'Untitled' }}
          </p>
          <p class="text-xs text-gray-500">
            {{ new Date(item.created_at).toLocaleDateString() }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template> 