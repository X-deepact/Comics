import http from "../api";
import type { APIPaginationResponse, APIResponse } from "../types";
import type {ComicDetailResponse, ComicRecommendResponse} from "./types";

interface ComicRequestBody {
  genre: string;
  progress: string;
  audience: string
  originalLanguage: string;
  locale: string;
  page: string;
}

const getComics = async <T = ComicDetailResponse>(body: ComicRequestBody): Promise<APIPaginationResponse<T>> => {
  try {
    const apiResponse = await http.get(`comic`, {
      params: {
        page_size: 12,
        page: body.page,
        genre_id: body.genre,
        progress: body.progress,
        audience: body.audience,
        language: body.locale,
        original_language: body.originalLanguage
      }
    });

    if (!apiResponse?.data.data) {
      return {
        data: null as unknown as Array<T>,
        total: 0,
        status: apiResponse.status,
        statusText: "No data received from API",
      }
    }

    return {
      data: apiResponse.data.data,
      total: apiResponse.data.pagination.total,
      status: apiResponse.status,
      statusText: apiResponse.statusText
    }

  } catch (e) {
    return {
      data: {} as Array<T>,
      total: 0,
      status: 500,
      statusText: "Failed to fetch comic list",
    }
  }
}

const getComicDetails = async <T = ComicDetailResponse>(id: string, locale: string): Promise<APIResponse<T>> => {
  try {
    const apiResponse = await http.get(`comic/${id}`, {
      headers: {
        'X-Language': locale
      }
    });

    if (!apiResponse?.data?.data) {
      return {
        data: null as unknown as T,
        status: apiResponse.status,
        code: apiResponse.data.code,
        statusText: "No data received from API",
      }
    }

    return {
      data: apiResponse.data.data,
      status: apiResponse.status,
      code: apiResponse.data.code,
      statusText: apiResponse.statusText
    }

  } catch (e) {
    return {
      data: {} as T,
      status: 500,
      code: 'ERROR',
      statusText: "Failed to fetch comic list",
    }
  }
}


const getRecentlyUpdated = async <T = ComicDetailResponse>(page: number = 1, language: string): Promise<APIPaginationResponse<T>> =>  {
  try {
    const apiResponse = await http.get(`comic/recent`, {
      params: {
        page_size: 12,
        page,
        language
      }
    });

    if (!apiResponse.data) {
      return {
        success: true,
        data: null as unknown as Array<T>,
        total: 0,
        status: 204,
        statusText: "No data received from API",
      }
    }

    return {
      success: true,
      data: apiResponse.data.data,
      total: apiResponse.data.pagination.total,
      status: apiResponse.status,
      statusText: apiResponse.statusText
    }
  } catch (e: any) {
    return {
      success: false,
      data: {} as Array<T>,
      total: 0,
      status: 500,
      statusText: e?.message || "Failed to fetch recently update list"
    }
  }
}

const getRecommendComics = async <T = ComicRecommendResponse[]>(language: string): Promise<APIResponse<T>> => {
  try {
    const apiResponse = await http.get(`recommend-comics/all`, {
      params: {
        language
      }
    });

    if (apiResponse?.data) {
      return {
        data: apiResponse.data.data,
        status: apiResponse.status,
        statusText: apiResponse.statusText
      }
    }

    return {
      data: null as unknown as T,
      status: 204,
      statusText: "No data received from API",
    }

  } catch (e) {
    return {
      data: {} as T,
      status: 500,
      statusText: "Failed to fetch recommend comic list",
    }
  }
}

const getComicsSearch = async <T = ComicDetailResponse>(
  search_keyword: string = "",
  page: number = 1,
  language: string,
  signal?: AbortSignal
): Promise<APIPaginationResponse<T>> => {
  try {
    const apiResponse = await http.get(`/comic/searchAll`, {
      params: {
        search_keyword,
        page_size: 12,
        page,
        language
      },
      signal
    });

    if (apiResponse?.data) {
      return {
        success: true,
        data: apiResponse.data.data,
        total: apiResponse.data.pagination.total,
        status: 200,
        statusText: apiResponse.statusText
      };
    }
    return {
      success: false,
      data: [] as Array<T>,
      total: 0,
      status: 204,
      statusText: "No data received from API",
    };
  } catch (e: any) {
    if (e.name === "AbortError") {
      return {
        success: false,
        data: null as unknown as Array<T>,
        total: 0,
        status: 499, // a common custom code for Client Closed Request
        statusText: "Request was aborted",
      };
    }
    return {
      success: false,
      data: null as unknown as Array<T>,
      total: 0,
      status: 500,
      statusText: "Failed to fetch comic list",
    };
  }
};


export default {
  getComics,
  getComicDetails,
  getRecentlyUpdated,
  getComicsSearch,
  getRecommendComics
};