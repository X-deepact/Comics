<template>
    <Dialog :open="chapterStore.chapteritemDialogIsOpen"
        @update:open="(value: boolean) => { chapterStore.chapteritemDialogIsOpen = value; }">
        <DialogContent class="min-w-[100rem] overflow-y-scroll max-h-screen">
            <Component title="ChapterItem" :pageSize="chapteritemStore.page_size" :searchIsEnable="true"
                @clickAdd="() => (chapteritemStore.createDialogIsOpen = true)" @clickRefresh="
                    () => {
                        chapteritemStore.getChapterItemsData();
                    }
                " @update:search="
                    (keyword: any) => {
                        chapteritemStore.searchKeyword = keyword;
                        chapteritemStore.getChapterItemsData();
                    }
                " @update:pageSize="
                    (pageSize: any) => {
                        chapteritemStore.page_size = pageSize;
                        chapteritemStore.getChapterItemsData();
                    }
                " />
            <TableComponent :subjects="chapteritemStore.chapteritemsData" :columns="ColumnNames"
                :isLoading="chapteritemStore.isLoading" :row="ChapterItemRow" @clickDelete="
                    (data: any) => {
                        chapteritemStore.selectedData = data;
                        chapteritemStore.deleteDialogIsOpen = true;
                    }
                " @clickUpdate="
                    (data: any) => {
                        chapteritemStore.selectedData = data;
                        chapteritemStore.updateDialogIsOpen = true;
                    }
                " @clickSorting="
                    (data: any) => {
                        chapteritemStore.sortBy = data;
                        chapteritemStore.setSorting();
                        chapteritemStore.getChapterItemsData();
                    }" />
            <Pagination :currentPage="chapteritemStore.current_page" :totalItems="chapteritemStore.totalItems"
                :totalPages="chapteritemStore.totalItems % chapteritemStore.page_size === 0
                    ? chapteritemStore.totalItems / chapteritemStore.page_size
                    : Math.ceil(chapteritemStore.totalItems / chapteritemStore.page_size)
                    " @update:currentPage="
                        (currentPage: any) => {
                            chapteritemStore.current_page = currentPage;
                            chapteritemStore.getChapterItemsData();
                        }
                    " />
            <CreateChapterItem />
            <UpdateChapterItem />
            <DeleteChapterItem />
        </DialogContent>
    </Dialog>
</template>
<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import Pagination from "@/lib/Pagination.vue";
import { useChapterStore } from "../../../../../stores/chapterStore";
import { ColumnNames } from "./columnHeader";
import ChapterItemRow from "./row.vue";
import CreateChapterItem from "./modal/CreateChapterItem.vue";
import UpdateChapterItem from "./modal/UpdateChapterItem.vue";
import DeleteChapterItem from "./modal/DeleteChapterItem.vue";
import { useChapterItemStore } from "../../../../../stores/chapteritemStore";
const chapterStore = useChapterStore();
const chapteritemStore = useChapterItemStore();
</script>
