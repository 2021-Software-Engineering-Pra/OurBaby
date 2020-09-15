const path=require('path')
const VueloaderPlugin=require('vue-loader/lib/plugin')
const MiniCssExtractPlugin = require("mini-css-extract-plugin")
module.exports={
    mode:'development',
    // mode:'production',
    entry:'./src/index.js',
    output:{
        // path:path.resolve(__dirname,'dist'),
        path:"D:\\code\\web\\shop\\web",
        filename:'main.js'
    },
    resolve:{
        alias:{
            'vue$':'vue/dist/vue.min.js'
        }
    },
    module:{
        rules:[
          
            {test:/\.css$/,use: [
                "style-loader",      
                "css-loader"
            ],},
            {test:/\.jpg|png|bmp|svg|gif$/,use:[{
                loader:'url-loader',
                options:{
                    limit:'5000',
                    name:'[name].[ext]',
                    outputPath:'./img/',
                    publicPath:'./img/'
                }
            }]},
            {test:/\.ttf|eot|woff|woff2$/,use:[{
                loader:'file-loader',
                options:{
                    name:'[name].[ext]',
                    outputPath:'./font/',
                    publicPath:'./font/'
                    }
            }],},
            {test:/\.vue$/,loader:'vue-loader'}
        ]
    },
    plugins:[
        new VueloaderPlugin(),
        new MiniCssExtractPlugin({
            　　filename: "./css/[name].css"
         　　 })

    ],
    devServer:{
        proxy:{
            '/api':{
                target:'http://localhost:9001',
                changeOrigin:true
            }
        },
        publicPath:"/",
    }
}