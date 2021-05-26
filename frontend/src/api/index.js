import axios from 'axios';

const ENDPOINTS = {
  memes: '/memes',
  meme: '/meme'
}

export default class ApiClient {
  constructor () {
    this.API = axios.create({
      baseURL: 'http://localhost:8080'
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

  async addMeme(title, file) {
    let response = null;
    const formData = new FormData();
    formData.append('title',title);
    formData.append('file',file);
    const config = {
      headers: {
        'content-type': 'multipart/form-data'
      }
    }
    try {
      response = await this.API.post(ENDPOINTS.meme, formData, config);
    } catch (e) {
      return null;
    }
    return response && response.status === 201 ? response.data : null;
  }

  async deleteMeme(id) {
    let response = null;
    try {
      response = await this.API.delete(`${ENDPOINTS.meme}/${id}`);
    } catch (e) {
      return null
    }
    return response && response.status === 200 ? response.data : null;
  }

}
