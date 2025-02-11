<template>
    <div class="p-8">
        <Component 
            title="Subject Management"
            :totalItems="filteredSubjects.length"
            :currentPage="currentPage"
            :pageSize="pageSize"
            :displayedPages="displayedPages"
            :onAddClick="handleAdd"
            @update:pageSize="pageSize = $event"
        />
        <Table 
            :subjects="paginatedSubjects" 
            :columns="columns"
            @edit="handleEdit" 
            @delete="handleDelete" 
        />
        <!-- Use the new Pagination component here -->
        <Pagination 
            :totalItems="filteredSubjects.length"
            :currentPage="currentPage"
            :totalPages="totalPages"
            :displayedPages="displayedPages"
        />
        
        <!-- Reusable Modal for adding a new subject -->
        <Modal 
            :isVisible="showAddModal" 
            title="Add New Subject" 
            description="Fill in the details for the new subject."
            :newSubject="newSubject"
            @update:isVisible="showAddModal = $event"
            @confirm="saveNewSubject"
        >
            <template #default="{ newSubject }">
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Name</Label>
                                <Input v-model="newSubject.name" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Create Time</Label>
                                <Input type="datetime-local" v-model="newSubject.createTime" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Operator</Label>
                                <Input v-model="newSubject.operator" class="col-span-3" />
                            </div>
                    </div>
                </div>
            </template>
        </Modal>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import Component from '@/lib/Component.vue'
import Table from '@/lib/Table.vue'
import Pagination from '@/lib/Pagination.vue' // Import the new Pagination component
import Modal from '@/lib/Modal.vue' // Import the reusable Modal component
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'

const subjects = ref([
    {
        name: 'The Hulk',
        createTime: '2014-12-24 23:12:00',
        operator: 'Stan Lee'
    }
    // Add more subject data as needed
])

const columns = ref([ // Define the columns structure
    { key: 'name', label: 'Name' },
    { key: 'createTime', label: 'Create Time' },
    { key: 'operator', label: 'Operator' }
    
])

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref('10')
const showAddModal = ref(false)
const newSubject = ref({
    name: '',
    createTime: '',
    operator: '',
    language: ''
})

const filteredSubjects = computed(() => {
    return subjects.value.filter(subject => 
        subject.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const totalPages = computed(() => Math.ceil(filteredSubjects.value.length / parseInt(pageSize.value)))

const paginatedSubjects = computed(() => {
    const start = (currentPage.value - 1) * parseInt(pageSize.value)
    return filteredSubjects.value.slice(start, start + parseInt(pageSize.value))
})

const displayedPages = computed(() => {
    const pages = []
    for (let i = 1; i <= totalPages.value; i++) {
        pages.push(i)
    }
    return pages
})

const handleAdd = () => {
    newSubject.value = { name: '', createTime: '', operator: '' }
    showAddModal.value = true
}

const saveNewSubject = () => {
    subjects.value.push({ ...newSubject.value })
    newSubject.value = { name: '', createTime: '', operator: '', language: '' }
    showAddModal.value = false
}

const handleEdit = (subject) => {
    console.log('Edit:', subject.name)
}

const handleDelete = (subject) => {
    subjects.value = subjects.value.filter(s => s.name !== subject.name)
    console.log('Delete:', subject.name)
}
</script>