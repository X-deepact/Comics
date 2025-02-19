<template>
    <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-semibold">{{ title }}</h1>
        <Button @click="onAddClick" class="bg-blue-500 hover:bg-blue-600 text-white">Add {{ title }}</Button>
    </div>
    <div class="flex justify-between items-center mb-4 gap-4">
        <div class="flex items-center gap-2">
            <Input 
                v-model="localSearchQuery"
                placeholder="Search..."
                class="w-[280px]"
                @input="handleSearchInput"
            />
        </div>
        <div class="flex items-center gap-4">
            <Select 
                :value="pageSize" 
                @update:value="handlePageSizeChange"
            >
                <SelectTrigger class="w-[120px]">
                    <SelectValue :placeholder="`${pageSize} / page`" />
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
import { ref, defineProps, defineEmits, watch } from 'vue'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'

const props = defineProps({
    title: String,
    totalItems: Number,
    currentPage: Number,
    pageSize: Number,
    displayedPages: Array,
    onAddClick: Function
})

const emit = defineEmits(['update:pageSize', 'search'])

const localSearchQuery = ref('')

// Debounce search input
let searchTimeout
const handleSearchInput = () => {
    clearTimeout(searchTimeout)
    searchTimeout = setTimeout(() => {
        emit('search', localSearchQuery.value)
    }, 300)
}

const handlePageSizeChange = (newSize) => {
    emit('update:pageSize', newSize)
}
</script>