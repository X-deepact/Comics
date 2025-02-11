<template>
    <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-semibold">{{ title }}</h1>
        <Button @click="onAddClick" class="bg-blue-500 hover:bg-blue-600 text-white">Add {{ title }}</Button>
    </div>
    <div class="flex justify-between items-center mb-4 gap-4">
        <div class="flex items-center gap-2">
            <Input 
                v-model="searchQuery"
                placeholder="Search..."
                class="w-[280px]"
            />
        </div>
        <div class="flex items-center gap-4">
            <Select :value="pageSize" @change="updatePageSize">
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
 
</template>

<script setup>
import { ref, defineProps, defineEmits } from 'vue'
import { Button, Input, Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui'

const props = defineProps({
    title: String,
    totalItems: Number,
    currentPage: Number,
    pageSize: Number,
    displayedPages: Array,
    onAddClick: Function
})

const emit = defineEmits(['update:pageSize'])

const searchQuery = ref('')

// Emit the new page size when it changes
const updatePageSize = (newSize) => {
    emit('update:pageSize', newSize)
}
</script>