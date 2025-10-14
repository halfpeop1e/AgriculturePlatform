const store = new Map<string, string>();

export const codeStore = {
  set(email: string, code: string) {
    store.set(email, code);
    setTimeout(() => store.delete(email), 5 * 60 * 1000); // 5分钟后自动清除
  },
  get(email: string) {
    return store.get(email);
  },
  delete(email: string) {
    store.delete(email);
  },
};
