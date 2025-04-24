<template>
  <ListLayout>
    <template #top>
      <DashboardTopSection :items="toppingList[0]?.comics" />
    </template>
    <template #main>
      <DashboardNav v-if="isLargeScreen" />

      <DashboardNewUpdate :items="recommendationsList" />

      <DashboardAds
        ad-image="src/assets/images/ad1.jpg"
      />

      <DashboardSlider
        slider-icon="src/assets/images/ic__dj.png"
        :slider-header="$t('dashboard.completeMasterpiece')"
        :items="completeMasterpieceList"
      />

      <DashboardSlider
        slider-icon="src/assets/images/ic__ss.png"
        :slider-header="$t('dashboard.fastestRising')"
        :items="fastestRisingList"
      />

      <DashboardSlider
        slider-icon="src/assets/images/ic__hz.png"
        :slider-header="$t('dashboard.newPublishing')"
        :items="newProductList"
      />

      <Section
        icon="src/assets/images/ic__zf.png"
        :text="$t('recentlyUpdate.pageTitle')"
        @navigate="router.push('/recently-updated')"
      >
        <template #main>
          <template v-if="isLoading">
            <ListComicSkeleton />
          </template>
          <div v-else>
            <div v-if="!currentRecentlyUpdated || currentRecentlyUpdated.length === 0">
              <p class="text-center text-lg">No comics found</p>
            </div>
            <ComicList v-else :data="currentRecentlyUpdated" />
          </div>
        </template>
      </Section>

      <DashboardAds
        ad-image="src/assets/images/ad3.jpg"
      />
    </template>
  </ListLayout>
  <ScrollToTopBtn />
</template>

<script lang="ts" setup>
import ListLayout from "@/layouts/ListLayout.vue";
import DashboardNav from "@/components/module/dashboard/DashboardNav.vue";
import DashboardNewUpdate from "@/components/module/dashboard/DashboardNewUpdate.vue";
import DashboardAds from "@/components/module/dashboard/DashboardAds.vue";
import DashboardSlider from "@/components/module/dashboard/DashboardSlider.vue";
import ComicList from "@/components/module/recently-update/ComicList.vue";
import ScrollToTopBtn from "@/components/common/ScrollToTopBtn.vue";
import ListComicSkeleton from "@/components/skeleton/ListComicSkeleton.vue";
import { useRecommend } from "@/stores/homepage/useRecommend.ts";
import { storeToRefs } from "pinia";
import DashboardTopSection from "@/components/module/dashboard/DashboardTopSection.vue";
import { useRecentlyUpdated } from "@/stores/recently-updated/useRecentlyUpdated.ts";
import { useRouter } from "vue-router";
import { onMounted } from "vue";
import { useResponsive } from "@/composables/useResponsive";

const { isLargeScreen } = useResponsive();
const router = useRouter();
const store = useRecommend();
const recentlyUpdatedStore = useRecentlyUpdated();
const {
  completeMasterpieceList,
  recommendationsList,
  fastestRisingList,
  newProductList,
  toppingList
} = storeToRefs(store);

onMounted(() => {
  store.fetchRecommendation();
  recentlyUpdatedStore.fetchComics();
})
const {
  currentRecentlyUpdated,
  isLoading
} = storeToRefs(recentlyUpdatedStore);
</script>
