import type { Updater } from "@tanstack/vue-table";
import type { Ref } from "vue";
import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { formatDistanceToNow } from "date-fns";
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function valueUpdater<T extends Updater<any>>(
  updaterOrValue: T,
  ref: Ref
) {
  ref.value =
    typeof updaterOrValue === "function"
      ? updaterOrValue(ref.value)
      : updaterOrValue;
}
export function formatDate(
  date: Date | string | number | null | undefined,
  includeTime = false
): string {
  if (!date) return "N/A";

  const parsedDate =
    typeof date === "string" || typeof date === "number"
      ? new Date(date)
      : date;

  if (isNaN(parsedDate.getTime())) {
    throw new Error("Invalid date input");
  }

  const day = String(parsedDate.getDate()).padStart(2, "0");
  const month = String(parsedDate.getMonth() + 1).padStart(2, "0");
  const year = parsedDate.getFullYear();

  if (includeTime) {
    const hours = String(parsedDate.getHours()).padStart(2, "0");
    const minutes = String(parsedDate.getMinutes()).padStart(2, "0");
    return `${day}/${month}/${year} ${hours}:${minutes}`;
  }

  return `${day}/${month}/${year}`;
}
export const getTimeAgo = (
  date: Date | string | number | null | undefined,
  addSuffix = true
): string => {
  if (!date) return "N/A";

  try {
    const parsedDate =
      typeof date === "string" || typeof date === "number"
        ? new Date(date)
        : date;

    return formatDistanceToNow(parsedDate, { addSuffix });
  } catch (error) {
    console.error("Invalid date provided to getTimeAgo:", error);
    return "Invalid date";
  }
};

export const langConverter = (lang: string): string => {
  switch (lang) {
    case "en":
      return "English";
    case "cn":
      return "Chinese";
    case "vi":
      return "Vietnamese";
    case "all":
      return "All";
  }
  return "";
};
