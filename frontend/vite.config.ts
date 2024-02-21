import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [svelte()],
    resolve: {
        alias: {
            $lib: path.resolve("./src/lib"),
            $components: path.resolve("./src/components"),
            $assets: path.resolve("./src/assets"),
            $wails: path.resolve("./wailsjs"),
        }
    }
});