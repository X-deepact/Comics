<template>
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
                    {{ pageNumber }}
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
</template>

<script setup>
import { defineProps } from 'vue'
import { Button } from '@/components/ui'

const props = defineProps({
    totalItems: Number,
    currentPage: Number,
    totalPages: Number,
    displayedPages: Array
})
</script>