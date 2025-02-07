<template>
    <div class="p-8">
        <h1 class="text-2xl font-semibold mb-6">Recommendation Management</h1>
        
        <!-- Add Recommendation Button -->
        <div class="flex justify-between items-center mb-4">
            <Button @click="handleAdd" class="bg-blue-500 hover:bg-blue-600 text-white">Add Recommendation</Button>
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
                <TableRow v-for="(recommendation, index) in recommendations" :key="index">
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
                Total {{ totalItems }} items
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

        <!-- Add your content here -->
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'

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
const totalItems = ref(recommendations.value.length) // Total count of recommendations

const totalPages = computed(() => Math.ceil(totalItems.value / parseInt(pageSize.value)))

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

const handleAdd = () => {
    console.log('Add new recommendation')
    // Logic to open a modal or form to add a new recommendation
}
</script> 