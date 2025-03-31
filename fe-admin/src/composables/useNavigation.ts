import { computed } from "vue";
import { useRoute } from "vue-router";
import type { NavigationItem } from "../types/navigation";

export function useNavigation() {
  const route = useRoute();

  const navigation = computed<NavigationItem[]>(() => [
    {
      title: "Comics Management",
      icon: "file",
      children: [
        {
          name: "Comics",
          href: "/dashboard/comics",
          active: route.path === "/dashboard/comics",
        },
        {
          name: "Genres",
          href: "/dashboard/subjects",
          active: route.path === "/dashboard/subjects",
        },
        {
          name: `Authors`,
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
          name: "Recommendations",
          href: "/dashboard/recommendation",
          active: route.path === "/dashboard/recommendation",
        },
      ],
    },
    {
      title: "Dramas Management",
      icon: "target",
      children: [
        {
          name: "Dramas",
          href: "/dashboard/dramas",
          active: route.path === "/dashboard/dramas",
        },
        {
          name: "Genres",
          href: "/dashboard/genres",
          active: route.path === "/dashboard/genres",
        },
      ],
    },
    {
      title: "System Management",
      icon: "settings",
      children: [
        {
          name: "Users",
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
