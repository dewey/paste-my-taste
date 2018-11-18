import axios from 'axios'

export default() => {
    return axios.create({
        baseURL: `http://localhost:9999`,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
}