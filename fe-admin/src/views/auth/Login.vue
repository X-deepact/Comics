<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../../stores/authStore";
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import Toast from "primevue/toast";
import Card from "primevue/card";
import { useToast } from "primevue/usetoast";

const username = ref("");
const password = ref("");
const authStore = useAuthStore();
const router = useRouter();
const toast = useToast();

const login = async () => {
  if (await authStore.login(username.value, password.value)) {
    router.push("/dashboard");
  } else {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Invalid credentials",
      life: 3000,
    });
  }
};
</script>

<template>
  <div class="flex items-center justify-center min-h-screen">
    <Card
      className="justify-self-center bg-gray-100 rounded-lg shadow-lg w-[20rem]"
    >
      <template #title>Login</template>
      <template #content>
        <div className="flex flex-col gap-5">
          <Toast />
          <InputText v-model="username" type="text" placeholder="Username" />
          <InputText
            v-model="password"
            type="password"
            placeholder="Password"
          />
          <Button
            type="submit"
            severity="secondary"
            label="Login"
            className="bg-blue-500"
            @click="login"
          ></Button>
        </div>
      </template>
    </Card>
  </div>
</template>
