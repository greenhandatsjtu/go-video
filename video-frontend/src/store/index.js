import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        videos: []
    },
    // 更新相关信息的操作
    mutations: {
        update(state, videos){
            state.videos = videos
        }
    }
})

export default store
