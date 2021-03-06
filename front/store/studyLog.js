export const state = () => ({
  isStart: false,
  currentId: -1
});

export const mutations = {
  isStartStateChange(state) {
    state.isStart = !state.isStart;
  },
  changeCurrentId(state, id) {
    state.currentId = state.currentId === -1 ? id : -1;
  }
};

export const getters = {
  getIsStart(state) {
    return state.isStart;
  },
  getId(state) {
    return state.currentId;
  }
};

export default {
  state,
  mutations,
  getters,
};