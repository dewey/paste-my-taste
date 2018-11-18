import axios from 'axios'

export default() => {
    return axios.create({
        baseURL: `/`,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
}