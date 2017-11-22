import mgr from './mgr';
require("../../assert/css/common.css");
new Vue({
    template:"<mgr></mgr>",
    el:"#root",
    components: { mgr }
})