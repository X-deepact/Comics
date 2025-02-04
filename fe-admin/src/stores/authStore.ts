import { defineStore } from "pinia";
import axios from "axios";
const API_URL = import.meta.env.VITE_API_URL;

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null as { id: number; username: string } | null,
    token: localStorage.getItem("token") || "",
  }),
  actions: {
    async login(username: string, password: string) {
      try {
        const { data } = await axios.post(`${API_URL}/login`, {
          username,
          password,
        });
        this.token = data.token;
        localStorage.setItem("token", data.token);
        axios.defaults.headers.common["Authorization"] = `Bearer ${data.token}`;
        return true;
      } catch (error) {
        console.error(error);
        return false;
      }
    },
    async fetchProfile() {
      try {
        const { data } = await axios.get(`${API_URL}/profile`, {
          withCredentials: true,
        });
        this.user = data.user;
      } catch (error) {
        console.error(error);
      }
    },
    logout() {
      this.user = null;
      this.token = "";
      localStorage.removeItem("token");
      delete axios.defaults.headers.common["Authorization"];
    },
  },
});
