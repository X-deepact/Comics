<script setup lang="ts">
import { ref } from 'vue';

const stats = ref([
    { 
        name: 'Total Accounts', 
        value: '2,543', 
        change: '+12.5%',
        icon: 'users'
    },
    { 
        name: 'Active Sessions', 
        value: '1,325', 
        change: '+5.2%',
        icon: 'activity'
    },
    { 
        name: 'Total Streams', 
        value: '845', 
        change: '+8.1%',
        icon: 'play'
    },
    { 
        name: 'Storage Used', 
        value: '2.4 TB', 
        change: '+2.3%',
        icon: 'database'
    },
]);
</script>

<template>
    <div class="p-8">
        <h1 class="text-2xl font-semibold mb-6">Dashboard Overview</h1>
        
        <!-- Stats Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            <div 
                v-for="stat in stats" 
                :key="stat.name"
                class="bg-white p-6 rounded-lg border"
            >   
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm text-gray-600">{{ stat.name }}</p>
                        <p class="text-2xl font-semibold mt-1">{{ stat.value }}</p>
                    </div>
                    <div class="h-12 w-12 bg-blue-50 rounded-full flex items-center justify-center">
                        <!-- Users Icon -->
                        <svg v-if="stat.icon === 'users'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-600">
                            <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                            <circle cx="9" cy="7" r="4"></circle>
                            <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                            <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                        </svg>
                        <!-- Activity Icon -->
                        <svg v-if="stat.icon === 'activity'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-600">
                            <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"></polyline>
                        </svg>
                        <!-- Play Icon -->
                        <svg v-if="stat.icon === 'play'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-600">
                            <polygon points="5 3 19 12 5 21 5 3"></polygon>
                        </svg>
                        <!-- Database Icon -->
                        <svg v-if="stat.icon === 'database'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-600">
                            <ellipse cx="12" cy="5" rx="9" ry="3"></ellipse>
                            <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"></path>
                            <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"></path>
                        </svg>
                    </div>
                </div>
                <div class="mt-4 flex items-center">
                    <span class="text-green-500 text-sm">{{ stat.change }}</span>
                    <span class="text-gray-500 text-sm ml-2">vs last month</span>
                </div>
            </div>
        </div>

        <!-- Quick Actions -->
        <div class="bg-white p-6 rounded-lg border">
            <h2 class="text-lg font-semibold mb-4">Quick Actions</h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                <button class="p-4 border rounded-lg hover:bg-gray-50 flex items-center gap-3">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                        <circle cx="8.5" cy="7" r="4"></circle>
                        <line x1="20" y1="8" x2="20" y2="14"></line>
                        <line x1="23" y1="11" x2="17" y2="11"></line>
                    </svg>
                    Add New Account
                </button>
                <button class="p-4 border rounded-lg hover:bg-gray-50 flex items-center gap-3">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <rect x="2" y="3" width="20" height="14" rx="2" ry="2"></rect>
                        <line x1="8" y1="21" x2="16" y2="21"></line>
                        <line x1="12" y1="17" x2="12" y2="21"></line>
                    </svg>
                    View Active Streams
                </button>
                <button class="p-4 border rounded-lg hover:bg-gray-50 flex items-center gap-3">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                        <polyline points="17 8 12 3 7 8"></polyline>
                        <line x1="12" y1="3" x2="12" y2="15"></line>
                    </svg>
                    Download Reports
                </button>
            </div>
        </div>
    </div>
</template>

<style src="@/styles/icons.css" />
