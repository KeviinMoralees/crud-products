import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '20s', target: 10 }, // sube a 10 usuarios
    { duration: '40s', target: 10 }, // mantiene 10 usuarios
    { duration: '20s', target: 0 },  // baja a 0
  ],
  thresholds: {
    http_req_failed: ['rate<0.05'],      // menos del 5% de errores
    http_req_duration: ['p(95)<500'],    // P95 menor a 500ms
  },
};

export default function () {
  const res = http.get('http://localhost:8080/api/productos');

  check(res, {
    'status es 200': (r) => r.status === 200,
    'respuesta menor a 500ms': (r) => r.timings.duration < 500,
  });

  sleep(1);
}
