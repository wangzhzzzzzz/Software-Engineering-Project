<template>
  <div class="notice">
	  <el-form  label-width="80px">
	    <el-form-item label="查询内容:">
	      <el-input v-model="name"></el-input>
	    </el-form-item>
	  </el-form>
	  <el-button type="primary" @click="search()" style="margin-left: 15px;">查询</el-button>
	   <el-button type="success" @click="AddNotice()" icon="el-icon-plus" style="position: absolute;right: 15px;">添加课程信息</el-button>
      <el-table border :fit="true" 
        :data="tableData.slice((currentPage - 1) * pageSize, currentPage * pageSize)">
        <el-table-column prop="CourseID" sortable label="CourseID">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.CourseID }}</span>
            <el-input
              type="text"
              v-model="row.CourseID"
              v-else
              placeholder="请输入ID"
            ></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="Name" sortable label="课程名称">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.Name }}</span>
            <el-input
              type="text"
              v-model="row.Name"
              v-else
              placeholder="请输入课程名称"
            ></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="Capacity" sortable label="课程容量">
           <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.Capacity }}</span>
            <el-input
              type="text"
              v-model="row.Capacity"
              v-else
              placeholder="请输入课程容量"
            ></el-input>
          </template>
        </el-table-column>
        
        <el-table-column prop="CapSelected" sortable label="已选人数">
          <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">{{ row.CapSelected }}</span>
            <el-input
              type="text"
              v-model="row.CapSelected"
              v-else
              placeholder=""
            ></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="startime" sortable label="创建时间">
           <template slot-scope="{ row, $index }">
            <span v-if="!showEdit[$index]">
              {{row.startime}}
            </span>
            <el-date-picker
              type="datetime"
              value-format="yyyy-MM-dd HH:mm:ss"
              v-model="row.startime"
              v-else
              placeholder="创建时间"
            >
            </el-date-picker>
          </template>
        </el-table-column>
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
    <Addcourse :dialog="dialog" @update="getProfile($event)" />
  </div>
</template>
<script>
import Addcourse from "../../components/addcourse/Addcourse.vue";
import pageNation from '../../components/pagenation/pageNation';
export default {
  name: "Notice",
  components: { Addcourse,pageNation }, 
  data() {
    return {
	  name:"",
      tableData: [
        {
          CourseID:1001,
          Name:"程序设计",
          Capacity:"120",
          CapSelected:"88",
          startime:"2022-3-01 8:12:35",
          // endtime:"未注销"
        },
        {
          CourseID:1002,
          Name:"离散数学",
          Capacity:"90",
          CapSelected:"70",
          startime:"2022-3-01 8:12:35",
          // endtime:"已注销"
        },
        
      ],
      showEdit: [], //控制页面修改显示隐藏
      currentPage: 1, //当前页数
      pageSize: 10, //每页的数据
      dialog: {
        show: false, //是否显示弹出框
        title: "", //弹出框的title
      },
    };
  },
  methods: {
    //分页
    pageChange (item) {
	      this.currentPage = item.page_index;
	      this.pageSize = item.page_limit;
    },
    AddNotice() {
      this.dialog = {
        show: true,
        title: "添加课程信息", 
      };
    },
    getProfile(data) {
      console.log(data)
    },
     //编辑后点击确定按钮
    submit(index, row) {
      console.log(index, row);
      this.$set(this.showEdit, index, false);
    },
    //点击删除
    del(index, row) {
      console.log(index, row);
    },
    //点击修改
    showUpdate(index) {
      this.showEdit[index] = true;
      //点击修改后改变状态切换为确定按钮
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