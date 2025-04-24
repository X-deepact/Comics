<template>
  <div class="relative group">
    <slot name="btn" />
    <div
      class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 px-2 py-1 rounded bg-black text-white
      text-xs opacity-0 group-hover:opacity-100 transition pointer-events-none z-20 whitespace-nowrap"
    >
      <slot name="tooltip" />
    </div>
  </div>
</template>
