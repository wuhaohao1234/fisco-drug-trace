<script setup>
import { ElMessage, ElMessageBox } from 'element-plus';
import _ from 'lodash';
import api from '/@/api/api';
import moment from 'moment';
import List from '/@/components/List.vue';
import { getCurrentInstance, ref } from 'vue';

const { proxy } = getCurrentInstance()

const defaultQuery = {
    id: "",

    page: 1,
    size: 10,
    total: 0

}

const query = ref(_.cloneDeep(defaultQuery))

const reset = () => {
    query.value = _.cloneDeep(defaultQuery)
}
// - 编号id
// - 厂家名称
// - 生产批次号
// - 生产日期
// - 有效期
// - 药品成分
// - 药品配方
// - 质量控制和检验结果
// - 运输途径
// - 运输条件
// - 药品存储条件
// - 药品存放地点
// - 药品接受时间
const tableColProps = [
{
      "prop": "id",
      "label": "编号",
      "width": "80px"
    },
    {
        "prop": "producer",
        "label": "厂家名称",
        "width": "120px"
    },
    {
        "prop": "batchNumber",
        "label": "生产批次号",
        "width": "120px"
    },
    {
        "prop": "productionDate",
        "label": "生产日期",
        "width": "120px",
        "type": "time"
    },
    {
        "prop": "expirationDate",
        "label": "有效期",
        "width": "120px",
        "type": "time"
    },
    {
        "prop": "drugComposition",
        "label": "药品成分",
        "width": "120px"
    },
    {
        "prop": "drugFormula",
        "label": "药品配方",
        "width": "120px"
    },
    {
        "prop": "qualityControlAndTestResults",
        "label": "质量控制和检验结果",
        "width": "200px"
    },
    {
        "prop": "flow",
        "label": "流程",
        "width": "120px"

    },
    {
        "prop": "transportationMethod",
        "label": "运输途径",
        "width": "120px"
    },
    {
        "prop": "transportationConditions",
        "label": "运输条件",
        "width": "120px"
    },
    {
        "prop": "dealer",
        "label": "经销商名称",
        "width": "120px"
    },
    {
        "prop": "drugStorageConditions",
        "label": "药品存储条件",
        "width": "120px"
    },
    {
        "prop": "drugStorageLocation",
        "label": "药品存放地点",
        "width": "120px"
    },
    {
        "prop": "drugAcceptanceTime",
        "label": "药品签收时间",
        "width": "120px",
        "type": "time"

    }
]


const tableData = ref([
])

const find = () => {
    console.log("find query: ", query.value);
    let id = 0;
    if (query.value.id != "") {
        id = Number(query.value.id)
    }
    api.drug.buyList({
        id: id,
        page: query.value.page,
        pageSize: query.value.size
    }).then((res) => {
        console.log("find res: ", res);
        tableData.value = res.data
        .filter(item => item.id != 0)
        .map((item) => {
            const data = JSON.parse(item.productionDateStr)
            let transportationData = {}
            if (item.transportationStr != "") {
                transportationData = JSON.parse(item.transportationStr)
            }
            return {
                ...item,
                ...data,
                ...transportationData
            }
        })
        query.value.total = res.total
    })
}

find()

const handleBuy = (row) => {
    ElMessageBox.confirm('确定购买吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        api.drug.buy({
            drugId: row.id,
        }).then((res) => {
            find()
            ElMessage.success('购买成功');
        })
    }).catch(() => {
        ElMessage.info('已取消购买');
    });
}

</script>
<template>
    <div class="min-h-full h-full">
        <List :page="query" :tableColProps="tableColProps" :tableDate="tableData"  @reset="reset" @find="find"  >
            <template #header>
                <el-form :inline="true" :model="query">
                    <el-form-item label="编号: ">
                        <el-input v-model="query.id"></el-input>
                    </el-form-item>
                </el-form>
            </template>
            <template #handler="{ row }">
                <el-button type="danger" @click="handleBuy(row)">购买</el-button>
            </template>
        </List>
      
    </div>
</template>