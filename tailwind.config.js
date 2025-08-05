module.exports = {
  content: ['./*.go', './*.templ'],
  darkMode: ['class', '.theme--dark'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui', 'sans-serif'],
      },
      fontSize: {
        xs: '0.7rem',
        sm: '0.8rem',
        '2xl': ['1.5rem']
      },
      colors: {
        lavender: '#fdf0f5',
        strongpink: '#7dd3fc',
        crimson: '#548ea8ff',
        garnet: '#0f172a'
      },
      borderWidth: {
        '05rem': '0.5rem' // you can define custom border widths in rem units
      },
      gap: {
        '48vw': '4.8vw'
      },
      typography: ({theme}) => ({
        /* for markdown html content */
        DEFAULT: {
          css: {
            '--tw-prose-invert-body': theme('colors.neutral[100]'),
            '--tw-prose-invert-quotes': theme('colors.neutral[100]'),
            '--tw-prose-headings': theme('colors.strongpink'),
            '--tw-prose-invert-headings': theme('colors.strongpink'),
            '--tw-prose-links': theme('colors.gray[700]'),
            '--tw-prose-invert-links': theme('colors.neutral[50]'),
            maxWidth: '100%',
            a: {
              'font-weight': 300
            }
          }
        }
      })
    }
  },
  plugins: [require('@tailwindcss/typography')]
}
