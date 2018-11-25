import Api from '@/services/Api'

export default {
    getPopularArtists(params) {
        return Api().get('/api/lastfm/' + params.username, {
            params: {
                limit: params.limit,
                period: params.period,
                linkArtists: params.linkArtists,
            }
        }).then(response => {
            return response.data
        })
    }
}