import { watch } from 'node:fs'

const entrypoints = ['./assets/scripts.js', './react/components/about_page.tsx']

const buildConfig = {
  entrypoints: entrypoints,
  outdir: './static',
  target: 'browser',
  format: 'esm',
  minify: true,
  externals: ['react', 'react-dom'],
}

const srcWatcher = watch(
  `${import.meta.dir}`,
  { recursive: true },
  (event, filename) => {
    Bun.build(buildConfig)

    console.log(`Detected ${event} in ${filename} ()`)
  }
)

await Bun.build(buildConfig)

console.log('Build complete.')
