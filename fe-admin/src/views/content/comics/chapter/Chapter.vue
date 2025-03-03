<template>
  <Dialog :open="comicStore.chapterDialogIsOpen"
    @update:open="(value: boolean) => { comicStore.chapterDialogIsOpen = value; }">
    <DialogContent class="min-w-[100rem] overflow-y-scroll max-h-screen">
      <Component title="Chapter" :pageSize="chapterStore.page_size" :searchIsEnable="true"
        @clickAdd="() => (chapterStore.createDialogIsOpen = true)" @clickRefresh="
          () => {
            chapterStore.getChapterData();
          }
        " @update:search="
          (keyword: any) => {
            chapterStore.searchKeyword = keyword;
            chapterStore.getChapterData();
          }
        " @update:pageSize="
          (pageSize: any) => {
            chapterStore.page_size = pageSize;
            chapterStore.getChapterData();
          }
        " />
      <TableComponent :subjects="chapterStore.chaptersData" :columns="ColumnNames" :isLoading="chapterStore.isLoading"
        :row="ChapterRow" @clickDelete="
          (data: any) => {
            chapterStore.selectedData = data;
            chapterStore.deleteDialogIsOpen = true;
          }
        " @clickUpdate="
          (data: any) => {
            chapterStore.selectedData = data;
            chapterStore.updateDialogIsOpen = true;
          }
        " @clickAction="
          (data: any) => {
            chapterStore.selectedData.id = data;
            chapteritemStore.getChapterItemsData();
            chapterStore.chapteritemDialogIsOpen = true;
          }" @clickSorting="
            (data: any) => {
              chapterStore.sortBy = data;
              chapterStore.setSorting();
              chapterStore.getChapterData();
            }" />
      <Pagination :currentPage="chapterStore.current_page" :totalItems="chapterStore.totalItems" :totalPages="chapterStore.totalItems % chapterStore.page_size === 0
        ? chapterStore.totalItems / chapterStore.page_size
        : Math.ceil(chapterStore.totalItems / chapterStore.page_size)
        " @update:currentPage="
          (currentPage: any) => {
            chapterStore.current_page = currentPage;
            chapterStore.getChapterData();
          }
        " />
      <CreateChapter />
      <UpdateChapter />
      <DeleteChapter />
      <ChapterItem />
    </DialogContent>
  </Dialog>
</template>
<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import Pagination from "@/lib/Pagination.vue";
import { useComicStore } from "../../../../stores/comicStore";
import { useChapterStore } from "../../../../stores/chapterStore";
import { ColumnNames } from "./columnHeader";
import ChapterRow from "./row.vue";
import CreateChapter from "./modal/CreateChapter.vue";
import UpdateChapter from "./modal/UpdateChapter.vue";
import DeleteChapter from "./modal/DeleteChapter.vue";
import ChapterItem from "./chapteritem/ChapterItem.vue";
import { useChapterItemStore } from "../../../../stores/chapteritemStore";
const comicStore = useComicStore();
const chapterStore = useChapterStore();
const chapteritemStore = useChapterItemStore();
</script>
