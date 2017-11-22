import user from './user';

require("../../assert/css/common.css");
new Vue({
    template:"<user></user>",
    el:"#root",
    components: { user }
})