<template>
    <div class="p-8">
        <Component 
            title="Ads"
            :pageSize="adsStore.page_size"
            @clickAdd="() => (adsStore.createDialogIsOpen = true)"
            @clickRefresh="() => adsStore.getAdsData()"
            @update:search="
                (keyword: string) => {
                    adsStore.searchKeyword = keyword;
                    adsStore.getAdsData();
                }
            "
            @update:pageSize="
                (pageSize: number) => {
                    adsStore.page_size = pageSize;
                    adsStore.getAdsData();
                }
            "
        />
        
        <TableComponent
            :subjects="adsStore.adsData"
            :columns="columns"
            :isLoading="adsStore.isLoading"
            :row="AdRow"
            @clickDelete="
                (data: any) => {
                    adsStore.selectedData = data;
                    adsStore.deleteDialogIsOpen = true;
                }
            "
            @clickUpdate="
                (data: any) => {
                    adsStore.selectedData = data;
                    adsStore.updateDialogIsOpen = true;
                }
            "
        />

        <Pagination
            :currentPage="adsStore.current_page"
            :totalItems="adsStore.totalItems"
            :totalPages="
                adsStore.totalItems % adsStore.page_size === 0
                    ? adsStore.totalItems / adsStore.page_size
                    : Math.ceil(adsStore.totalItems / adsStore.page_size)
            "
            @update:currentPage="
                (currentPage: number) => {
                    adsStore.current_page = currentPage;
                    adsStore.getAdsData();
                }
            "
        />

        <CreateAd />
        <UpdateAd />
        <DeleteAd />
    </div>
</template>

<script setup lang="ts">
import { useAdsStore } from "@/stores/adsStore";
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import CreateAd from "./modal/CreateAd.vue";
import UpdateAd from "./modal/UpdateAd.vue";
import DeleteAd from "./modal/DeleteAd.vue";
import AdRow from "./row.vue";

const adsStore = useAdsStore();
adsStore.getAdsData();

const columns = {
    id: { label: "ID", key: "id" },
    title: { label: "Title", key: "title" },
    image: { label: "Image", key: "image" },
    active_from: { label: "Start Date", key: "active_from" },
    active_to: { label: "End Date", key: "active_to" },
    type: { label: "Type", key: "type" },
    direct_url: { label: "URL", key: "direct_url" },
    status: { label: "Status", key: "status" },
    action: { label: "Action", key: "action" },
};
</script>