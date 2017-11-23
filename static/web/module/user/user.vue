<style scoped>
    .pj{
        width: 200px;
        height: 230px;
        overflow: auto;
    }
    .search{
        flex: 1;
        margin-left: 20px;
        height: 230px;
        overflow: auto;
    }
    .log{
       width: 100%;
    }
    .main{
        display: flex;
    }
    .addpj{
        margin-top: -6px;
    }
</style>

<template>
    <div>
        <mainHeader :header-title="title"></mainHeader>
        <div class="container">
            <div class="main">
                <section class="pj section">
                    <h4>
                        项目列表
                        <span class="addpj btn btn-link fr" data-toggle="modal" data-target="#addDlg">添加</span>
                    </h4>
                    <div class="content">
                        <ul>
                            <li v-for="p in ps">
                                <a :href="'/web/dist/project.html?p='+p.Id+'&u='+uid">{{p.Name}}</a>
                            </li>
                        </ul>
                    </div>
                </section>

                <section class="section search">
                    <h4>数据搜索</h4>
                    <div>
                        123
                    </div>
                </section>
            </div>
            <section class="section log">
                <h4>警报日志</h4>
            </section>
        </div>

        <!-- 添加项目的对话框 -->
        <div class="modal fade" id="addDlg" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title" id="myModalLabel">添加项目</h4>
                    </div>
                    <div class="modal-body">
                        <div class="form-group">
                            <label for="newPjName">请输入项目名称</label>
                            <input v-model="newPjName" class="form-control" id="newPjName" placeholder="project name">
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                        <button @click="onAddProject" type="button" class="btn btn-primary">Submit</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import mainHeader from '../../assert/vue-compoment/mainHeader.vue'
    var common = require("../../assert/js/common");
    export default {
        components: {
            mainHeader
        },
        methods:{
            onAddProject:function(){
                common.get("/addProject?pjname=" + this.newPjName+"&u="+this.uid).then(function(data){
                    common.showTip("添加成功");
                    $('#addDlg').modal('toggle');
                   setTimeout(function(){
                       window.location.reload();
                   },500);
                }).catch(function(resp){
                    common.handleErr(resp)
                });
            }
        },
        data(){
            return {
                title:"",
                ps:[],
                newPjName:"",
                uid:common.getQueryString("u"),
            }
        },
        created:function(){
            var thiz = this;
            common.get("/getUserData?u=" + this.uid).then(function(data){
                thiz.ps = data.Projects || [];
                thiz.title = data.PageTitle;
            }).catch(function(resp){
                common.handleErr(resp)
            });
        },
    };
</script>