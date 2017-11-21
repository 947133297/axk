import index from './index';
require("../../assert/css/common.css");

new Vue({
    template:"<index></index>",
    el:"#root",
    components: { index }
})