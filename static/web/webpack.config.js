const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin')
module.exports = {
    entry:  {
        index: __dirname + "/module/index/index.js"
    },
    output: {
        path: __dirname + "/dist",
        filename: "js/bundle/[name]/[hash].js",
    },
    plugins: [
        new CleanWebpackPlugin(['dist']),
        new HtmlWebpackPlugin({
            template: __dirname + "/assert/template/tmp.html",
            chunks: ['index'],
            title:'首页'
        })
    ],
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            }
        ]
    },
    resolve: {
        extensions: [".vue",".js",".json"],
    },
}