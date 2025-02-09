<template>
    <div class="p-8">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-semibold">Comics Management</h1>
            <Button @click="handleAdd" class="bg-blue-500 hover:bg-blue-600 text-white">Add Comic</Button>
        </div>

        <div class="flex justify-between items-center mb-4 gap-4">
            <div class="flex items-center gap-2">
                <Input 
                    v-model="searchQuery"
                    placeholder="Search comics..."
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
                    <TableHead>introduction</TableHead>
                    <TableHead>update time</TableHead>
                    <TableHead>subject</TableHead>
                    <TableHead>progress</TableHead>
                    <TableHead>status</TableHead>
                    <TableHead>operate</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(comic, index) in comics" :key="index">
                    <TableCell>{{ index + 1 }}</TableCell>
                    <TableCell>{{ comic.title }}</TableCell>
                    <TableCell>{{ comic.author }}</TableCell>
                    <TableCell>
                        <img :src="comic.cover" alt="cover" class="w-12 h-12 object-cover"/>
                    </TableCell>
                    <TableCell>{{ comic.introduction }}</TableCell>
                    <TableCell>{{ comic.updateTime }}</TableCell>
                    <TableCell>
                        <div class="flex flex-wrap gap-1">
                            <span v-for="(tag, i) in comic.subject" 
                                  :key="i" 
                                  class="px-2 py-1 text-xs rounded-full bg-gray-100">
                                {{ tag }}
                            </span>
                        </div>
                    </TableCell>
                    <TableCell>{{ comic.progress }}</TableCell>
                    <TableCell>
                        <label class="switch">
                            <input 
                                type="checkbox" 
                                v-model="comic.status" 
                                @change="handleStatusChange(comic)"
                            />
                            <span class="slider"></span>
                        </label>
                    </TableCell>
                    <TableCell>
                        <div class="flex space-x-2">
                            <Button variant="ghost" size="sm" @click="handleChapters(comic)">
                                chapters
                            </Button>
                            <Button variant="ghost" size="sm" @click="handleEdit(comic)">
                                edit
                            </Button>
                            <Button variant="ghost" size="sm" class="text-red-500 hover:text-red-700" @click="handleDelete(comic)">
                                delete
                            </Button>
                        </div>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>

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
                                <Label class="text-right">Subject</Label>
                                <div class="col-span-3">
                                    <Select v-model="newComic.subject" multiple>
                                        <SelectTrigger>
                                            <SelectValue 
                                                :placeholder="newComic.subject.length ? '' : 'Select subjects'" 
                                                :value="newComic.subject.join(', ')" 
                                            />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem v-for="subject in availableSubjects" :key="subject" :value="subject">
                                                {{ subject }}
                                            </SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Original Language</Label>
                                <div class="col-span-3">
                                    <Select v-model="newComic.language">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select language" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <Input 
                                                v-model="languageSearchQuery" 
                                                @input="handleSearchInput" 
                                                placeholder="Search language..." 
                                                class="p-2 mb-2"
                                            />
                                            <SelectItem v-for="language in filteredCountries" :key="language.value" :value="language.value">
                                                {{ language.label }}
                                            </SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Progress</Label>
                                <div class="col-span-3">
                                    <Select v-model="newComic.progress">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select progress" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="ongoing">Ongoing</SelectItem>
                                            <SelectItem value="stopped">Completed</SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Audience</Label>
                                <div class="col-span-3">
                                    <Select v-model="newComic.audience">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select audience" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="all">Boy's Comics</SelectItem>
                                            <SelectItem value="teen">Girl's Comics</SelectItem>
                                            <SelectItem value="teen">Children's Comics</SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Language</Label>
                                <div class="col-span-3">
                                    <Select v-model="newComic.language">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select language" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem v-for="lang in languages" 
                                                       :key="lang.value" 
                                                       :value="lang.value">
                                                {{ lang.label }}
                                            </SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Update Time</Label>
                                <div class="col-span-3">
                                    <Input 
                                        type="datetime-local" 
                                        v-model="newComic.updateTime"
                                        class="w-full"
                                    />
                                </div>
                            </div>
                        </div>
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Cover</Label>
                                <div class="col-span-3">
                                    <Input type="file" accept="image/*" @change="handleCoverUpload" />
                                    <div v-if="newComic.cover" class="mt-2">
                                        <img :src="newComic.cover" alt="Preview" class="w-32 h-32 object-cover" />
                                    </div>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Introduction</Label>
                                <Textarea 
                                    v-model="newComic.introduction" 
                                    class="col-span-3"
                                    rows="6" 
                                />
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
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Switch } from '@/components/ui/switch'
import { ref, computed, watch } from 'vue'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { X } from 'lucide-vue-next'

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref('10')
const totalItems = ref(100) // Replace with actual total count

const comics = ref([
    {
        title: 'The hulk',
        author: 'Stan lee',
        cover: 'https://avatars.githubusercontent.com/u/186627792?v=4',
        introduction: 'At the end of the 1960s...',
        updateTime: '2014-12-24 23:12:00',
        subject: ['suspense', 'thriller'],
        progress: 'ongoing',
        status: true,
        isOnline: true
    }
    // Add more comic data as needed
])

const totalPages = computed(() => Math.ceil(totalItems.value / parseInt(pageSize.value)))

const displayedPages = computed(() => {
    const pages = []
    const maxVisiblePages = 7
    
    if (totalPages.value <= maxVisiblePages) {
        // If total pages is less than max visible, show all pages
        for (let i = 1; i <= totalPages.value; i++) {
            pages.push(i)
        }
    } else {
        // Always show first page
        pages.push(1)
        
        if (currentPage.value > 3) {
            pages.push('...')
        }
        
        // Calculate middle pages
        let start = Math.max(2, currentPage.value - 1)
        let end = Math.min(totalPages.value - 1, currentPage.value + 1)
        
        // Adjust if near the beginning or end
        if (currentPage.value <= 3) {
            end = 4
        }
        if (currentPage.value >= totalPages.value - 2) {
            start = totalPages.value - 3
        }
        
        for (let i = start; i <= end; i++) {
            pages.push(i)
        }
        
        if (currentPage.value < totalPages.value - 2) {
            pages.push('...')
        }
        
        // Always show last page
        pages.push(totalPages.value)
    }
    
    return pages
})

const showAddModal = ref(false)
const languages = [
    { value: 'en', label: 'English' },
    { value: 'zh', label: 'Chinese' },
    { value: 'vi', label: 'Vietnamese' },
]
const newComic = ref({
    title: '',
    author: '',
    introduction: '',
    subject: [],
    language: '',
    progress: '',
    status: true,
    cover: null,
    updateTime: '',
    audience: ''
})

const countries = ref([])

const languageSearchQuery = ref('')
const debounceTimeout = ref(null)

const allCountries = [
    { value: 'japan', label: 'Japanese' },
    { value: 'korea', label: 'Korean' },
    { value: 'china', label: 'Chinese' },
    { value: 'hong_kong', label: 'Chinese (Hong Kong & Taiwan)' },
    { value: 'english', label: 'English' },
]

const filteredCountries = computed(() => {
    return allCountries.filter(language => 
        language.label.toLowerCase().includes(languageSearchQuery.value.toLowerCase())
    )
})

// Debounce function
const handleSearchInput = (event) => {
    clearTimeout(debounceTimeout.value)
    debounceTimeout.value = setTimeout(() => {
        languageSearchQuery.value = event.target.value
    }, 300) // Adjust the delay as needed
}

const handleAdd = () => {
    showAddModal.value = true
}

const handleChapters = (comic) => {
    console.log('Chapters for:', comic.title)
}

const handleEdit = (comic) => {
    console.log('Edit:', comic.title)
}

const handleDelete = (comic) => {
    console.log('Delete:', comic.title)
}

const saveNewComic = () => {
    // Add your logic to save the new comic here
    console.log('Saving new comic:', newComic.value)
    // Reset form and close modal
    newComic.value = {
        title: '',
        author: '',
        introduction: '',
        subject: [],
        language: '',
        progress: '',
        status: true,
        cover: null,
        updateTime: '',
        audience: ''
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

const handleStatusChange = (comic) => {
    console.log('Status change triggered');
    console.log('Comic:', comic.title, 'New Status:', comic.status);
}

watch(newComic.subject, (newValue) => {
    console.log('Selected subjects:', newValue);
});

const availableSubjects = ['Action', 'Adventure', 'Comedy', 'Drama', 'Fantasy', 'Horror', 'Romance', 'Sci-Fi', 'Thriller'];
</script>

<style scoped>
.switch {
    position: relative;
    display: inline-block;
    width: 40px;
    height: 20px;
}

.switch input {
    opacity: 0;
    width: 0;
    height: 0;
}

.slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #ccc;
    transition: .4s;
    border-radius: 20px;
}

.slider:before {
    position: absolute;
    content: "";
    height: 20px;
    width: 20px;
    /* left: 1px; */
    bottom: 1px;
    background-color: white;
    transition: .4s;
    border-radius: 50%;
}

input:checked + .slider {
    background-color: #2196F3;
}

input:checked + .slider:before {
    transform: translateX(26px);
}
</style>