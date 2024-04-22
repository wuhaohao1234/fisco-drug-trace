import { createRouter, createWebHistory } from "vue-router"

export const routerList = [
  {
    path: "/login",
    name: "Login",
    component: () => import("/@/views/login.vue"),
  },
  {
    path: "/",
    redirect: "/login",
    component: () => import("/@/layout/index.vue"),
    children: [

      {
        path: "user/list",
        name: "UserList",
        component: () => import("/@/views/user/index.vue"),
      },
      {
        path: "info",
        name: "Info",
        component: () => import("/@/views/user/info.vue"),
      },
    ]
  },
  {
    path: "/drug",
    redirect: "/drug/manufacturer",
    component: () => import("/@/layout/index.vue"),
    children: [
      {
        path: "manufacturer",
        name: "DrugManufacturer",
        component: () => import("/@/views/drug/manufacturer.vue"),
      },
      {
        path: "dealer",
        name: "DrugDealer",
        component: () => import("/@/views/drug/dealer.vue"),
      },
      {
        path: "user",
        name: "DrugUser",
        component: () => import("/@/views/drug/user.vue"),
      },
      {
        path: "buy",
        name: "DrugBuy",
        component: () => import("/@/views/drug/buy.vue"),
      }
    ]
  },
]



const router = createRouter({
  history: createWebHistory(),
  routes: routerList
})


export default router
