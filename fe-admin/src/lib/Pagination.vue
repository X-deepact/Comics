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
                @click="handlePageChange(currentPage - 1)"
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
                    @click="handlePageChange(pageNumber)"
                >
                    {{ pageNumber }}
                </Button>
            </div>
            <Button 
                variant="outline" 
                size="sm"
                :disabled="currentPage === totalPages"
                @click="handlePageChange(currentPage + 1)"
            >
                Next
            </Button>
        </div>
    </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
import { Button } from '@/components/ui'

const props = defineProps({
    totalItems: Number,
    currentPage: Number,
    totalPages: Number,
    displayedPages: Array
})

const emit = defineEmits(['update:currentPage'])

const handlePageChange = (newPage) => {
    if (newPage >= 1 && newPage <= props.totalPages) {
        emit('update:currentPage', newPage)
    }
}
</script>