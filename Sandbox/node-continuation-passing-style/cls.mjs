import http from 'node:http';
import { createNamespace } from 'continuation-local-storage';

const ns = createNamespace("my-namespace");
function useCtx() {
return ns.get('id');
}
let idSeq = 0;
http.createServer((req, res) => {
  ns.run(() => {
    console.log(ctx)
    ns.set('id', idSeq++);
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
}).listen(8082);