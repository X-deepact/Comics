import {computed, onMounted, onUnmounted, ref} from "vue";

export function useResponsive() {
    const windowWidth = ref(window.innerWidth);

    const is2xs = ref(false);
    const isXs = ref(false);
    const isSm = ref(false);
    const isMd = ref(false);
    const isLg = ref(false);
    const isXl = ref(false);
    const is2xl = ref(false);

    const isLargeScreen = computed(() => {
        return isLg.value || isXl.value || is2xl.value;
    });

    const isSmallAndMiddleScreen = computed(() => {
        return isXs.value || isSm.value || isMd.value;
    });

    const isExtraSmallAndSmallScreen = computed(() => {
        return isXs.value || isSm.value;
    });

    const isDoubleExtraSmallAndExtraSmallScreen = computed(() => {
        return is2xs.value || isXs.value;
    });

    const updateValues = () => {
        windowWidth.value = window.innerWidth;

        is2xs.value = windowWidth.value < 375;
        isXs.value = windowWidth.value >= 375 && windowWidth.value < 640;
        isSm.value = windowWidth.value >= 640 && windowWidth.value < 768;
        isMd.value = windowWidth.value >= 768 && windowWidth.value < 1024;
        isLg.value = windowWidth.value >= 1024 && windowWidth.value < 1280;
        isXl.value = windowWidth.value >= 1280 && windowWidth.value < 1536;
        is2xl.value = windowWidth.value >= 1536;
    };

    onMounted(() => {
        window.addEventListener('resize', updateValues);
        updateValues();
    });

    onUnmounted(() => {
        window.removeEventListener('resize', updateValues);
    });

    return {
        windowWidth,
        is2xs,
        isXs,
        isSm,
        isMd,
        isLg,
        isXl,
        is2xl,
        isLargeScreen,
        isSmallAndMiddleScreen,
        isExtraSmallAndSmallScreen,
        isDoubleExtraSmallAndExtraSmallScreen
    }
}