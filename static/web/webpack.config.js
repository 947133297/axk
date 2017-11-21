const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const CommonsChunkPlugin = require("webpack/lib/optimize/CommonsChunkPlugin");
module.exports = {
    entry:  {
        index: __dirname + "/module/index/index.js"
    },
    output: {
        path: __dirname + "/dist",
        filename: "js/bundle/[name]/[hash].js",
    },
    devtool: 'eval-source-map',
    plugins: [
        new CleanWebpackPlugin(['dist']),

        new HtmlWebpackPlugin({
            template: __dirname + "/assert/template/tmp.html",
            chunks: ['index',"commonChunk"],
            title:'首页'
        }),
        new CommonsChunkPlugin('commonChunk'),
    ],
    module: {
        rules: [
            { test: /\.vue$/, loader: 'vue-loader' },
            { test: /\.s?css$/, loader: 'style-loader!css-loader!sass-loader' },
        ]
    },
    resolve: {
        extensions: [".vue",".js",".json"],
    },
}