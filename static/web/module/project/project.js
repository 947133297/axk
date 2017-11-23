import project from './project';

require("../../assert/css/common.css");
new Vue({
    template:"<project></project>",
    el:"#root",
    components: { project }
})