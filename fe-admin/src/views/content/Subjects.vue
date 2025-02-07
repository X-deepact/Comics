<template>
    <div class="p-8">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-semibold">Subject Management</h1>
            <Button @click="showAddModal = true" class="bg-blue-500 hover:bg-blue-600 text-white">Add Subject</Button>
        </div>

        <div class="flex justify-between items-center mb-4 gap-4">
            <div class="flex items-center gap-2">
                <Input 
                    v-model="searchQuery"
                    placeholder="Search subjects..."
                    class="w-[280px]"
                />
            </div>
            <div class="flex items-center gap-4">
                <Select v-model="pageSize">
                    <SelectTrigger class="w-[120px]">
                        <SelectValue placeholder="Select size" />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                        <SelectItem value="10">10 / page</SelectItem>
                        <SelectItem value="20">20 / page</SelectItem>
                        <SelectItem value="30">30 / page</SelectItem>
                        <SelectItem value="50">50 / page</SelectItem>
                    </SelectContent>
                </Select>
            </div>
        </div>

        <Table class="border rounded-lg">
            <TableHeader>
                <TableRow>
                    <TableHead>no.</TableHead>
                    <TableHead>name</TableHead>
                    <TableHead>createtime</TableHead>
                    <TableHead>operator</TableHead>
                    <TableHead>operate</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(subject, index) in paginatedSubjects" :key="index">
                    <TableCell>{{ index + 1 }}</TableCell>
                    <TableCell>{{ subject.name }}</TableCell>
                    <TableCell>{{ subject.createTime }}</TableCell>
                    <TableCell>{{ subject.operator }}</TableCell>
                    <TableCell>
                        <div class="flex space-x-2">
                            <Button variant="ghost" size="sm" @click="handleEdit(subject)">
                                edit
                            </Button>
                            <Button variant="ghost" size="sm" class="text-red-500 hover:text-red-700" @click="handleDelete(subject)">
                                delete
                            </Button>
                        </div>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>

        <div class="flex items-center justify-between mt-4">
            <div class="text-sm text-gray-500">
                Total {{ filteredSubjects.length }} items
            </div>
            <div class="flex items-center gap-2">
                <Button 
                    variant="outline" 
                    size="sm"
                    :disabled="currentPage === 1"
                    @click="currentPage--"
                >
                    Previous
                </Button>
                <div class="flex items-center gap-1">
                    <Button 
                        v-for="pageNumber in displayedPages" 
                        :key="pageNumber"
                        variant="outline"
                        size="sm"
                        :class="{ 'bg-primary text-primary-foreground': pageNumber === currentPage }"
                        @click="currentPage = pageNumber"
                    >
                        {{ pageNumber === '...' ? pageNumber : pageNumber }}
                    </Button>
                </div>
                <Button 
                    variant="outline" 
                    size="sm"
                    :disabled="currentPage === totalPages"
                    @click="currentPage++"
                >
                    Next
                </Button>
            </div>
        </div>

        <Dialog v-model:open="showAddModal" class="fixed inset-0 z-50">
            <DialogContent class="fixed top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] bg-white p-6 rounded-lg shadow-lg sm:max-w-[800px] w-full">
                <button 
                    class="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                    @click="showAddModal = false"
                >
                    <X class="h-4 w-4" />
                    <span class="sr-only">Close</span>
                </button>
                <DialogHeader>
                    <DialogTitle>Add New Subject</DialogTitle>
                    <DialogDescription>
                        Fill in the details for the new subject. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Name</Label>
                                <Input v-model="newSubject.name" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Create Time</Label>
                                <Input v-model="newSubject.createTime" type="datetime-local" class="col-span-3" />
                            </div>
                        </div>
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Operator</Label>
                                <Input v-model="newSubject.operator" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Language</Label>
                                <div class="col-span-3">
                                    <Select v-model="newSubject.language">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select language" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="en">English</SelectItem>
                                            <SelectItem value="zh">Chinese</SelectItem>
                                            <SelectItem value="vi">Vietnamese</SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
                <DialogFooter>
                    <Button variant="outline" @click="showAddModal = false">Cancel</Button>
                    <Button @click="saveNewSubject">Save</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    </div>
</template>

<script setup>
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { ref, computed } from 'vue'

const subjects = ref([
    {
        name: 'funny',
        createTime: '2014-12-24 23:12:00',
        operator: 'adminname'
    }
    // Add more subject data as needed
])

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref('10')
const showAddModal = ref(false)
const newSubject = ref({
    name: '',
    createTime: '',
    operator: '',
    language: '' // Added language property
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

const handleAdd = () => {
    showAddModal.value = true
}

const saveNewSubject = () => {
    subjects.value.push({ ...newSubject.value }) // Add the new subject to the list
    newSubject.value = { name: '', createTime: '', operator: '', language: '' } // Reset the form
    showAddModal.value = false // Close the modal
}

const handleEdit = (subject) => {
    console.log('Edit:', subject.name)
}

const handleDelete = (subject) => {
    console.log('Delete:', subject.name)
}
</script> 