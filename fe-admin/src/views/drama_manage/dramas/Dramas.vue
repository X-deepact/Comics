<template>
    <div class="p-8">
        <Component title="Drama" :pageSize="dramaStore.page_size" :searchIsEnable="false"
            @clickAdd="() => (dramaStore.createDialogIsOpen = true)" @clickRefresh="
                () => {
                    dramaStore.getDramaData();
                }
            " @update:search="
                (keyword: any) => {
                    dramaStore.searchKeyword = keyword;
                    dramaStore.getDramaData();
                }
            " @update:pageSize="
                (pageSize: any) => {
                    dramaStore.page_size = pageSize;
                    dramaStore.getDramaData();
                }
            " />

        <TableComponent :subjects="dramaStore.dramaData" :columns="ColumnNames" :isLoading="dramaStore.isLoading"
            :row="DramaRow" @clickDelete="
                (data: any) => {
                    dramaStore.selectedData = data;
                    dramaStore.deleteDialogIsOpen = true;
                }
            " @clickUpdate="
                (data: any) => {
                    dramaStore.selectedData.id = data.id;
                    dramaStore.selectedData.thumb = data.thumb;
                    dramaStore.selectedData.release_date = data.release_date.substring(0, 10);
                    dramaStore.selectedData.genres = data.genres;
                    dramaStore.selectedData.translations = data.translations;
                    dramaStore.updateDialogIsOpen = true;
                }" @clickSorting="(data: any) => {
                    dramaStore.sortBy = data;
                    dramaStore.setSorting();
                    dramaStore.getDramaData();
                }" @clickActive="
                    (data: any) => {
                        dramaStore.selectedData.id = data.id;
                        dramaStore.updateActiveDialogIsOpen = true;
                    }" @clickAction="(data: any) => {
                        dramaStore.selectedData.id = data.id;
                        episodeStore.getEpisodeData();
                        dramaStore.episodeDialogIsOpen = true
                    }" />

        <Pagination :currentPage="dramaStore.current_page" :totalItems="dramaStore.totalItems" :totalPages="dramaStore.totalItems % dramaStore.page_size === 0
            ? dramaStore.totalItems / dramaStore.page_size
            : Math.ceil(dramaStore.totalItems / dramaStore.page_size)
            " @update:currentPage="
                (currentPage: any) => {
                    dramaStore.current_page = currentPage;
                    dramaStore.getDramaData();
                }
            " />
        <CreateDrama />
        <UpdateDrama />
        <DeleteDrama />
        <UpdateActiveDrama />
        <Episode />
    </div>
</template>

<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { useDramaStore } from "../../../stores/dramaStore";
import { ColumnNames } from "./columnHeader";
import CreateDrama from "./modal/CreateDrama.vue";
import DeleteDrama from "./modal/DeleteDrama.vue";
import UpdateDrama from "./modal/UpdateDrama.vue";
import UpdateActiveDrama from "./modal/UpdateActiveDrama.vue";
import DramaRow from "./row.vue";
import Episode from "./episodes/Episodes.vue";
import { useEpisodeStore } from "../../../stores/episodeStore";
// Define columns structure
const dramaStore = useDramaStore();
dramaStore.getDramaData();
const episodeStore = useEpisodeStore();
</script>