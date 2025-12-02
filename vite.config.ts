import { sentrySvelteKit } from '@sentry/sveltekit';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  esbuild: {
    drop: ['console', 'debugger'],
  },
  plugins: [
    sentrySvelteKit({
      sourceMapsUploadOptions: {
        org: 'arizonabebras',
        project: 'shumi-space',
      },
    }),
    tailwindcss(),
    sveltekit(),
  ],
});
