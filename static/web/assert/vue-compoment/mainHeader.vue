<style scoped>
    .header-wrapper{
        width: 100%;
        height: 55px;
        background-color: #fff;
        box-shadow: 0px 1px 5px #888888;
        line-height: 55px;
        text-align: center;
        font-weight: bold;
    }
</style>

<template>
    <div class="header-wrapper">
        <header class="container">
            {{headerTitle}}
        </header>
    </div>
</template>

<script>
    export default {
        props: ['headerTitle'],
    }
</script>