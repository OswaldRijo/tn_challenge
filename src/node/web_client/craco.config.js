const path = require(`path`);

module.exports = {
  devServer: {
    port: 3000
  },
  webpack: {
    alias: {
      '@': path.resolve(__dirname, 'src/')
    }
  }
};
