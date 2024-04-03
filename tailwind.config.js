/** @type {import('tailwindcss').Config} */

module.exports = {
    content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
    theme: {
        container: {
            center: true,
            padding: {
                DEFAULT: "1rem",
                mobile: "2rem",
                tablet: "4rem",
                desktop: "5rem",
            },
        },
    },
}

