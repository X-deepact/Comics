<template>
  <div class="p-8">
    <Component
      title="Genre"
      :pageSize="genreStore.page_size"
      @clickAdd="() => (genreStore.createDialogIsOpen = true)"
      @clickRefresh="
        () => {
          genreStore.getGenreData();
        }
      "
      @update:search="
        (keyword: any) => {
          genreStore.searchKeyword = keyword;
          genreStore.getGenreData();
        }
      "
      @update:pageSize="
        (pageSize: any) => {
          genreStore.page_size = pageSize;
          genreStore.getGenreData();
        }
      "
    />
    <TableComponent
      :subjects="genreStore.genreData"
      :columns="ColumnNames"
      :isLoading="genreStore.isLoading"
      :row="GenreRow"
      @clickDelete="
        (data: any) => {
          genreStore.selectedData = data;
          genreStore.deleteDialogIsOpen = true;
        }
      "
      @clickUpdate="
        (data: any) => {
          genreStore.selectedData = data;
          genreStore.updateDialogIsOpen = true;
        }
      "
    />

    <Pagination
      :currentPage="genreStore.current_page"
      :totalItems="genreStore.totalItems"
      :totalPages="
        genreStore.totalItems % genreStore.page_size === 0
          ? genreStore.totalItems / genreStore.page_size
          : Math.ceil(genreStore.totalItems / genreStore.page_size)
      "
      @update:currentPage="
        (currentPage: any) => {
          genreStore.current_page = currentPage;
          genreStore.getGenreData();
        }
      "
    />
    <CreateGenre />
    <UpdateGenre />
    <DeleteGenre />
  </div>
</template>

<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { ColumnNames } from "./columnHeader";
import CreateGenre from "./modal/CreateGenre.vue";
import DeleteGenre from "./modal/DeleteGenre.vue";
import UpdateGenre from "./modal/UpdateGenre.vue";
import { useGenreStore } from "../../../stores/genreStore";
import GenreRow from "./row.vue";
const genreStore = useGenreStore();
genreStore.getGenreData();
</script>
