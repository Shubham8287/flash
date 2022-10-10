import http from 'k6/http';
import { SharedArray } from 'k6/data';
import { sleep, check, randomSeed } from 'k6';

const BASE_URL = 'http://localhost:8080'

export const options = {
    stages: [
      { duration: '1s', target: 2000 },
      { duration: '20s', target: 2000 }
    ],
  };

  const testQueries = new SharedArray('users', function () {
    const f = JSON.parse(open('./test_data.json'));
    return f;
  });
  
export default function () {
  const rndIndex = Math.floor((Math.random()*10000000)%testQueries.length)
  const res = http.get(BASE_URL+'/find?prefix='+ testQueries[rndIndex]);
  //const res = http.get(BASE_URL+'/isAlive');
  check(res, { 'status was 200': (r) => isFine(JSON.parse(r.body), testQueries[rndIndex]) });
 // check(res, { 'status was 200': (r) => true });
  sleep(1);
}

function isFine(body, prefix) {
  if(body.prefix != prefix) {
    return false;
  }
   let allMatches = true 
   body.matches.map(m => {
          allMatches &= m.startsWith(prefix);
    })
    return allMatches;

}