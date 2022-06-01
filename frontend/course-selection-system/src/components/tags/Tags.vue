<template>
  <div class="tags-wrap">
    <div class="tags" v-if="showTags">
        <ul class="tags-ul">
            <li class="tags-li" v-for="(item,index) in tagsList" :class="{'active': isActive(item.path)}" :key="index">
                <router-link :to="item.path" class="tags-li-title">
                    {{item.name}}
                </router-link>
                <span class="tags-li-icon" @click="closeTags(index)"><i class="el-icon-close"></i></span>
            </li>
        </ul>
        <div class="tags-close-box">
            <el-dropdown @command="handleTags">
                <el-button type="primary">
                    标签选项<i class="el-icon-arrow-down el-icon--right"></i>
                </el-button>
                <el-dropdown-menu size="small" slot="dropdown">
                    <el-dropdown-item command="other">关闭其他</el-dropdown-item>
                    <el-dropdown-item command="all">关闭所有</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
  </div>
</template>
<script>
    import bus from '../../config/bus';
    export default {
        name:'Tags',
        data() {
            return {
                tagsList: []
            }
        },
        methods: {
            isActive(path) {
                return path == this.$route.fullPath;
            },
            // 关闭单个标签
            closeTags(index) {
                const delItem = this.tagsList.splice(index, 1)[0];
                const item = this.tagsList[index] ? this.tagsList[index] : this.tagsList[index - 1];
                if (item) {
                    delItem.path ==this.$route.fullPath && this.$router.push(item.path);
                }else{
                    this.$router.push('/info');
                }
            },
            // 关闭全部标签
            closeAll(){
                this.tagsList = [];
                this.$router.push('/info');
            },
            // 关闭其他标签
            closeOther(){
                const curItem = this.tagsList.filter(item => {
                    return item.path ==this.$route.fullPath;
                })
                this.tagsList = curItem;
            },
            // 设置标签
            setTags(route){
                const isExist = this.tagsList.some(item => {
                    return item.path == route.fullPath;
                })
                !isExist && this.tagsList.push({
                    name: route.meta.name,
                    path: route.fullPath
                })
                bus.$emit('tags', this.tagsList);
            },
            handleTags(command){
                command == 'other' ? this.closeOther() : this.closeAll();
                command == 'all' ? this.closeAll() : this.closeOther();
            }
        },
        computed: {
            showTags() {
                return this.tagsList.length > 0;
            }
        },
        watch:{
            $route(newValue, oldValue){
                this.setTags(newValue);
            }
        },
        mounted(){
            this.setTags(this.$route);
        }
    }

</script>
<style>
.tags{
	width: 100%;
    position:relative;
    top: -18px;
    height:40px;
    line-height:40px;
    overflow:hidden;
	background-color: #FFFFFF;
    box-shadow:0 3px 3px #ddd;
}
.tags-ul{
    width:80%;
    display: flex;
    height:100%;
    margin-left: -36px;
}
.tags-li{
    float: left;
    display: flex;
    border-radius: 3px;
    font-size: 0.2rem;
    margin: 1px;
    cursor: pointer;
    height: 40px;
    line-height: 40px;
    border: 1px solid #e9eaec;
    background: #fff;
    padding: 0 8px 0 8px;  
    color: #666;
    margin-top:-18px;
	margin-left: -4px;
}
.tags-li.active{
    background:#4098ff;
    color:#fff;
}
.tags-li.active .tags-li-title{
    color:#fff;
}
.tags-li-title{
    text-decoration: none;
    display:inline-block;
    color:#000;
    vertical-align:middle;
    font-size:12px;
}
.tags-li-icon{
    display:inline-block;
    vertical-align:middle;
    font-size:0.5rem;
    margin-left:3px;
}
.tags-close-box{
    position:absolute;
    right:28px;
    top:0px;
    width:90px;
    padding-left:12px;
    box-shadow:0 5px 20px #ccc;
}

</style>
