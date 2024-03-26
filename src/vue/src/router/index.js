import {createRouter, createWebHistory} from 'vue-router'
import login from '../components/login'
import register from '../components/register'
import policeQuery from '../components/policeQuery'
import policeAdd from '../components/policeAdd'
import houseQuery from '../components/houseQuery'
import houseAdd from '../components/houseAdd'
import creditQuery from '../components/creditQuery'
import creditAdd from '../components/creditAdd'
import contractQuery from '../components/contractQuery'
import contractAdd from '../components/contractAdd'
import transactionQuery from '../components/transactionQuery'
import transactionAdd from '../components/transactionAdd'
import userFrame from '../components/userFrame'

const routes = [
    {
        path: '/',
        name: 'login',
        component: login
    },
    {
        path: '/register',
        name: 'register',
        component: register
    },
    {
        path: '/policeQuery',
        name: 'policeQuery',
        component: policeQuery
    },
    {
        path: '/policeAdd',
        name: 'policeAdd',
        component: policeAdd
    },
    {
        path: '/houseQuery',
        name: 'houseQuery',
        component: houseQuery
    },
    {
        path: '/houseAdd',
        name: 'houseAdd',
        component: houseAdd
    },
    {
        path: '/creditQuery',
        name: 'creditQuery',
        component: creditQuery
    },
    {
        path: '/creditAdd',
        name: 'creditAdd',
        component: creditAdd
    },
    {
        path: '/user',
        component: userFrame,
        children: [
            {
                path: 'contractQuery',
                name: 'contractQuery',
                component: contractQuery
            },
            {
                path: 'contractAdd',
                name: 'contractAdd',
                component: contractAdd
            },
            {
                path: 'transactionQuery',
                name: 'transactionQuery',
                component: transactionQuery
            },
            {
                path: 'transactionAdd',
                name: 'transactionAdd',
                component: transactionAdd
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router