<template>
    <div class="p-8">
        <Component 
            title="Recommendation"
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

        <Pagination 
            :totalItems="filteredAds.length"
            :currentPage="currentPage"
            :totalPages="totalPages"
            :displayedPages="displayedPages"
        />
        
        <Modal 
            :isVisible="showAddModal" 
            title="Add New Recommend" 
            description="Fill in the details for the new ad."
            :newSubject="newAd"
            @update:isVisible="showAddModal = $event"
            @confirm="saveNewAd"
        >
            <template #default="{ newSubject }">
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-2 gap-4">
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Title</Label>
                                <Input v-model="newAd.title" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Image Upload</Label>
                                <Input type="file" @change="handleFileUpload" class="col-span-3" />
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Display Time</Label>
                                <div class="col-span-3 space-y-2">
                                    <Input type="datetime-local" v-model="newAd.displayStartAt" />
                                </div>
                            </div>
                            <div class="grid grid-cols-4 items-center gap-4">
                                <Label class="text-right">Location</Label>
                                <div class="col-span-3">
                                    <Select v-model="newAd.inOffSite" class="col-span-3">
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select location" />
                                        </SelectTrigger>
                                        <SelectContent class="bg-white">
                                            <SelectItem value="in-site">In-Site</SelectItem>
                                            <SelectItem value="off-site">Off-Site</SelectItem>
                                        </SelectContent>
                                    </Select>
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
import Component from '@/lib/Component.vue'
import Table from '@/lib/Table.vue'
import Pagination from '@/lib/Pagination.vue'
import Modal from '@/lib/Modal.vue'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'

const ads = ref([
    {
        title: 'ads1',
        pic: 'https://example.com/image1.jpg',
        displayStartAt: '2014-12-24 23:12:00', 
        displayEndAt: '2014-12-25 23:12:00',
        inOffSite: 'in-site',
        jumpLink: '/comics/ss',
        status: true
    }
    // Add more ads as needed
])

const columns = ref([
    { key: 'title', label: 'Title' },
    { key: 'pic', label: 'Picture' }, 
    { key: 'displayStartAt', label: 'Start Time' },
    { key: 'displayEndAt', label: 'End Time' },
    { key: 'inOffSite', label: 'Location' },
    { key: 'jumpLink', label: 'Link' },
    { key: 'status', label: 'Status' }
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
    status: true
})

const filteredAds = computed(() => {
    return ads.value.filter(ad => 
        ad.title.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
})

const totalPages = computed(() => Math.ceil(filteredAds.value.length / parseInt(pageSize.value)))

const paginatedAds = computed(() => {
    const start = (currentPage.value - 1) * parseInt(pageSize.value)
    return filteredAds.value.slice(start, start + parseInt(pageSize.value))
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
        status: true
    }
    showAddModal.value = true
}

const saveNewAd = () => {
    ads.value.push({ ...newAd.value })
    showAddModal.value = false
}

const handleEdit = (ad) => {
    console.log('Edit:', ad.title)
}

const handleDelete = (ad) => {
    ads.value = ads.value.filter(a => a.title !== ad.title)
    console.log('Delete:', ad.title)
}

const handleFileUpload = (event) => {
    const file = event.target.files[0]
    if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
            newAd.value.pic = e.target.result
        }
        reader.readAsDataURL(file)
    }
}
</script>