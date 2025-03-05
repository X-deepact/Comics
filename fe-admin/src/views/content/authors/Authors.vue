<template>
  <div class="p-8">
    <Component title="Author" :pageSize="authorStore.page_size"
      @clickAdd="() => (authorStore.createDialogIsOpen = true)" @clickRefresh="
        () => {
          authorStore.getAuthorData();
        }
      " @update:search="
        (keyword: any) => {
          authorStore.searchKeyword = keyword;
          authorStore.getAuthorData();
        }
      " @update:pageSize="
        (pageSize: any) => {
          authorStore.page_size = pageSize;
          authorStore.getAuthorData();
        }
      " />

    <TableComponent :subjects="authorStore.authorData" :columns="ColumnNames" :isLoading="authorStore.isLoading"
      :row="AuthorRow" @clickDelete="
        (data: any) => {
          authorStore.selectedData = data;
          authorStore.deleteDialogIsOpen = true;
        }
      " @clickUpdate="
        (data: any) => {
          authorStore.selectedData.id = data.id;
          authorStore.selectedData.name = data.name;
          authorStore.selectedData.biography = data.biography;
          const date = new Date(data.birth_day);
          authorStore.selectedData.birth_day = new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
          authorStore.updateDialogIsOpen = true;
        }
      " @clickSorting="
        (data: any) => {
          authorStore.sortBy = data;
          authorStore.setSorting();
          authorStore.getAuthorData();
        }" />

    <Pagination :currentPage="authorStore.current_page" :totalItems="authorStore.totalItems" :totalPages="authorStore.totalItems % authorStore.page_size === 0
      ? authorStore.totalItems / authorStore.page_size
      : Math.ceil(authorStore.totalItems / authorStore.page_size)
      " @update:currentPage="
        (currentPage: any) => {
          authorStore.current_page = currentPage;
          authorStore.getAuthorData();
        }
      " />
    <CreateAuthor />
    <UpdateAuthor />
    <DeleteAuthor />
  </div>
</template>

<script setup lang="ts">
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { useAuthorStore } from "../../../stores/authorStore";
import { ColumnNames } from "./columnHeader";
import CreateAuthor from "./modal/CreateAuthor.vue";
import DeleteAuthor from "./modal/DeleteAuthor.vue";
import UpdateAuthor from "./modal/UpdateAuthor.vue";
import AuthorRow from "./row.vue";
import { CalendarDate } from "@internationalized/date";
// Define columns structure
const authorStore = useAuthorStore();
authorStore.getAuthorData();
</script>
