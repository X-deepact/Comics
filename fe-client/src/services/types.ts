export type APIResponse<T> = {
	statusText: string;
	data: T;
	status?: number;
	code?: string;
}

export type APIPaginationResponse<T> = {
	success?: boolean
	data: Array<T>;
	total: number;
	status?: number;
	statusText?: string;
}
