<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo_login.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">Gin-Vue-Admin</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form :model="registerForm" :rules="rules" ref="registerForm" @keyup.enter.native="submitForm">
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="registerForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input :type="lock === 'lock' ? 'password' : 'text'" placeholder="请输入密码" v-model="registerForm.password">
              <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item prop="rePassword" style="position: relative">
            <el-input :type="lock === 'lock' ? 'password' : 'text'" placeholder="请再次输入密码" v-model="registerForm.rePassword">
              <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%">注 册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import { register } from "@/request/user";

export default {
  name: "Register",
  data() {
    const ratioPassword = (rule, value, callback) => {
      if (value != this.registerForm.password) {
        return callback(new Error("两次密码不同"));
      } else {
        callback();
      }
    };
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
    return {
      lock: "lock",
      registerForm: {
        username: "",
        password: "",
        rePassword: "",
      },
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }],
        rePassword: [{ validator: ratioPassword, trigger: "blur" }],
      },
    };
  },
  methods: {
    submitForm() {
      this.$refs.registerForm.validate((v) => {
        if (v) {
          register(this.registerForm).then((res) => {
            if (res.code == 0) {
              this.$message({
                type: "success",
                message: res.msg,
                showClose: true,
              });
              this.$router.push({ name: "login" });
            }
          });
        } else {
          this.$message({
            type: "error",
            message: "请正确填写注册信息",
            showClose: true,
          });
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
  },
};
</script>

<style scoped lang="scss">
.login-register-box {
  height: 100vh;

  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top: 5vh;

    .logo {
      height: 35vh;
      width: 35vh;
    }
  }
}

.link-icon {
  width: 20px;
  min-width: 20px;
  height: 20px;
  border-radius: 10px;
}

.vPic {
  width: 33%;
  height: 38px;
  float: right !important;

  img {
    cursor: pointer;
    vertical-align: middle;
  }
}

.logo_login {
  width: 100px;
}

#userLayout.user-layout-wrapper {
  height: 100%;
  position: relative;

  & .mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }

  .container {
    width: 100%;
    min-height: 100%;
    background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
    background-size: 100%;
    padding: 110px 0 144px;

    a {
      text-decoration: none;
    }

    .top {
      text-align: center;
      margin-top: -40px;

      .header {
        height: 44px;
        line-height: 44px;
        margin-bottom: 30px;

        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }

        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }

        .title {
          font-size: 33px;
          color: rgba(0, 0, 0, 0.85);
          font-family: Avenir, "Helvetica Neue", Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
      }
    }

    .main {
      min-width: 260px;
      width: 368px;
      margin: 0 auto;
    }

    .footer {
      position: relative;
      width: 100%;
      margin: 40px 0 0 0;
      text-align: center;

      .links {
        margin-bottom: 8px;
        font-size: 14px;

        a {
          color: rgba(0, 0, 0, 0.45);
          transition: all 0.3s;

          & :not(:last-child) {
            margin-right: 80px;
          }
        }
      }
      .copyright {
        color: rgba(0, 0, 0, 0.45);
        font-size: 14px;
      }
    }
  }
}
</style>