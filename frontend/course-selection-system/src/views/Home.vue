<template>
  <div class="home">
	 <el-container :style="{height: 100 + 'vh' }">
		<el-menu background-color="#304156" text-color="#fff" active-text-color="#09F7F7" default-active="1"
			class="el-menu-vertical-demo" :collapse="isCollapse" unique-opened router :default-active="$route.path">
			<div>
				<div class="biaoti" v-if="check">
					<span style="font-size: 18px;color: white">学生选课管理系统</span>
				</div>
				<div class="biaoti" v-else>
					<span style="font-size: 18px;color: white">系统</span>
				</div>
			</div>
			<el-submenu v-for="item in menu" :index="item.id" :key="item.id">
				<template slot="title">
					<i :class="item.icon" style="color: #fff"></i>
					<span slot="title">{{item.name}}</span>
				</template>
				<el-menu-item-group v-for="(items,index) in item.submenu" :key="index">
					<el-menu-item :index="items.to" style="margin-top: -13px;">{{items.name}}</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
		</el-menu>
		<el-container>
			<el-header>
				<div class="icons" @click="toggleCollapse">
					<i class="el-icon-s-unfold" v-if="isCollapse"></i>
					<i class="el-icon-s-fold"></i>
				</div>
				<div class="mbx">
					<el-breadcrumb separator="/">
					  <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
					  <el-breadcrumb-item v-for="(item, indexs) in breadlist" :key="'info'+indexs">
						  {{item.pathname}}
					  </el-breadcrumb-item>
					  <el-breadcrumb-item v-for="(item, index) in breadlist" :key="index" :to="item.to">
					     {{item.name}}
					  </el-breadcrumb-item>
					</el-breadcrumb>
				</div>
				<div class="headeright">
					<span style="padding: 18px;">欢迎，<span style="color: blue;">{{user}}</span></span>
					<el-dropdown @command="setDialogInfo">
						<span class="el-dropdown-link">
						<span class="titles">用户设置</span>
						<i class="el-icon-arrow-down el-icon--right" style="color: #000000;font-size: 18px;"></i>
						</span>
						<el-dropdown-menu slot="dropdown">
							<el-dropdown-item command="info">个人信息</el-dropdown-item>
							<el-dropdown-item command="logout">退出登录</el-dropdown-item>
					   </el-dropdown-menu>
					</el-dropdown>
				</div>
			</el-header>
			<el-main>
				<Tags/>
				<div class="maincontent">
					<keep-alive>
						<router-view></router-view>
					</keep-alive>
				</div>
			</el-main>
		</el-container>
	</el-container>
  </div>
</template>
<script>
import menu from "../config/menu-config.js"
import Tags from "../components/tags/Tags"
export default {
	name: 'Home',
	components:{Tags},
	data() {
		return {
		  isCollapse: false,
		  user:"",
		  check:true,
		  menu:menu,
		  breadlist: []
		};
	},
	 watch: {
	 // 监听路由
		 $route(val) {
		  // 调用获取路由数组方法
		  this.getbreadlist(val);
		 }
	 },
	mounted() {
		let user=window.localStorage.getItem("user");
		this.user=user;
	},
	methods: {	
		 getbreadlist(val) {
			 let obj={
				 to:val.meta.comp,
				 name:val.meta.name,
				 pathname:val.meta.pathname
			 }
			 let matched=[]
			 matched.push(obj)
			 this.breadlist = matched;
		 },
		toggleCollapse() {
		  this.isCollapse = !this.isCollapse
		  if(this.isCollapse==false){
			  this.check=true;
		  }else{
			  this.check=false;
		  }
		},
		  // 下拉框选择，点击个人信息后调整到对应的函数执行
		      setDialogInfo(comItem) {
		        switch (comItem) {
		          case "info":
		            this.showInfoList();
		            break;
		          case "logout":
		            this.logout();
		            break;
		        }
		      },
		      //展示个人信息
		      showInfoList() {
		        this.$alert(
				  `<div>
						名称：${this.user}
						<br/>
						身份：管理员
				  </div>`,
		          '个人名片', {
		            dangerouslyUseHTMLString: true
		          });
		      },
			  logout(){
				 this.$confirm('确定退出登录吗?', '提示', {
					confirmButtonText: '确定',
					cancelButtonText: '取消',
					type: 'warning'
				  }).then(() => { 
					//点击确定的操作(调用接口)
					window.localStorage.clear();
					this.$router.push('/login')
				  }).catch(() => {
					return;
				  });
			  },
	}
}
</script>
<style scoped="scoped">
.titles{
    font-size: 20px;
    color: #000;
}
.mbx{
	padding: 20px;
	font-size: 20px;
}
.headeright{
	font-size: 20px;
	position: absolute;
	right: 15px;
}
.maincontent{
	overflow-y: scroll;
	height: calc(100vh-70px);
	padding: 15px;
}
::-webkit-scrollbar{
	display:none;
}
.biaoti{
	background-color: #42B983;
	width: 100%;
	height: 60px;
	line-height: 60px;
	text-align: center;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
	width: 200px;
	min-height: 400px;
}

.el-header{
	background-color: #FFFFFF;
	color: #333;
	line-height: 60px;
	display: flex;
	float: left;
	padding: 0 !important;
}
.el-header i {
	font-size: 25px;
}
.el-main {
	background-color: #F1F4F5;
	padding: 20px 4px;
}
.icons {
	width: 25px;
	height: 60px;
	padding: 0 20px;
	text-align: center
}
.icons:hover {
	background: rgb(248, 248, 248);
	transition: 1s;
	cursor: pointer;
}
</style>
