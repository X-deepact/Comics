<template>
  <div class="change-password-container">
    <h2 class="text-lg font-semibold">Change Password</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="current-password">* Current Password</label>
        <input type="password" id="current-password" v-model="currentPassword" required />
      </div>
      <div class="form-group">
        <label for="new-password">* New Password</label>
        <input type="password" id="new-password" v-model="newPassword" required />
      </div>
      <div class="form-group">
        <label for="repeat-password">* Repeat Password</label>
        <input type="password" id="repeat-password" v-model="repeatPassword" required />
      </div>
      <div class="form-actions">
        <button type="submit" class="submit-button">Save</button>
        <button type="button" class="reset-button" @click="resetForm">Reset</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const currentPassword = ref('');
const newPassword = ref('');
const repeatPassword = ref('');

const handleSubmit = () => {
  // Handle form submission logic here
  console.log('Current Password:', currentPassword.value);
  console.log('New Password:', newPassword.value);
  console.log('Repeat Password:', repeatPassword.value);
};

const resetForm = () => {
  currentPassword.value = '';
  newPassword.value = '';
  repeatPassword.value = '';
};
</script>

<style scoped>
.change-password-container {
  max-width: 400px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #fff;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.form-actions {
  display: flex;
  justify-content: space-between;
}

.submit-button {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 15px;
  border-radius: 4px;
  cursor: pointer;
}

.reset-button {
  background-color: #f8f9fa;
  color: #000;
  border: 1px solid #ccc;
  padding: 10px 15px;
  border-radius: 4px;
  cursor: pointer;
}
</style> 