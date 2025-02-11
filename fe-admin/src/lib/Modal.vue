<template>
    <Dialog :open="isVisible" @update:open="closeModal" class="fixed inset-0 z-50">
        <DialogContent class="fixed top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] bg-white p-6 rounded-lg shadow-lg sm:max-w-[800px] w-full">
            <button 
                class="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                @click="closeModal"
            >
                <X class="h-4 w-4" />
                <span class="sr-only">Close</span>
            </button>
            <DialogHeader>
                <DialogTitle>{{ title }}</DialogTitle>
                <DialogDescription>{{ description }}</DialogDescription>
            </DialogHeader>
            <div class="modal-content">
                <slot :newSubject="newSubject"></slot>
            </div>
            <DialogFooter>
                <Button variant="outline" @click="closeModal">Cancel</Button>
                <Button @click="confirmAction">Save</Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>

<script setup>
import { defineProps, defineEmits, ref } from 'vue'
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { X } from 'lucide-vue-next'

const props = defineProps({
    isVisible: Boolean,
    title: String,
    description: String,
    newSubject: Object
})

const emit = defineEmits(['update:isVisible', 'confirm'])

const closeModal = () => {
    emit('update:isVisible', false)
}

const confirmAction = () => {
    emit('confirm')
    closeModal()
}
</script> 