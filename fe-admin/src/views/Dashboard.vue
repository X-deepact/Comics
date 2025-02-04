<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useAuthStore } from "../stores/authStore";
import { useRouter } from "vue-router";
import { Button } from "@/components/ui/button";
import { Sheet, SheetContent } from "@/components/ui/sheet";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

const authStore = useAuthStore();
const router = useRouter();
const sidebarVisible = ref(false);

const sidebarItems = [
    {
        label: 'Dashboard',
        route: '/dashboard'
    },
    // Add more sidebar menu items as needed
];

onMounted(async () => {
    if (!authStore.token) {
        router.push("/login");
    } else {
        await authStore.fetchProfile();
    }
});
</script>

<template>
    <div class="layout-wrapper">
        <header class="fixed top-0 left-0 right-0 border-b bg-background z-50">
            <div class="flex h-16 items-center px-4">
                <Button variant="ghost" size="icon" @click="sidebarVisible = true">
                    <i class="i-lucide-menu h-6 w-6" />
                </Button>
                <div class="ml-auto flex items-center space-x-4">
                    <DropdownMenu>
                        <DropdownMenuTrigger as-child>
                            <Button variant="ghost" class="relative h-8 w-8 rounded-full">
                                <i class="i-lucide-user h-6 w-6" />
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent class="w-56" align="end" forceMount>
                            <DropdownMenuLabel class="font-normal">
                                <div class="flex flex-col space-y-1">
                                    <p class="text-sm font-medium leading-none">{{ authStore.user?.username }}</p>
                                </div>
                            </DropdownMenuLabel>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem>
                                <i class="i-lucide-user mr-2 h-4 w-4" />
                                <span>Profile</span>
                            </DropdownMenuItem>
                            <DropdownMenuItem>
                                <i class="i-lucide-settings mr-2 h-4 w-4" />
                                <span>Settings</span>
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem @click="authStore.logout">
                                <i class="i-lucide-log-out mr-2 h-4 w-4" />
                                <span>Log out</span>
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </div>
            </div>
        </header>

        <Sheet v-model:open="sidebarVisible" side="left">
            <SheetContent class="w-[250px] sm:w-[300px]" :side="'left'">
                <div class="px-1">
                    <h3 class="text-lg font-semibold mb-4">Admin Menu</h3>
                    <nav class="space-y-1">
                        <router-link 
                            v-for="item in sidebarItems" 
                            :key="item.route"
                            :to="item.route"
                            class="flex items-center px-3 py-2 text-sm rounded-md hover:bg-accent hover:text-accent-foreground"
                            @click="sidebarVisible = false"
                        >
                            {{ item.label }}
                        </router-link>
                    </nav>
                </div>
            </SheetContent>
        </Sheet>

        <main class="main-content pt-16 px-4">
            <h2 class="text-2xl font-bold mb-4">Dashboard</h2>
            <p v-if="authStore.user">Welcome, {{ authStore.user.username }}</p>
        </main>
    </div>
</template>

<style scoped>
.layout-wrapper {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
}

.main-content {
    flex: 1;
    position: relative;
}
</style>
