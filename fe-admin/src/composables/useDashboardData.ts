import { ref, computed } from 'vue';
import { useAdStore } from '@/stores/adStore';
import { useComicStore } from '@/stores/comicStore';
import { useDramaStore } from '@/stores/dramaStore';
import { useUserStore } from '@/stores/userStore';

export function useDashboardData() {
  const adStore = useAdStore();
  const comicStore = useComicStore();
  const dramaStore = useDramaStore();
  const userStore = useUserStore();

  const loading = ref({
    stats: true,
    comics: true,
    users: true,
    dramas: true,
    ads: true
  });

  const stats = computed(() => [
    {
      name: 'Total Comics',
      value: comicStore.totalItems,
      icon: 'BookOpen',
    },
    {
      name: 'Total Users',
      value: userStore.totalItems,
      icon: 'Users',
    },
    {
      name: 'Total Dramas',
      value: dramaStore.totalItems,
      icon: 'Film',
    },
    {
      name: 'Total Ads',
      value: adStore.totalItems,
      icon: 'RectangleGroup',
    },
  ]);

  const recentComics = computed(() => comicStore.comicData.slice(0, 5));
  const recentUsers = computed(() => userStore.userData.slice(0, 5));
  const recentDramas = computed(() => dramaStore.dramaData.slice(0, 5));
  const recentAds = computed(() => adStore.adData.slice(0, 5));

  // For the content growth chart
  const contentStats = computed(() => ({
    comics: comicStore.totalItems,
    users: userStore.totalItems,
    dramas: dramaStore.totalItems,
    ads: adStore.totalItems,
  }));

  async function fetchStats() {
    loading.value.stats = true;
    try {
      await Promise.all([
        comicStore.getComicData(),
        userStore.getUserData(),
        dramaStore.getDramaData(),
        adStore.getAdData(),
      ]);
    } finally {
      loading.value.stats = false;
    }
  }

  async function fetchRecentData() {
    loading.value.comics = true;
    loading.value.users = true;
    loading.value.dramas = true;
    loading.value.ads = true;

    try {
      await Promise.all([
        comicStore.getComicData().finally(() => loading.value.comics = false),
        userStore.getUserData().finally(() => loading.value.users = false),
        dramaStore.getDramaData().finally(() => loading.value.dramas = false),
        adStore.getAdData().finally(() => loading.value.ads = false),
      ]);
    } catch (error) {
      console.error('Error fetching recent data:', error);
    }
  }

  return {
    stats,
    recentComics,
    recentUsers,
    recentDramas,
    recentAds,
    contentStats,
    loading,
    fetchStats,
    fetchRecentData,
  };
} 