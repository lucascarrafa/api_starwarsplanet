import http from 'k6/http';
import { sleep } from 'k6';

export default function() {
  var url = 'http://ip172-18-0-95-bqul5kdim9m000bosnu0-3333.direct.labs.play-with-docker.com/add';
  var payload = JSON.stringify({
	"nome":"Endor",
	"clima":"arid",
	"terreno":"desert",
  });

  var params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
  http.get('http://ip172-18-0-95-bqul5kdim9m000bosnu0-3333.direct.labs.play-with-docker.com/lista');
  sleep(1);
}
