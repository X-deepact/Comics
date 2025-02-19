<template>
    <table class="min-w-full border-collapse border border-gray-300">
        <thead>
            <tr>
                <th v-for="column in columns" :key="column.key" class="border border-gray-300 p-2">{{ column.label }}</th>
                <th class="border border-gray-300 p-2">Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="subject in subjects" :key="subject.name">
                <td v-for="column in columns" :key="column.key" class="border border-gray-300 p-2">{{ subject[column.key] }}</td>
                <td class="border border-gray-300 p-2">
                    <Button @click="$emit('edit', subject)">Edit</Button>
                    <Button @click="$emit('delete', subject)">Delete</Button>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script setup>
import { defineProps } from 'vue'
import { Button } from '@/components/ui'

const props = defineProps({
    subjects: Array,
    columns: Array
})
</script>