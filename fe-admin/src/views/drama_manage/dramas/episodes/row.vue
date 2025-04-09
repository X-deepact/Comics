<template>
    <TableRow>
        <TableCell>
            <button @click="
                () => {
                    $emit('clickAction', data);
                }
            ">
                <Play />
            </button>
        </TableCell>
        <TableCell>
            {{ data.number }}
        </TableCell>
        <TableCell>
            <Badge variant="secondary" class="text-sm">{{
                data.created_by_name
            }}</Badge>
            <div>
                <Label class="text-xs">{{
                    formatDateSafe(data.created_at, true)
                }}</Label>
                <p className="text-xs text-muted-foreground">
                    {{ getTimeAgoSafe(data.created_at) }}
                </p>
            </div>
        </TableCell>
        <TableCell>
            <Badge variant="secondary" class="text-sm">{{
                data.updated_by_name
            }}</Badge>
            <div>
                <Label class="text-xs">{{
                    formatDateSafe(data.updated_at, true)
                }}</Label>
                <p className="text-xs text-muted-foreground">
                    {{ getTimeAgoSafe(data.updated_at) }}
                </p>
            </div>
        </TableCell>
        <TableCell>
            <Switch v-model="data.active" :checked="data.active" @click="
                () => {
                    $emit('clickActive', data);
                }
            " />
        </TableCell>
        <TableCell>
            <div class="flex gap-3">
                <Button variant="outline" size="icon" @click="
                    () => {
                        $emit('clickUpdate', data);
                    }
                ">
                    <Pencil />
                </Button>
                <Button variant="destructive" size="icon" @click="
                    () => {
                        $emit('clickDelete', data);
                    }
                ">
                    <Trash2 />
                </Button>
            </div>
        </TableCell>
    </TableRow>
</template>

<script setup lang="ts">
import { formatDate, getTimeAgo } from "@/lib/utils";
import { TableCell, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Trash2, Pencil, Play } from "lucide-vue-next";
import { Badge } from "@/components/ui/badge";
import { Switch } from "@/components/ui/switch";
import type { Episode } from "../../../../stores/episodeStore";
defineProps<{
    data: Episode;
}>();
defineEmits(["clickDelete", "clickUpdate", "clickActive", "clickAction"]);
const formatDateSafe = (
    date: Date | string | number,
    includeTime: boolean
): string => {
    try {
        return formatDate(date, includeTime);
    } catch (error) {
        console.error("Date formatting error:", error);
        return "Invalid date";
    }
};

const getTimeAgoSafe = (date: Date | string | number): string => {
    try {
        return getTimeAgo(date);
    } catch (error) {
        console.error("Time ago formatting error:", error);
        return "Invalid date";
    }
};
</script>