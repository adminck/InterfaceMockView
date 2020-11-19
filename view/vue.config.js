'use strict'

const path = require('path')


function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  outputDir: '../dist',
  assetsDir: 'static',
  configureWebpack: {
    //    @路径走src文件夹
    resolve: {
      alias: {
        '@': resolve('src')
      }
    }
  },
}
