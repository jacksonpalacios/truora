import axios from 'axios'
import configService from './config'

axios.defaults.baseURL = configService.apiUrl

const goApi = axios

export default goApi
