import { defineConfig, loadEnv } from 'vite';
import preact from '@preact/preset-vite'

export default ({ mode }) => {
  // loadEnv(mode, process.cwd()) will load the .env files depending on the mode
  // import.meta.env.VITE_BASE_APP available here with: process.env.VITE_BASE_APP
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [preact()],
    base: process.env.VITE_BASE_APP,
    build: {
      emptyOutDir: true,
      manifest: true,
      sourcemap: true,
      assetsDir: './', // put the assets next to the index.html file
    },
  });
};