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

        <CreateAd @upload-image="handleImageUpload" />
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
import { onMounted, ref } from 'vue';

const adStore = useAdStore();
const isUploading = ref(false);

const handleSort = async (key: string) => {
    if (key === 'updated_at') {
        adStore.updateSort(key);
        await adStore.getAdData();
    }
};

const handleImageUpload = async (file: File) => {
    try {
        isUploading.value = true;
        const filename = await adStore.uploadImage(file);
        if (filename) {
            // Update the image in the create ad form
            adStore.selectedData.image = filename;
        }
    } catch (error) {
        console.error('Error uploading image:', error);
    } finally {
        isUploading.value = false;
    }
};

onMounted(() => {
    adStore.getAdData();
});
</script>