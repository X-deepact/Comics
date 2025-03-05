<template>
  <div class="p-8">
    <Component
      title="Users"
      :pageSize="userStore.page_size"
      :displayedPages="displayedPages"
      @clickAdd="() => (userStore.createDialogIsOpen = true)"
      @clickRefresh="
        () => {
          userStore.getUserData();
        }
      "
      @update:search="
        (keyword: any) => {
          userStore.searchKeyword = keyword;
          userStore.getUserData();
        }
      "
      @update:pageSize="
        (pageSize: any) => {
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
      @clickDelete="
        (data: any) => {
          userStore.selectedData = data;
          userStore.deleteDialogIsOpen = true;
        }
      "
      @clickUpdate="
        (data: any) => {
          userStore.selectedData = data;
          userStore.updateDialogIsOpen = true;
        }
      "
      @clickUpdateUserStatus="(data) => { 
        console.log('data');
        userStore.selectedData = data;
        userStore.updateStatusDialogIsOpen = true; 
      }"
    />

    <Pagination
      :currentPage="userStore.current_page"
      :totalItems="userStore.totalItems"
      :totalPages="
        userStore.totalItems % userStore.page_size === 0
          ? userStore.totalItems / userStore.page_size
          : Math.ceil(userStore.totalItems / userStore.page_size)
      "
      @update:currentPage="
        (currentPage: any) => {
          userStore.current_page = currentPage;
          userStore.getUserData();
        }
      "
    />
  </div>
  <CreateUser />
  <UpdateUser />
  <DeleteUser />
  <UpdateStatus />
</template>

<script setup>
import { ref, computed } from "vue";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import Component from "@/lib/Component.vue";
import Pagination from "@/lib/Pagination.vue";
import { ColumnNames } from "./columnHeader";
// import Modal from '@/lib/Modal.vue'
import { Label } from "@/components/ui/label";
import TableComponent from "@/lib/Table.vue";
import CreateUser from "./modal/CreateUser.vue";
import UpdateUser from "./modal/UpdateUser.vue";
import DeleteUser from "./modal/DeleteUser.vue";
import UpdateStatus from "./modal/UpdateStatus.vue";
import { useUserStore } from "../../stores/userStore";
import UserRow from "./row.vue";

const userStore = useUserStore();
userStore.getUserData();

const users = ref([
  {
    username: "username1",
    createTime: "2014-12-24 23:12:00",
    updateTime: "2014-12-25 23:12:00",
    operator: "admin",
    status: true,
  },
  {
    username: "username2",
    createTime: "2014-12-24 23:12:00",
    updateTime: "2014-12-25 23:12:00",
    operator: "username1",
    status: true,
  },
  {
    username: "username3",
    createTime: "2014-12-24 23:12:00",
    updateTime: "2014-12-25 23:12:00",
    operator: "username2",
    status: false,
  },
]);
</script>
