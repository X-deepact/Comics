<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogTitle,
    DialogHeader,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { ref } from "vue";
import loadingImg from "@/assets/loading.svg";
import { useToast } from "@/components/ui/toast/use-toast";
import { useDramaStore } from "../../../../../stores/dramaStore";
import { useEpisodeStore } from "../../../../../stores/episodeStore";
import M3uPlayer from "../../../../../lib/Player.vue"
import { langConverter } from "../../../../../lib/utils";
const { toast } = useToast();
const drameStore = useDramaStore();
const episodeStore = useEpisodeStore()
const previewUrl = ref<string | null>(null);
const isLoading = ref(false);
const episode = ref({
    subtitle_en: null as string | null,
    subtitle_zh: null as string | null,
    subtitle_vi: null as string | null,
});
const resetEpisode = () => {
    episode.value = {
        subtitle_en: null as string | null,
        subtitle_zh: null as string | null,
        subtitle_vi: null as string | null,
    };
    if (previewUrl.value) {
        URL.revokeObjectURL(previewUrl.value);
        previewUrl.value = null;
    }
};
const video = ref<File | null>(null);
const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
        // Clear previous preview URL if it exists
        if (previewUrl.value) {
            URL.revokeObjectURL(previewUrl.value);
        }
        video.value = target.files[0];
        // Create new preview URL
        previewUrl.value = URL.createObjectURL(target.files[0]);
    }
};
const checkForm = () => {
    if (!video.value) {
        toast({
            description: "Fill in all data.",
            variant: "destructive",
        });
        return false;
    }
    return true;
}
const en = ref<File | null>(null);
const zh = ref<File | null>(null);
const vi = ref<File | null>(null);
const handleSubtitleChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
        const file = target.files[0];
        const fieldName = target.name;
        if (!file || file.type !== "text/vtt") return;
        switch (fieldName) {
            case "subtitle_en":
                en.value = file;
                break;
            case "subtitle_zh":
                zh.value = file;
                break;
            case "subtitle_vi":
                vi.value = file;
                break;
        }
    }
}
const formattedSubtitles = (datas: any) => {
    if (Array.isArray(datas)) {
        return datas.map((track: any) => ({
            language: langConverter(track.language),
            kind: 'subtitles',
            src: track.url
        }));
    } else {
        return [];
    }
}
</script>
<template>
    <Dialog :open="episodeStore.updateDialogIsOpen"
        @update:open="(value: boolean) => { episodeStore.updateDialogIsOpen = value; }">
        <DialogContent class="overflow-y-auto max-h-full">
            <DialogHeader>
                <DialogTitle>Update Episode</DialogTitle>
            </DialogHeader>
            <div class="w-full h-auto" v-if="!previewUrl">
                <M3uPlayer :videoSource="episodeStore.selectedData.video"
                    :subtitles="formattedSubtitles(episodeStore.selectedData.subtitles)" />
            </div>
            <div class="flex justify-center w-full h-auto items-center" v-if="previewUrl">
                <video :src="previewUrl" controls class="justify-self-center h-auto" />
            </div>
            <Input type="file" placeholder="Video" class="justify-self-center w-[50%]" @change="handleFileChange"
                accept="video/mp4" />

            <div class="flex items-center gap-4">
                <Label class="text-center w-1/4">Number</Label>
                <Input v-model="episodeStore.selectedData.number" placeholder="Number" type="number" />
            </div>
            <div class="flex items-center gap-4 whitespace-nowrap">
                <Label class="text-center w-1/4">English Subtitle</Label>
                <Input type="file" placeholder="Subtitle" class="w-full" name="subtitle_en"
                    @change="handleSubtitleChange" accept=".vtt" />
            </div>

            <div class="flex items-center gap-4 whitespace-nowrap">
                <Label class="text-center w-1/4">Chinese Subtitle</Label>
                <Input type="file" placeholder="Subtitle" class="w-full" name="subtitle_zh"
                    @change="handleSubtitleChange" accept=".vtt" />
            </div>

            <div class="flex items-center gap-4 whitespace-nowrap">
                <Label class="text-center w-1/4">Vietnamese Subtitle</Label>
                <Input type="file" placeholder="Subtitle" class="w-full" name="subtitle_vi"
                    @change="handleSubtitleChange" accept=".vtt" />
            </div>
            <DialogFooter class="sm:justify-end">
                <Button variant="secondary" @click="episodeStore.updateDialogIsOpen = false">
                    Close
                </Button>
                <Button :disabled="isLoading" @click="
                    async () => {
                        if (checkForm()) {
                            isLoading = true;
                            let videoURL;
                            if (video) {
                                videoURL = await episodeStore.uploadVideo(video, drameStore.selectedData.id);
                            } else { videoURL = null }
                            if (en != null)
                                episode.subtitle_en = await episodeStore.uploadSubtitle(en);
                            if (zh != null)
                                episode.subtitle_zh = await episodeStore.uploadSubtitle(zh);
                            if (vi != null)
                                episode.subtitle_vi = await episodeStore.uploadSubtitle(vi);
                            const subtitles = [];
                            if (episode.subtitle_en != null) subtitles.push({ language: 'en', url: episode.subtitle_en });
                            if (episode.subtitle_zh != null) subtitles.push({ language: 'zh', url: episode.subtitle_zh });
                            if (episode.subtitle_vi != null) subtitles.push({ language: 'vi', url: episode.subtitle_vi });

                            const data = {
                                id: episodeStore.selectedData.id,
                                drama_id: drameStore.selectedData.id,
                                number: episodeStore.selectedData.number,
                                video: videoURL,
                                subtitles: subtitles,
                            };
                            await episodeStore.updateEpisode(data);
                            await episodeStore.getEpisodeData();
                            isLoading = false;
                            episodeStore.updateDialogIsOpen = false;
                            resetEpisode();
                        }
                    }">
                    <img v-if="isLoading" :src="loadingImg" size="icon" />Update</Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>
