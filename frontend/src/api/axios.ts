import axios from 'axios';

const API = axios.create({
  baseURL: 'http://localhost:8081',
  withCredentials: false,
});

export default API;
