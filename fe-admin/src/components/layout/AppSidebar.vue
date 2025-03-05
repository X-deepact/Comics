<script setup lang="ts">
import { type NavigationItem } from '../../types/navigation';
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const props = defineProps<{
    navigation: NavigationItem[]
}>();

const route = useRoute();

// Track collapsed state for each section
const collapsedSections = ref<{ [key: string]: boolean }>({});

// Toggle collapse for a section
const toggleSection = (title: string) => {
    collapsedSections.value[title] = !collapsedSections.value[title];
};

// Initialize all sections as expanded (not collapsed)
props.navigation.forEach(section => {
    collapsedSections.value[section.title] = false;
});
</script>

<template>
    <aside class="w-[260px] fixed top-16 bottom-0 left-0 bg-white border-r overflow-y-auto">
        <nav class="py-4">
            <div class="space-y-1">
                <div v-for="(section, index) in navigation" :key="section.title">
                    <!-- Section Header -->
                    <div 
                        class="flex items-center justify-between px-4 py-2 text-sm font-medium text-gray-900 cursor-pointer hover:bg-gray-50"
                        @click="toggleSection(section.title)"
                    >
                        <div class="flex items-center gap-2">
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                width="16" 
                                height="16" 
                                viewBox="0 0 24 24" 
                                fill="none" 
                                stroke="currentColor" 
                                stroke-width="2" 
                                stroke-linecap="round" 
                                stroke-linejoin="round"
                            >
                                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                            </svg>
                            {{ section.title }}
                        </div>
                        <svg 
                            xmlns="http://www.w3.org/2000/svg" 
                            width="16" 
                            height="16" 
                            viewBox="0 0 24 24" 
                            fill="none" 
                            stroke="currentColor" 
                            stroke-width="2" 
                            stroke-linecap="round" 
                            stroke-linejoin="round" 
                            :class="{'transform rotate-180': !collapsedSections[section.title]}"
                            class="transition-transform duration-200"
                        >
                            <polyline points="6 9 12 15 18 9"></polyline>
                        </svg>
                    </div>

                    <!-- Section Items -->
                    <div 
                        v-show="!collapsedSections[section.title]"
                        class="mt-1 space-y-1 relative"
                    >
                        <div class="absolute left-5 top-0 bottom-0 w-px bg-gray-200"></div>
                        <router-link
                            v-for="item in section.children"
                            :key="item.name"
                            :to="item.href"
                            class="block px-2 py-2 text-sm text-gray-600 hover:bg-blue-100 transition-colors active:bg-blue-100 w-[200px] ml-10 rounded"
                            :class="{
                                'bg-blue-600 text-white font-medium': route.path === item.href
                            }"
                        >
                            {{ item.name }}
                        </router-link>
                    </div>
                </div>
            </div>
        </nav>
    </aside>
</template>

<style scoped>
.transform {
    transition: transform 0.2s ease;
}
</style>