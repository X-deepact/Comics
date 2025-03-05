<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50">
        <div class="text-center">
            <h1 class="text-6xl font-bold text-gray-900">404</h1>
            <p class="text-xl text-gray-600 mt-4">Page not found</p>
            <router-link 
                to="/dashboard" 
                class="mt-6 inline-block px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
                Return to Dashboard
            </router-link>
        </div>
    </div>
</template>