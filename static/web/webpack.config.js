const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require('webpack');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const CommonsChunkPlugin = require("webpack/lib/optimize/CommonsChunkPlugin");
var config = {
    output: {
        path: __dirname + "/dist",
        filename: "js/bundle/[name]/[chunkhash].js",
    },
    // devtool: 'eval-source-map',
    plugins: [
        new CleanWebpackPlugin(['dist']),
        new CommonsChunkPlugin('commonChunk'),
        new webpack.optimize.UglifyJsPlugin()
    ],
    module: {
        rules: [
            { test: /\.vue$/, loader: 'vue-loader' },
            { test: /\.js/, loader: 'babel-loader' , exclude:  /node_modules/},
            { test: /\.s?css$/, loader: 'style-loader!css-loader!sass-loader' },
        ]
    },
    resolve: {
        extensions: [".vue",".js",".json"],
    },
};
// 如果没有使用babel-loader，则在ie11中运行会报错

var modules = [
    ['index'],
    ['mgr'],
    ['user'],
    ['project'],
];
function getEntrys(){
    var o = {};
    modules.forEach(function(ele){
        o[ele[0]] = __dirname + "/module/"+ele[0]+"/"+ele[0]+".js"
    });
    return o;
}


function HtmlWebpackPlugins(){
    return modules.map(function(arr){
        return new HtmlWebpackPlugin({
            template: __dirname + "/assert/template/tmp.html",
            chunks: ['commonChunk',...arr],
            filename:arr[0]+".html"
        })
    });
}

config.entry = getEntrys();
config.plugins.push(...(HtmlWebpackPlugins()));

module.exports = config;