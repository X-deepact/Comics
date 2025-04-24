export interface CategoryButton {
  name: string;
  group: string;
}

export const getProgressesList = (): CategoryButton[] => {
  return [
    { name: "all", group: "progress" },
    { name: "ongoing", group: "progress" },
    { name: "completed", group: "progress" },
  ];
};

export const getAudiencesList = (): CategoryButton[] => {
  return [
    { name: "all", group: "audience" },
    { name: "boys", group: "audience" },
    { name: "girls", group: "audience" },
    { name: "children", group: "audience" },
  ];
};

export const getOriginalLanguageList = (): CategoryButton[] => {
  return [
    { name: "all", group: "originalLanguage" },
    { name: "jp", group: "originalLanguage" },
    { name: "ko", group: "originalLanguage" },
    { name: "zh-cn", group: "originalLanguage" },
    { name: "zh-tw", group: "originalLanguage" },
    { name: "en", group: "originalLanguage" },
  ];
};

export const getReleaseYearList = (): CategoryButton[] => {
  return [
    { name: "all", group: "releaseYear" },
    { name: "2023", group: "releaseYear" },
    { name: "2024", group: "releaseYear" },
    { name: "2025", group: "releaseYear" },
  ];
};