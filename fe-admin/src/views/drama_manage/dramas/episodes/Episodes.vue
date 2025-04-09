<template>
    <Dialog :open="dramaStore.episodeDialogIsOpen"
        @update:open="(value: boolean) => { dramaStore.episodeDialogIsOpen = value; }">
        <DialogContent class="min-w-[80%] overflow-y-scroll max-h-screen">
            <Component title="Episode" :pageSize="episodeStore.page_size" :searchIsEnable="true"
                @clickAdd="() => (episodeStore.createDialogIsOpen = true)" @clickRefresh="
                    () => {
                        episodeStore.getEpisodeData();
                    }
                " @update:pageSize="
                    (pageSize: any) => {
                        episodeStore.page_size = pageSize;
                        episodeStore.getEpisodeData();
                    }
                " />
            <TableComponent :subjects="episodeStore.episodeData" :columns="ColumnNames"
                :isLoading="episodeStore.isLoading" :row="EpisodeRow" @clickDelete="
                    (data: any) => {
                        episodeStore.selectedData = data;
                        episodeStore.deleteDialogIsOpen = true;
                    }
                " @clickUpdate="
                    (data: any) => {
                        episodeStore.selectedData.id = data.id;
                        episodeStore.selectedData.number = data.number;
                        episodeStore.selectedData.active = data.active;
                        episodeStore.selectedData.drama_id = data.drama_id;
                        episodeStore.selectedData.video = data.video;
                        episodeStore.updateDialogIsOpen = true;
                    }
                " @clickActive="(data: any) => {
                    episodeStore.selectedData = data;
                    episodeStore.updateActiveDialogIsOpen = true
                }" @clickSorting="
                    (data: any) => {
                        episodeStore.sortBy = data;
                        episodeStore.setSorting();
                        episodeStore.getEpisodeData();
                    }" @clickAction="(data: any) => {
                        episodeStore.selectedData = data;
                        episodeStore.playDialogIsOpen = true
                    }" />
            <Pagination :currentPage="episodeStore.current_page" :totalItems="episodeStore.totalItems" :totalPages="episodeStore.totalItems % episodeStore.page_size === 0
                ? episodeStore.totalItems / episodeStore.page_size
                : Math.ceil(episodeStore.totalItems / episodeStore.page_size)
                " @update:currentPage="
                    (currentPage: any) => {
                        episodeStore.current_page = currentPage;
                        episodeStore.getEpisodeData();
                    }
                " />
            <CreateEpisode />
            <UpdateEpisode />
            <DeleteEpisode />
            <UpdateActiveEpisode />
            <PlayEpisode />
        </DialogContent>
    </Dialog>
</template>
<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import Pagination from "@/lib/Pagination.vue";
import { ColumnNames } from "./columnHeader";
import EpisodeRow from "./row.vue";
import CreateEpisode from "./modal/CreateEpisode.vue";
import UpdateEpisode from "./modal/UpdateEpisode.vue";
import DeleteEpisode from "./modal/DeleteEpisode.vue";
import { useDramaStore } from "../../../../stores/dramaStore";
import { useEpisodeStore } from "../../../../stores/episodeStore";
import UpdateActiveEpisode from "./modal/UpdateActiveEpisode.vue";
import PlayEpisode from "./modal/PlayEpisode.vue";

const dramaStore = useDramaStore();
const episodeStore = useEpisodeStore();
</script>
