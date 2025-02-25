<script setup lang="ts">
import { useRoute } from "vue-router";
import { useNavigation } from "../composables/useNavigation";
import AppHeader from "../components/layout/AppHeader.vue";
import AppSidebar from "../components/layout/AppSidebar.vue";

const route = useRoute();
const { navigation } = useNavigation();
</script>

<template>
    <div class="min-h-screen">
        <AppHeader />
        
        <div class="flex">
            <AppSidebar :navigation="navigation" />
            
            <main class="flex-1 flex flex-col ml-[260px]">
                <!-- Breadcrumb -->
                <div class="flex items-center gap-2 px-6 py-4 text-sm bg-white border-b sticky top-16 z-10">
                    <router-link 
                        to="/dashboard" 
                        class="text-gray-500 hover:text-gray-700"
                    >
                        Home
                    </router-link>
                    <span class="text-gray-400">/</span>
                    <span class="text-gray-900 capitalize">
                        {{ route.name?.toString().replace(/-/g, ' ') }} 
                    </span>
                </div>

                <!-- Content area -->
                <div class="flex-1 w-full px-6 py-6 mt-16">
                        <router-view></router-view>
                </div>
            </main>
        </div>
    </div>
</template>
