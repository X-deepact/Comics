<template>
    <div class="p-8">
        <h1 class="text-2xl font-semibold mb-6">Subject Management</h1>
        
        <div class="flex justify-between items-center mb-4">
            <Button @click="showAddModal = true" class="bg-blue-500 hover:bg-blue-600 text-white">Add Comic</Button>
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
                <TableRow v-for="(subject, index) in subjects" :key="index">
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
                    <DialogTitle>Add New Comic</DialogTitle>
                    <DialogDescription>
                        Fill in the details for the new comic. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Title</Label>
                                <Input v-model="newComic.title" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Author</Label>
                                <Input v-model="newComic.author" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Introduction</Label>
                                <Textarea 
                                    v-model="newComic.introduction" 
                                    class="col-span-3"
                                    rows="6" 
                                />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Cover</Label>
                                <div class="col-span-3">
                                    <Input type="file" accept="image/*" @change="handleCoverUpload" />
                                    <div v-if="newComic.cover" class="mt-2">
                                        <img :src="newComic.cover" alt="Preview" class="w-32 h-32 object-cover" />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <DialogFooter>
                    <Button variant="outline" @click="showAddModal = false">Cancel</Button>
                    <Button @click="saveNewComic">Save</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    </div>
</template>

<script setup>
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Select } from '@/components/ui/select'
import { Switch } from '@/components/ui/switch'
import { ref } from 'vue'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { X } from 'lucide-vue-next'

const subjects = ref([
    {
        name: 'funny',
        createTime: '2014-12-24 23:12:00',
        operator: 'adminname'
    }
    // Add more subject data as needed
])

const showAddModal = ref(false)
const newComic = ref({
    title: '',
    author: '',
    introduction: '',
    cover: null,
})

const handleEdit = (subject) => {
    console.log('Edit:', subject.name)
}

const handleDelete = (subject) => {
    console.log('Delete:', subject.name)
}

const saveNewComic = () => {
    // Add your logic to save the new comic here
    console.log('Saving new comic:', newComic.value)
    // Reset form and close modal
    newComic.value = {
        title: '',
        author: '',
        introduction: '',
        cover: null,
    }
    showAddModal.value = false
}

const handleCoverUpload = (event) => {
    const file = event.target.files[0]
    if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
            newComic.value.cover = e.target.result
        }
        reader.readAsDataURL(file)
    }
}
</script> 