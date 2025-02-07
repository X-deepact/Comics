<template>
    <div class="p-8">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-2xl font-semibold">User Management</h1>
            <Button @click="handleNewUser" class="bg-blue-500 hover:bg-blue-600 text-white">New User</Button>
        </div>

        <div class="flex justify-between items-center mb-4 gap-4">
            <div class="flex items-center gap-2">
                <Input 
                    v-model="searchQuery"
                    placeholder="Search users..."
                    class="w-[280px]"
                />
            </div>
            <div class="flex items-center gap-4">
                <Select v-model="pageSize" @change="handlePageSizeChange">
                    <SelectTrigger class="w-[120px]">
                        <SelectValue :placeholder="pageSize ? '' : 'Select size'" />
                    </SelectTrigger>
                    <SelectContent class="bg-white">
                        <SelectItem value="10">10 / page</SelectItem>
                        <SelectItem value="20">20 / page</SelectItem>
                        <SelectItem value="50">50 / page</SelectItem>
                    </SelectContent>
                </Select>
            </div>
        </div>

        <Table class="border rounded-lg">
            <TableHeader>
                <TableRow>
                    <TableHead>no.</TableHead>
                    <TableHead>username</TableHead>
                    <TableHead>create time</TableHead>
                    <TableHead>update time</TableHead>
                    <TableHead>Operator</TableHead>
                    <TableHead>status</TableHead>
                    <TableHead>operate</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow v-for="(user, index) in paginatedUsers" :key="index">
                    <TableCell>{{ index + 1 }}</TableCell>
                    <TableCell>{{ user.username }}</TableCell>
                    <TableCell>{{ user.createTime }}</TableCell>
                    <TableCell>{{ user.updateTime }}</TableCell>
                    <TableCell>{{ user.operator }}</TableCell>
                    <TableCell>
                        <Switch 
                            v-model="user.status" 
                            @update:modelValue="handleStatusChange(user)"
                            class="data-[state=checked]:bg-primary"
                        >
                            <span class="sr-only">Toggle status</span>
                        </Switch>
                    </TableCell>
                    <TableCell>
                        <div class="flex space-x-2">
                            <Button variant="ghost" size="sm" @click="handleEdit(user)">
                                edit
                            </Button>
                            <Button variant="ghost" size="sm" class="text-red-500 hover:text-red-700" @click="handleDelete(user)">
                                delete
                            </Button>
                        </div>
                    </TableCell>
                </TableRow>
            </TableBody>
        </Table>

        <div class="flex items-center justify-between mt-4">
            <div class="text-sm text-gray-500">
                Total {{ totalUsers }} users
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
                        {{ pageNumber === '...' ? pageNumber : pageNumber }}
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

        <Dialog v-model:open="showAddModal" class="fixed inset-0 z-50">
            <DialogContent class="fixed top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] bg-white p-6 rounded-lg shadow-lg sm:max-w-[800px] w-full">
                <button 
                    class="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                    @click="showAddModal = false"
                >
                    <X class="h-4 w-4" />
                    <span class="sr-only">Close</span>
                </button>
                <DialogHeader>
                    <DialogTitle>Add New User</DialogTitle>
                    <DialogDescription>
                        Fill in the details for the new user. Click save when you're done.
                    </DialogDescription>
                </DialogHeader>
                <div class="grid gap-4 py-4">
                    <!-- Add your form fields for new user here -->
                </div>
                <DialogFooter>
                    <Button variant="outline" @click="showAddModal = false">Cancel</Button>
                    <Button @click="saveNewUser">Save</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    </div>
</template>

<script setup>
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Switch } from '@/components/ui/switch'
import { ref, computed } from 'vue'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { X } from 'lucide-vue-next'

const searchQuery = ref('')
const users = ref([
    {
        username: 'username1',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'admin',
        status: true
    },
    {
        username: 'username2',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'username1',
        status: true
    },
    {
        username: 'username3',
        createTime: '2014-12-24 23:12:00',
        updateTime: '2014-12-25 23:12:00',
        operator: 'username2',
        status: false
    }
    // Add more user data as needed
])

const totalUsers = computed(() => users.value.length)

const pageSize = ref(10) // Default page size
const currentPage = ref(1) // Current page

const paginatedUsers = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value
    return users.value.slice(start, start + pageSize.value)
})

const totalPages = computed(() => Math.ceil(users.value.length / pageSize.value))

const handleEdit = (user) => {
    console.log('Edit:', user.username)
}

const handleDelete = (user) => {
    console.log('Delete:', user.username)
}

const handleStatusChange = (user) => {
    console.log('Status changed:', user.username, user.status)
}

const handlePageChange = (page) => {
    currentPage.value = page
}

const showAddModal = ref(false)

const handleNewUser = () => {
    showAddModal.value = true
}

const saveNewUser = () => {
    // Add your logic to save the new user here
    console.log('Saving new user')
    // Reset form and close modal
    showAddModal.value = false
}

// New method to handle page size changes
const handlePageSizeChange = (size) => {
    pageSize.value = parseInt(size) // Ensure size is an integer
    currentPage.value = 1 // Reset to first page when page size changes
}
</script> 