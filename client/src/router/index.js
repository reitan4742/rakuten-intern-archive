import { createRouter, createWebHistory } from 'vue-router';

import Login from '/src/pages/Login.vue';
import Signup from '/src/pages/Signup.vue';
import Home from '/src/pages/Home.vue';
import Receipt from '/src/pages/Receipt.vue';
import ReceiptResult from '/src/pages/Receiptresult.vue';
import FoodLossReducePerson from '/src/pages/FoodLossReduce/Person.vue';
import FoodLossReduceAll from '/src/pages/FoodLossReduce/All.vue';
import History from '/src/pages/History.vue';

const routes = [
    {
      path: '/signup',
      name: 'Signup',
      component: Signup,
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
    },
    {
      path: '/home',
      name: 'Home',
      component: Home,
    },
    {
      path: '/receipt',
      name: 'Receipt',
      component: Receipt,
    },
    {
      path: '/receiptresult',
      name: 'ReceiptResult',
      component: ReceiptResult,
    },
    {
      path: '/foodlossreduce/person',
      name: 'FoodLossReducePerson',
      component: FoodLossReducePerson,
    },
    {
      path: '/foodlossreduce/all',
      name: 'FoodLossReduceAll',
      component: FoodLossReduceAll,
    },
    {
      path: '/history',
      name: 'History',
      component: History,
    }
  ];
  
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;