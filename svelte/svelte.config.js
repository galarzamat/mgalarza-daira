import { vitePreprocess } from "@sveltejs/vite-plugin-svelte"; // Importa vitePreprocess una sola vez
import adapter from "@sveltejs/adapter-auto";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
  },
  preprocess: [vitePreprocess({})], // Configura preprocess aquí dentro de config
};

export default config; // Exporta la configuración
