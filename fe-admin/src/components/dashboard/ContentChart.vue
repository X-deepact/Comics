<script setup lang="ts">
import { computed } from 'vue';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js';
import { Line } from 'vue-chartjs';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

const props = defineProps<{
  loading: boolean;
  contentStats: {
    comics: number;
    users: number;
    dramas: number;
    ads: number;
  };
}>();

const chartData = computed(() => ({
  labels: ['Comics', 'Users', 'Dramas', 'Ads'],
  datasets: [
    {
      label: 'Content Count',
      data: [
        props.contentStats.comics,
        props.contentStats.users,
        props.contentStats.dramas,
        props.contentStats.ads
      ],
      borderColor: 'rgb(75, 192, 192)',
      backgroundColor: 'rgba(75, 192, 192, 0.5)',
      tension: 0.4,
      fill: true
    }
  ]
}));

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      position: 'top' as const,
    },
    title: {
      display: false
    }
  },
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        stepSize: 1
      }
    }
  }
};
</script>

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <h3 class="text-lg font-semibold mb-4">Content Overview</h3>
    
    <div v-if="loading" class="h-64 flex items-center justify-center">
      <div class="animate-pulse text-gray-400">Loading chart data...</div>
    </div>
    
    <div v-else class="h-64">
      <Line
        :data="chartData"
        :options="chartOptions"
      />
    </div>
  </div>
</template> 