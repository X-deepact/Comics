export interface ChapterDetailResponse {
	active: boolean;
	active_from: number;
	comic_id: number;
	comic_name: string;
	cover: boolean;
	created_at: number;
	id: number;
	items?: ChapterItem[];
	name: string;
	next_chapter: number | null;
	number?: number;
	prev_chapter: number | null;
}

interface ChapterItem {
	active: boolean;
	active_from: number;
	created_at: number;
	image: string;
	page: number;
}

