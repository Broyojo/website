// @ts-check
import { defineConfig } from 'astro/config';
import rehypeKatex from "rehype-katex";
import remarkMath from "remark-math";

// https://astro.build/config
export default defineConfig({
    markdown: {
        shikiConfig: {
            theme: "dark-plus",
            wrap: true,
            defaultColor: false,
        },

        remarkPlugins: [remarkMath],
        rehypePlugins: [rehypeKatex],
    }
});
