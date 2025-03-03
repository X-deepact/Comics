<template>
  <div class="p-8">
    <Component title="Comic" :pageSize="comicStore.page_size" :searchIsEnable="false"
      @clickAdd="() => (comicStore.createDialogIsOpen = true)" @clickRefresh="
        () => {
          comicStore.getComicData();
        }
      " @update:search="
        (keyword: any) => {
          comicStore.searchKeyword = keyword;
          comicStore.getComicData();
        }
      " @update:pageSize="
        (pageSize: any) => {
          comicStore.page_size = pageSize;
          comicStore.getComicData();
        }
      " />

    <TableComponent :subjects="comicStore.comicData" :columns="ColumnNames" :isLoading="comicStore.isLoading"
      :row="ComicRow" @clickDelete="
        (data: any) => {
          comicStore.selectedData = data;
          comicStore.deleteDialogIsOpen = true;
        }
      " @clickUpdate="
        (data: any) => {
          comicStore.selectedData = data;
          comicStore.updateDialogIsOpen = true;
        }
      " @clickAction="
        (data: any) => {
          comicStore.selectedData.id = data;
          chapterStore.getChapterData();
          comicStore.chapterDialogIsOpen = true;
        }" @clickSorting="
          (data: any) => {
            comicStore.sortBy = data;
            comicStore.setSorting();
            comicStore.getComicData();
          }" />

    <Pagination :currentPage="comicStore.current_page" :totalItems="comicStore.totalItems" :totalPages="comicStore.totalItems % comicStore.page_size === 0
      ? comicStore.totalItems / comicStore.page_size
      : Math.ceil(comicStore.totalItems / comicStore.page_size)
      " @update:currentPage="
        (currentPage: any) => {
          comicStore.current_page = currentPage;
          comicStore.getComicData();
        }
      " />
    <CreateComic />
    <UpdateComic />
    <DeleteComic />
    <Chapter />
  </div>
</template>

<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { useComicStore } from "../../../stores/comicStore";
import { ColumnNames } from "./columnHeader";
import CreateComic from "./modal/CreateComic.vue";
import DeleteComic from "./modal/DeleteComic.vue";
import UpdateComic from "./modal/UpdateComic.vue";
import Chapter from "./chapter/Chapter.vue";
import ComicRow from "./row.vue";
import { useChapterStore } from "../../../stores/chapterStore";
// Define columns structure
const comicStore = useComicStore();
comicStore.getComicData();
const chapterStore = useChapterStore();
</script>
