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
                @click="$emit('update:currentPage', currentPage - 1)"
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
                    @click="$emit('update:currentPage', pageNumber)"
                >
                    {{ pageNumber }}
                </Button>
            </div>
            <Button 
                variant="outline" 
                size="sm"
                :disabled="currentPage === totalPages"
                @click="$emit('update:currentPage', currentPage + 1)"
            >
                Next
            </Button>
        </div>
    </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
import { Button } from '@/components/ui'

defineProps({
    totalItems: {
        type: Number,
        required: true
    },
    currentPage: {
        type: Number,
        required: true
    },
    totalPages: {
        type: Number,
        required: true
    },
    displayedPages: {
        type: Array,
        required: true
    }
})

defineEmits(['update:currentPage'])
</script>