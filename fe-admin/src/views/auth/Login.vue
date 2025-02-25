<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../../stores/authStore";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useToast } from "@/components/ui/toast/use-toast";
import CardFooter from "../../components/ui/card/CardFooter.vue";
import { Label } from "@/components/ui/label";
const { toast } = useToast();

const username = ref("");
const password = ref("");
const authStore = useAuthStore();
const router = useRouter();

const login = async () => {
  if (await authStore.login(username.value, password.value)) {
    router.push("/dashboard");
  } else {
    toast({
      title: "Error",
      variant: "destructive",
      description: "Invalid credentials.",
    });
  }
};
const handleSubmit = (e: Event) => {
  e.preventDefault();
  login();
};
</script>

<template>
  <div class="flex items-center justify-center min-h-screen">
    <Card class="justify-self-center w-[25rem]">
      <form @submit="handleSubmit">
        <CardHeader>
          <CardTitle>Login</CardTitle>
        </CardHeader>
        <CardContent>
          <div class="grid items-center w-full gap-4">
            <div class="flex flex-col gap-1 space-y-1.5">
              <Label for="username">Username:</Label>
              <Input v-model="username" type="text" />
            </div>
            <div class="flex flex-col gap-1 space-y-1.5">
              <Label for="password">Password:</Label>
              <Input v-model="password" type="password" />
            </div>
          </div>
        </CardContent>
        <CardFooter class="flex justify-end">
          <Button type="submit" variant="outline" class="bg-blue-500 text-white"
            >Login</Button
          >
        </CardFooter>
      </form>
    </Card>
  </div>
</template>
