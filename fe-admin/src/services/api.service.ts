import axios from 'axios'

const BASE_URL = 'http://localhost:8080/api'

export class ApiService {
  protected axios: any
  protected endpoint: string

  constructor(resource: string) {
    this.endpoint = `${BASE_URL}/${resource}`
    this.axios = axios.create({
      baseURL: BASE_URL,
      headers: {
        'Content-Type': 'application/json',
      }
    })

    // Add request interceptor to include auth token
    this.axios.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('token')
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      }
    )

    // Add response interceptor for error handling
    this.axios.interceptors.response.use(
      response => response,
      error => {
        if (error.response?.status === 401) {
          // Handle unauthorized access - redirect to login
          localStorage.removeItem('token')
          window.location.href = '/login'
        }
        console.error('API Error:', error.response?.data || error.message)
        return Promise.reject(error)
      }
    )
  }

  async create(data: any) {
    return this.axios.post(`${this.endpoint}/create`, data)
  }

  async update(id: string, data: any) {
    return this.axios.put(`${this.endpoint}/update/${id}`, data)
  }

  async delete(id: string) {
    return this.axios.delete(`${this.endpoint}/delete/${id}`)
  }

  async getAll(params?: any) {
    return this.axios.get(`${this.endpoint}/list`, { params })
  }

  async getById(id: string) {
    return this.axios.get(`${this.endpoint}/${id}`)
  }
} 