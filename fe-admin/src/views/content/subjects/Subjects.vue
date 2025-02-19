<template>
  <div class="p-8">
    <Component
      title="Subject Management"
      :pageSize="subjectStore.page_size"
      @clickAdd="() => (subjectStore.createDialogIsOpen = true)"
      @clickRefresh="() => { subjectStore.getSubjectData(); }"
      @update:search="(keyword) => {
        subjectStore.searchKeyword = keyword;
        subjectStore.getSubjectData();
      }"
      @update:pageSize="(pageSize) => {
        subjectStore.page_size = pageSize;
        subjectStore.getSubjectData();
      }"
    />

    <TableComponent
      :subjects="subjectStore.subjectData"
      :columns="ColumnNames"
      :isLoading="subjectStore.isLoading"
      @clickDelete="(data) => {
        subjectStore.selectedData = data;
        subjectStore.deleteDialogIsOpen = true;
      }"
      @clickUpdate="(data) => {
        subjectStore.selectedData = data;
        subjectStore.updateDialogIsOpen = true;
      }"
    />

    <Pagination
      :currentPage="subjectStore.current_page"
      :totalItems="subjectStore.totalItems"
      :totalPages="
        subjectStore.totalItems % subjectStore.page_size === 0
          ? subjectStore.totalItems / subjectStore.page_size
          : Math.ceil(subjectStore.totalItems / subjectStore.page_size)
      "
      @update:currentPage="(currentPage) => {
        subjectStore.current_page = currentPage;
        subjectStore.getSubjectData();
      }"
    />
    <CreateSubject />
    <UpdateSubject />
    <DeleteSubject />
  </div>
</template>

<script setup>
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import Pagination from "@/lib/Pagination.vue";
import { useSubjectStore } from "../../../stores/subjectStore";
import { ColumnNames } from "./columnHeader";
import CreateSubject from "./modal/CreateSubject.vue";
import DeleteSubject from "./modal/DeleteSubject.vue";
import UpdateSubject from "./modal/UpdateSubject.vue";

const subjectStore = useSubjectStore();
subjectStore.getSubjectData();
</script> 