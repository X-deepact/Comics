<template>
  <Section icon="src/assets/images/ic__ff.png"
           :hide-btn="true"
           text="Recommendations"
  >
    <template #main>
      <template v-if="!featuredItem">
        {{"Loading Skeleton"}}
      </template>
      <div v-else class="grid grid-cols-2 sm:grid-cols-4 lg:grid-cols-6 gap-2 mx-3 md:mx-0">
        <div
            v-for="(item, index) in props.items"
            :key="index"
        >
          <ComicCardWithBanner :item>
            <template #comicDescription>
              <DescriptionCard
                  :id="item.id"
                  :title="item.name"
                  :description="item.description === '' ? 'No description' : item.description"
                  items=""
              />
            </template>
          </ComicCardWithBanner>
        </div>
      </div>
    </template>
  </Section>
</template>

<script setup lang="ts">
import ComicCardWithBanner from "@/components/module/recently-update/ComicCardWithBanner.vue";
import { defineComponent, h, computed } from "vue";
import type {ComicDetailResponse} from "@/services/comics/types.ts";
import { useRouter } from "vue-router";
import Section from "@/components/common/Section.vue";

const router = useRouter();

const props = defineProps<{
  items: ComicDetailResponse[] | undefined;
}>();

const hasItems = computed(() => Array.isArray(props.items) && props.items.length > 0);

const featuredItem = computed(() => {
  if (!hasItems.value) return undefined;
  return props.items![0];
});

const DescriptionCard = defineComponent({
  name: 'DescriptionCard',
  props: {
    id: { type: Number, default: '' },
    title: { type: String, default: '' },
    description: { type: String, default: '' },
  },
  setup(props) {
    return () => h('div', { class: 'w-full mx-2' }, [
      h('div', {
        class: 'mt-1 font-bold hover:text-[#ff620e] w-full justify-start overflow-hidden whitespace-nowrap text-ellipsis cursor-pointer',
        onClick: () => onNavigateToComicDetails(props.id),
      }, props.title),
      h('div', { class: 'text-sm leading-[20px] text-black/50 dark:text-white my-1 overflow-hidden whitespace-nowrap text-ellipsis' }, props.description),
    ]);
  }
});

const onNavigateToComicDetails = (id: number) => {
  window.scrollTo({ top: 0, behavior: "smooth" });
  router.push({
    name: "comic-detail",
    params: {
      comicId: id.toString(),
    },
  });
};
</script>
