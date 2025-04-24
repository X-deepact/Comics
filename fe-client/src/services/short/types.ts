import type { ChapterDetailResponse } from "@/services/chapters/type.ts";
import type { GenreResponse } from "@/services/shortGenres/types.ts";

export interface EpisodeResponse {
    created_at?: string;
    id?: number;
    number?: number;
    updated_at: string;
    video?: string;
}

export interface GenresForShortDramaResponse {
    created_at?: string;
    id: number;
    language?: string;
    name?: string;
    position: number;
    translated_name: string;
    updated_at?: string;
}

export interface ShortDramaDetailResponse {
    created_at: string;
    description: string;
    episodes: EpisodeResponse[];
    genres: GenresForShortDramaResponse[];
    id: number;
    language: number;
    name: string;
    release_date: string;
    thumb: string;
    translated_name: string;
    updated_at: string;
}

export interface ShortDetailResponse {
    id: number;
    name: string;
    code: string;
    language: string;
    audience: string;
    description: string;
    cover: string;
    active: boolean;
    created_at?: number;
    latest_date?: string;
    chapters?: ChapterDetailResponse[];
    genres?: GenreResponse[];
    updated_at?: number
    authors?: [
        {
            name: string
        }
    ]
}
