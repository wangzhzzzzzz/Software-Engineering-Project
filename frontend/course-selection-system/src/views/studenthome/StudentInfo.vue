<template>
	<el-card class="box-card">
  <div slot="header" class="clearfix">
    <span>姓名：{{Name}}</span>
  </div>
  <div class="demo-image">
  <div class="block" v-for="fit in fits" :key="fit">
    <!-- <span class="demonstration">{{ fit }}</span> -->
    <el-image
      style="width: 100px; height: 105px"
      :src="url"
      :fit="fit"></el-image>
  </div>
  </div>

  <div class="text item">学号：{{NetID}}</div>
  <div class="text item">班级：6班</div>
  <div class="text item">学院：计算机学院</div>
  <div class="text item">专业：计算机科学与技术</div>
  <div class="text item">账号创建时间：{{StartTime}}</div>
</el-card>
</template>

<script>
	 export default {
	    data() {
	      return {
	        NetID: "12345",
          Name: "哈哈",
          StartTime: "2022-06-01 21:02:14",
          fits: ['cover'],
          url: 'https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/avator.png',
	      };
	    },
      created(){
        this.initData();
      },
	    methods: {
	      async initData(){
          const api = '/api/v1/auth/whoami';
          this.axios.get(api,{}).then((response)=>{
                console.log(response);
                if(response.status === 200){
                    //获得成功响应返回的数据
                    // console.log(response.data.Data.MemberList);
                    // this.tableData = response.data.Data.MemberList;
                    // console.log(this.tableData);
                    console.log(response.data.Data);
                    this.Name = response.data.Data.Name;
                    this.NetID = response.data.Data.NetID;
                    this.StartTime = response.data.Data.StartTime;
                }
            },(error)=>{
                console.log('err:', err.response.data.msg);
            });
        },
	    }
	  }
</script>

<style>
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

  .box-card {
    width: 480px;
  }
</style>
