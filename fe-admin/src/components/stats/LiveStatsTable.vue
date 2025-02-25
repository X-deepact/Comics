<script setup lang="ts">
import { ref } from 'vue';
import type { TableHeader } from '@/types/table';

const props = defineProps<{
    headers: TableHeader[]
}>();

const pageSize = ref(5);
const currentPage = ref(1);
const searchQuery = ref('');
</script>

<template>
    <div>
        <div class="p-4 border-b">
            <input
                v-model="searchQuery"
                type="text"
                placeholder="Search by title..."
                class="w-full px-4 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
        </div>
        <div class="overflow-x-auto">
            <table class="w-full">
                <thead>
                    <tr class="border-b">
                        <th
                            v-for="header in headers"
                            :key="header.title"
                            class="px-6 py-3 text-left text-sm font-medium text-gray-900"
                        >
                            <div class="flex items-center gap-2">
                                {{ header.title }}
                                <svg 
                                    v-if="header.sortable" 
                                    xmlns="http://www.w3.org/2000/svg" 
                                    width="16" 
                                    height="16" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round"
                                >
                                    <path d="m6 9 6 6 6-6"/>
                                </svg>
                            </div>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td colspan="5" class="px-6 py-8 text-center text-gray-500">
                            No results.
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="flex items-center justify-between p-4 border-t">
            <select 
                v-model="pageSize" 
                class="px-3 py-1.5 border rounded bg-white"
            >
                <option :value="5">5</option>
                <option :value="10">10</option>
                <option :value="20">20</option>
            </select>
            <div class="flex items-center gap-3">
                <span class="text-sm text-gray-600">Page</span>
                <input 
                    type="number" 
                    v-model="currentPage" 
                    class="w-16 px-3 py-1.5 border rounded text-center" 
                />
                <span class="text-sm text-gray-600">of 0</span>
            </div>
        </div>
    </div>
</template>