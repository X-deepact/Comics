<template>
    <div class="p-8">
        <Component 
            title="Comics Management"
            :totalItems="filteredComics.length"
            :currentPage="currentPage"
            :pageSize="pageSize"
            :displayedPages="displayedPages"
            :onAddClick="handleAdd"
            @update:pageSize="pageSize = $event"
        />
        
        <TableComponent 
            :subjects="paginatedComics" 
            :columns="columns"
            @edit="handleEdit" 
            @delete="handleDelete" 
        />
        
        <Pagination 
            :totalItems="filteredComics.length"
            :currentPage="currentPage"
            :totalPages="totalPages"
            :displayedPages="displayedPages"
        />
        
        <Modal 
            :isVisible="showAddModal" 
            title="Add New Comic" 
            description="Fill in the details for the new comic."
            :newSubject="newComic"
            @update:isVisible="showAddModal = $event"
            @confirm="saveNewComic"
        >
            <template #default="{ newSubject }">
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
            </template>
        </Modal>
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
import Component from '@/lib/Component.vue'
import TableComponent from '@/lib/Table.vue'
import Pagination from '@/lib/Pagination.vue'
import Modal from '@/lib/Modal.vue'

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

const saveNewComic = async () => {
    // Prepare the comic data for the API request
    const comicData = {
        active: true,
        audience: newComic.value.audience,
        code: '', // Add logic to generate or assign a code if necessary
        cover: newComic.value.cover,
        description: newComic.value.introduction,
        language: newComic.value.language,
        title: newComic.value.title
    };

    try {
        // Make the API call to create a new comic
        const response = await fetch('http://localhost:8080/api/comic/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(comicData),
        });

        if (!response.ok) {
            throw new Error('Failed to create comic');
        }

        const result = await response.json();
        console.log('Comic created successfully:', result);
        
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
        };
        showAddModal.value = false;
    } catch (error) {
        console.error('Error creating comic:', error);
        // Handle error (e.g., show a notification to the user)
    }
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
    console.log('Status changed:', comic.title, comic.status)
}

watch(newComic.subject, (newValue) => {
    console.log('Selected subjects:', newValue);
});

const availableSubjects = ['Action', 'Adventure', 'Comedy', 'Drama', 'Fantasy', 'Horror', 'Romance', 'Sci-Fi', 'Thriller'];

const filteredComics = computed(() => {
    return comics.value.filter(comic => 
        comic.title.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const paginatedComics = computed(() => {
    const start = (currentPage.value - 1) * parseInt(pageSize.value)
    return filteredComics.value.slice(start, start + parseInt(pageSize.value))
})
</script>