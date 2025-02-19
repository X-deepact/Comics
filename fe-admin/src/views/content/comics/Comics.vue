<template>
  <div class="p-8">
    <Component
      title="Comic"
      :pageSize="comicStore.page_size"
      @clickAdd="() => (comicStore.createDialogIsOpen = true)"
      @clickRefresh="
        () => {
          comicStore.getComicData();
        }
      "
      @update:search="
        (keyword) => {
          comicStore.searchKeyword = keyword;
          comicStore.getComicData();
        }
      "
      @update:pageSize="
        (pageSize) => {
          comicStore.page_size = pageSize;
          comicStore.getComicData();
        }
      "
    />

    <TableComponent
      :subjects="comicStore.comicData"
      :columns="ColumnNames"
      :isLoading="comicStore.isLoading"
      @clickDelete="
        (data) => {
          comicStore.selectedData = data;
          comicStore.deleteDialogIsOpen = true;
        }
      "
      @clickUpdate="
        (data) => {
          comicStore.selectedData = data;
          comicStore.updateDialogIsOpen = true;
        }
      "
    />

    <Pagination
      :currentPage="comicStore.current_page"
      :totalItems="comicStore.totalItems"
      :totalPages="
        comicStore.totalItems % comicStore.page_size === 0
          ? comicStore.totalItems / comicStore.page_size
          : Math.ceil(comicStore.totalItems / comicStore.page_size)
      "
      @update:currentPage="
        (currentPage) => {
          comicStore.current_page = currentPage;
          comicStore.getComicData();
        }
      "
    />
    <CreateComic />
    <UpdateComic />
    <DeleteComic />
  </div>
</template>

<script setup>
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { useComicStore } from "../../../stores/comicStore";
import { ColumnNames } from "./columnHeader";
import CreateComic from "./modal/CreateComic.vue";
import DeleteComic from "./modal/DeleteComic.vue";
import UpdateComic from "./modal/UpdateComic.vue";
// Define columns structure
const comicStore = useComicStore();
comicStore.getComicData();
</script>
