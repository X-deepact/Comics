<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import loadingImg from "@/assets/loading.svg";
import { ref } from "vue";
import { useChapterStore } from "../../../../../stores/chapterStore";
import { useComicStore } from "../../../../../stores/comicStore";
const comicStore = useComicStore();
const isLoading = ref(false);
const chapterStore = useChapterStore();
</script>
<template>
    <Dialog :open="chapterStore.updateDialogIsOpen" @update:open="
        (value: boolean) => {
            chapterStore.updateDialogIsOpen = value;
        }
    ">
        <DialogContent>
            <DialogHeader>
                <DialogTitle>Update Chapter</DialogTitle>
            </DialogHeader>
            <div class="flex items-center gap-4">
                <Label class="text-center w-1/4">Name</Label>
                <Input v-model="chapterStore.selectedData.name" placeholder="Name" />
            </div>
            <div class="flex items-center gap-4">
                <Label class="text-center w-1/4">Number</Label>
                <Input v-model="chapterStore.selectedData.number" placeholder="Number" type="number" />
            </div>
            <div class="flex items-center gap-4">
                <Label class="text-center w-1/4">Active</Label>
                <Select v-model="chapterStore.selectedData.active">
                    <SelectTrigger>
                        <SelectValue :value="chapterStore.selectedData.active.toString()" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            <SelectItem :value="true">On</SelectItem>
                            <SelectItem :value="false">Off</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                </Select>
            </div>
            <div class="flex items-center gap-4">
                <Label class="text-center w-1/4">Cover</Label>
                <Select v-model="chapterStore.selectedData.cover">
                    <SelectTrigger>
                        <SelectValue :value="chapterStore.selectedData.cover.toString()" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            <SelectItem :value="true">On</SelectItem>
                            <SelectItem :value="false">Off</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                </Select>
            </div>
            <DialogFooter class="sm:justify-end">
                <Button variant="secondary" @click="chapterStore.updateDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        isLoading = true;
                        await chapterStore.updateChapter({
                            id: chapterStore.selectedData.id,
                            name: chapterStore.selectedData.name,
                            number: chapterStore.selectedData.number,
                            active: chapterStore.selectedData.active,
                            cover: chapterStore.selectedData.cover,
                            comic_id: comicStore.selectedData.id,
                        });
                        isLoading = false;
                        chapterStore.updateDialogIsOpen = false;
                        chapterStore.getChapterData();
                    }
                ">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />Update
                </Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
