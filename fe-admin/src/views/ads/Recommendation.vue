<template>
    <div class="p-8">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-semibold">Recommendation Management</h1>
            <Button @click="handleAdd" class="bg-blue-500 hover:bg-blue-600 text-white">Add</Button>
        </div>

        <div class="flex justify-between items-center mb-4 gap-4">
            <div class="flex items-center gap-2">
                <Input 
                    v-model="searchQuery"
                    placeholder="Search title..."
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
                    <TableHead>title</TableHead>
                    <TableHead>author</TableHead>
                    <TableHead>cover</TableHead>
                    <TableHead>display start at</TableHead>
                    <TableHead>display end at</TableHead>
                    <TableHead>position</TableHead>
                    <TableHead>operate</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(recommendation, index) in filteredRecommendations" :key="index">
                    <TableCell>{{ index + 1 }}</TableCell>
                    <TableCell>{{ recommendation.title }}</TableCell>
                    <TableCell>{{ recommendation.author }}</TableCell>
                    <TableCell>
                        <img :src="recommendation.cover" alt="cover" class="w-12 h-12 object-cover"/>
                    </TableCell>
                    <TableCell>{{ recommendation.displayStartAt }}</TableCell>
                    <TableCell>{{ recommendation.displayEndAt }}</TableCell>
                    <TableCell>{{ recommendation.position }}</TableCell>
                    <TableCell>
                        <div class="flex space-x-2">
                            <Button variant="ghost" size="sm" @click="handleEdit(recommendation)">
                                edit
                            </Button>
                            <Button variant="ghost" size="sm" class="text-red-500 hover:text-red-700" @click="handleDelete(recommendation)">
                                delete
                            </Button>
                        </div>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>

        <!-- Pagination -->
        <div class="flex items-center justify-between mt-4">
            <div class="text-sm text-gray-500">
                Total {{ filteredRecommendations.length }} items
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

        <!-- Add your modal here -->
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
                    <DialogTitle>Add New Recommendation</DialogTitle>
                    <DialogDescription>
                        Fill in the details for the new recommendation. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">

                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Title</Label>
                                <Input v-model="newRecommendation.title" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Author</Label>
                                <Input v-model="newRecommendation.author" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Language</Label>
                                <div class="col-span-3">
                                    <Select v-model="newRecommendation.language">
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
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Cover Image</Label>
                                <Input type="file" @change="handleFileUpload" class="col-span-3" />
                            </div>
                        </div>
                    <div class="space-y-4">
                        <div class="grid grid-cols-4 items-center gap-4">
                            <Label class="text-right">Display Start At</Label>
                            <Input type="datetime-local" v-model="newRecommendation.displayStartAt" class="col-span-3" />
                        </div>
                        <div class="grid grid-cols-4 items-center gap-4">
                            <Label class="text-right">Display End At</Label>
                            <Input type="datetime-local" v-model="newRecommendation.displayEndAt" class="col-span-3" />
                        </div>
                        <div class="grid grid-cols-4 items-center gap-4">
                            <Label class="text-right">Position</Label>
                            <Input v-model="newRecommendation.position" class="col-span-3" />
                        </div>
                    </div>
                    </div>
                </div>
                <DialogFooter>
                    <Button variant="outline" @click="showAddModal = false">Cancel</Button>
                    <Button @click="saveNewRecommendation">Save</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Input } from '@/components/ui/input'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from '@/components/ui/dialog'
import { X } from 'lucide-vue-next'
import { Label } from '@/components/ui/label'

const recommendations = ref([
    {
        title: 'The Hulk',
        author: 'Stan Lee',
        cover: 'https://example.com/hulk.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-24 23:12:00',
        position: 'top, fastest-growing'
    },
    {
        title: 'The Red',
        author: 'Stan Lee',
        cover: 'https://example.com/red.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-24 23:12:00',
        position: 'recommended'
    },
    {
        title: 'The Blue',
        author: 'Stan Lee',
        cover: 'https://example.com/blue.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-24 23:12:00',
        position: 'fastest-growing'
    }
    // Add more recommendations as needed
])

const currentPage = ref(1)
const pageSize = ref('10')
const searchQuery = ref('')
const totalItems = ref(recommendations.value.length) // Total count of recommendations

const filteredRecommendations = computed(() => {
    return recommendations.value.filter(recommendation => 
        recommendation.title.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const totalPages = computed(() => Math.ceil(filteredRecommendations.value.length / parseInt(pageSize.value)))

const displayedPages = computed(() => {
    const pages = []
    const maxVisiblePages = 7
    
    if (totalPages.value <= maxVisiblePages) {
        for (let i = 1; i <= totalPages.value; i++) {
            pages.push(i)
        }
    } else {
        pages.push(1)
        if (currentPage.value > 3) pages.push('...')
        let start = Math.max(2, currentPage.value - 1)
        let end = Math.min(totalPages.value - 1, currentPage.value + 1)
        if (currentPage.value <= 3) end = 4
        if (currentPage.value >= totalPages.value - 2) start = totalPages.value - 3
        for (let i = start; i <= end; i++) pages.push(i)
        if (currentPage.value < totalPages.value - 2) pages.push('...')
        pages.push(totalPages.value)
    }
    return pages
})

const handleEdit = (recommendation) => {
    console.log('Edit:', recommendation.title)
}

const handleDelete = (recommendation) => {
    console.log('Delete:', recommendation.title)
}

const showAddModal = ref(false)
const newRecommendation = ref({
    title: '',
    author: '',
    cover: '',
    displayStartAt: '',
    displayEndAt: '',
    position: '',
    // Add other fields as necessary
})

const handleAdd = () => {
    showAddModal.value = true
}

const saveNewRecommendation = () => {
    console.log('Saving new recommendation:', newRecommendation.value)
    // Logic to save the new recommendation
    // Reset form and close modal
    newRecommendation.value = {
        title: '',
        author: '',
        cover: '',
        displayStartAt: '',
        displayEndAt: '',
        position: '',
        // Reset other fields as necessary
    }
    showAddModal.value = false
}
</script> 