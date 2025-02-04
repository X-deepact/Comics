import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL + "/auth";

export interface LoginResponse {
  token: string;
}

export const login = async (
  email: string,
  password: string
): Promise<LoginResponse> => {
  const response = await axios.post<LoginResponse>(`${API_URL}/login`, {
    email,
    password,
  });
  localStorage.setItem("token", response.data.token);
  return response.data;
};

export const logout = () => {
  localStorage.removeItem("token");
};

export const getToken = (): string | null => {
  return localStorage.getItem("token");
};

export const isAuthenticated = (): boolean => {
  return !!getToken();
};
