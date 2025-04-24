import type {ChapterDetailResponse} from "@/services/chapters/type.ts";
import type {APIResponse} from "@/services/types.ts";
import http from "../api";

const getChapterDetail = async <T = ChapterDetailResponse>(id: string): Promise<APIResponse<T>> => {
  try {
    const apiResponse = await http.get(`chapter/${id}`)

    if (!apiResponse.data) {
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
  getChapterDetail
}
