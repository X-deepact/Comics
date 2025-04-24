<template>
  <footer class="bg-gray-800 text-white text-sm w-full">
    <div class="container mx-auto px-4 py-4">
      <div v-if="!isDoubleExtraSmallAndExtraSmallScreen" class="top-section flex justify-between items-center mb-8">
        <div class="top-left flex items-center">
          <div class="logo-container mr-4">
            <Image
              src="/src/assets/images/footer_logo.png"
              alt="Logo"
              class="h-12"
            />
          </div>
          <div class="email flex flex-col">
            <div>
              {{ $t("footer.submissionEmail") }}：<a
                href="#"
                class="hover:text-orange-400"
                >888888@qq.com</a
              >
            </div>
            <div class="flex mt-2">
              <div
                v-for="(button, index) in footerButtons"
                :key="index"
                class="mr-4"
                @click="handleButtonClick(button.tab, 'about')"
              >
                <Button
                  :label="$t(button.label)"
                  class="hover:text-orange-400 no-underline"
                />
              </div>
            </div>
          </div>
        </div>

        <div class="customer-service flex items-center justify-center flex-1">
          <div class="service-left flex items-center">
            <span class="support mr-2"
              >{{ $t("footer.onlineCustomerService") }}：</span
            >
            <Image
              src="/src/assets/images/support.png"
              alt="Image"
              class="h-8"
            />
          </div>
          <div class="service-right flex items-center ml-8">
            <span class="noti mr-2">{{ $t("footer.QQGroup") }}：</span>
            <Image src="/src/assets/images/noti.png" alt="Image" class="h-8" />
          </div>
        </div>

        <!-- Empty Spacer to Balance the Layout -->
        <div class="top-right flex-1"></div>
      </div>

      <!-- Bottom Section -->
      <div class="bottom-section border-t border-gray-700 pt-4 text-center">
        <div
          class="footer-links flex flex-col sm:flex-row justify-center items-center mb-4 space-x-2"
        >
          <template v-for="(button, index) in footerLinks" :key="index">
            <Button
              :label="$t(button.label)"
              class="footer-link p-button-text hover:text-orange-400"
              @click="handleButtonClick(button.tab, 'about')"
            />
            <span
              v-if="index < footerLinks.length - 1"
              class="text-gray-500 m-4"
            >
              |
            </span>
          </template>
        </div>

        <p class="footer-info mb-2">
          {{ $t("footer.onlineBrowsingPlatform") }}
        </p>
        <p class="copyright">
          Copyright ©2010-2020 {{ $t("footer.allRightServed") }}
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import {useResponsive} from "@/composables/useResponsive.ts";

const router = useRouter();
const { isDoubleExtraSmallAndExtraSmallScreen } = useResponsive();

const footerButtons = [
  { label: "footer.workRequirements", tab: "workRequirements" },
  { label: "footer.comicsSigning", tab: "comicsSigning" },
  { label: "footer.reviewMechanism", tab: "contentModeration" },
  { label: "footer.welfareSystem", tab: "welfareSystem" },
];
const footerLinks = [
  { label: "footer.aboutUs", tab: "about" },
  { label: "footer.contactUs", tab: "contact" },
  { label: "footer.sitemap", tab: "sitemap" },
  { label: "footer.legalStatement", tab: "legalStatement" },
  { label: "footer.contentModeration", tab: "contentModeration" },
  { label: "footer.websiteStatistics", tab: "websiteStatistics" },
];

const handleButtonClick = (tab: string, action: string) => {
  router.push({ path: action, query: { tab: tab } });
};
</script>

<style scoped>
/* Custom styles if needed */
</style>
