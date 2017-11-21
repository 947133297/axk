function post(url,form){
    return new Promise(function(resolve,reject){
        Vue.http.post(url, form).then(function(resp){
            console.log(resp)
            if(resp.body.code == 0){
                resolve(resp.body.msg);
            }else{
                reject(resp.body.msg);
            }
        });
    });
}

function get(url,form){
    return new Promise(function(resolve,reject){
        Vue.http.get(url, form).then(function(resp){
            if(resp.body.code == 0){
                resolve(resp.body.msg);
            }else{
                reject(resp.body.msg);
            }
        });
    });
}

module.exports = {
    post,
    get
};