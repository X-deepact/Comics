<template>
    <div class="p-8">
        <Component 
            title="Recommendation"
            :pageSize="recommendStore.page_size"
            @clickAdd="() => (recommendStore.createDialogIsOpen = true)"
            @clickRefresh="() => { recommendStore.getRecommendData() }"
            @update:search="(keyword: string) => {
                recommendStore.searchKeyword = keyword;
                recommendStore.getRecommendData();
            }"
            @update:pageSize="(pageSize: number) => {
                recommendStore.page_size = pageSize;
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
        />

        <Pagination 
            :currentPage="recommendStore.current_page"
            :totalItems="recommendStore.totalItems"
            :totalPages="recommendStore.totalItems % recommendStore.page_size === 0
                ? recommendStore.totalItems / recommendStore.page_size
                : Math.ceil(recommendStore.totalItems / recommendStore.page_size)"
            @update:currentPage="(currentPage: any) => {
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
recommendStore.getRecommendData()
</script>