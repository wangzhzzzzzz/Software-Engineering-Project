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
  },
  {
    path: '/studenthome',
    name: 'StudentHome',
    component:() => import('../views/StudentHome.vue'),
	children:[
		{
			path: '/studentinfo',
			name: 'studentinfo',
			component:() => import('../views/studenthome/StudentInfo.vue'),
			meta:{comp:"/studentinfo",pathname:"用户信息管理",name:"个人信息管理"}
		},
		{
			path: '/studentcourse',
			name: 'studentcourse',
			component:() => import('../views/studenthome/StudentCourse.vue'),
			meta:{comp:"/studentcourse",pathname:"选课信息管理",name:"选课信息管理"}
		},
	]
  },
  {
    path: '/teacherhome',
    name: 'TeacherHome',
    component:() => import('../views/TeacherHome.vue'),
	children:[
		{
			path: '/teacherinfo',
			name: 'teacherinfo',
			component:() => import('../views/teacherhome/TeacherInfo.vue'),
			meta:{comp:"/teacherinfo",pathname:"用户信息管理",name:"个人信息管理"}
		},
		{
			path: '/teachercourse',
			name: 'teachercourse',
			component:() => import('../views/teacherhome/TeacherCourse.vue'),
			meta:{comp:"/teachercourse",pathname:"选课信息管理",name:"选课信息管理"}
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
