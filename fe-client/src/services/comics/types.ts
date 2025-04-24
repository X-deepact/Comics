import type { ChapterDetailResponse } from "@/services/chapters/type.ts";
import type { GenreResponse } from "@/services/genres/types.ts";

export interface ComicDetailResponse {
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

export interface ComicRecommendResponse {
    id: number;
    title: string;
    cover: string;
    position: string;
    active_from: string;
    active_to: string;
    created_at: string;
    updated_at: string;
    comics: ComicDetailResponse[];
}
