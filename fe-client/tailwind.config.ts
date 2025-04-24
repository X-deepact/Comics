export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
	  screens: {
		  // max-width: 989px
		  'max-989': { max: '989px' },
		  
		  // min-width: 990px
		  'min-990': { min: '990px' },
		  
		  // min-width: 1240px
		  'min-1240': { min: '1240px' },
		  
		  // min-width: 1500px
		  'min-1500': { min: '1500px' },
          'xs': '375px',
          'sm': '640px',
          'md': '768px',
          'lg': '1024px',
          'xl': '1280px',
          '2xl': '1536px',
	  }
  },
  plugins: [],
  safelist: [
    {
      pattern: /bg-+/,
    },
  ]
}