<style scoped>
    .container{
        border: 1px solid #ccc;
        width: 400px;
        box-shadow: 10px 10px 5px #888888;
        background: #fff;
        margin-top: 50px;
        padding: 20px 20px;
        border-radius: 5px;
        display: inline-block;
        text-align: left;
    }
    .wrapper{
        height: 100%;
        width: 100%;
        position: fixed;
        display:-webkit-box;
        -webkit-box-align:center;
        -webkit-box-pack:center;
        text-align: center;
    }
    .vk{
        cursor: pointer;
        border:1px solid #ccc;
        height: 35px;
    }
    .vk-input{
        width: 100px;
        display: inline-block;
    }
</style>
<template>
    <div class="wrapper">
        <div class="container">
            <form action="" v-if="loginPage" v-on:submit.prevent="login">
                <div class="form-group">
                    <label for="loginAccount">账号</label>
                    <input v-model="Account" class="form-control" id="loginAccount" placeholder="Account" maxlength="20">
                </div>
                <div class="form-group">
                    <label for="loginPwd">密码</label>
                    <input v-model="Pwd" type="password" class="form-control" id="loginPwd" placeholder="Password" maxlength="20">
                </div>
                <div class="checkbox fr">
                    <label>
                        <input type="checkbox"> 记住我
                    </label>
                </div>
                <div class="clearfix"></div>
                <button type="submit" class="btn btn-default fr">登录</button>
                <button  type="button" class="btn btn-link fr" @click="change()">去注册</button>
            </form>
            <form v-else>
                <div class="form-group">
                    <label for="registAccount">注册账号</label>
                    <input v-model="RegAccount"  class="form-control" id="registAccount" placeholder="Account" maxlength="20">
                </div>
                <div class="form-group">
                    <label for="registPwd">密码</label>
                    <input v-model="RegPwd" type="password" class="form-control" id="registPwd" placeholder="Password" maxlength="20">
                </div>
                <div class="form-group">
                    <label for="repPwd">确认密码</label>
                    <input v-model="RPwd" type="password" class="form-control" id="repPwd" placeholder="retpype Password" maxlength="20">
                </div>
                <div class="form-group">
                    <label for="repPwd">验证码</label>
                    <br>
                    <input type="text" class="form-control vk-input" name="code" v-model="code">
                    <img :src="vk" alt="验证码" @click="changeVk()" class="vk">
                </div>
                <div class="clearfix"></div>
                <button type="button" class="btn btn-default fr" @click="regist()">注册</button>
                <button  type="button" class="btn btn-link fr" @click="change()">去登录</button>
            </form>
        </div>
    </div>
</template>

<script>
    var http = require("../../assert/js/common");
    export default {
        methods: {
            change: function () {
                this.loginPage=!this.loginPage;
            },
            regist,
            login,
            changeVk
        },
        created: function () {
            this.changeVk();
        },
        data () {
            return {
                loginPage:true,
                Account:"",
                Pwd:"",
                RegAccount:"",
                RegPwd:"",
                RPwd:"",
                vk:"",
                code:""
            }
        }
    }

    function regist(){
        if(this.RegAccount.length === 0){
            showTip("注册账号不能为空");
            return ;
        }
        if(this.RegPwd.length === 0){
            showTip("密码不能为空");
            return ;
        }
        if(this.RPwd !== this.RegPwd){
            showTip("两次输入密码不一致")
            return ;
        }
        if(this.code.length === 0){
            showTip("请填验证码");
            return ;
        }
        var form = {
            Account:this.RegAccount,
            Pwd:this.RegPwd,
            Code:this.code.toLowerCase()
        };
        var v = this;
        http.post('/register',form)
            .then(function(){
                showTip("注册成功 ");
                v.change();
                v.code = v.RegPwd = v.RPwd = v.RegAccount= ""
            })
            .catch(function(data){
                showTip(data.msg)
            });
    }

    function login(){
        if(this.Account.length === 0){
            showTip("登录账号不能为空");
            return ;
        }
        if(this.Pwd.length === 0){
            showTip("密码不能为空");
            return ;
        }
        var form = {
            Account:this.Account,
            Pwd:this.Pwd
        };
        http.post('/login',form)
            .then(function(data){
                if(data.msg == "user"){
                    window.location.href = "/web/dist/user.html"
                }else {
                    window.location.href = "/web/dist/mgr.html"
                }
            })
            .catch(function(data){
                showTip("密码错误或账号不存在");
                console.log(data.msg)
            });
    }
    function showTip(msg){
        alert(msg)
    }

    function changeVk(){
        this.vk = "/vk?r=" + Math.floor(Math.random() * ( 1000 + 1))
    }
</script>