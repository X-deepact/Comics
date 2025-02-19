<template>
  <Dialog
    :open="subjectStore.updateDialogIsOpen"
    @update:open="(value) => { subjectStore.updateDialogIsOpen = value; }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Update Subject</DialogTitle>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Name</Label>
          <Input 
            v-model="subjectStore.selectedData.name" 
            placeholder="name" 
          />
        </div>
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Create Time</Label>
          <Input 
            type="datetime-local" 
            v-model="subjectStore.selectedData.createTime"
          />
        </div>
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Operator</Label>
          <Input 
            v-model="subjectStore.selectedData.operator" 
            placeholder="operator" 
          />
        </div>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="subjectStore.updateDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="async () => {
            await subjectStore.updateSubject({
              id: subjectStore.selectedData.id,
              name: subjectStore.selectedData.name,
              createTime: subjectStore.selectedData.createTime,
              operator: subjectStore.selectedData.operator,
            });
            subjectStore.updateDialogIsOpen = false;
            subjectStore.getSubjectData();
          }"
        >
          Update
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { useSubjectStore } from "../../../../stores/subjectStore";

const subjectStore = useSubjectStore();
</script> 