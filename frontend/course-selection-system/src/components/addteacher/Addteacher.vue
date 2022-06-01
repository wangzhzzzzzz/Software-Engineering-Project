<template>
    <div class="dialog">
        <el-dialog 
          :title="dialog.title"
		  :close-on-click-modal="false"
		  :close-on-press-escape="false"
		  :modal-append-to-body="false"
		  :visible.sync="dialog.show">
            <el-form label-width="100px" :rules="rule">
			  <el-form-item label="姓名" prop="Name">
			    <el-input v-model="Name"></el-input>
			  </el-form-item>
              
			  <el-form-item label="教工号" prop="NetID">
			    <el-input v-model="NetID"></el-input>
			  </el-form-item>

              <el-form-item label="密码" prop="password">
			    <el-input v-model="password"></el-input>
			  </el-form-item>
              <!-- <el-form-item label="是否注销：" prop="IsDeleted">
                 <el-radio-group  v-model="IsDeleted">
                    <el-radio-button  v-for="item in tanchutype" :key="item.label" :label="item.value"></el-radio-button>
                </el-radio-group>
            </el-form-item> -->
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
    name:"Addstudent",
    // 获取父亲传过来的值
    props:{
	    dialog:Object
	},
    data(){
        return{
            Name:"",
            NetID:"",
            password:"",
            IsDeleted:"",
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
            if(this.Name.length < 2 || this.Name.length > 5){
                alert("姓名长度不符合要求");
            }
            if(this.NetID.length < 8 || this.NetID.length > 20){
                alert("教工号长度不符合要求");
            }
            if(this.password < 8 || this.password > 20){
                alert("密码长度不符合要求");
            }
            const api = '/api/v1/member/create';
            let params = 'nickname=' + this.Name +'&username=' + this.NetID + '&password=' + this.password + '&user_type=3';
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
            
        },
        handleClose() {
            this.dialog.show=false
        }
    },
   
}
</script>
<style scoped>

</style>