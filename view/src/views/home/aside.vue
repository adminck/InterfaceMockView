<template>
  <div>
    <el-scrollbar style="height:calc(100vh - 64px)">
      <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="true"
          :default-active="active"
          @select="selectMenuItem"
          active-text-color="#fff"
          class="el-menu-vertical"
          text-color="rgb(191, 203, 217)"
          unique-opened
        >
          <el-menu-item index="person">
            <i class="el-icon-location"></i>
            <span>首页</span>
          </el-menu-item>
          <el-submenu index="1">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>接口域名管理</span>
            </template>
            <el-menu-item index="person">
              <i class="el-icon-document"></i>
              <span>接口管理</span>
            </el-menu-item>
            <el-menu-item index="Domain">
              <i class="el-icon-folder"></i>
              <span>域名管理</span>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  name: "Aside",
  data() {
    return {
      active: "",
      isCollapse: false
    };
  },
  methods: {
    ...mapMutations("history", ["addHistory"]),
    selectMenuItem(index) {
      if (index === this.$route.name) return;
      this.$router.push({ name: index });
    }
  },
  created() {
    this.active = this.$route.name;
    let screenWidth = document.body.clientWidth;
    if (screenWidth < 1000) {
      this.isCollapse = !this.isCollapse;
    }
    this.$bus.on("collapse", item => {
      this.isCollapse = item;
    });
  },
  watch: {
    $route() {
      this.active = this.$route.name;
    }
  },
  beforeDestroy() {
    this.$bus.off("collapse");
  }
};
</script>

<style lang="scss">
.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>