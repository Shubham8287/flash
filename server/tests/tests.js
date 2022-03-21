import http from 'k6/http';
import { sleep, check, randomSeed } from 'k6';

export const options = {
    stages: [
      { duration: '2s', target: 2 },
      { duration: '5s', target: 10 },
      { duration: '2s', target: 0 },
    ],
  };

export default function () {
    randomSeed(123456789);
    const rnd = Math.random();
    console.log(rnd);
  const res = http.get('http://localhost:8080/isAlive');
  console.log("yo,", res.body);
  //check(res, { 'status was 200': (r) => isFine(r.body, prefix) });
  sleep(1);
}
