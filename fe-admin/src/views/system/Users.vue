<template>
    <div class="p-8">
        <Component 
            title="Users"
            :totalItems="filteredUsers.length"
            :currentPage="currentPage"
            :pageSize="pageSize"
            :displayedPages="displayedPages"
            :onAddClick="handleAdd"
            @update:pageSize="pageSize = $event"
        />
        
        <Table 
            :subjects="paginatedUsers" 
            :columns="columns"
            @edit="handleEdit" 
            @delete="handleDelete" 
        />

        <Pagination 
            :totalItems="filteredUsers.length"
            :currentPage="currentPage"
            :totalPages="totalPages"
            :displayedPages="displayedPages"
        />
        
        <Modal 
            :isVisible="showAddModal" 
            title="Add New User" 
            description="Fill in the details for the new user."
            :newSubject="newUser"
            @update:isVisible="showAddModal = $event"
            @confirm="saveNewUser"
        >
            <template #default="{ newSubject }">
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Username</Label>
                                <Input v-model="newUser.username" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Create Time</Label>
                                <Input type="datetime-local" v-model="newUser.createTime" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Update Time</Label>
                                <Input type="datetime-local" v-model="newUser.updateTime" class="col-span-3" />
                            </div>
                        </div>
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Operator</Label>
                                <Input v-model="newUser.operator" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Status</Label>
                                <div class="col-span-3">
                                    <Select v-model="newUser.status">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select status" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem :value="true">Active</SelectItem>
                                            <SelectItem :value="false">Inactive</SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </Modal>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { Input } from '@/components/ui/input'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import Component from '@/lib/Component.vue'
import Pagination from '@/lib/Pagination.vue'
import { Label } from '@/components/ui/label'
import Table from '@/lib/Table.vue'

const columns = [
    { key: 'username', label: 'Username' },
    { key: 'createTime', label: 'Create Time' },
    { key: 'updateTime', label: 'Update Time' },
    { key: 'operator', label: 'Operator' },
    { key: 'status', label: 'Status' }
]

const users = ref([
    {
        username: 'username1',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'admin',
        status: true
    },
    {
        username: 'username2',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'username1',
        status: true
    },
    {
        username: 'username3',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'username2',
        status: false
    }
])

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref('10')
const showAddModal = ref(false)
const newUser = ref({
    username: '',
    createTime: '',
    updateTime: '',
    operator: '',
    status: true
})

const filteredUsers = computed(() => {
    return users.value.filter(user => 
        user.username.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const totalPages = computed(() => Math.ceil(filteredUsers.value.length / parseInt(pageSize.value)))

const paginatedUsers = computed(() => {
    const start = (currentPage.value - 1) * Number(pageSize.value)
    const end = start + Number(pageSize.value)
    return filteredUsers.value.slice(start, end)
})

const displayedPages = computed(() => {
    const pages = []
    for (let i = 1; i <= totalPages.value; i++) {
        pages.push(i)
    }
    return pages
})

const handleAdd = () => {
    newUser.value = {
        username: '',
        createTime: '',
        updateTime: '',
        operator: '',
        status: true
    }
    showAddModal.value = true
}

const saveNewUser = () => {
    console.log('Saving new user:', newUser.value)
    users.value.push({ ...newUser.value })
    showAddModal.value = false
}

const handleEdit = (user) => {
    console.log('Edit:', user.username)
}

const handleDelete = (user) => {
    users.value = users.value.filter(u => u.username !== user.username)
    console.log('Delete:', user.username)
}
</script>