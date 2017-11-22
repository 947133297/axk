function post(url,form){
    return new Promise(function(resolve,reject){
        Vue.http.post(url, form).then(function(resp){
            console.log(resp)
            if(resp.body.code == 0){
                resolve(resp.body);
            }else{
                reject(resp.body);
            }
        });
    });
}

function get(url,form){
    return new Promise(function(resolve,reject){
        Vue.http.get(url, form).then(function(resp){
            if(resp.body.code == 0){
                resolve(resp.body);
            }else{
                reject(resp.body);
            }
        });
    });
}

function handleErr(body){
    if(body.code == "2"){
        window.location.href = "/web/dist/"
    }else if(body.code == "3"){
        window.location.href = "/static/no-auth.html"
    }
}
function showTip(msg){
    alert(msg)
}
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]); return null;
}
module.exports = {
    post,
    get,
    handleErr,
    getQueryString,
    showTip
};