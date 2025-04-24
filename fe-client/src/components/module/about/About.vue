<template>
  <div class="card flex flex-col justify-center items-center mt-5">
    <Tabs :value="activeTab" v-model="activeTab" class="w-full">
      <div class="w-full flex justify-center">
        <TabList
          class="flex gap-4 sm:gap-10 justify-start sm:justify-center overflow-x-auto whitespace-nowrap scrollbar-hide px-2 sm:px-0"
        >
          <Tab
            value="about"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.aboutUs") }}
          </Tab>
          <Tab
            value="contact"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.contactUs") }}
          </Tab>
          <Tab
            value="sitemap"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.siteMap") }}
          </Tab>
          <Tab
            value="legalStatement"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.legalNotice") }}
          </Tab>
          <Tab
            value="workRequirements"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.workRequirements") }}
          </Tab>
          <Tab
            value="comicsSigning"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.comicSigning") }}
          </Tab>
          <Tab
            value="contentModeration"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.auditMechanism") }}
          </Tab>
          <Tab
            value="welfareSystem"
            as="div"
            class="tab px-4 py-2 cursor-pointer transition-colors duration-200 text-gray-700 shrink-0"
          >
            {{ $t("message.welfareSystem") }}
          </Tab>
        </TabList>
      </div>      
      <TabPanels>
        <TabPanel value="about">
          <AboutUs />
        </TabPanel>
        <TabPanel value="contact">
          <ContactUs />
        </TabPanel>
        <TabPanel value="sitemap">
          <Sitemap />
        </TabPanel>
        <TabPanel value="legalStatement">
          <LegalStatement />
        </TabPanel>
        <TabPanel value="workRequirements">
          <WorkRequirement />
        </TabPanel>
        <TabPanel value="comicsSigning">
          <ComicSigning />
        </TabPanel>
        <TabPanel value="contentModeration">
          <ContentModeration />
        </TabPanel>
        <TabPanel value="welfareSystem">
          <WelfareSystem />
        </TabPanel>
      </TabPanels>
    </Tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { useRoute } from "vue-router";
import AboutUs from "@/components/module/about/AboutUs.vue";
import ContactUs from "@/components/module/about/ContactUs.vue";
import Sitemap from "@/components/module/about/Sitemap.vue";
import LegalStatement from "@/components/module/about/LegalStatement.vue";
import WorkRequirement from "@/components/module/about/WorkRequirement.vue";
import ComicSigning from "@/components/module/about/ComicSigning.vue";
import ContentModeration from "@/components/module/about/ContentModeration.vue";
import WelfareSystem from "@/components/module/about/WelfareSystem.vue";

const route = useRoute();

const activeTab = ref(route.query.tab || "about");

watch(
  () => route.query.tab,
  (newTab) => {
    activeTab.value = newTab || "about";
  }
);
</script>
<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.p-tab-active {
  color: orange !important;
}
</style>
