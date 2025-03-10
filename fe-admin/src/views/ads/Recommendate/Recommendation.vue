<template>
    <div class="p-8">
        <Component 
            title="Recommendation"
            :pageSize="recommendStore.page_size"
            :isLoading="recommendStore.isLoading"
            @clickAdd="() => (recommendStore.createDialogIsOpen = true)"
            @clickRefresh="() => { recommendStore.getRecommendData() }"
            @update:search="(keyword: string) => recommendStore.handleSearch(keyword)"
            @update:pageSize="(pageSize: number) => {
                recommendStore.page_size = pageSize;
                recommendStore.current_page = 1;
                recommendStore.getRecommendData();
            }"
        />
        
        <Table 
            :subjects="recommendStore.recommendData" 
            :columns="columns"
            :isLoading="recommendStore.isLoading"
            :row="RecommendRow"
            @clickDelete="(data: any) => {
                recommendStore.selectedData = data;
                recommendStore.deleteDialogIsOpen = true;
            }"
            @clickUpdate="(data: any) => {
                recommendStore.selectedData = data;
                recommendStore.updateDialogIsOpen = true;
            }"
            @clickSorting="(sortKey: string) => recommendStore.handleSort(sortKey)"
        />

        <Pagination 
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
import { useRecommendStore } from '@/stores/recommendStore'
import { columns } from './columnHeader'
import RecommendRow from './row.vue'

const recommendStore = useRecommendStore()

// Initial data load
recommendStore.getRecommendData()
</script>