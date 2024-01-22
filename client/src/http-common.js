import axios from 'axios';
axios.defaults.withCredentials = true
const HTTP = axios.create({
    baseURL: `http://localhost:8080/root`,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8',
        'Access-Control-Allow-Credentials': true,
    }
})
export default HTTP