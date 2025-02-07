<template>
    <div class="p-8">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-semibold">Ads Management</h1>
            <Button @click="showAddModal = true" class="bg-blue-500 hover:bg-blue-600 text-white">Add Ad</Button>
        </div>

        <div class="flex justify-between items-center mb-4 gap-4">
            <div class="flex items-center gap-2">
                <Input 
                    v-model="searchQuery"
                    placeholder="Search ads..."
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
                    <TableHead>pic</TableHead>
                    <TableHead>display start at</TableHead>
                    <TableHead>display end at</TableHead>
                    <TableHead>in/off-site</TableHead>
                    <TableHead>jump link</TableHead>
                    <TableHead>status</TableHead>
                    <TableHead>operate</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(ad, index) in ads" :key="index">
                    <TableCell>{{ index + 1 }}</TableCell>
                    <TableCell>{{ ad.title }}</TableCell>
                    <TableCell>
                        <img :src="ad.pic" alt="pic" class="w-12 h-12 object-cover"/>
                    </TableCell>
                    <TableCell>{{ ad.displayStartAt }}</TableCell>
                    <TableCell>{{ ad.displayEndAt }}</TableCell>
                    <TableCell>{{ ad.inOffSite }}</TableCell>
                    <TableCell>{{ ad.jumpLink }}</TableCell>
                    <TableCell>
                        <Switch 
                            v-model="ad.status" 
                            @update:modelValue="handleStatusChange(ad)"
                            class="data-[state=checked]:bg-primary"
                        >
                            <span class="sr-only">Toggle status</span>
                        </Switch>
                    </TableCell>
                    <TableCell>
                        <div class="flex space-x-2">
                            <Button variant="ghost" size="sm" @click="handleEdit(ad)">
                                edit
                            </Button>
                            <Button variant="ghost" size="sm" class="text-red-500 hover:text-red-700" @click="handleDelete(ad)">
                                delete
                            </Button>
                        </div>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>

        <div class="flex items-center justify-between mt-4">
            <div class="text-sm text-gray-500">
                Total {{ totalAds }} ads
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
                    <DialogTitle>Add New Ad</DialogTitle>
                    <DialogDescription>
                        Fill in the details for the new ad. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
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
                <DialogFooter>
                    <Button variant="outline" @click="showAddModal = false">Cancel</Button>
                    <Button @click="saveNewAd">Save</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    </div>
</template>

<script setup>
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { Input } from '@/components/ui/input'
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select'
import { ref, computed } from 'vue'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from '@/components/ui/dialog'
import { X } from 'lucide-vue-next'
import { Label } from '@/components/ui/label'

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
const totalAds = ref(100) // Replace with actual total count

const displayedPages = computed(() => {
    const pages = []
    const maxVisiblePages = 7
    const totalPages = Math.ceil(totalAds.value / parseInt(pageSize.value))

    // ... logic for pagination ...

    return pages
})

const handleEdit = (ad) => {
    console.log('Edit:', ad.title)
}

const handleDelete = (ad) => {
    console.log('Delete:', ad.title)
}

const handleStatusChange = (ad) => {
    console.log('Status changed:', ad.title, ad.status)
}

const showAddModal = ref(false)
const newAd = ref({
    title: '',
    pic: '',
    displayStartAt: '',
    displayEndAt: '',
    inOffSite: '',
    jumpLink: ''
})

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

const saveNewAd = () => {
    // Logic to save the new ad
    console.log('Saving new ad:', newAd.value)
    // Reset form and close modal
    newAd.value = {
        title: '',
        pic: '',
        displayStartAt: '',
        displayEndAt: '',
        inOffSite: '',
        jumpLink: ''
    }
    showAddModal.value = false
}
</script> 