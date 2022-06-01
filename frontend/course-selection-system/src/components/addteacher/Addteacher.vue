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
              <el-form-item label="是否注销：" prop="IsDeleted">
                 <el-radio-group  v-model="IsDeleted">
                    <el-radio-button  v-for="item in tanchutype" :key="item.label" :label="item.value"></el-radio-button>
                </el-radio-group>
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
    name:"Addstudent",
    // 获取父亲传过来的值
    props:{
	    dialog:Object
	},
    data(){
        return{
            dialogVisible: false,
            checkedtanchu:"",
            title:"",
            content:"",
            tanchucount:"",
            tanchutype:[
                {
                    label:1,
                    value:'已注销'
                },{
                    label:2,
                    value:'未注销'
                }
           ],
           rule:{
               checkedtanchu:[{require:true,trigger:'blur'}],
               title:[{require:true,message:"请输入标题",trigger:"blur"}]
           }
        }
    },
     methods: {
        handleAddAndEdit(){
             //todo
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