<template>
  <div id="app" @click="clicked">
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { store } from "@/store/index";
import { Message } from "element-ui";
export default {
  name: "app",
  components: {},
  methods: {
    ...mapActions("user", ["LoginOut"]),
    clicked() {
      const token = store.getters["user/token"];
      if (token) {
        const expiresAt = store.getters["user/expiresAt"];
        const nowUnix = new Date().getTime();
        const hasExpires = expiresAt - nowUnix < 0;
        if (hasExpires) {
          Message({
            showClose: true,
            message: "登录超时",
            type: "error",
          });
          this.LoginOut();
        }
      }
    },
  },
};
</script>

<style lang="scss">
@import "@/style/main.scss";
@import "@/style/base.scss";
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}
</style>
