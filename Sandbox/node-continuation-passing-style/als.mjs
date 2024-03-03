import http from 'node:http';
import { AsyncLocalStorage } from 'node:async_hooks';

const asyncLocalStorage = new AsyncLocalStorage();
function useCtx(msg) {
  return asyncLocalStorage.getStore();
}
let idSeq = 0;
http.createServer((req, res) => {
  asyncLocalStorage.run(idSeq++, () => {
    setImmediate(() => {
      const recursiveWork = (n)=>{
        if (n>0){
          useCtx("haha")
          recursiveWork(n-1);
        }
      }
      recursiveWork(5000);
      res.end();
    });
  });
}).listen(8080);