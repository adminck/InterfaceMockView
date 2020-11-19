<template>
  <div class="Param" style="position:relative;">
    <el-tabs v-model="editableTabsValue" type="card" @tab-remove="removeTab" :before-leave="beforeLeave">
      <el-tab-pane v-for="item in ApiJsonInfoList" :key="item.ID" :label=" 'Parameter-' +item.ID" :name="item.ID + ''" closable />
      <el-tab-pane key="add" name="add">
        <span slot="label" style="padding: 8px;font-size:20px;font-weight:bold;"> + </span>
      </el-tab-pane>
    </el-tabs>
    <div v-if="ApiJsonInfo">
      <el-form ref="form" :model="ApiJsonInfo" label-width="120px" size="mini">
        <el-form-item label="参数类型：">
          <el-select v-model="ApiJsonInfo.ParamType" placeholder="参数类型">
            <el-option v-for="(item,index) in ParamType" :label="index" :value="item" :key="index"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="参数：">
          <el-input type="textarea" v-model="ApiJsonInfo.Parameter"></el-input>
        </el-form-item>
        <el-form-item label="返回接口文件：">
          <el-button type="text" @click="outerVisible = true">{{ ApiJsonInfo.JsonFilePath }}</el-button>
        </el-form-item>
        <el-form-item label="是否启动校验：">
          <el-switch v-model="ApiJsonInfo.IsOpen"></el-switch>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="update">保存修改</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-dialog title="json文件编辑" @opened="OpenDialog" :visible.sync="outerVisible">
      <div id="jsoneditor" />
    </el-dialog>
  </div>
</template>

<script>
import { ApiJsonInfo } from '@/request/api'
import JSONEditor from 'jsoneditor'
import 'jsoneditor/dist/jsoneditor.css'

export default {
  props: {
    ApiId: Number
  },
  data() {
    return {
      editableTabsValue: 'default',
      ApiJsonInfoList: [],
      ApiJsonInfo: undefined,
      ParamType: { "Param": 0, "Raw": 1, "FormData": 2, "不校验返回": 8 },
      outerVisible: false,
      jsoneditor:undefined,
    }
  },
  methods: {
    addTab() {
      let index = this.ApiJsonInfoList.push({ ID: '' + (this.ApiJsonInfoList.length + 1), ParamType: 2, Parameter: '', JsonFilePath: 'default.json', IsOpen: false, New: true });
      this.editableTabsValue = (this.ApiJsonInfoList.length + 1) + ''
      this.ApiJsonInfo = this.ApiJsonInfoList[index - 1]
    },
    removeTab(targetName) {
      let targetData;
      for (let index = 0; index < this.ApiJsonInfoList.length; index++) {
        if (targetName === this.ApiJsonInfoList[index].ID + '') {
          targetData = this.ApiJsonInfoList[index]
          if (targetData.New == undefined) {
            this.$confirm('删除接口配置, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              ApiJsonInfo.DeleteApiJsonInfo(targetData).then(res => {
                this.$message({ type: 'info', message: res.msg });
                if (res.code === 0) {
                  this.GetApiJsonData()
                }
              })
            })
          }else{
            this.ApiJsonInfoList.splice(index, 1);
            if (this.ApiJsonInfoList.length >= 1){
              this.ApiJsonInfo = this.ApiJsonInfoList[this.ApiJsonInfoList.length - 1]
            }else{
              this.ApiJsonInfo = undefined
            }
          }
        }
      }
    },
    beforeLeave(currentName, oldName) {
      if (currentName == "add") {
        if (this.ApiId == 0){
          this.$message({ type: 'error', message: "请先选择接口", showClose: true })
          return false
        }
        this.addTab()
        return false
      } else {
        for (let index = 0; index < this.ApiJsonInfoList.length; index++) {
        if (currentName === this.ApiJsonInfoList[index].ID + '') {
          this.ApiJsonInfo = this.ApiJsonInfoList[index]
          this.editableTabsValue = this.ApiJsonInfoList[index].ID + ''
        }
      }
      }
    },
    GetApiJsonData() {
      let data = new FormData()
      data.append("ApiID", this.ApiId)
      ApiJsonInfo.GetApiJsonlist(data).then(res => {
        if (res.code != 0) {
          this.$message({ type: 'error', message: res.msg, showClose: true })
          return
        }
        this.ApiJsonInfoList = res.data
        if(res.data.length >= 1){
          this.ApiJsonInfo = this.ApiJsonInfoList[this.ApiJsonInfoList.length - 1]
          this.editableTabsValue = (res.data[res.data.length-1].ID) + ''
        }else{
          this.ApiJsonInfo = undefined
        }
      })
    },
    update() {
      if (targetData.New === undefined) {
        let jsondata = this.resultInfo
        let InsertData = { ParamType: this.ApiJsonInfo.ParamType, Parameter: this.ApiJsonInfo.Parameter, JsonFilePath: this.ApiJsonInfo.JsonFilePath, IsOpen: this.ApiJsonInfo.IsOpen }
        InsertData.ApiID = this.ApiId
        let data = {json:jsondata,Data:InsertData}
        ApiJsonInfo.InsertApiJsonInfo(data).then((res) => {
          if (res.code == 0) {
            this.$message({ type: "success", message: res.msg, showClose: true, });
            this.GetApiJsonData();
          }
        });
      } else {
        let jsondata = this.resultInfo
        let UpdateData = { ID: parseInt(this.ApiJsonInfo.ID), ParamType: this.ApiJsonInfo.ParamType, Parameter: this.ApiJsonInfo.Parameter, JsonFilePath: this.ApiJsonInfo.JsonFilePath, IsOpen: this.ApiJsonInfo.IsOpen }
        UpdateData.ApiID = this.ApiId
        let data = {json:jsondata,Data:InsertData}
        ApiJsonInfo.UpdateApiJsonInfo(data).then((res) => {
          if (res.code == 0) {
            this.$message({ type: "success", message: res.msg, showClose: true, });
            this.GetApiJsonData();
          }
        });
      }
    },
    OpenDialog () {
      if(this.jsoneditor === undefined){
        this.jsoneditor = new JSONEditor(document.getElementById("jsoneditor"), {mode: 'text',modes: ['view', 'form', 'text', 'code', 'tree']})
      }
      if (this.ApiJsonInfo.New === undefined){
        ApiJsonInfo.GetJSON({JsonFilePath:this.ApiJsonInfo.JsonFilePath}).then((res) => {
          this.jsoneditor.set(res.data)
        })
      }else {
         this.jsoneditor.set(`{}`)
      }
    }
  },
  watch: {
    ApiId: function () {
      this.GetApiJsonData()
    },
    ApiJsonInfo : function(newApiJsonInfo, oldApiJsonInfo) {
      if(newApiJsonInfo != undefined){
        this.editableTabsValue = this.ApiJsonInfo.ID + ''
      }
    }
  },
}
</script>

<style>
.Param .el-dialog__body {
  padding: 0px;
}
.Param .el-dialog__header {
  padding: 10px;
}
.Param .el-dialog__headerbtn {
  top: 15px;
}

#jsoneditor {
  height: 700px;
  width:  -webkit-fill-available;
}
</style>