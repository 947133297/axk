<style scoped>
    .nav{
        display: inline-block;
        vertical-align: top;
    }
    .nav a{
        display: block;
        background-color: transparent;
        padding: 5px;
        border-right:1px solid #ccc;
        text-decoration: none;
    }
    .nav a.router-link-active{
        border:1px solid #ccc;
        background-color: #fff;
    }
    .content{
        border:  1px solid #ccc;
        border-left : none;
        flex: 1;
        display: inline-block;
        background-color: #fff;
        padding: 5px;
    }
    .wrapper{
        margin-top: 10px;
        display: flex;
    }
</style>

<template>
    <div>
        <mainHeader :header-title="title"></mainHeader>
        <div class="container wrapper">
            <nav class="nav">
                <router-link to="/ax">项目列表</router-link>
                <router-link to="/axx">警报日志</router-link>
            </nav>
            <section class="content">
                <router-view></router-view>
            </section>
        </div>
    </div>
</template>

<script>
    import mainHeader from '../../assert/vue-compoment/mainHeader.vue'
    import ax from './subPages/ax.vue'
    import VueRouter from 'vue-router'
    Vue.use(VueRouter);
    var http = require("../../assert/js/common");

    const routes = [
        { path: '/ax', component: ax },
        { path: '/axx', component: ax },
    ];

    export default {
        router:new VueRouter({
            routes
        }),
        components: {
            mainHeader
        },
        created:function(){
//            var thiz = this;
//            http.get("/getMgrData").then(function(data){
//                thiz.title = data.PageTitle;
//                thiz.uList = data.UserList;
//            }).catch(function(resp){
//                http.handleErr(resp)
//            });
        },
        data(){
            return {
                title:"用户页面",
            }
        }
    };
</script>