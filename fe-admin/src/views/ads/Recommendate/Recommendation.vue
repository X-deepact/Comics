<template>
    <div class="p-8">
        <Component 
            title="Recommendation"
            :pageSize="recommendStore.page_size"
            :search-enabled="true"
            :search-value="recommendStore.searchKeyword"
            @clickAdd="() => (recommendStore.createDialogIsOpen = true)"
            @clickRefresh="() => { recommendStore.getRecommendData(); }"
            @update:search="handleSearch"
            @update:pageSize="(pageSize: number) => {
                recommendStore.page_size = pageSize;
                recommendStore.getRecommendData();
            }"
        >
            <template #search>
                <div class="flex items-center space-x-2">
                    <Input
                        type="search"
                        :value="recommendStore.searchKeyword"
                        @input="(e: Event) => handleSearch((e.target as HTMLInputElement).value)"
                        placeholder="Search by title..."
                        class="h-9 w-[150px] lg:w-[250px]"
                    />
                </div>
            </template>
        </Component>
        
        <Table 
            :subjects="recommendStore.recommendData" 
            :columns="columns"
            :isLoading="recommendStore.isLoading"
            :row="RecommendRow"
            @clickDelete="(data: Recommend) => {
                recommendStore.selectedData = data;
                recommendStore.deleteDialogIsOpen = true;
            }"
            @clickUpdate="(data: Recommend) => {
                recommendStore.selectedData = data;
                recommendStore.updateDialogIsOpen = true;
            }"
            @clickSorting="handleSort"
        />

        <Pagination 
            v-if="recommendStore.totalItems > 0"
            :currentPage="recommendStore.current_page"
            :totalItems="recommendStore.totalItems"
            :totalPages="Math.ceil(recommendStore.totalItems / recommendStore.page_size)"
            @update:currentPage="(currentPage: number) => {
                recommendStore.current_page = currentPage;
                recommendStore.getRecommendData();
            }"
        />
        
        <CreateModal v-if="recommendStore.createDialogIsOpen" />
        <UpdateModal v-if="recommendStore.updateDialogIsOpen" />
        <DeleteModal v-if="recommendStore.deleteDialogIsOpen" />
    </div>
</template>

<script setup lang="ts">
import Component from '@/lib/Component.vue'
import Table from '@/lib/Table.vue'
import Pagination from '@/lib/Pagination.vue'
import CreateModal from './modal/CreateModal.vue'
import UpdateModal from './modal/UpdateModal.vue'
import DeleteModal from './modal/DeleteModal.vue'
import { useRecommendStore, type Recommend } from '@/stores/recommendStore'
import { columns } from './columnHeader'
import RecommendRow from './row.vue'
import { onMounted } from 'vue'
import { Input } from "@/components/ui/input"

const recommendStore = useRecommendStore()

// Add debounce utility
const debounce = (fn: Function, delay: number) => {
    let timeoutId: ReturnType<typeof setTimeout>;
    return (...args: any[]) => {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => fn(...args), delay);
    };
};

// Handle search with debounce
const handleSearch = debounce((value: string) => {
    recommendStore.searchKeyword = value;
    recommendStore.current_page = 1; // Reset to first page when searching
    recommendStore.getRecommendData();
}, 300);

const handleSort = async (key: string) => {
    if (key === 'updated_at') {
        recommendStore.sortBy = key;
        recommendStore.sortDirection = recommendStore.sortDirection === 'asc' ? 'desc' : 'asc';
        await recommendStore.getRecommendData();
    }
};

onMounted(() => {
    recommendStore.getRecommendData();
})
</script>