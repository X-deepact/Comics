<template>
  <div class="change-password-container">
    <h2 class="text-lg font-semibold mb-4">Change Password</h2>
    
    <!-- Add error/success messages -->
    <div v-if="errorMessage" class="error-message mb-4">
      {{ errorMessage }}
    </div>
    <div v-if="successMessage" class="success-message mb-4">
      {{ successMessage }}
    </div>

    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="current-password">Current Password</label>
        <div class="password-input-wrapper">
          <input 
            :type="showCurrentPassword ? 'text' : 'password'"
            id="current-password" 
            v-model="currentPassword" 
            required 
          />
          <button 
            type="button" 
            class="toggle-password"
            @click="showCurrentPassword = !showCurrentPassword"
          >
            {{ showCurrentPassword ? '🔓' : '🔒' }}
          </button>
        </div>
      </div>

      <div class="form-group">
        <label for="new-password">New Password</label>
        <div class="password-input-wrapper">
          <input 
            :type="showNewPassword ? 'text' : 'password'"
            id="new-password" 
            v-model="newPassword" 
            required 
            @input="validatePassword"
          />
          <button 
            type="button" 
            class="toggle-password"
            @click="showNewPassword = !showNewPassword"
          >
            {{ showNewPassword ? '🔓' : '🔒' }}
          </button>
        </div>
        <!-- Password requirements -->
        <div class="password-requirements">
          <p class="text-sm">Password must contain:</p>
          <ul class="text-sm">
            <li :class="{ 'requirement-met': hasMinLength }">At least 8 characters</li>
            <li :class="{ 'requirement-met': hasUpperCase }">One uppercase letter</li>
            <li :class="{ 'requirement-met': hasLowerCase }">One lowercase letter</li>
            <li :class="{ 'requirement-met': hasNumber }">One number</li>
          </ul>
        </div>
      </div>

      <div class="form-group">
        <label for="repeat-password">Confirm Password</label>
        <div class="password-input-wrapper">
          <input 
            :type="showRepeatPassword ? 'text' : 'password'"
            id="repeat-password" 
            v-model="repeatPassword" 
            required 
          />
          <button 
            type="button" 
            class="toggle-password"
            @click="showRepeatPassword = !showRepeatPassword"
          >
            {{ showRepeatPassword ? '🔓' : '🔒' }}
          </button>
        </div>
        <span v-if="!passwordsMatch && repeatPassword" class="error-text">
          Passwords do not match
        </span>
      </div>

      <div class="form-actions">
        <button 
          type="submit" 
          class="submit-button"
          :disabled="!isFormValid"
        >
          Save Changes
        </button>
        <button type="button" class="reset-button" @click="resetForm">
          Reset
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const currentPassword = ref('');
const newPassword = ref('');
const repeatPassword = ref('');
const errorMessage = ref('');
const successMessage = ref('');
const showCurrentPassword = ref(false);
const showNewPassword = ref(false);
const showRepeatPassword = ref(false);

// Password validation
const hasMinLength = computed(() => newPassword.value.length >= 8);
const hasUpperCase = computed(() => /[A-Z]/.test(newPassword.value));
const hasLowerCase = computed(() => /[a-z]/.test(newPassword.value));
const hasNumber = computed(() => /\d/.test(newPassword.value));
const passwordsMatch = computed(() => newPassword.value === repeatPassword.value);

const isFormValid = computed(() => {
  return currentPassword.value &&
         hasMinLength.value &&
         hasUpperCase.value &&
         hasLowerCase.value &&
         hasNumber.value &&
         passwordsMatch.value;
});

const validatePassword = () => {
  errorMessage.value = '';
  successMessage.value = '';
};

const handleSubmit = async () => {
  try {
    errorMessage.value = '';
    successMessage.value = '';
    
    if (!isFormValid.value) {
      errorMessage.value = 'Please ensure all password requirements are met.';
      return;
    }

    // Add your API call here
    // await changePassword({ ... })
    
    successMessage.value = 'Password successfully updated!';
    resetForm();
  } catch (error) {
    errorMessage.value = 'Failed to update password. Please try again.';
  }
};

const resetForm = () => {
  currentPassword.value = '';
  newPassword.value = '';
  repeatPassword.value = '';
  errorMessage.value = '';
  showCurrentPassword.value = false;
  showNewPassword.value = false;
  showRepeatPassword.value = false;
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

.password-input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.toggle-password {
  position: absolute;
  right: 10px;
  background: none;
  border: none;
  cursor: pointer;
}

.password-requirements {
  margin-top: 8px;
  padding: 8px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.password-requirements ul {
  list-style: none;
  padding-left: 0;
  margin: 5px 0 0;
}

.password-requirements li {
  color: #dc3545;
  margin: 2px 0;
}

.password-requirements li::before {
  content: '❌';
  margin-right: 5px;
}

.requirement-met {
  color: #28a745 !important;
}

.requirement-met::before {
  content: '✅' !important;
}

.error-message {
  background-color: #fff3f3;
  color: #dc3545;
  padding: 10px;
  border-radius: 4px;
  border: 1px solid #dc3545;
}

.success-message {
  background-color: #f0fff4;
  color: #28a745;
  padding: 10px;
  border-radius: 4px;
  border: 1px solid #28a745;
}

.error-text {
  color: #dc3545;
  font-size: 0.875rem;
  margin-top: 4px;
  display: block;
}

.submit-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style> 