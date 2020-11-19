<template>
  <el-container>
    <el-main>
      <el-card shadow="never">
        <div slot="header" class="clearfix">
          <el-button type="primary" icon="el-icon-plus" @click="showDialog()">添加</el-button>
        </div>
        <el-table :data="tableData.filter( (data) => !search || data.Domain.toLowerCase().includes(search.toLowerCase()))" border style="width: 100%">
          <el-table-column prop="Domain" label="域名"></el-table-column>
          <el-table-column prop="CrtFilePath" label="Crt文件" ></el-table-column>
          <el-table-column prop="KeyFilePath" label="Key文件"></el-table-column>
          <el-table-column prop="HostAgent" align="center" label="是否反向代理" width="120">
            <template slot-scope="scope">
              <el-switch v-model="scope.row.IsHostAgent" disabled></el-switch>
            </template>
          </el-table-column>
          <el-table-column prop="HostAgent" label="反向代理HOSTS" ></el-table-column>
          <el-table-column prop="UpdatedUser" label="最后修改用户"></el-table-column>
          <el-table-column prop="UpdatedAt" label="最后修改时间"></el-table-column>
          <el-table-column align="right" width="180">
            <template slot="header">
              <el-input v-model="search" size="mini" placeholder="输入关键字搜索" />
            </template>
            <template slot-scope="scope">
              <el-button size="mini" @click="handleEdit(scope.$index, scope.row)" style="margin-right: 10px">Edit</el-button>
              <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-main>
    <el-dialog title="添加修改域名信息" :visible.sync="IsDialog" width="30%">
      <el-form ref="form" :model="formdata" label-width="150px" size="mini" label-position="left">
        <el-form-item label="域名：">
          <el-input v-model="formdata.Domain"></el-input>
        </el-form-item>
        <el-form-item label="Crt文件：">
          <el-upload action="#" :multiple="false" :before-upload="CrtUploadFile" :show-file-list="false" accept=".crt">
            <el-button size="small" v-if="formdata.CrtFilePath" type="text">{{ formdata.CrtFilePath }}</el-button>
            <el-button slot="trigger" size="small" type="primary" icon="el-icon-upload">选取文件</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="Key文件：">
          <el-upload action="#" :multiple="false" :before-upload="KeyUploadFile" :show-file-list="false" accept=".key,.pem">
            <el-button size="small" v-if="formdata.KeyFilePath" type="text">{{ formdata.KeyFilePath }}</el-button>
            <el-button slot="trigger" size="small" type="primary" icon="el-icon-upload">选取文件</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="是否启动反向代理：">
          <el-switch v-model="formdata.IsHostAgent"></el-switch>
        </el-form-item>
        <el-form-item label="反向代理HOSTS：">
          <el-input v-model="formdata.HostAgent"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="update">确 定</el-button>
        <el-button @click="IsDialog = false">取 消</el-button>
      </span>
    </el-dialog>
  </el-container>
</template>

<script>
import { Domain } from "@/request/api";
import { Message } from "element-ui";
import { store } from "@/store/index";

export default {
  name: "domain",
  data() {
    return {
      tableData: [],
      search: "",
      IsDialog: false,
      formdata: { Domain: "", CrtFilePath: "", KeyFilePath: "", IsHostAgent: false, HostAgent: "" },
      filelist: { CrtFile: null, KeyFile: null },
      iscreate: true,
    };
  },
  methods: {
    handleEdit(index, row) {
      this.formdata = row;
      this.iscreate = false;
      this.IsDialog = true;
    },
    handleDelete(index, row) {
      this.$confirm("此操作将删除该记录, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        Domain.DeleteDomain(row).then((res) => {
          this.$message({
            type: "success",
            message: res.msg,
          });
          this.getTableData();
        });
      });
    },
    showDialog() {
      this.formdata = { Domain: "", CrtFilePath: "", KeyFilePath: "", IsHostAgent: false, HostAgent: "" };
      this.iscreate = true;
      this.IsDialog = true;
    },
    getTableData() {
      Domain.GetDomainlist().then((res) => {
        this.tableData = res.data;
      });
    },
    update() {
      let data = new window.FormData()
      data.append("CrtFile",this.filelist.CrtFile)
      data.append("KeyFile",this.filelist.KeyFile)
      data.append("Domain",JSON.stringify(this.formdata))

      if (this.iscreate) {
        Domain.InsertDomain(data).then((res) => {
          if (res.code == 0) {
            this.$message( { type: "success", message: res.msg, showClose: true,} );
            this.getTableData();
          }
        });
      } else {
        Domain.UpdateDomain(data).then((res) => {
          if (res.code == 0) {
            this.$message( { type: "success", message: res.msg, showClose: true,} );
            this.getTableData();
          }
        });
      }
      this.IsDialog = false;
    },
    CrtUploadFile(file) {
      console.log(file);
      this.formdata.CrtFilePath = file.name;
      this.filelist.CrtFile = file;
      return false;
    },
    KeyUploadFile(file) {
      console.log(file);
      this.formdata.KeyFilePath = file.name;
      this.filelist.KeyFile = file;
      return false;
    },
  },
  created() {
    this.getTableData();
  }
};
</script>

<style lang="scss">
.domain-search {
  height: 10%;
  width: 100%;
}
</style>