<style scoped>
    .list{
        width: 200px;
        height: 230px;
        overflow: auto;
    }
</style>

<template>
    <div>
        <mainHeader :header-title="title"></mainHeader>
        <div class="container">
            <section class="list section">
                <h4>
                    检测区域
                    <span class="btn btn-link fr" data-toggle="modal" data-target="#addDlg">添加</span>
                </h4>
                <div class="content">
                    {{sectionList}}
                    <!--<ul>-->
                        <!--<li v-for="p in ps">-->
                            <!--<a :href="'/web/dist/project.html?p='+p.Id+'&u='+uid">{{p.Name}}</a>-->
                        <!--</li>-->
                    <!--</ul>-->
                </div>
            </section>
        </div>

        <div class="modal fade" id="addDlg" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title" id="myModalLabel">添加检测区域</h4>
                    </div>
                    <form v-on:submit.prevent="onAddSection">
                        <div class="modal-body">
                            <div class="form-group">
                                <label for="sectionName">请输入检测区域名称</label>
                                <input v-model="addSectionName" class="form-control" id="sectionName" placeholder="section name">
                            </div>
                            <div class="form-group">
                                <label for="sectionFile">上传施工图</label>
                                <input v-on:change="fileChanged"  id="sectionFile" type="file">
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                            <button   type="submit" class="btn btn-primary">Submit</button>
                        </div>
                    </form>
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
            fileChanged:function(evt){
                var files = evt.target.files;
                if (!files || files.length === 0) {
                    return;
                }
                this.file = files[0];
            },
            onAddSection:function(){
                var formData = new FormData();
                formData.append('file', this.file);
                formData.append('Name', this.addSectionName);
                common.post('/addWatchSection?p='+this.pid + "&u="+this.uid,formData)
                    .then(function(data){
                        console.log(data);
                        common.showTip("添加成功 ");
                        $('#addDlg').modal('toggle');
                        setTimeout(function(){
                            window.location.reload();
                        },500);
                    })
                    .catch(function(data){
                        common.showTip(data.msg)
                    });
            }
        },
        data(){
            return {
                title:"",
                addSectionName:"",
                uid:common.getQueryString("u"),
                pid:common.getQueryString("p"),
                sectionList:[]
            }
        },
        created:function(){
            var thiz = this;
            common.get("/getProjectData?u=" + this.uid + "&p="+this.pid).then(function(data){
                thiz.title = "项目主页：" + data.Name;
                thiz.sectionList = data.SectionList;
                // 这个接口要返回区域列表
                console.log(data)
            }).catch(function(resp){
                common.handleErr(resp)
            });
        },
    };

</script>