export interface ComicRequest {
  active: boolean;
  audience: string;
  code: string;
  cover: string;
  description: string;
  language: string;
  title: string;
}

export interface ComicResponse {
  active: boolean;
  code: string;
  cover: string;
  created_at: string;
  created_by: number;
  description: string;
  id: number;
  title: string;
  updated_at: string;
  updated_by: number;
}

export interface ApiError {
  message: string;
} 