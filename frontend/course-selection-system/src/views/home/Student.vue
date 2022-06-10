<template>
  <div class="notice">
	  <el-form  label-width="80px">
	    <el-form-item label="查询内容:">
	      <el-input v-model="name"></el-input>
	    </el-form-item>
	  </el-form>
	  <el-button type="primary" @click="search()" style="margin-left: 15px;">查询</el-button>
	   <el-button type="success" @click="AddNotice()" icon="el-icon-plus" style="position: absolute;right: 15px;">添加学生信息</el-button>
      <el-table border :fit="true" 
        :data="tableData.slice((currentPage - 1) * pageSize, currentPage * pageSize)">
        <el-table-column prop="UserID" sortable label="UserID">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.UserID }}</span>
            <el-input
              type="text"
              v-model="row.UserID"
              v-else
              placeholder="请输入UserID"
            ></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="Name" sortable label="姓名">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.Name }}</span>
            <el-input
              type="text"
              v-model="row.Name"
              v-else
              placeholder="请输入姓名"
            ></el-input>
          </template>
        </el-table-column>
        
        <el-table-column prop="NetID" sortable label="学号">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.NetID }}</span>
            <el-input
              type="text"
              v-model="row.NetID"
              v-else
              placeholder="请输入学号"
            ></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="StartTime" sortable label="创建时间">
           <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">
              {{row.StartTime}}
            </span>
            <el-date-picker
              type="datetime"
              value-format="yyyy-MM-dd HH:mm:ss"
              v-model="row.StartTime"
              v-else
              placeholder="请输入生效开始时间"
            >
            </el-date-picker>
          </template>
        </el-table-column>
        <!-- <el-table-column prop="IsDeleted" sortable label="是否注销">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">
              {{row.IsDeleted}}
            </span>
            <el-date-picker
              type="datetime"
              value-format="yyyy-MM-dd HH:mm:ss"
              v-model="row.IsDeleted"
              v-else
              placeholder="是否已注销"
            >
            </el-date-picker>
          </template>
        </el-table-column> -->
          <el-table-column
          header-align="center"
          align="center"
           width="200"
          label="操作"
        >
          <template slot-scope="{ row, $index }">
            <el-button
              v-if="!showEdit[$index]"
              @click="showUpdate($index, row)"
              type="primary"
            >
              <i class="el-icon-edit"></i>修改
            </el-button>
            <el-button
              v-if="!showEdit[$index]"
              @click="del($index, row)"
              type="danger"
            >
              <i class="el-icon-delete"></i>删除
            </el-button>
            <el-button
              v-if="showEdit[$index]"
              @click="submit($index, row)"
              type="success"
              >确定</el-button
            >
            <el-button
              v-if="showEdit[$index]"
              @click="cancelUpdate($index)"
              type="warning"
              >取消</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <pageNation :total="tableData.length" :page-size="pageSize" @pageChange="pageChange"  :page_index="currentPage" />
    <Addstudent :dialog="dialog" @update="getProfile($event)" />
  </div>
</template>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
<script>
import Addstudent from "../../components/addstudent/Addstudent.vue";
import pageNation from '../../components/pagenation/pageNation';
export default {
  name: "Notice",
  components: { Addstudent,pageNation }, 
  data() {
    return {
	  name:"",
    // Name:"",
    // NetID:"",
    // IsDeleted:"",
      tableData: [
        // {
        //   UserID:1001,
        //   Name:"吴浩岚",
        //   NetID:"19335209",
        //   StartTime:"2022-3-01 8:12:35",
        //   IsDeleted:"未注销"
        // },
        // {
        //   UserID:1002,
        //   Name:"李凌一",
        //   NetID:"19335102",
        //   StartTime:"2022-3-01 8:12:35",
        //   IsDeleted:"已注销"
        // },
        
      ],
      showEdit: [], //控制页面修改显示隐藏
      currentPage: 1, //当前页数
      pageSize: 20, //每页的数据
      dialog: {
        show: false, //是否显示弹出框
        title: "", //弹出框的title
      },
    };
  },
  created(){
      this.initData();
  },
  methods: {
    async initData(){
      const api = '/api/v1/student/list';
      this.axios.get(api,{
        params:{
          Offset:this.currentPage,
          Limit:this.pageSize,
        } 
    }).then((response)=>{
            console.log(response);
            if(response.status === 200){
                //获得成功响应返回的数据
                console.log(response.data.Data.MemberList);
                // let arr = ;
                // this.tableData.length = 0;
                // response.data.Data.MemberList.forEach(function(item){
                //   this.tableData.push(item);
                // });
                this.tableData = response.data.Data.MemberList;
                console.log(this.tableData);
                // for(int i = 0; i < response.data.Data.MemberList.length; i++){

                // }
            }
        },(error)=>{
            console.log('err:', err.response.data.msg);
        });
    },
    //分页
    pageChange (item) {
	      this.currentPage = item.page_index;
	      this.pageSize = item.page_limit;
    },
    AddNotice() {
      this.dialog = {
        show: true,
        title: "添加学生信息", 
      };
    },
    getProfile(data) {
      console.log(data)
    },
     //编辑后点击确定按钮
    submit(index, row) {-
      console.log("update confirm");
      console.log(index, row);

      //调用post方法
      const api = '/api/v1/member/update';
      let params = 'userID=' + row.UserID +'&name=' + row.Name + '&NetID=' + row.NetID;
      console.log(params);
      this.axios.post(api, 
                      params,
                      {headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'}}
      ).then((res) => {
          // this.dialog.show=false;
          // location.reload();
          console.log(res);
      }).catch((err) => {
          console.log('err:', err.response.data.msg);
      });
      console.log('addStudent');


      this.$set(this.showEdit, index, false);
    },
    //点击删除
    del(index, row) {
      console.log(index, row);
      //调用post方法
      const api = '/api/v1/member/delete';
      let params = 'userID=' + row.UserID;
      console.log(params);
      this.axios.post(api, 
                      params,
                      {headers: {'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'}}
      ).then((res) => {
          // this.dialog.show=false;
          // location.reload();
          console.log(res);
          location.reload();
      }).catch((err) => {
          console.log('err:', err.response.data.msg);
      });
      console.log('deleteStudent');
      
    },
    //点击修改
    showUpdate(index) {
      //点击修改后改变状态切换为确定按钮
      this.showEdit[index] = true;

      this.$set(this.showEdit, index, true); //这里要用$set方法，否则页面状态不更新
    },
    //取消修改
    cancelUpdate(index) {
      this.$confirm("取消修改？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$set(this.showEdit, index, false);
        })
        .catch(() => {});
    },
  },
};
</script>
<style scoped>
.notice .el-form{
  display: flex;
  float: left;
}
</style>