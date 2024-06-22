import { defineConfig } from "astro/config";
import mdx from "@astrojs/mdx";

import sitemap from "@astrojs/sitemap";

// https://astro.build/config
export default defineConfig({
    site: "https://axyut.github.io",
    base: "/playgo",
    integrations: [mdx(), sitemap()],
});
