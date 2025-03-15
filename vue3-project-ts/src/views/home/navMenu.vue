<template>
  <el-aside width="200px" class="aside"> </el-aside>
</template>

<script lang="ts" setup>
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from "@element-plus/icons-vue";
import { ref, onMounted } from "vue";
import { getMenu } from "../../api";

const navMenu = ref([]);
onMounted(async () => {
  let res = await getMenu();
  console.log(res);
});
</script>