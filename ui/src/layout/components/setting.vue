<template>
  <el-form
    ref="form"
    label-width="120px"
    label-position="left"
    style="padding: 0 20px"
  >
    <el-alert
      title="以下配置可实时预览，开发者可在 config/index.js 中配置默认值，只有开发环境下开放布局设置"
      type="error"
      :closable="false"
    ></el-alert>
    <el-divider></el-divider>
    <el-form-item :label="$t('usercenter.nightmode')">
      <el-switch
        v-model="theme"
        active-value="dark"
        inactive-value="default"
      ></el-switch>
    </el-form-item>
    <el-form-item :label="$t('usercenter.language')">
      <el-select v-model="lang">
        <el-option label="简体中文" value="zh_CN"></el-option>
        <el-option label="English" value="en"></el-option>
      </el-select>
    </el-form-item>
    <el-divider></el-divider>
    <el-form-item label="主题颜色">
      <el-color-picker v-model="colorPrimary" :predefine="colorList"
        >></el-color-picker
      >
    </el-form-item>
    <el-divider></el-divider>
    <el-form-item label="框架布局">
      <el-select v-model="layout" placeholder="请选择">
        <el-option label="默认" value="default"></el-option>
        <el-option label="通栏" value="header"></el-option>
        <el-option label="经典" value="menu"></el-option>node.js:381Browserslist: caniuse-lite is outdated. Please run: npx browserslist@latest --update-db Why you should do it regularly: https://github.com/browserslist/browserslist#browsers-data-updating
				logger.js:36 INFO  Starting development server...
				output.js:82
























				output.js:102 DONE  Compiled successfully in 119994ms                                                                        22:51:10
				serve.js:262  App running at:
				serve.js:263  - Local:   http://localhost:8005
				serve.js:265  - Network: http://192.168.1.7:8005
				serve.js:287  Note that the development build is not optimized.
				serve.js:288  To create a production build, run yarn build.
        <el-option label="功能坞" value="dock"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="折叠菜单">
      <el-switch v-model="menuIsCollapse"></el-switch>
    </el-form-item>
    <el-form-item label="标签栏">
      <el-switch v-model="layoutTags"></el-switch>
    </el-form-item>
    <el-divider></el-divider>
  </el-form>
</template>

<script>
import colorTool from "@/utils/color";

export default {
  data() {
    return {
      layout: this.$store.state.global.layout,
      menuIsCollapse: this.$store.state.global.menuIsCollapse,
      layoutTags: this.$store.state.global.layoutTags,
      lang: this.$TOOL.data.get("APP_LANG") || this.$CONFIG.LANG,
      theme: this.$TOOL.data.get("APP_THEME") || "default",
      colorList: [
        "#0960bd",
        "#409EFF",
        "#009688",
        "#536dfe",
        "#ff5c93",
        "#c62f2f",
        "#fd726d",
      ],
      colorPrimary:
        this.$TOOL.data.get("APP_COLOR") || this.$CONFIG.COLOR || "#409EFF",
    };
  },
  watch: {
    layout(val) {
      this.$store.commit("SET_layout", val);
      this.$TOOL.data.set("APP_LAYOUT", val);
    },
    menuIsCollapse() {
      this.$store.commit("TOGGLE_menuIsCollapse");
    },
    layoutTags() {
      this.$store.commit("TOGGLE_layoutTags");
    },
    theme(val) {
      val === 'dark' ? document.documentElement.classList.add("dark") : document.documentElement.classList.remove("dark")
      this.$TOOL.data.set("APP_THEME", val);
    },
    lang(val) {
      this.$i18n.locale = val;
      this.$TOOL.data.set("APP_LANG", val);
    },
    colorPrimary(val) {
      document.documentElement.style.setProperty("--el-color-primary", val);
      for (let i = 1; i <= 9; i++) {
        document.documentElement.style.setProperty(
          `--el-color-primary-light-${i}`,
          colorTool.lighten(val, i / 10)
        );
      }
      document.documentElement.style.setProperty(
        `--el-color-primary-darken-1`,
        colorTool.darken(val, 0.1)
      );
      this.$TOOL.data.set("APP_COLOR", val);
    },
  },
};
</script>

<style>
</style>
