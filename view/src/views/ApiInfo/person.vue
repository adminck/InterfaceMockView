<template>
  <el-row :gutter="20">
    <el-col :span="7">
      <el-card shadow="never">
        <div slot="header" class="clearfix">
          <span>接口列表</span>
          <el-button style="float: right; padding: 3px 0" type="text" @click="outerVisible = true">新增接口
          </el-button>
        </div>
        <el-input placeholder="输入关键字进行过滤" v-model="filterText"></el-input>
        <el-tree :data="data" :props="{label: 'Name'}" :filter-node-method="filterNode" ref="tree" style="margin-top: 10px" highlight-current @node-click="handleNodeClick">
          <el-tooltip class="item" effect="dark" placement="bottom" slot-scope="{ node, data }">
            <div slot="content">
              {{ "接口名称：" + data.Name }}<br />
              {{ "接口地址：" + data.Path }}<br />
              {{ "接口域名：" + data.Domain }}
            </div>
            <span class="custom-tree-node">
              <span>{{ node.label }}</span>
              <span>
                <el-button type="primary" size="mini" @click.stop="() => Update(node, data)">修 改</el-button>
                <el-button type="danger" size="mini" @click.stop="() => remove(node, data)">删 除</el-button>
              </span>
            </span>
          </el-tooltip>
        </el-tree>
      </el-card>
    </el-col>
    <el-col :span="17">
      <el-card shadow="never">
        <div slot="header" class="clearfix">
          <span>接口策略信息</span>
        </div>
        <apijson v-bind:ApiId="checkedApiId" />
      </el-card>
    </el-col>
    <el-dialog title="新增API 接口" :visible.sync="outerVisible" width="20%">
      <el-form ref="form" :model="apiinfo" label-width="100px" size="mini">
        <el-form-item label="接口名称：">
          <el-input v-model="apiinfo.Name"></el-input>
        </el-form-item>
        <el-form-item label="接口地址：">
          <el-input v-model="apiinfo.Path"></el-input>
        </el-form-item>
        <el-form-item label="接口域名：">
          <el-input v-model="apiinfo.Domain"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="SetApiInfo">确定</el-button>
        <el-button @click="outerVisible = false">取 消</el-button>
      </div>
    </el-dialog>
  </el-row>
</template>
<script>
import { Message } from 'element-ui';
import { store } from '@/store/index'
import { ApiInfo, Domain } from '@/request/api'
import apijson from './apijsoninfo/apijson.vue'


export default {
  name: 'Person',
  data() {
    return {
      data: [],
      apiinfo: { Name: '', Path: '', Domain: ''},
      filterText: '',
      outerVisible: false,
      iscreate: true,
      checkedApiId : 0, 
    }
  },
  components: {
    apijson
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val);
    }
  },
  methods: {
    filterNode(value, data) {
      if (!value) return true;
      return data.Name.indexOf(value) !== -1 || data.Path.indexOf(value) !== -1 || data.Domain.indexOf(value) !== -1 ;
    },
    addapi() {
      this.iscreate = true
      this.apiinfo = { Name: '', Path: '', Domain: '', IsOpen: false }
      this.outerVisible = true
    },
    getapiData() {
      ApiInfo.GetApiInfolist().then(res => {
        if (res.code != 0) {
          this.$message({ type: 'error', message: res.msg, showClose: true })
          return
        }
        this.data = res.data
        this.checkedApiId = 0
      })
    },
    remove(node, data) {
      this.$confirm('删除接口' + data.Name + ', 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        ApiInfo.DeleteApiInfo(data).then(res => {
          this.$message({ type: 'info', message: res.msg });
          this.getapiData()
        })
      }).catch(() => {
        this.$message({ type: 'info', message: '已取消删除' });
      });
    },
    handleNodeClick(data) {
      this.checkedApiId = data.ID
    },
    Update(node, data) {
      this.iscreate = false
      this.outerVisible = true
      this.apiinfo = data
    },
    SetApiInfo() {
      if (this.iscreate) {
        ApiInfo.InsertApiInfo(this.apiinfo).then((res) => {
          if (res.code == 0) {
            this.$message({ type: "success", message: res.msg, showClose: true, });
            this.getapiData();
          }
        });
      } else {
        ApiInfo.UpdateApiInfo(this.apiinfo).then((res) => {
          if (res.code == 0) {
            this.$message({ type: "success", message: res.msg, showClose: true, });
            this.getapiData();
          }
        });
      }
      this.outerVisible = false;
    }
  },
  created() {
    this.getapiData()
  }
}
</script>
<style lang="scss">
.clearfix:after {
  clear: both;
}

.el-tree-node__content {
  height: 40px;
}

.el-card__body {
  padding: 0px;
  margin: 10px;
}

.el-tree--highlight-current .el-tree-node.is-current > .el-tree-node__content {
  background-color: #8abeec;
}

.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 16px;
  padding-right: 8px;
}

.el-select {
  width: 100%;
}
</style>