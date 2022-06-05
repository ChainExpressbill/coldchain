module.exports = {
  content: ['./public/**/*.html', './src/**/*.{js,jsx,ts,tsx}'],
  important: false,
  theme: {
    extend: {
      colors: {
        main: '#022a68',
      },
      boxShadow: {
        basic: '1px 2px 5px 0px rgba(0, 0, 0, 0.8)',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
};
