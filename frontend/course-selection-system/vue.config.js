module.exports = {
    devServer: {
      open: true,  // 自动打开浏览器
      host: 'localhost',
      port: 8080,
      https: false,
      //以上的ip和端口是我们本机的;下面为需要跨域的
      proxy: {
        // 配置跨域
        '/api': {
          target: 'http://175.178.98.212:8080/',  // 后端接口地址
          ws: true,
          changOrigin: true,  // 允许跨域
          pathRewrite: {
            // '^/api': ''  // 请求路径重定向，/api/xxx => /xxx
          }
        }
      }
    }
  }