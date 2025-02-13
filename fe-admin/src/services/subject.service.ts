import { ApiService } from './api.service'

class SubjectService extends ApiService {
  constructor() {
    super('subject')
  }
}

export const subjectService = new SubjectService() 