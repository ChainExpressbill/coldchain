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
      gridTemplateColumns: {
        order: '1.7fr 0.8fr 0.8fr 1.4fr 0.3fr 1fr',
      },
    },
  },
  plugins: [require('@tailwindcss/typography')],
};
