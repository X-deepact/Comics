import type {APIPaginationResponse, APIResponse} from "../types";
import type {AxiosRequestConfig} from "axios";
import http from "../api";

const get = async <T = unknown>(
    path: string,
    options?: {
        locale?: string;
        config?: AxiosRequestConfig;
        params?: any
    }
): Promise<APIResponse<T>> => {
    try {
        const {locale, config, params} = options || {};
        const apiResponse = await http.get(path, {
            params,
            ...config,
            headers: {
                'X-Language': locale,
                ...(config?.headers || {})
            }
        });

        const isText = config?.responseType === 'text';

        return {
            data: isText ? apiResponse.data : apiResponse.data?.data ?? null,
            status: apiResponse.status,
            code: isText ? 'OK' : apiResponse.data?.code ?? 'NO_DATA',
            statusText: apiResponse.statusText,
        };
    } catch (e) {
        return {
            data: {} as T,
            status: 500,
            code: 'ERROR',
            statusText: "Failed to fetch data",
        };
    }
}

const getList = async <T = unknown>(
    path: string,
    options?: {
        locale?: string,
        config?: AxiosRequestConfig,
        params?: any
    }
): Promise<APIPaginationResponse<T>> => {
    try {
        const {locale, config, params} = options || {};
        const apiResponse = await http.get(path, {
            params,
            ...config,
            headers: {
                'X-Language': locale,
                ...(config?.headers || {})
            }
        });

        if (!apiResponse?.data.data) {
            return {
                data: [] as unknown as Array<T>,
                total: 0,
                status: apiResponse.status,
                statusText: "No data received from API",
            }
        }

        return {
            data: apiResponse.data.data,
            total: apiResponse.data.pagination ? apiResponse.data.pagination.total : 0,
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


export default {
    get,
    getList
}
