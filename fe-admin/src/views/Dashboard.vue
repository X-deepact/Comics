<script setup lang="ts">
import { onMounted } from 'vue';
import StatCard from '@/components/dashboard/StatCard.vue';
import StatCardSkeleton from '@/components/dashboard/StatCardSkeleton.vue';
import ContentChart from '@/components/dashboard/ContentChart.vue';
import RecentItemsList from '@/components/dashboard/RecentItemsList.vue';
import { useDashboardData } from '@/composables/useDashboardData';

const {
  stats,
  recentComics,
  recentUsers,
  recentDramas,
  recentAds,
  contentStats,
  loading,
  fetchStats,
  fetchRecentData
} = useDashboardData();

onMounted(() => {
  fetchStats();
  fetchRecentData();
});
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-semibold mb-6">Dashboard Overview</h1>
    
    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <template v-if="loading.stats">
        <stat-card-skeleton v-for="i in 4" :key="i" />
      </template>
      <template v-else>
        <stat-card 
          v-for="stat in stats" 
          :key="stat.name" 
          :stat="stat" 
        />
      </template>
    </div>

    <!-- Content Growth Chart -->
    <content-chart
      :loading="loading.comics || loading.users || loading.dramas || loading.ads"
      :content-stats="contentStats"
    />

    <!-- Recent Items Grid -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mt-8">
      <recent-items-list
        title="Recent Comics"
        :loading="loading.comics"
        :items="recentComics"
        type="comics"
      />
      <recent-items-list
        title="Recent Users"
        :loading="loading.users"
        :items="recentUsers"
        type="users"
      />
      <recent-items-list
        title="Recent Dramas"
        :loading="loading.dramas"
        :items="recentDramas"
        type="dramas"
      />
      <recent-items-list
        title="Recent Ads"
        :loading="loading.ads"
        :items="recentAds"
        type="ads"
      />
    </div>
  </div>
</template>

<style src="@/styles/icons.css" />
