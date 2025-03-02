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
  date: Date | string | number,
  includeTime = false
): string {
  const parsedDate =
    typeof date === "string" || typeof date === "number"
      ? new Date(date)
      : date;

  if (isNaN(parsedDate.getTime())) {
    throw new Error("Invalid date input");
  }
  const day = String(parsedDate.getDate()).padStart(2, "0");
  const month = String(parsedDate.getMonth() + 1).padStart(2, "0"); // Months are 0-based
  const year = parsedDate.getFullYear();
  if (includeTime) {
    const min = String(parsedDate.getMinutes()).padStart(2, "0");
    const time = String(parsedDate.getHours()).padStart(2, "0");
    return `${day}/${month}/${year} ${time}:${min}`;
  }

  return `${day}/${month}/${year}`;
}
export const getTimeAgo = (
  date: Date | string | number,
  addSuffix = true
): string => {
  try {
    // Convert the input to a valid Date object
    const parsedDate =
      typeof date === "string" || typeof date === "number"
        ? new Date(date)
        : date;

    // Return the "time ago" formatted string
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
