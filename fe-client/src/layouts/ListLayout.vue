<template>
  <div
    class="
      sm:w-[600px]
      md:w-[700px]
      lg:w-[1000px]
      xl:w-[1200px]
      2xl:w-[1500px]
      m-auto
      mb-10
    "
  >
    <div class="w-full h-fit my-7 text-[20px] leading-6">
      <slot name="top" />
    </div>
    <slot name="main" />
    <div class="w-full my-10 m-auto ">
      <slot name="bottom" />
    </div>
  </div>
</template>

