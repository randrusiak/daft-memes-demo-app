import axios from 'axios';

const ENDPOINTS = {
  memes: '/memes'
}

export default class ApiClient {
  constructor () {
    this.API = axios.create({
      baseURL: 'localhost:8080'
    })
  }

  async getMemes() {
    let response = null;
    try {
      response = await this.API.get(ENDPOINTS.memes);
    } catch (e) {
      throw e;
    }
    return response && response.status === 200 ? response.data : null;
  }

}
