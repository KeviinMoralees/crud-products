import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 50 },    // calentamiento
    { duration: '30s', target: 100 },   // carga baja
    { duration: '1m', target: 250 },    // carga media
    { duration: '1m', target: 500 },    // carga fuerte
    { duration: '1m', target: 500 },    // sostener 500 usuarios
    { duration: '30s', target: 100 },   // bajada controlada
    { duration: '30s', target: 0 },     // finalizar
  ],
  thresholds: {
    http_req_failed: ['rate<0.05'],      // máximo 5% de errores
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
