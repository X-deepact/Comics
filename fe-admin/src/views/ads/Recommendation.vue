<template>
    <div class="p-8">
        <Component 
            title="Recommendation"
            :totalItems="recommendStore.totalItems"
            :currentPage="recommendStore.current_page"
            :pageSize="recommendStore.page_size"
            :displayedPages="displayedPages"
            @clickAdd="() => (recommendStore.createDialogIsOpen = true)"
            @clickRefresh="() => recommendStore.getRecommendData()"
            @update:pageSize="
                (pageSize: number) => {
                    recommendStore.page_size = pageSize;
                    recommendStore.getRecommendData();
                }
            "
            @update:search="
                (keyword: string) => {
                    recommendStore.searchKeyword = keyword;
                    recommendStore.getRecommendData();
                }
            "
        />
        
        <TableComponent 
            :subjects="recommendStore.recommendData" 
            :columns="ColumnNames"
            :isLoading="recommendStore.isLoading"
            :row="RecommendRow"
            @clickDelete="
                (data: any) => {
                    recommendStore.selectedData = data;
                    recommendStore.deleteDialogIsOpen = true;
                }
            "
            @clickUpdate="
                (data: any) => {
                    recommendStore.setSelectedDataForUpdate(data);
                    recommendStore.updateDialogIsOpen = true;
                }
            "
        />

        <Pagination 
            :totalItems="recommendStore.totalItems"
            :currentPage="recommendStore.current_page"
            :totalPages="
                recommendStore.totalItems % recommendStore.page_size === 0
                    ? recommendStore.totalItems / recommendStore.page_size
                    : Math.ceil(recommendStore.totalItems / recommendStore.page_size)
            "
            @update:currentPage="
                (currentPage: number) => {
                    recommendStore.current_page = currentPage;
                    recommendStore.getRecommendData();
                }
            "
        />
        
        <CreateRecommend />
        <UpdateRecommend />
        <DeleteRecommend />
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Component from '@/lib/Component.vue'
import TableComponent from '@/lib/Table.vue'
import Pagination from '@/lib/Pagination.vue'
import { useRecommendStore } from "@/stores/recommendStore";
import { ColumnNames } from './columnHeader'
import RecommendRow from './row.vue'
import CreateRecommend from './modal/CreateRecommend.vue'
import UpdateRecommend from './modal/UpdateRecommend.vue'
import DeleteRecommend from './modal/DeleteRecommend.vue'

const recommendStore = useRecommendStore()
recommendStore.getRecommendData()

const displayedPages = computed(() => {
    const pages = []
    for (let i = 1; i <= Math.ceil(recommendStore.totalItems / recommendStore.page_size); i++) {
        pages.push(i)
    }
    return pages
})
</script>