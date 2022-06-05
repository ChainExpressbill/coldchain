module.exports = {
  content: ['./public/**/*.html', './src/**/*.{js,jsx,ts,tsx}'],
  important: false,
  theme: {
    extend: {
      colors: {
        main: '#022a68',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
};
