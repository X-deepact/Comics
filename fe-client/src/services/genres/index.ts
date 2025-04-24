import http from "../api";
import type { GenreResponse } from "@/services/genres/types.ts";
import type {APIResponse} from "@/services/types.ts";

interface GenreRequestBody {
  isHomePage: boolean;
  locale: string;
}

const getGenres = async <T = GenreResponse[]>(body: GenreRequestBody): Promise<APIResponse<T>> => {
  try {
    const apiResponse = await http.get('/genre/all', {
      params: {
        isHomePage: body.isHomePage,
        language: body.locale,
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
      statusText: "Failed to fetch genres list",
    }
  }
};

export default {
  getGenres,
}