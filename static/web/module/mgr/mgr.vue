<style>
    .main{
        width: 200px;
        overflow: auto;
        display: inline-block;
        border: 1px solid #ccc;
        background-color: #fff;
        margin-top: 10px;
        padding: 5px;
        border-radius: 5px;
    }
</style>

<template>
    <div>
        <mainHeader :header-title="title"></mainHeader>
        <div class="container">
           <div class="main">
               <h4> 用户列表 </h4>
               <ul>
                   <li v-for="u in uList">
                       <a target="_blank" :href="'/web/dist/user.html?u=' + u.Id">{{u.Account}}</a>
                   </li>
               </ul>
           </div>
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