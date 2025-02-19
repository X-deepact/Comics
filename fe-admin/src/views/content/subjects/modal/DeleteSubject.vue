<template>
  <Dialog
    :open="subjectStore.deleteDialogIsOpen"
    @update:open="(value) => { subjectStore.deleteDialogIsOpen = value; }"
  >
    <DialogContent>
      <DialogHeader>Do you really want to delete this subject?</DialogHeader>
      <DialogFooter class="justify-end gap-4">
        <Button variant="secondary" @click="subjectStore.deleteDialogIsOpen = false">
          Close
        </Button>
        <Button
          @click="async () => {
            await subjectStore.deleteSubject(subjectStore.selectedData.id);
            subjectStore.deleteDialogIsOpen = false;
            subjectStore.getSubjectData();
          }"
        >Delete</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader } from "@/components/ui/dialog";
import { useSubjectStore } from "../../../../stores/subjectStore";

const subjectStore = useSubjectStore();
</script> 