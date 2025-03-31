<template>
  <div class="p-8">
    <Component 
      title="Genres"
      :pageSize="genreStore.page_size"
      @clickAdd="() => (genreStore.createDialogIsOpen = true)"
      @clickRefresh="() => { genreStore.getGenreData(); }"
      @update:search="(keyword: string) => {
        genreStore.searchKeyword = keyword;
        genreStore.getGenreData();
      }"
      @update:pageSize="(pageSize: number) => {
        genreStore.page_size = pageSize;
        genreStore.getGenreData();
      }"
    />
    
    <TableComponent
      :subjects="genreStore.genreData"
      :columns="ColumnNames"
      :isLoading="genreStore.isLoading"
      :row="GenreRow"
      class="w-full"
      @clickSorting="handleSort"
      @clickDelete="(data: Genre) => {
        genreStore.selectedData = data;
        genreStore.deleteDialogIsOpen = true;
      }"
      @clickUpdate="(data: Genre) => {
        genreStore.selectedData = { ...data };
        genreStore.updateDialogIsOpen = true;
      }"
    >
    </TableComponent>

    <Pagination
      v-if="genreStore.totalItems > 0"
      :currentPage="genreStore.current_page"
      :totalItems="genreStore.totalItems"
      :totalPages="Math.ceil(genreStore.totalItems / genreStore.page_size)"
      @update:currentPage="(currentPage: number) => {
        genreStore.current_page = currentPage;
        genreStore.getGenreData();
      }"
    />

    <CreateGenre />
    <UpdateGenre />
    <DeleteGenre />
  </div>
</template>

<script setup lang="ts">
import { Genre, useDramaGenreStore } from '@/stores/drama_genreStore';
import Component from '@/lib/Component.vue';
import TableComponent from '@/lib/Table.vue';
import Pagination from '@/lib/Pagination.vue';
import CreateGenre from './modal/CreateGenre.vue';
import DeleteGenre from './modal/DeleteGenre.vue';
import UpdateGenre from './modal/UpdateGenre.vue';
import GenreRow from './row.vue';
import { ColumnNames } from './columnHeader';
import { onMounted } from 'vue';

const genreStore = useDramaGenreStore();

const handleSort = async (key: string) => {
  genreStore.updateSort(key);
  await genreStore.getGenreData();
};

onMounted(() => {
  genreStore.getGenreData();
});
</script>