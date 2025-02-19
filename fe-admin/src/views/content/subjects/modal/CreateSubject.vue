<template>
  <Dialog
    :open="subjectStore.createDialogIsOpen"
    @update:open="(value) => { subjectStore.createDialogIsOpen = value; }"
  >
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create Subject</DialogTitle>
      </DialogHeader>
      <div class="grid gap-4 py-4">
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Name</Label>
          <Input v-model="subject.name" placeholder="name" />
        </div>
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Create Time</Label>
          <Input type="datetime-local" v-model="subject.createTime" />
        </div>
        <div class="flex items-center gap-4">
          <Label class="text-right w-1/4">Operator</Label>
          <Input v-model="subject.operator" placeholder="operator" />
        </div>
      </div>
      <DialogFooter class="sm:justify-end">
        <Button
          variant="secondary"
          @click="subjectStore.createDialogIsOpen = false"
        >
          Close
        </Button>
        <Button
          @click="async () => {
            await subjectStore.createSubject(subject);
            subjectStore.createDialogIsOpen = false;
            subjectStore.getSubjectData();
            resetSubject();
          }"
        >Add</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { useSubjectStore } from "../../../../stores/subjectStore";

const subjectStore = useSubjectStore();

const subject = ref({
  name: "",
  createTime: "",
  operator: "",
});

const resetSubject = () => {
  subject.value = {
    name: "",
    createTime: "",
    operator: "",
  };
};
</script> 