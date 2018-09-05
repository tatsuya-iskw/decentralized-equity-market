const webpack = require('webpack');
const WebpackDevServer = require('webpack-dev-server');
const webpackConfig = require('../webpack.config');
const port = process.env.PORT || 9999;

new WebpackDevServer(webpack(webpackConfig), {
  // publicPath: webpackConfig.output.publicPath,
  publicPath: '/',
  hot: true,
  contentBase: 'src/app',
  historyApiFallback: true
}).listen(port, 'localhost', (err, result) => {
  if (err) console.log(err);
  else {
    console.log('Webpack Dev Server listening at localhost:${port}');
  }
});
