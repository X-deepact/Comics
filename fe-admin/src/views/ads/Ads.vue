<template>
    <div class="p-8">
        <h1 class="text-2xl font-semibold mb-6">Ads Management</h1>
        
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
    </div>
</template>

<script setup>
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { ref } from 'vue'

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

const handleEdit = (ad) => {
    console.log('Edit:', ad.title)
}

const handleDelete = (ad) => {
    console.log('Delete:', ad.title)
}

const handleStatusChange = (ad) => {
    console.log('Status changed:', ad.title, ad.status)
}
</script> 