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
  padding: 10px 12px;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s ease;
  background-color: #f8fafc;
}

input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-actions {
  display: flex;
  justify-content: space-between;
}

.submit-button {
  background-color: #3b82f6;
  color: white;
  border: none;
  padding: 6px 20px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: background-color 0.2s ease;
}

.submit-button:hover {
  background-color: #2563eb;
}

.reset-button {
  background-color: #f1f5f9;
  color: #64748b;
  border: 1px solid #e2e8f0;
  padding: 6px 20px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}

.reset-button:hover {
  background-color: #e2e8f0;
  color: #475569;
}
</style> 