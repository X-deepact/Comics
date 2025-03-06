<template>
  <div class="p-8">
    <Component 
      title="Users" 
      :pageSize="userStore.page_size" 
      @clickAdd="() => (userStore.createDialogIsOpen = true)" 
      @clickRefresh="() => { userStore.getUserData(); }"
      @update:search="handleSearch"
      @update:pageSize="
        (pageSize) => {
          userStore.page_size = pageSize;
          userStore.getUserData();
        }
      "
    />
    <TableComponent 
      :subjects="userStore.userData" 
      :columns="ColumnNames" 
      :row="UserRow"
      :isLoading="userStore.isLoading" 
      @clickDelete="(data) => {
        userStore.selectedData = data;
        userStore.deleteDialogIsOpen = true;
      }"
      @clickUpdate="(data) => {
        userStore.selectedData = data;
        userStore.updateDialogIsOpen = true;
      }"
      @clickActive="(data) => {
        userStore.selectedData = data;
        userStore.updateStatusDialogIsOpen = true;
      }"
    />
    <Pagination :currentPage="userStore.current_page" :totalItems="userStore.totalItems" :totalPages="userStore.totalItems % userStore.page_size === 0
      ? userStore.totalItems / userStore.page_size
      : Math.ceil(userStore.totalItems / userStore.page_size)
      " @update:currentPage="
        (currentPage: any) => {
          userStore.current_page = currentPage;
          userStore.getUserData();
        }
      " />
  </div>
  <CreateUser />
  <UpdateUser />
  <DeleteUser />
  <UpdateStatus />
</template>

<script setup>
import { ref } from "vue";
import { useUserStore } from "../../stores/userStore";
import Component from "@/lib/Component.vue";
import TableComponent from "@/lib/Table.vue";
import { ColumnNames } from "./columnHeader";
import UserRow from "./row.vue";
import Pagination from "@/lib/Pagination.vue";
import CreateUser from "./modal/CreateUser.vue";
import UpdateUser from "./modal/UpdateUser.vue";
import DeleteUser from "./modal/DeleteUser.vue";
import UpdateStatus from "./modal/UpdateStatus.vue";

const userStore = useUserStore();

// Initialize the data
userStore.getUserData();

const handleSearch = (value) => {
  userStore.searchKeyword = value;
  userStore.current_page = 1;
  userStore.getUserData();
};
</script>
