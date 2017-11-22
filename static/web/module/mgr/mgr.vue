<style>
</style>

<template>
    <div>
        <mainHeader :header-title="title"></mainHeader>
        <div class="container">
            <div> 用户列表 </div>
            <ul>
                <li v-for="u in uList">
                    <a :href="'/web/dist/user.html?account=' + u">{{u}}</a>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
    import mainHeader from '../../assert/vue-compoment/mainHeader.vue'
    var http = require("../../assert/js/common");

    export default {
        components: {
            mainHeader
        },
        created:function(){
            var thiz = this;
            http.get("/getMgrData").then(function(data){
                thiz.title = data.PageTitle;
                thiz.uList = data.UserList;
            }).catch(function(resp){
                http.handleErr(resp)
            });
        },
        data(){
            return {
                title:"",
                uList:[]
            }
        }
    }
</script>