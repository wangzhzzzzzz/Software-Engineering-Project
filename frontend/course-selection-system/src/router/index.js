import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
	redirect:"/login"
  },	
  {
    path: '/login',
  	component:() => import('../views/Login.vue'),
  },	
  {
    path: '/home',
    name: 'Home',
    component:() => import('../views/Home.vue'),
	children:[
		{
			path: '/student',
			name: 'student',
			component:() => import('../views/home/Student.vue'),
			meta:{comp:"/student",pathname:"用户信息管理",name:"学生信息管理"}
		},
		{
			path: '/teacher',
			name: 'teacher',
			component:() => import('../views/home/Teacher.vue'),
			meta:{comp:"/teacher",pathname:"用户信息管理",name:"教师信息管理"}
		},
		{
			path: '/course',
			name: 'course',
			component:() => import('../views/home/Course.vue'),
			meta:{comp:"/course",pathname:"课程信息管理",name:"课程信息管理"}
		},
	]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
const VueRouterPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (to) {
  return VueRouterPush.call(this, to).catch(err => err)
}
export default router
