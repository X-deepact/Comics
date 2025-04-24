import http from "../api";
import type { AdvertisementResponse } from "@/services/ads/types.ts";
import type { APIResponse } from "@/services/types.ts";

const getAds = async <T = AdvertisementResponse[]>(language: string): Promise<APIResponse<T>> => {
	try {
		const apiResponse = await http.get('/ads/comic', {
			params: {
			  language
			}
		  });

		if (!apiResponse?.data.data) {
			return {
				data: null as unknown as T,
				status: apiResponse.status,
				statusText: "No data received from API",
			}
		}

		return {
			data: apiResponse.data.data,
			status: apiResponse.status,
			statusText: apiResponse.statusText
		}

	} catch (e) {
		return {
			data: {} as T,
			status: 500,
			statusText: "Failed to fetch comic list",
		}
	}
}

export default {
	getAds
}