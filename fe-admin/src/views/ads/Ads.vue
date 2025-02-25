<template>
    <div class="p-8">
        <Component 
            title="Ads"
            :totalItems="filteredAds.length"
            :currentPage="currentPage"
            :pageSize="pageSize"
            :displayedPages="displayedPages"
            :onAddClick="handleAdd"
            @update:pageSize="pageSize = $event"
        />
        
        <Table 
            :subjects="paginatedAds" 
            :columns="columns"
            @edit="handleEdit" 
            @delete="handleDelete" 
        />

        <!-- Use the new Pagination component here -->
        <Pagination 
            :totalItems="filteredAds.length"
            :currentPage="currentPage"
            :totalPages="totalPages"
            :displayedPages="displayedPages"
        />
        
        <!-- Reusable Modal for adding a new ad -->
        <Modal 
            :isVisible="showAddModal" 
            title="Add New Ad" 
            description="Fill in the details for the new ad."
            :newSubject="newAd"
            @update:isVisible="showAddModal = $event"
            @confirm="saveNewAd"
        >
            <template #default="{ newSubject }">
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Title</Label>
                                <Input v-model="newAd.title" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Image Upload</Label>
                                <Input type="file" @change="handleFileUpload" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Display Start At</Label>
                                <Input type="datetime-local" v-model="newAd.displayStartAt" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Display End At</Label>
                                <Input type="datetime-local" v-model="newAd.displayEndAt" class="col-span-3" />
                            </div>
                        </div>
                        <div class="space-y-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">In/Off Site</Label>
                                <div class="col-span-3">
                                    <Select v-model="newAd.inOffSite">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select option" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="in-site">In-Site</SelectItem>
                                            <SelectItem value="off-site">Off-Site</SelectItem>
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Language</Label>
                                <div class="col-span-3">
                                    <Select v-model="newAd.language">
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
                                <Label class="text-right">Jump Link</Label>
                                <Input v-model="newAd.jumpLink" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Status</Label>
                                <div class="col-span-3">
                                    <Select v-model="newAd.status" class="col-span-3">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select status" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="active">display</SelectItem>
                                            <SelectItem value="inactive">not display</SelectItem>
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
// import { Table } from '@/lib/Table.vue'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { Input } from '@/components/ui/input'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import Component from '@/lib/Component.vue'
import Pagination from '@/lib/Pagination.vue' // Import the new Pagination component
import Modal from '@/lib/Modal.vue' // Import the reusable Modal component
import { Label } from '@/components/ui/label'
import Table from '@/lib/Table.vue'

const columns = [
    { key: 'title', label: 'Title' },
    { key: 'pic', label: 'Picture' },
    { key: 'displayStartAt', label: 'Start Date' },
    { key: 'displayEndAt', label: 'End Date' },
    { key: 'inOffSite', label: 'Location' },
    { key: 'jumpLink', label: 'Link' },
    { key: 'status', label: 'Status' }
]

const ads = ref([
    {
        title: 'ads1',
        pic: 'https://example.com/image1.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-25 23:12:00',
        inOffSite: 'in-site',
        jumpLink: '/comics/ss',
        status: true
    },
    {
        title: 'ads2',
        pic: 'https://example.com/image2.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-25 23:12:00',
        inOffSite: 'off-site',
        jumpLink: 'www.google.com',
        status: true
    },
    {
        title: 'ads3',
        pic: 'https://example.com/image3.jpg',
        displayStartAt: '2014-12-24 23:12:00',
        displayEndAt: '2014-12-25 23:12:00',
        inOffSite: 'in-site',
        jumpLink: '/comics/ss',
        status: true
    }
    // Add more ad data as needed
])

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref('10')
const showAddModal = ref(false)
const newAd = ref({
    title: '',
    pic: '',
    displayStartAt: '',
    displayEndAt: '',
    inOffSite: '',
    jumpLink: '',
    language: '',
    status: 'active'
})

const filteredAds = computed(() => {
    return ads.value.filter(ad => 
        ad.title.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const totalPages = computed(() => Math.ceil(filteredAds.value.length / parseInt(pageSize.value)))

const paginatedAds = computed(() => {
    const start = (currentPage.value - 1) * Number(pageSize.value)
    const end = start + Number(pageSize.value)
    return filteredAds.value.slice(start, end)
})

const displayedPages = computed(() => {
    const pages = []
    for (let i = 1; i <= totalPages.value; i++) {
        pages.push(i)
    }
    return pages
})

const handleAdd = () => {
    newAd.value = {
        title: '',
        pic: '',
        displayStartAt: '',
        displayEndAt: '',
        inOffSite: '',
        jumpLink: '',
        language: '',
        status: 'active'
    }
    showAddModal.value = true
}

const saveNewAd = () => {
    // Logic to save the new ad
    console.log('Saving new ad:', newAd.value)
    ads.value.push({ ...newAd.value }) // Add the new ad to the list
    showAddModal.value = false
}

const handleEdit = (ad) => {
    // Handle edit logic
}

const handleDelete = (ad) => {
    // Handle delete logic
    ads.value = ads.value.filter(a => a.title !== ad.title)
    console.log('Delete:', ad.title)
}

const handleFileUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
            newAd.value.pic = e.target.result; // Set the image URL to the base64 string
        };
        reader.readAsDataURL(file);
    }
}
</script>