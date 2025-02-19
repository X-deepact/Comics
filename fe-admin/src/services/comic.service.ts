import { ApiService } from './api.service'
import type { ComicRequest, ComicResponse } from '../types/comic'

class ComicService extends ApiService {
  constructor() {
    super('comic')
  }

  async create(data: ComicRequest): Promise<ComicResponse> {
    try {
      console.log('Sending request to:', `${this.endpoint}/create`)
      console.log('Request data:', data)
      
      const response = await this.axios.post<ComicResponse>('/comic/create', data)
      
      console.log('Response:', response.data)
      return response.data
    } catch (error: any) {
      console.error('API Error Details:', {
        status: error.response?.status,
        data: error.response?.data,
        message: error.message
      })

      if (error.response?.status === 400) {
        throw new Error(error.response.data.message || 'Invalid request')
      } else if (error.response?.status === 401) {
        throw new Error('User not authenticated')
      } else if (error.response?.status === 500) {
        throw new Error('Server error occurred')
      }
      throw new Error(`Failed to create comic: ${error.message}`)
    }
  }
}

export const comicService = new ComicService()