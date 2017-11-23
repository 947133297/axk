<style scoped>
    .list{
        width: 200px;
        height: 230px;
        overflow: auto;
    }
    .graph-wrapper{
        position: relative;
        max-width: 800px;
    }
    .point{
        position: absolute;
        cursor: pointer;
    }
    .menu{
        position: absolute;
        border: 1px solid #ccc;
    }
    .menu-item{
        cursor: pointer;
        padding: 5px;
        background-color: #fff;
    }
    .menu-item:hover{
        background-color: #cccccc;
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
                    <ul>
                        <li v-for="p in ps">
                            <a :href="'/web/dist/project.html?p='+p.Id+'&u='+uid">{{p.Name}}</a>
                        </li>
                    </ul>
                </div>
            </section>
        </div>

        <div class="modal fade" id="addDlg" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
            <div class="modal-dialog modal-lg" role="document">
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
                            <div class="loading" v-if="showLoading">
                                上传中...
                            </div>
                            <div class="form-group" v-if="showGraph">
                                <div class="graph-wrapper" @click="onGraphWrapperClick">
                                    <img style="width: 100%" :src="graph" alt="施工图">
                                    <img v-on:click.stop="clickPoint(p,$event)" class="point" v-for="p in points" :src="'/static/img/'+(p.normal?'green':'red')+'_pos.png'" :style="'top:'+(p.y-57)+'px;left:'+(p.x-32)+'px'">
                                    <div v-if="showMenu" class="menu" :style="'top:'+(menu.y)+'px;left:'+(menu.x)+'px'">
                                        <div class="menu-item" @click.stop="onMenuDelete">删除</div>
                                        <div class="menu-item" @click.stop="onMenuConfig">配置</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                            <button type="submit" class="btn btn-primary">Submit</button>
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
            fileChanged,
            onGraphWrapperClick,
            onAddSection,
            clickPoint,
            onMenuDelete,
            onMenuConfig,
        },
        data(){
            return {
                title:"",
                addSectionName:"",
                uid:common.getQueryString("u"),
                pid:common.getQueryString("p"),
                sectionList:[],
                graph:"",
                showLoading:false,
                showGraph:false,
                points:[],
                showMenu:false,
                menu:{x:0,y:0}
            }
        },
        created:function(){
            var thiz = this;
            forTest(thiz);
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

    function forTest(vur) {
        vur.graph = "http://localhost:12306/static/upload/65aee94e824c50fdb8163fc58b03149e..png"
        vur.showGraph = true;
    }

    function onGraphWrapperClick(evt){
        addPoint(this,evt.offsetX,evt.offsetY)
    }
    function onAddSection(){
        var formData = new FormData();
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
    function fileChanged(){
        var thiz = this;
        var files = evt.target.files;
        if (!files || files.length === 0) {
            return;
        }
        var formData = new FormData();
        formData.append('file', files[0]);
        thiz.showLoading = true;
        common.post('/upload?p='+this.pid + "&u="+this.uid,formData)
            .then(function(data){
                thiz.graph = '/static/upload/' + data.msg;
                thiz.showGraph = true;
                thiz.showLoading = false;
            })
            .catch(function(data){
                common.showTip(data.msg)
            });
    }
    function addPoint(vur,x,y){
        if(vur.showMenu){
            vur.showMenu = false;
            return ;
        }
        var point = new Point(x,y,vur);
        vur.points.push(point);
    }
    function clickPoint(point,evt){
        curPoint = point;
        this.showMenu = true;
        this.menu.x = evt.offsetX + point.x - 32;
        this.menu.y = evt.offsetY + point.y - 57;
    }
    var curPoint;
    function Point(x,y,vur){
        this.x = x;
        this.y = y;
        this.normal = true;
        this.vueInstance = vur;
        this.index = vur.points.length;
        this.delete = function(){
            vur.points.splice(this.index,1);
            for(let i = this.index;i<vur.points.length;i++){
                vur.points[i].index --;
            }
        }
    }

    function onMenuDelete(){
        curPoint.delete();
        this.showMenu = false;
    }
    function onMenuConfig(){
        alert("config")
    }
</script>