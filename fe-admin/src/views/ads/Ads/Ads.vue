<template>
    <div class="p-8">
        <Component 
            title="Ads"
            :pageSize="adStore.page_size"
            @clickAdd="() => (adStore.createDialogIsOpen = true)"
            @clickRefresh="() => { adStore.getAdData(); }"
            @update:search="(keyword: string) => {
                adStore.searchKeyword = keyword;
                adStore.getAdData();
            }"
            @update:pageSize="(pageSize: number) => {
                adStore.page_size = pageSize;
                adStore.getAdData();
            }"
        />
        
        <TableComponent
            :subjects="adStore.adData"
            :columns="ColumnNames"
            :isLoading="adStore.isLoading"
            :row="AdRow"
            @clickSorting="handleSort"
            @clickDelete="(data: Ad) => {
                adStore.selectedData = data;
                adStore.deleteDialogIsOpen = true;
            }"
            @clickUpdate="(data: Ad) => {
                adStore.selectedData = data;
                adStore.updateDialogIsOpen = true;
            }"
        />

        <Pagination
            v-if="adStore.totalItems > 0"
            :currentPage="adStore.current_page"
            :totalItems="adStore.totalItems"
            :totalPages="Math.ceil(adStore.totalItems / adStore.page_size)"
            @update:currentPage="(currentPage: number) => {
                adStore.current_page = currentPage;
                adStore.getAdData();
            }"
        />

        <CreateAd />
        <UpdateAd />
        <DeleteAd />
    </div>
</template>

<script setup lang="ts">
import { useAdStore } from '@/stores/adStore';
import Component from '@/lib/Component.vue';
import TableComponent from '@/lib/Table.vue';
import Pagination from '@/lib/Pagination.vue';
import CreateAd from './modal/CreateAd.vue';
import DeleteAd from './modal/DeleteAd.vue';
import UpdateAd from './modal/UpdateAd.vue';
import AdRow from './row.vue';
import { ColumnNames } from './columnHeader';
import { onMounted } from 'vue';

const adStore = useAdStore();

const handleSort = (key: string) => {
    console.log('Sort clicked for key:', key);
    console.log('Current sort state:', adStore.sort);
    
    if (key === 'updated_at' || key === 'updated_by') {
        console.log(`Attempting to sort by ${key}`);
        adStore.updateSort(key);
        console.log('New sort state:', adStore.sort);
    }
};

onMounted(() => {
    adStore.getAdData();
});
</script>