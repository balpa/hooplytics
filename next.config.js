/** @type {import('next').NextConfig} */
const nextConfig = {
  async redirects() {
      return [
        {
          source: '/',
          destination: '/home',
          permanent: true,
        },
      ]
    },
}

const envConfig = {
  NEXT_PUBLIC_TEST_KEY: process.env.NEXT_PUBLIC_TEST_KEY,
}

module.exports = {
  ...nextConfig,
  env: envConfig
}
