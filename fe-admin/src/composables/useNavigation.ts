import { computed } from "vue";
import { useRoute } from "vue-router";
import type { NavigationItem } from "../types/navigation";

export function useNavigation() {
  const route = useRoute();

  const navigation = computed<NavigationItem[]>(() => [
    {
      title: "Content Manage",
      icon: "file",
      children: [
        {
          name: "Comics Manage",
          href: "/dashboard/comics",
          active: route.path === "/dashboard/comics",
        },
        {
          name: "Genre Manage",
          href: "/dashboard/subjects",
          active: route.path === "/dashboard/subjects",
        },
        {
          name: `Author Manage`,
          href: "/dashboard/authors",
          active: route.path === `/dashboard/authors`,
        },
      ],
    },
    {
      title: "Ads and Recommendation",
      icon: "target",
      children: [
        {
          name: "Ads",
          href: "/dashboard/ads",
          active: route.path === "/dashboard/ads",
        },
        {
          name: "Recommendation",
          href: "/dashboard/recommendation",
          active: route.path === "/dashboard/recommendation",
        },
      ],
    },
    {
      title: "System Manage",
      icon: "settings",
      children: [
        {
          name: "User",
          href: "/dashboard/users",
          active: route.path === "/dashboard/users",
        },
      ],
    },
  ]);

  return {
    navigation,
  };
}
