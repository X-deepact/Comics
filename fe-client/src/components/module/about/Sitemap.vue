<template>
  <div class="m-10">
    <div class="mb-10">
      <h3 class="mb-2 font-bold">功能模块</h3>
      <div class="flex flex-wrap space-x-4">
        <span
          v-for="module in modules"
          :key="module"
          class="text-sm hover:text-orange-500 cursor-pointer leading-7"
          >{{ module }}</span
        >
      </div>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">题材</h3>
      <div class="flex flex-wrap space-x-4">
        <span
          v-for="theme in themes"
          :key="theme"
          class="text-sm hover:text-orange-500 cursor-pointer leading-7"
          >{{ theme }}</span
        >
      </div>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">VIP漫画</h3>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">付费漫画</h3>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">完结大作</h3>
      <div class="flex flex-wrap space-x-4">
        <span
          v-for="completeMasterpiece in completeMasterpieces"
          :key="completeMasterpiece"
          class="text-sm hover:text-orange-500 cursor-pointer leading-7"
          >{{ completeMasterpiece }}</span
        >
      </div>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">我是读者</h3>
      <div class="flex flex-wrap space-x-4">
        <span
          v-for="reader in readers"
          :key="reader"
          class="text-sm hover:text-orange-500 cursor-pointer leading-7"
          >{{ reader }}</span
        >
      </div>
    </div>
    <div class="mb-10">
      <h3 class="mb-2 font-bold">我是作者</h3>
      <div class="flex flex-wrap space-x-4">
        <span
          v-for="author in authors"
          :key="author"
          class="text-sm hover:text-orange-500 cursor-pointer leading-7"
          >{{ author }}</span
        >
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";

const modules = ref(["首页", "更新", "排行", "分类"]);

const themes = ref([
  "热血",
  "冒险",
  "科幻",
  "霸总",
  "玄幻",
  "校园",
  "修真",
  "搞笑",
  "穿越",
  "后宫",
  "职美",
  "恋爱",
  "悬疑",
  "恐怖",
  "战争",
  "动作",
  "同人",
  "竞技",
  "励志",
  "架空",
  "灵异",
  "百合",
  "古风",
  "生活",
  "真人",
  "都市",
  "其他",
  "爱情",
  "武侠",
  "格斗",
  "侦探",
  "推理",
  "剧悚",
  "其它",
  "总裁",
  "奇幻",
  "韩漫",
  "纯爱",
  "日常",
  "少年",
  "大女主",
  "异能",
  "战斗",
  "欧式宫廷",
  "治愈",
  "魔幻",
  "欢乐向",
  "轻小说",
  "奇幻爱情",
  "心机",
  "腹黑",
  "女强",
  "穿书",
  "神鬼",
  "历史",
  "官廷",
  "唯美",
  "冒险热血",
  "重生",
  "系统",
  "少女",
  "彩虹",
  "豪门总裁",
  "惊悚",
  "傲娇",
  "戏精",
  "双喜冤家",
  "游戏",
  "武侠格斗",
  "老实人",
  "复仇",
  "逆袭",
  "生活漫画",
  "都市脑洞",
  "修仙",
  "侦探推理",
  "欧风",
  "强者归来",
  "西方",
  "娱乐圈",
  "偏执",
  "双向暗恋",
  "双穿",
  "炸毛",
  "久别重逢",
  "狼人",
  "年下",
  "救赎",
  "橘里橘气",
  "言情",
  "ABO",
  "替身梗",
  "甜宠",
  "强制爱",
  "颜控",
  "独宠",
  "一见钟情",
  "男神",
  "兽人",
  "养娃",
]);

const completeMasterpieces = ref([
  "家庭教师",
  "我独自升级",
  "猫之茗",
  "总裁在上",
  "家有女友",
  "血族禁域",
  "幽游白书",
  "骑士幻想夜",
  "阎王不高兴",
  "家兄又在作死",
]);

const readers = ref([
  "读者中心",
  "我的收藏",
  "阅读历史",
  "我的购买",
  "我的评论",
  "消费记录",
]);

const authors = ref([
  "作者中心",
  "我的作品",
  "新增漫画",
  "作者认证",
  "结算中心",
  "读者评论",
]);
</script>
