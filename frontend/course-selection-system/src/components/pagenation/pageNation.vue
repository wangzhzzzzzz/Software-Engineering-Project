<template>
  <div class="pagenation">
    <el-pagination class="text-center"  
     @size-change="handleSizeChange" 
     @current-change="handleCurrentChange"
     :current-page="page_index"
     :page-sizes="pageSizes" 
     :page-size="page.limit"
      background
      style="text-align: right;margin:12px 5px;"
     layout="total, sizes, prev, pager, next, jumper" 
     :total="total">
    </el-pagination>
   </div>
</template>

<script>
export default {
  props: {
    total: {
      type: Number
    }, // 总条数
    page_index: {
				type: Number
		},
  },
  data() {
    return {
      pageSizes:[5,10,20,40,60,80,100], 
      page: {
        page_index: 1,
        page_limit: 10
      }
    };
  },
  methods: {
    // // 每页查看条数变化
    handleSizeChange(val) {
      this.page.page_limit= val;
      this.$emit('pageChange', this.page);
    },
    // 当前页码变化
    handleCurrentChange(val) {
      this.page.page_index= val;
      this.$emit('pageChange', this.page);
    }
  },

}
</script>
<style>
/* 全局样式 */
.el-pagination .el-select .el-input{
    width: 100px;
    margin: -12px 5px;
}
</style>