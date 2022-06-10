<template>
    <div class="dialog">
        <el-dialog 
          :title="dialog.title"
		  :close-on-click-modal="false"
		  :close-on-press-escape="false"
		  :modal-append-to-body="false"
		  :visible.sync="dialog.show">
            <el-form label-width="100px" :rules="rule">
			  <el-form-item label="课程名称" prop="CourseName">
			    <el-input v-model="CourseName"></el-input>
			  </el-form-item>
              <el-form-item label="教工号" prop="TeacherID">
			    <el-input v-model="TeacherID"></el-input>
			  </el-form-item>
              <el-form-item label="教师名称" prop="TeacherName">
			    <el-input v-model="TeacherName"></el-input>
			  </el-form-item>
			  <el-form-item label="人数限制" prop="Capacity">
			    <el-input v-model="Capacity"></el-input>
			  </el-form-item>
		  </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="handleClose()">取 消</el-button>
                <el-button type="primary" @click="handleAddAndEdit()">创建</el-button>
            </span>
        </el-dialog>
    </div>
</template>
<script>
export default {
    name:"Addcourse",
    // 获取父亲传过来的值
    props:{
	    dialog:Object
	},
    data(){
        return{
            CourseName: "",
            TeacherID: "",
            TeacherName: "",
            Capacity: "",
            dialogVisible: false,
            checkedtanchu:"",
            title:"",
            content:"",
            tanchucount:"",
           rule:{
               checkedtanchu:[{require:true,trigger:'blur'}],
               title:[{require:true,message:"请输入标题",trigger:"blur"}]
           }
        }
    },
     methods: {
        handleAddAndEdit(){
            //todo
            if(this.TeacherName.length < 2 || this.TeacherName.length > 5){
                alert("姓名长度不符合要求");
                this.dialog.show=false;
                return ;
            }
            if(this.TeacherID.length < 8 || this.TeacherID.length > 20){
                alert("教工号长度不符合要求");
                this.dialog.show=false;
                return;
            }
            const api = '/api/v1/course/add';
            let params = 'CourseName=' + this.CourseName +'&NetID=' + this.TeacherID + '&TeacherName=' + this.TeacherName + '&Capacity=' + this.Capacity;
            console.log(params);
            this.axios.post(api, 
                            params,
                            {headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'}}
            ).then((res) => {
                this.dialog.show=false;
                location.reload();
            }).catch((err) => {
                console.log('err:', err.response.data.msg);
            });
            console.log('addStudent');
            this.dialog.show=false
        },
        handleClose() {
            this.dialog.show=false
        }
    },
   
}
</script>
<style scoped>

</style>