<script setup>
import { ElMessage, ElMessageBox } from 'element-plus';
import _ from 'lodash';
import moment from 'moment';
import api from '/@/api/api';
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
// 经销商名称
// - 运输途径
// - 运输条件
const tableColProps = [
    {
        "prop": "id",
        "label": "编号",
        "width": "80px"
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
]

const formProps = [
    {
        "prop": "batchNumber",
        "label": "生产批次号",

    },
    {
        "prop": "expirationDate",
        "label": "有效期",
        "type": "date"
    },
    {
        "prop": "drugComposition",
        "label": "药品成分",

    },
    {
        "prop": "drugFormula",
        "label": "药品配方",

    },
    {
        "prop": "qualityControlAndTestResults",
        "label": "质量控制和检验结果",

    },
    {
        "prop": "dealer",
        "label": "上传图片",
        "type": "upload"
    }
]

const formRules = {
    batchNumber: [
        { required: true, message: '请输入生产批次号', trigger: 'blur' },
    ],
    productionDate: [
        { required: true, message: '请选择生产日期', trigger: 'blur' },
    ],
    expirationDate: [
        { required: true, message: '请选择有效期', trigger: 'blur' },
    ],
    drugComposition: [
        { required: true, message: '请输入药品成分', trigger: 'blur' },
    ],
    drugFormula: [
        { required: true, message: '请输入药品配方', trigger: 'blur' },
    ],
    qualityControlAndTestResults: [
        { required: true, message: '请输入质量控制和检验结果', trigger: 'blur' },
    ],
    dealer: [
        { required: true, message: '请输入经销商名称', trigger: 'blur' },
    ],
    transportationMethod: [
        { required: true, message: '请输入运输途径', trigger: 'blur' },
    ],
    transportationConditions: [
        { required: true, message: '请输入运输条件', trigger: 'blur' },
    ],
}

const tableData = ref([
])

const find = () => {
    console.log("find query: ", query.value);
    let id = 0;
    if (query.value.id != "") {
        id = Number(query.value.id)
    }
    api.drug.list({
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


const add = (item) => {
    console.log(item);
    api.drug.add({
        productionDateStr: JSON.stringify(item),
    }).then((res) => {
        ElMessage.success('添加成功');
        find()
    })
}

const saleVisible = ref(false)
const saleForm = ref({
    dealer: "",
    transportationMethod: "",
    transportationConditions: "",
})

const dealerList = ref([])

const handleSale = (row) => {
    api.user.dealer().then((res) => {
        dealerList.value = res.data
        saleVisible.value = true
        saleForm.value = {
            id: row.id,
            dealer: "",
            transportationMethod: "",
            transportationConditions: "",
        }
    })


}

const sale = () => {
    if (saleForm.value.dealer == "") {
        ElMessage.error("经销商名称不能为空")
        return
    }
    if (saleForm.value.transportationMethod == "") {
        ElMessage.error("运输途径不能为空")
        return
    }
    if (saleForm.value.transportationConditions == "") {
        ElMessage.error("运输条件不能为空")
        return
    }
    ElMessageBox.confirm('确定销售吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {

        api.drug.sale({
            drugId: saleForm.value.id,
            dealer: Number(saleForm.value.dealer),
            transportationStr: JSON.stringify({
                transportationMethod: saleForm.value.transportationMethod,
                transportationConditions: saleForm.value.transportationConditions,
            }),
        }).then((res) => {
            find()
            ElMessage.success('销售成功');
            saleVisible.value = false
        })
    }).catch(() => {
        ElMessage.info('已取消销售');
    });
}

</script>
<template>
    <div class="min-h-full h-full">
        <List :page="query" :tableColProps="tableColProps" :tableDate="tableData" :formProps="formProps"
            :formRules="formRules" @reset="reset" @find="find" :showOperate="true" :showAdd="true" @add="add">
            <template #header>
                <el-form :inline="true" :model="query">
                    <el-form-item label="编号: ">
                        <el-input v-model="query.id"></el-input>
                    </el-form-item>
                </el-form>
            </template>
            <template #handler="{ row }">
                <el-button type="success" v-if="row.flow == '生产商'" @click="handleSale(row)">销售</el-button>
            </template>
        </List>
        <el-dialog title="销售" v-model="saleVisible" style="width: 500px;">
            <el-form :model="saleForm" label-width="100px">
                <el-form-item label="经销商名称: " prop="dealer">
                    <el-select v-model="saleForm.dealer" placeholder="请选择经销商名称">
                        <el-option v-for="item in dealerList" :key="item.ID" :label="item.Nickname"
                            :value="item.ID">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="运输途径: " prop="transportationMethod">
                    <el-input v-model="saleForm.transportationMethod" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="运输条件: " prop="transportationConditions">
                    <el-input v-model="saleForm.transportationConditions" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div class="flex justify-center w-full">
                <el-button type="primary" @click="sale" class="w-50">确定</el-button>
            </div>
        </el-dialog>
    </div>
</template>