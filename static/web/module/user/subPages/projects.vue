<style>

</style>

<template>
    <div>
        项目列表
        <ul>
            <li v-for="p in ps">
                <a target="_blank" :href="'#?u=' + p.Id">{{p.Name}}</a>
            </li>
        </ul>
    </div>
</template>

<script>
    var common = require("../../../assert/js/common");
    export default {
        data(){
            return {
                ps:[]
            }
        },
        created:function(){
            var thiz = this;
            common.get("/getUserData?u=" + common.getQueryString("u")).then(function(data){
                thiz.ps = data.Projects;
            }).catch(function(resp){
                common.handleErr(resp)
            });
        },
    }
</script>