'use strict'

import autocannon from 'autocannon'

// async/await
async function foo () {
  const [result, resultCls] = await Promise.all([autocannon({
    url: 'http://localhost:8080',
    connections: 100, 
    pipelining: 1,
    duration: 10
  }), 
  autocannon({
    url: 'http://localhost:8082',
    connections: 100, 
    pipelining: 1,
    duration: 10
  })])
  
  const averageLatency = {
    als: result.latency.average,
    cls: resultCls.latency.average
  }
  console.log({averageLatency})
}
foo()