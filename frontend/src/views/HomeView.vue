<template>

  <div class="common-layout">
    <el-container>
      <el-header>
        <NavBar />
      </el-header>
      <el-main>
        <el-table :data="filterTableData" style="width: 100%">
          <el-table-column label="ID" prop="ID" width="100" />
          <el-table-column label="书籍名称" prop="title" />
          <el-table-column label="作者" prop="author" width="200" />
          <el-table-column label="库存" prop="quantity" width="100" />
          <el-table-column label="操作" width="200">
            <template #default="scope">
              <el-button size="small" @click="handleEdit(scope.$index, scope.row)">
                Edit
              </el-button>
              <el-button
                size="small"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
              >
                Delete
              </el-button>
            </template>
          </el-table-column>
          <el-table-column align="right">
            <template #header>
              <el-input
                v-model="keyword"
                style="max-width: 600px"
                placeholder="请输入检索关键字"
                class="input-with-select"
                size="large"
              >
                <template #prepend>
                  <el-select v-model="searchType" style="width: 115px" size="large">
                    <el-option label="书籍名称" :value="1" />
                    <el-option label="作者" :value="2" />
                  </el-select>
                </template>
                <template #append>
                  <el-button :icon="Search" />
                </template>
              </el-input>
            </template>
            <template #default="scope">
              {{scope.row.CreatedAt}}
            </template>
            
          </el-table-column>
        </el-table>
        <el-pagination layout="prev, pager, next" :total="total" />
      </el-main>
    </el-container>
  </div>

</template>

<script setup lang="ts">
import { GetBookList } from '@/service/book'

import {ref, onMounted} from 'vue';
import { Search } from '@element-plus/icons-vue'
import NavBar from '@/components/NavBar.vue'

const keyword = ref("");
const searchType = ref(1);
const filterTableData = ref([]);
const total = ref(0);

const handleEdit = (index, row) => {

}

const handleDelete = (index, row) => {
  
}

const queryData = async () => {
  const resp = await GetBookList({
    page: 1,
    page_size: 30,
    keyword: keyword.value,
    searchType: searchType.value,
  })
  if (resp.status == 200) {
    filterTableData.value = resp.data.books
    total.value = resp.data.total
  }
}

onMounted(() => {
  console.log("onMounted")
  queryData()
})

</script>