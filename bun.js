import { watch } from 'node:fs'

const buildConfig = {
  entrypoints: ['./react/components/about_page.tsx', './assets/scripts.js'],
  outdir: './static',
  target: 'browser',
  format: 'esm',
  minify: true,
  externals: ['react', 'react-dom'],
}

const srcWatcher = watch(
  `${import.meta.dir}`,
  { recursive: false },
  (event, filename) => {
    Bun.build(buildConfig)

    console.log(`Detected ${event} in ${filename} ()`)
  }
)

await Bun.build(buildConfig)

console.log('Build complete.')
